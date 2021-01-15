package matcher

import (
	"testing"
)

func Test_loadFromYamlFile(t *testing.T) {
	type args struct {
		contents []byte
	}
	tests := []struct {
		name      string
		args      args
		wantRules int
		wantErr   bool
	}{
		{
			name:      "Should return no rules from an empty file",
			args:      args{contents: []byte("")},
			wantRules: 0,
			wantErr:   true,
		},
		{
			name:      "Should error from an invalid file",
			args:      args{contents: []byte("/some/garbage")},
			wantRules: 0,
			wantErr:   true,
		},
		{
			name: "Should not error if some valid rules found",
			args: args{contents: []byte(`
- name: Movies
  match: "ext(:video)not(:episode)"
  type: Movie
  store: ":title (:year)/:title"
  with: [dummy]

- name: Movies
  type: Movie
`)},
			wantRules: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadFromYamlFile(tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadFromYamlFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantRules {
				t.Errorf("loadFromYamlFile() count = %v, want %v", len(got), tt.wantRules)
			}
		})
	}
}
