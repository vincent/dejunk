package matcher

import (
	"testing"
)

func Test_testFromExpression(t *testing.T) {
	type args struct {
		p    *pattern
		item *ScrapItem
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should match an item that looks like an episode, using is()",
			args: args{
				p:    &pattern{Function: "is", Params: []string{":episode"}},
				item: &ScrapItem{SourcePath: "an episode s1e01.avi"},
			},
			want: true,
		},
		{
			name: "Should match an item that looks like a big number episode, using is()",
			args: args{
				p:    &pattern{Function: "is", Params: []string{":episode"}},
				item: &ScrapItem{SourcePath: "an episode s1e018765.avi"},
			},
			want: true,
		},
		{
			name: "Should not match an item that does not looks like an episode, using is()",
			args: args{
				p:    &pattern{Function: "is", Params: []string{":episode"}},
				item: &ScrapItem{SourcePath: "an episode 01.avi"},
			},
			want: false,
		},
		{
			name: "Should not match an item that looks like an episode, using not()",
			args: args{
				p:    &pattern{Function: "not", Params: []string{":episode"}},
				item: &ScrapItem{SourcePath: "an episode s1e01.avi"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testFromExpression(tt.args.p)(tt.args.item); got != tt.want {
				t.Errorf("testFromExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
