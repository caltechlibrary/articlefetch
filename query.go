package articlefetch

import (
	"fmt"
)

// QueryURL takes a set of query terms and returns the url with query
// string 
func QueryURL(hostname string, query string, sort string) string {
	return fmt.Sprintf("https://%s/search?q=%s&l=list&p=1&s=10&sort=%s", hostname, query, sort)
}

