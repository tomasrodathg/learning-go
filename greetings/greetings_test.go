package greetings

import (
	"errors"
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *regexp.Regexp
		wantErr error
	}{
		{"testing valid name", "Tomas", regexp.MustCompile(`\b+Tomas+\b`), nil},
		{"testing empty name", "", regexp.MustCompile(``), errors.New("empty name")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Hello(tt.input)

			if (err != nil) != (tt.wantErr != nil) || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("got error %v want error %v", err.Error(), tt.wantErr.Error())
				return
			}
			if !tt.want.MatchString(res) {
				t.Errorf("got %s want %s", res, tt.want)
			}
		})
	}
}
