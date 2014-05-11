package superspades

//import "fmt"

const (
    invalidUnlessForced = iota
    valid
)

// Returns the winning move between four cards. leading is zero-index
// of first card played.
func winningCardNum(moves [4]int, leading uint8) uint8 {
    h, i := leading, (leading+1) % 4

    for j:=0; j<3; i, j = (i+1) % 4, j + 1 {
        if beats(moves[i], moves[h]) {
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

    return moveType(card, hand, spadesBroken, leading, lastSuit) == valid
}

func moveType(card int, hand []int, spadesBroken bool, leading bool, lastSuit int) int {
    if !leading && (card/13 != lastSuit) {
        return invalidUnlessForced
    }

    if !spadesBroken && (card/13 == 3) {
        return invalidUnlessForced
    }

    return valid
}

func forced(hand []int, spadesBroken bool, leading bool, lastSuit int) bool {
    for _, card := range hand {
        if moveType(card, hand, spadesBroken, leading, lastSuit) == valid {
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

