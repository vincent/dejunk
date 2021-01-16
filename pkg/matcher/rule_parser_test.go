package matcher

import (
	"reflect"
	"strings"
	"testing"
)

func Test_replaceREPlaceholders(t *testing.T) {
	type args struct {
		terms        []string
		removeOthers bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Should return the input terms when no placeholders present",
			args: args{terms: []string{"some", "terms"}, removeOthers: false},
			want: []string{"some", "terms"},
		},
		{
			name: "Should return empty when no placeholders present",
			args: args{terms: []string{"some", "terms"}, removeOthers: true},
			want: []string{},
		},
		{
			name: "Should return replaced terms when some placeholders present",
			args: args{terms: []string{":audio", "terms"}, removeOthers: false},
			want: append([]string{"(" + strings.Join(AudioFileExts, "|") + ")"}, "terms"),
		},
		{
			name: "Should return only replaced terms when some placeholders present",
			args: args{terms: []string{":audio", "terms"}, removeOthers: true},
			want: []string{"(" + strings.Join(AudioFileExts, "|") + ")"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceREPlaceholders(tt.args.terms, tt.args.removeOthers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceREPlaceholders() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
