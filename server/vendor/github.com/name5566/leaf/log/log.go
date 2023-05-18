package log

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// levels
const (
	debugLevel   = 0
	releaseLevel = 1
	warnLevel = 2
	errorLevel   = 3
	fatalLevel   = 4
)

const (
	printDebugLevel   = "\x1b[36m[debug] "
	printReleaseLevel = "\x1b[32m[relea] "
	printWarnLevel    = "\x1b[33m[warn ] "
	printErrorLevel   = "\x1b[31m[error] "
	printFatalLevel   = "\x1b[35m[fatal] "
)

type Logger struct {
	level        int
	baseLogger   *log.Logger
	errorLogger  *log.Logger
	stdOutLogger *log.Logger
	logStdOut    bool
}

func New(strLevel string, pathname string, logConsole bool) (*Logger, error) {
	// level
	var level int
	switch strings.ToLower(strLevel) {
	case "debug":
		level = debugLevel
	case "release":
		level = releaseLevel
	case "warn":
		level = warnLevel
	case "error":
		level = errorLevel
	case "fatal":
		level = fatalLevel
	default:
		return nil, errors.New("unknown level: " + strLevel)
	}

	// new
	logger := new(Logger)
	logger.level = level
	logger.baseLogger = newLogger("2006-01-02.log", pathname)
	logger.errorLogger = newLogger("2006-01-02_error.log", pathname)

	logger.logStdOut = logConsole
	logger.stdOutLogger = newLogger("", "")

	return logger, nil
}

func newLogger(filename string, pathname string) *log.Logger {
	if pathname != "" {
		//now := time.Now()
		filename := fmt.Sprintf(filename)

		file, err := NewFile(path.Join(pathname, filename), nil)

		if err != nil {
			return nil
		}
		return log.New(file, "", log.LstdFlags)
	} else {
		return log.New(os.Stdout, "", log.LstdFlags)
	}
}

// It's dangerous to call the method on logging
func (logger *Logger) Close() {

	logger.baseLogger = nil
}

func (logger *Logger) doPrintf(level int, printLevel string, format string, a ...interface{}) {
	if level < logger.level {
		return
	}
	if logger.baseLogger == nil {
		panic("logger closed")
	}

	format = printLevel + format + "\x1b[0m"
	if level == errorLevel || level == fatalLevel {
		logger.baseLogger.Printf(format, a...)
		logger.errorLogger.Printf(format, a...)
	} else {
		logger.baseLogger.Printf(format, a...)
		if logger.logStdOut {
			logger.stdOutLogger.Printf(format, a...)
		}
	}

	if level == fatalLevel {
		os.Exit(1)
	}
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	logger.doPrintf(debugLevel, printDebugLevel, format, a...)
}

func (logger *Logger) Release(format string, a ...interface{}) {
	logger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
}
func (logger *Logger) Warn(format string, a ...interface{}) {
	logger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	logger.doPrintf(errorLevel, printErrorLevel, format, a...)
}

func (logger *Logger) Fatal(format string, a ...interface{}) {
	logger.doPrintf(fatalLevel, printFatalLevel, format, a...)
}

var gLogger, _ = New("debug", "", false)

// It's dangerous to call the method on logging
func Export(logger *Logger) {
	if logger != nil {
		gLogger = logger
	}
}

func Debug(format string, a ...interface{}) {
	gLogger.Debug(format, a...)
}

func Release(format string, a ...interface{}) {
	gLogger.Release(format, a...)
}
func Warn(format string, a ...interface{}) {
	gLogger.Warn(format, a...)
}

func Error(format string, a ...interface{}) {
	gLogger.Error(format, a...)
}

func Fatal(format string, a ...interface{}) {
	gLogger.Fatal(format, a...)
}

func Close() {
	gLogger.Close()
}
