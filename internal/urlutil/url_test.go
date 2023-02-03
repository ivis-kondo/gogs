package urlutil_test

import (
	"testing"

	"github.com/NII-DG/gogs/internal/urlutil"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePath(t *testing.T) {
	tests := []struct {
		url      string
		newPath  string
		want     string
		hasError bool
	}{
		{
			url:      "https://sample.ac.jp/path1/path2?query1=value1&query1=value1#frag1",
			newPath:  "newpath1/newpath2/newpath3",
			want:     "https://sample.ac.jp/newpath1/newpath2/newpath3?query1=value1&query1=value1#frag1",
			hasError: false,
		},
		{
			url:      "https://sample.ac.jp/path1/path2?query1=value1&query1=value1#frag1",
			newPath:  "/newpath1/newpath2/newpath3",
			want:     "https://sample.ac.jp/newpath1/newpath2/newpath3?query1=value1&query1=value1#frag1",
			hasError: false,
		},
		{
			url:      "https://sample.ac.jp/path1/path2?query1=value1&query1=value1#frag1",
			newPath:  "/newpath1/newpath2/newpath3/",
			want:     "https://sample.ac.jp/newpath1/newpath2/newpath3?query1=value1&query1=value1#frag1",
			hasError: false,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result, err := urlutil.UpdatePath(test.url, test.newPath)
			if err != nil && test.hasError {

			} else if err != nil && !test.hasError {
				t.Errorf("Failure Test. url : %s, newPath : %s, want : %s, hasError : %v", test.url, test.newPath, test.want, test.hasError)
			} else if !test.hasError {
				assert.Equal(t, test.want, result)
			}

		})
	}
}
