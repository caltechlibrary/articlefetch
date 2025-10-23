package articlefetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// FeedsURL takes a set of query terms and returns the url with query
// string 
//
//~~~
//  feedsUrl := feedsURL("authors.library.caltech.edu", "Grubbs, Robert")
//~~~
func FeedsURL(clpid string) string {
	return fmt.Sprintf("https://feeds.library.caltech.edu/people/%s/article.json", clpid)
}

// FeedsRdmIds retrieves the JSON object from Feeds
func FeedsRdmIds(u string) ([]string, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to retrieve %s, %s", u, res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	objs := []map[string]interface{}{}
	if err := json.Unmarshal(body, &objs); err != nil {
		return nil, err
	}
	if len(objs) == 0 {
		return nil, fmt.Errorf("not records returned")
	}
	rdmIds := []string{}
	for _, obj := range objs {
		if val, ok := obj["collection_id"].(string); ok {
			rdmIds = append(rdmIds, val)
		} else if val, ok := obj["id"].(string); ok {
			parts := strings.Split(val, ":")
			if len(parts) == 2 {
				rdmIds = append(rdmIds, parts[1])
			} else {
				rdmIds = append(rdmIds, val)				
			}
		}
	}
	return rdmIds, nil
}

func RdmRecordURL(hostname string, id string) string {
	return fmt.Sprintf("https://%s/api/records/%s/files", hostname, id)
}

func RdmFetchJSON(u string) ([]byte, time.Duration, error) {
	var targetTime time.Time
	res, err := http.Get(u)
	tReset := res.Header.Get("X-RateLimit-Reset")
	if tReset != "" {
		unixTime, _ := strconv.ParseInt(tReset, 10, 64)
		targetTime = time.Unix(unixTime, 0)
		fmt.Fprintf(os.Stderr, "Reset will happen at %s\n", targetTime.Format(time.RFC822Z))
	}
	if err != nil {
		return nil, time.Until(targetTime), err
	}
	duration := time.Until(targetTime)
	if res.StatusCode != 200 {
		return nil, duration, fmt.Errorf("failed to retrieve %s, %s", u, res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, duration, err
	}
	res.Body.Close()
	return body, duration, nil
}

func RdmPdfURLs(src []byte) ([]string, error) {
	//fmt.Printf("DEBUG obj retrieved -> %s\n", src)
	contentUrls := []string{}
	obj := map[string]interface{}{}
	if err := JSONUnmarshal(src, &obj); err != nil {
		return nil, fmt.Errorf("failed to unmarhsal object %s\n", err)
	}
	if entries, ok := obj["entries"].([]map[string]interface{}); ok {
		for _, entry := range entries {
			if mimetype, ok := entry["mimetype"].(string); ok && mimetype == "application/pdf" {
				if links, ok := entry["links"].(map[string]string); ok {
					if contentUrl, ok := links["content"]; ok {
						fmt.Printf("DEBUG contentUrl: %s\n", contentUrl)
						contentUrls = append(contentUrls, contentUrl)
					}
				}
			}
		}
	}
	return contentUrls, nil
}