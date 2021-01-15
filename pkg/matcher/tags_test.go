package matcher

import "testing"

func TestTagsToPath(t *testing.T) {
	type args struct {
		pattern string
		tags    Tags
	}
	tests := []struct {
		name string
		args args
		want string
		ok   bool
	}{
		{
			name: "Without tags it should return the input",
			args: args{pattern: "/a/path/witout/tags", tags: Tags{}},
			want: "/a/path/witout/tags", ok: true,
		},
		{
			name: "With some tags it should return false",
			args: args{pattern: "/a/:path/with/:some/tags", tags: Tags{"path": "A"}},
			want: "/a/A/with//tags", ok: false,
		},
		{
			name: "With all tags it should return true",
			args: args{pattern: "/a/:path/with/:some/tags", tags: Tags{"path": "A", "some": "B"}},
			want: "/a/A/with/B/tags", ok: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := TagsToPath(tt.args.pattern, tt.args.tags)
			if got != tt.want {
				t.Errorf("TagsToPath() got result = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("TagsToPath() got ok = %v, want %v", ok, tt.ok)
			}
		})
	}
}
