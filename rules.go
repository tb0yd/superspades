package superspades

//import "fmt"

const (
    INVALID_UNLESS_FORCED = iota
    VALID
)

// Returns the winning move between four cards. leading is zero-index
// of first card played.
func winningCardNum(moves [4]Card, leading uint8) uint8 {
    h, i := leading, (leading+1) % 4

    for j:=0; j<3; i, j = (i+1) % 4, j + 1 {
        if beats(int(cardToInt(moves[i])), int(cardToInt(moves[h]))) {
            h = i
        }
    }

    return h
}


func canPlayMove(card int, hand []int, spadesBroken bool, leading bool, lastSuit int) bool {
    isInHand := false

    for _, c := range hand {
        if c == card {
            isInHand = true
        }
    }

    if !isInHand {
        return false
    }

    if forced(hand, spadesBroken, leading, lastSuit) {
        return true
    }

    return moveType(card, hand, spadesBroken, leading, lastSuit) == VALID
}

func moveType(card int, hand []int, spadesBroken bool, leading bool, lastSuit int) int {
    if !leading && (card/13 != lastSuit) {
        return INVALID_UNLESS_FORCED
    }

    if !spadesBroken && (card/13 == 3) {
        return INVALID_UNLESS_FORCED
    }

    return VALID
}

func forced(hand []int, spadesBroken bool, leading bool, lastSuit int) bool {
    for _, card := range hand {
        if moveType(card, hand, spadesBroken, leading, lastSuit) == VALID {
            return false
        }
    }

    return true
}

func (g Game) spadesBroken() bool {
    for i := 39; i < 52; i++ {
        if g.cardPositions[i] > cardInHand4 {
            return true
        }
    }

    return false
}

func (g Game) leading() bool {
    for i := 0; i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            return false
        }
    }

    return true
}

func (g Game) lastSuit() int {
    var oxOnTable [4]int

    for i := 0; i < 52; i++ {
        if g.cardPositions[i] >= cardOnTableForPlayer1 && g.cardPositions[i] <= cardOnTableForPlayer4 {
            oxOnTable[g.cardPositions[i] - cardOnTableForPlayer1] = i
        }
    }

    for i,j := (g.CurrentPlayer+1) % 4,0; j < 4; i,j = (i+1) % 4, j + 1 {
        zxPrevPlayer := (i+3) % 4

        if oxOnTable[zxPrevPlayer] == 0 && oxOnTable[i] != 0 {
            return (oxOnTable[i] - 1) / 13
        }
    }

    return 4
}

