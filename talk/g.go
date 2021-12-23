package talk

type RobitI interface {
	Send(interface{}) error
}