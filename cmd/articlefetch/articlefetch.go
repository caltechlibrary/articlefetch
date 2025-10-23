package main


import (
	"flag"
	"fmt"
	"os"
	"path"

	// 3rd Party library
	"github.com/caltechlibrary/articlefetch"
)

func main() {
	appName := path.Base(os.Args[0])
	version, licenseText, releaseDate, releaseHash := articlefetch.Version, articlefetch.LicenseText, articlefetch.ReleaseDate, articlefetch.ReleaseHash
	showHelp, showVersion, showLicense, fmtHelp := false, false, false, articlefetch.FmtHelp
	helpText := articlefetch.HelpText
	flag.BoolVar(&showHelp, "help", showHelp, "display help")
	flag.BoolVar(&showVersion, "version", showVersion, "display version")
	flag.BoolVar(&showLicense, "license", showLicense, "display license")
	flag.Parse()
	args := flag.Args()

	// handle Help
	if showHelp {
		fmt.Printf("%s\n", fmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}

	// handle License
	if showLicense {
		fmt.Printf("%s\n", licenseText)
		os.Exit(0)
	}

	// handle Version
	if showVersion {
		fmt.Printf("%s %s\n", version, releaseHash)
		os.Exit(0)
	}

	// handle missing hostname and query
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "missing hostname or query")
		os.Exit(1)
	}
	hostname, query := args[0], args[1]
	os.Exit(articlefetch.Run(os.Stdin, os.Stdout, os.Stderr, appName, hostname, query))
}