package articlefetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
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

func RdmFetchJSON(u string) ([]byte, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return body, nil
}

func RdmPdfURLs(src []byte) ([]string, error) {
	contentUrls := []string{}
	obj := map[string]interface{}{}
	if err := JSONUnmarshal(src, &obj); err != nil {
		return nil, fmt.Errorf("failed to unmarhsal object %s\n", err)
	}
	// DEBUG
	//src, _ = JSONMarshalIndent(obj, "", "    ") // DEBUG
	//fmt.Printf("DEBUG obj retrieved -> %s\n", src) // DEBUG
	if entries, ok := obj["entries"].([]interface{}); ok {
		//fmt.Printf("DEBUG entries (%T): %+v\n", entries)
		for _, val := range entries {
			entry := val.(map[string]interface{})
			//src, _ = JSONMarshalIndent(entry, "", "    ") // DEBUG
			//fmt.Printf("DEBUG obj entries -> %s\n", src) // DEBUG
			if mimetype, ok := entry["mimetype"].(string); ok && mimetype == "application/pdf" {
				//fmt.Printf("DEBUG entry.mimetype -> %s\n", mimetype) // DEBUG
				if links, ok := entry["links"].(map[string]interface{}); ok {
					//src, _ = JSONMarshalIndent(links, "", "    ") // DEBUG
					//fmt.Printf("DEBUG links %s\n", src) // DEBUG
					if contentUrl, ok := links["content"].(string); ok {
						//fmt.Printf("DEBUG contentUrl: %s\n", contentUrl)
						contentUrls = append(contentUrls, contentUrl)
					}
				}
			}
		}
	}
	return contentUrls, nil
}

func RdmGetFilenameFromContentURL(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return path.Base(strings.TrimSuffix(u.Path, "/content")), nil
}

func RdmRetrieveFile(u string) ([]byte, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return body, nil
}