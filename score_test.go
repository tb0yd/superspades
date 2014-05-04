package superspades

import "testing"

func TestNewScore(t *testing.T) {
    var s Score

    s = NewScore(153)
    if s.Books != 150 {
        t.Errorf("expected 150 main, got %d\b", s.Books)
    }
    if s.Bags != 3 {
        t.Errorf("expected 3 bags, got %d\b", s.Bags)
    }

    s = NewScore(-47)
    if s.Books != -40 {
        t.Errorf("expected -40 main, got %d\b", s.Books)
    }
    if s.Bags != 7 {
        t.Errorf("expected 7 bags, got %d\b", s.Bags)
    }
}

func TestScore(t *testing.T) {
    var s Score

    s = NewScore(0)
    s = s.Add(5, 5, 0, 14)
    if s.ToInt() != 150 {
        t.Errorf("expected 150, got %d\n", s.ToInt())
    }

    s = NewScore(0)
    s = s.Add(3, 14, 3, 2)
    if s.ToInt() != -84 {
        t.Errorf("expected -84, got %d\n", s.ToInt())
    }

    s = NewScore(0)
    s = s.Add(2, 4, 0, 14)
    if s.ToInt() != 60 {
        t.Errorf("expected %d, got %d\n", 60, s.ToInt())
    }

    s = NewScore(0)
    s = s.Add(2, 14, 2, 4)
    if s.ToInt() != -60 {
        t.Errorf("expected %d, got %d\n", -60, s.ToInt())
    }

    s = NewScore(0)
    s = s.Add(5, 1, 7, 14)
    if s.ToInt() != -191 {
        t.Errorf("expected -191, got %d\n", s.ToInt())
    }
}
