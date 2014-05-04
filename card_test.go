package superspades

import (
   // "fmt"
    "github.com/jteeuwen/deck"
    "testing"
)

func TestAceOfSpades(t *testing.T) {
    var c Card

    c = card{deck.NewCard(deck.Spades, 0)}

    if !c.SameSuit(card{deck.NewCard(deck.Spades, 0)}) {
        t.Errorf("wasn't same suit")
    }

    if !c.IsSpades() {
        t.Errorf("wasn't spades")
    }

    if (card{}).InOrder(c.Order()) != c {
        t.Errorf("order didn't return c.")
    }

    for i := 0; i < 4; i++ {
        for j := 0; j < 13; j++ {
            if i != 3 && j != 0 {
                c2 := deck.NewCard(deck.Suit(i), uint8(j))

                if c.Order() == (card{c2}).Order() {
                    t.Errorf("order wasn't unique")
                }

                if !c.Beats(card{c2}) {
                    t.Errorf("didn't beat %s", c2.String())
                }
            }
        }
    }
}
