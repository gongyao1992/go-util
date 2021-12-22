package dingtalk

type RobitI interface {
	Send(interface{}) error
}