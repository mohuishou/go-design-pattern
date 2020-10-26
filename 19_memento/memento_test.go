package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	in := &InputText{}
	snapshots := []*Snapshot{}

	tests := []struct {
		input string
		want  string
	}{
		{
			input: ":list",
			want:  "",
		},
		{
			input: "hello",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
		{
			input: "world",
			want:  "",
		},
		{
			input: ":list",
			want:  "helloworld",
		},
		{
			input: ":undo",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			switch tt.input {
			case ":list":
				assert.Equal(t, tt.want, in.GetText())
			case ":undo":
				in.Restore(snapshots[len(snapshots)-1])
				snapshots = snapshots[:len(snapshots)-1]
			default:
				snapshots = append(snapshots, in.Snapshot())
				in.Append(tt.input)
			}
		})
	}
}
