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

	if score := bowling.Score(); score != 0 {
		t.Errorf("Score() = %d, want 0", score)
	}
}
