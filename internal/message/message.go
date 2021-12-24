package message

type Message interface {
	PrettyJson() []byte
}
