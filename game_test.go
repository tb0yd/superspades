package superspades

import "testing"

func TestGameState(t *testing.T) {
    g := NewGame()

    if state, _ := g.State(); state != GameNotDealt {
        t.Errorf("should be GameNotDealt")
    }

    g = g.Deal()

    if state, _ := g.State(); state != GameNoBids {
        t.Errorf("should be GameNoBids")
    }
}
