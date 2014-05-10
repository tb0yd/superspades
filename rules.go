package superspades

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
