package articlefetch

const (
	HelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] HOSTNAME QUERY_STRING

# DESCRIPTION

Take the HOSTNAME and QUERY_STRING values, retrieve the results from RDM and then using the results
retrieve the PDFs.

# OPTIONS

-help
: display this help page

-version
: display version info

-license
: display license

# EXAMPLE

~~~shell
{app_name} authors.library.caltech.edu "Grubbs, Robert"
~~~

`
)