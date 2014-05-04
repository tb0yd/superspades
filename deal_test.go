package superspades

import "testing"

func TestDeal(t *testing.T) {
    g := NewGame()
    p := [4]int{}

    g = g.Deal()

    for i := 0; i < 52; i++ {
        p[g.cardPositions[i] - cardInHand1]++
    }

    for i := 0; i < 4; i++ {
        if p[i] != 13 {
            t.Errorf("player doesn't have 13")
        }
    }
}
