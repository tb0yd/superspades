package superspades

import "testing"

func TestNewGame(t *testing.T) {
    if NewGame() != (Game{}) {
        t.Errorf("NewGame() != Game{}")
    }
}
