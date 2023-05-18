package msg

type Login struct {
	MsgId     string
	Username  string
	Password  string
	TokenText string
	Time      int
}
