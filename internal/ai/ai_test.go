package ai

import "testing"

func TestChooseMove(t *testing.T) {
	tests := []struct {
		name       string
		cells      []string
		wantMove   int
		wantReason string
	}{
		{
			name:       "AI wins top row",
			cells:      []string{"O", "O", "", "X", "X", "", "", "", ""},
			wantMove:   3,
			wantReason: "win",
		},
		{
			name:       "AI blocks X",
			cells:      []string{"X", "X", "", "", "", "", "", "", "O"},
			wantMove:   3,
			wantReason: "block",
		},
		{
			name:       "AI chooses center",
			cells:      []string{"", "", "", "", "", "", "", "", ""},
			wantMove:   5,
			wantReason: "center",
		},
	}

	for _, test := range tests {
		move, reason := ChooseMove(test.cells)

		if move != test.wantMove || reason != test.wantReason {
			t.Fatalf(
				"%s: got move=%d reason=%s, want move=%d reason=%s",
				test.name,
				move,
				reason,
				test.wantMove,
				test.wantReason,
			)
		}
	}
}