package superspades

import "math"

// Struct representing a team's overall score, split into number of Books
// and number of Bags.
type Score struct {
    Books int
    Bags int
}

// Return the scores of a game for teams 1 and 2, respectively. Returns 2
// zeroed scores if the game isn't over yet.
func (g Game) Scores() (s1 Score, s2 Score) {
    if !g.IsOver() {
        return
    }

    s1 = Score{}.Add(g.Books[0], g.Bids[0], g.Books[2], g.Bids[2])
    s2 = Score{}.Add(g.Books[1], g.Bids[1], g.Books[3], g.Bids[3])

    return
}

// Add the result of a game for a single team to the overall score of
// that team. You must pass in the individual bids and numbers of books
// for both players on the team. It doesn't matter what order the two
// players are added in.
func (s Score) Add(got1 uint8, bid1 uint8, got2 uint8, bid2 uint8) Score {
    bagBuster := func() {
        if s.Bags >= 10 {
            s.Books -= 100
            s.Bags = s.Bags - 10
        }
    }

    if bid1 == 14 {
        if got1 == 0 {
            s.Books += 100
        } else {
            s.Books -= 100
        }
        bid1 = 0
    }

    if bid2 == 14 {
        if got2 == 0 {
            s.Books += 100
        } else {
            s.Books -= 100
        }
        bid2 = 0
    }


    if (got1 + got2) >= (bid1 + bid2) {
        s.Bags = s.Bags + int((got1 + got2) - (bid1 + bid2))
        s.Books = s.Books + int((bid1 + bid2) * 10)
    } else {
        s.Books = s.Books - int((bid1 + bid2) * 10)
    }

    bagBuster()

    return s
}

// Display as an integer. An example of the format:
//
//    153  -> (15)(3) -> 15 books, 3 bags.
//    -47  -> (-4)(7) -> -4 books, 7 bags.
//
// Note: you can't add two scores together by adding their ToInt() values. You
// have to use Add().
func (s Score) ToInt() int {
    if s.Books > 0 {
        return s.Books + s.Bags
    }

    return (s.Books) - s.Bags
}

// Create a Score from the integer display format. Format is the same as in ToInt().
func NewScore(i int) Score {
    return Score{i/10*10, int(math.Abs(float64(i%10)))}
}
