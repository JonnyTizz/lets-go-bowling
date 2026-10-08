package bowling

import "testing"

func TestVeryBoringGame(t *testing.T) {
	bowling := NewBowling()

	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)
	bowling.Roll(0)

	expected := 0
	if score := bowling.Score(); score != expected {
		t.Errorf("Score() = %d, want %d", score, expected)
	}
}
