package articlefetch

const (
	HelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] RDM_HOSTNAME CLPID

# DESCRIPTION

Use the CLPID provided to retreive a list of article from feeds, then use the
RDM_HOSTNAME to retrieve the PDFs for the articles found.

# OPTIONS

-help
: display this help page

-version
: display version info

-license
: display license

# EXAMPLE

~~~shell
{app_name} authors.library.caltech.edu Grubbs-R-H
~~~

`
)