package f1lib

import "testing"

func TestWhosBest(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test Whos best",
			args: args{name: "f1"},
			want: "f1 is the best!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhosBest(tt.args.name); got != tt.want {
				t.Errorf("WhosBest() = %v, want %v", got, tt.want)
			}
		})
	}
}
