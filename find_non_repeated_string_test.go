package homework

import "testing"

func TestFindNonRepeatedString(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"happy", args{file:"words.txt"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findNonRepeatedString(tt.args.file)
		})
	}
}