package game

type testPlayer struct {
}

func (t testPlayer) GetMove(moves [][2]string) (string, string) {
	return "", ""
}

func (t testPlayer) State() map[string]any {
	return map[string]any{}
}

func NewTestState(fen string) *State {
	s := LoadFEN(fen, testPlayer{}, testPlayer{})

	return s
}
