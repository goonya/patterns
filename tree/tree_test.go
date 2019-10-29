package tree

import (
	"testing"
)

func TestTree_RemoveEmptyBranches(t *testing.T) {
	tests := []struct {
		name string
		got  Tree
		want Tree
	}{
		{
			"",
			Tree{
				{"1", "#", "o"},
			},
			Tree{},
		},
		{
			"",
			Tree{
				{"1", "#", "o"},
				{"11", "1", "o"},
				{"111", "11", "o"},
				{"1111", "111", "o"},
			},
			Tree{},
		},
		{
			"",
			Tree{
				{"1", "#", "o"},
				{"11", "1", "o"},
				{"111", "11", "o"},
				{"1111", "111", "d"},
			},
			Tree{
				{"1", "#", "o"},
				{"11", "1", "o"},
				{"111", "11", "o"},
				{"1111", "111", "d"},
			},
		},
		{
			"",
			Tree{
				{"1", "#", "o"},
				{"11", "1", "o"},
				{"12", "1", "o"},
				{"121", "12", "o"},
				{"1211", "121", "d"},
				{"13", "1", "o"},
				{"131", "13", "o"},
				{"1311", "131", "o"},
			},
			Tree{
				{"1", "#", "o"},
				{"12", "1", "o"},
				{"121", "12", "o"},
				{"1211", "121", "d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.RemoveEmptyBranches(); tt.got.String() != tt.want.String() {
				t.Errorf("\ngot: %s\nwant: %s", tt.got, tt.want)
			}
		})
	}
}
