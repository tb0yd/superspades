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

