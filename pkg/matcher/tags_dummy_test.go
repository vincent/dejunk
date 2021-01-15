package matcher

import (
	"reflect"
	"testing"
)

func Test_fillYearTagFromName(t *testing.T) {
	type args struct {
		name string
		tags Tags
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Does not fill year tag if no 4 digits is present",
			args: args{name: "A string", tags: Tags{}},
			want: false,
		},
		{
			name: "Fill year tag if a 4 digits is present",
			args: args{name: "A string with 2020", tags: Tags{}},
			want: true,
		},
		{
			name: "Fill year tag if a 4 digits is present with some more space",
			args: args{name: "A string with 2020  ", tags: Tags{}},
			want: true,
		},
		{
			name: "Fill year tag if a 4 digits is present with some more chars",
			args: args{name: "A string with 2020azerty", tags: Tags{}},
			want: true,
		},
		{
			name: "Fill year tag if a 4+ digit is present",
			args: args{name: "A string with 202000", tags: Tags{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillYearTagFromName(tt.args.name, tt.args.tags); got != tt.want {
				t.Errorf("fillYearTagFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillSeasonEpisodeTagsFromName(t *testing.T) {
	type args struct {
		name string
		tags Tags
	}
	tests := []struct {
		name   string
		args   args
		ok     bool
		result string
	}{
		{
			name:   "Fill SeasonEpisode from a matching string",
			args:   args{name: "A string with s1e1", tags: Tags{}},
			result: "s1e1",
			ok:     true,
		},
		{
			name:   "Fill SeasonEpisode from a matching string",
			args:   args{name: "A string with S1E1", tags: Tags{}},
			result: "S1E1",
			ok:     true,
		},
		{
			name:   "Fill SeasonEpisode from a matching string",
			args:   args{name: "A string with S01E01", tags: Tags{}},
			result: "S01E01",
			ok:     true,
		},
		{
			name:   "Fill SeasonEpisode from a matching string",
			args:   args{name: "A string with 1x04", tags: Tags{}},
			result: "1x04",
			ok:     true,
		},
		{
			name:   "Returns false from a non matching string",
			args:   args{name: "A string", tags: Tags{}},
			result: "",
			ok:     false,
		},
		{
			name:   "Returns false from a non matching string",
			args:   args{name: "A string with 1xx04", tags: Tags{}},
			result: "",
			ok:     false,
		},
		{
			name:   "Returns false from a non matching string",
			args:   args{name: "A string with SE01", tags: Tags{}},
			result: "",
			ok:     false,
		},
		{
			name:   "Returns false from a non matching string",
			args:   args{name: "A string with S01EA", tags: Tags{}},
			result: "",
			ok:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := fillSeasonEpisodeTagsFromName(tt.args.name, tt.args.tags)
			if ok != tt.ok {
				t.Errorf("fillSeasonEpisodeTagsFromName() got = %v, want %v", ok, tt.ok)
			}
			if !reflect.DeepEqual(result, tt.result) {
				t.Errorf("fillSeasonEpisodeTagsFromName() got1 = %v, want %v", result, tt.result)
			}
		})
	}
}
