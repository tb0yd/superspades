package superspades

import (
   // "fmt"
    "testing"
)

func TestAceOfSpades(t *testing.T) {
    //the highest value card
    aceOfSpadesNo := 51
    aceOfSpades   :=  Card{"♠", "A"}

    if !sameSuit(aceOfSpadesNo, 51) {
        t.Errorf("wasn't same suit")
    }

    if !isSpades(aceOfSpadesNo) {
        t.Errorf("wasn't spades")
    }

    if intToCard(aceOfSpadesNo) != aceOfSpades {
        t.Errorf("cardToInt didn't return c.")
    }

    for i := 0; i < 51; i++ {
        if cardToInt(aceOfSpades) == cardToInt(intToCard(i)) {
            t.Errorf("cardToInt wasn't unique")
        }

        if !beats(aceOfSpadesNo, i) {
            t.Errorf("didn't beat %s", intToCard(i))
        }
    }

    if aceOfSpades.Suit != "♠" || aceOfSpades.Value != "A" {
        t.Errorf("aceOfSpades was %s %s\n", aceOfSpades.Suit, aceOfSpades.Value)
    }
}
