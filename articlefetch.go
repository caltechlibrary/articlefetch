package articlefetch

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func Run(in io.Reader, out io.Writer, eout io.Writer, appName string, hostname string, clpid string) int {
	// NOTE: URL encode our query string
	feedsUrl := FeedsURL(clpid)
	rdmIds, err := FeedsRdmIds(feedsUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}
	tot := len(rdmIds)
	retrieved := 0
	fmt.Printf("retrieving %d records\n", tot)
	pdfToRetrieve := []string{}
	for i, id := range rdmIds {
		rdmUrl := RdmRecordURL(hostname, id)
		src, duration, err := RdmFetchJSON(rdmUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}
		time.Sleep(duration)

		pdfUrls, err := RdmPdfURLs(src)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to find a pdfUrls %q, %s\n", rdmUrl, err)
			continue
		}
		if len(pdfUrls) > 0 {
			fmt.Printf("DEBUG pdfUrls (%d)\n\t%+v\n", len(pdfUrls), strings.Join(pdfUrls, "\n\t"))
			pdfToRetrieve = append(pdfToRetrieve, pdfUrls...)
			time.Sleep(10 * time.Second)
		}
		retrieved += 1
		if (i % 5) == 0{
			fmt.Printf("%d/%d processed\n", i+1, tot)
		}
	}
	fmt.Printf("%d/%d retrieved\n", retrieved, tot)
	fmt.Printf("Retrieve the following URL:\n\n%s\n\n", strings.Join(pdfToRetrieve, "\n\t"))
	return 0
}