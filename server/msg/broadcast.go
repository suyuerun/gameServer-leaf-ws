package msg

type Broadcast struct {
	MsgId   string
	Message string `json:"message"`
	Time    int
}
