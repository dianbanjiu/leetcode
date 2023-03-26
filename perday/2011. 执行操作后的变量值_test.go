package perday

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{operations: []string{"--X","X++","X++"}},
			want: 1,
		},
		{
			name: "2",
			args: args{operations: []string{"++X","++X","X++"}},
			want: 3,
		},
		{
			name: "3",
			args: args{operations: []string{"X++","++X","--X","X--"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FinalValueAfterOperations(tt.args.operations); got != tt.want {
				t.Errorf("FinalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
