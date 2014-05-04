package superspades

import (
    "time"
    "math/rand"
)

// Deal cards to all 4 players evenly.
func (g Game) Deal() Game {
    rand.Seed(time.Now().UTC().UnixNano())

    for i, c := range rand.Perm(52) {
        g.cardPositions[c] = uint8((i / 13) + cardInHand1) // offset by constant value
    }

    return g
}
