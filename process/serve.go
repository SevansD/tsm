package process

type messageType int32

const (
	MessageType_Info     messageType = 0
	MessageType_Warning  messageType = 1
	MessageType_Error    messageType = 2
	MessageType_Critical messageType = 3
)

type Client struct {
	uid  string
	host string
}

type ClientMessage struct {
	clientUID  string `json:client_uid`
	clientHost string `json:client_host`
	message    struct {
		messageType messageType `json:type`
		messageText string      `json:text`
	} `json:message`
}

type TelegrammMessage struct {
	message string
}