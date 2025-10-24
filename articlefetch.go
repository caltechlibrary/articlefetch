package articlefetch

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	//"strings"
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
	t0 := time.Now()
	iTime := time.Now()
	reportProgress := false
	retrieved := 0
	waitTime := 5 * time.Second
	fmt.Printf("processing %d records\n", tot)
	for i, id := range rdmIds {
		if i > 0 {
			if iTime, reportProgress = CheckWaitInterval(iTime, waitTime); reportProgress {
				fmt.Printf("%s | next id %q\n", ProgressETA(t0, i, tot), id)
				time.Sleep(waitTime)
			}
		}
		rdmUrl := RdmRecordURL(hostname, id)
		src, err := RdmFetchJSON(rdmUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}

		pdfUrls, err := RdmPdfURLs(src)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to find a pdfUrls %q, %s\n", rdmUrl, err)
			continue
		}
		if len(pdfUrls) > 0 {
			// Make a directory for {clpid}/{rdmid}
			saveDir := filepath.Join(clpid, id)
			if _, err := os.Stat(saveDir); err != nil {
				os.MkdirAll(saveDir, 0775)
			}
			// For each PDF create a directory for the RDM record id
			for i, pdfUrl := range pdfUrls {
				fName, err := RdmGetFilenameFromContentURL(pdfUrl)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to extract filename (file no. %d for %s) %s, %s\n", i, id, pdfUrl, err)
				}
				// Retrieve and write out the PDF to dir
				if src, err := RdmRetrieveFile(pdfUrl); err != nil {
					fmt.Fprintf(os.Stderr, "failed to retrieve file %s, %s\n", pdfUrl, err)
				} else {
					fName = filepath.Join(saveDir, fName)
					if err := os.WriteFile(fName, src, 0664); err != nil {
						fmt.Fprintf(os.Stderr, "failed to write %s, %s\n", fName, err)
					}
				}
			}
		}
		retrieved += 1
	}
	fmt.Printf("%d of %d records processed\n", retrieved, tot)
	return 0
}