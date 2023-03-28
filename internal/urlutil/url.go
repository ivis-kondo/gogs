package urlutil

import (
	"fmt"
	"net/url"
)

func UpdatePath(rawURL string, newPath string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failure parsing url. <%s>", rawURL)
	}
	u.Path = newPath
	return u.String(), nil
}
