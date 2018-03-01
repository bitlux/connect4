package board

import "testing"

func TestBoard_Move(t *testing.T) {
	b := New(2, 3)
	tests := []struct {
		move    int
		wantErr bool
	}{
		{0, true},
		{4, true},
		{5, true},
		{1, false},
		{2, false},
		{3, false},
		{1, false},
		{1, true},
		{1, true},
		{2, false},
	}
	for _, tt := range tests {
		if err := b.Move(tt.move, 1); (err != nil) != tt.wantErr {
			b.Print()
			t.Errorf("Board.Move() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}
