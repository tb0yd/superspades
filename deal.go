package superspades

import (
    "time"
    "math/rand"
)

// Deal cards to all 4 players evenly. Random number generator is automatically seeded.
func (g Game) ShuffleAndDeal() Game {
    rand.Seed(time.Now().UTC().UnixNano())
    return g.Deal()
}

// Deal cards to all 4 players evenly. User is responsible for seeding random number generator.
func (g Game) Deal() Game {
    for i, c := range rand.Perm(52) {
        g.cardPositions[c] = uint8((i / 13) + cardInHand1) // offset by constant value
    }

    return g
}
