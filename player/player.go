package player

type Player interface {
	GetMove([][2]string) (string, string)
	State() map[string]any
}
