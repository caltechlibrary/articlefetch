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

Use the CLPID is used to retrieve a list of articles from <feeds.library.caltech.edu>. The
resulting list of articles is used to retrieve the metadata for the articles from the RDM
repository indicated by RDM_HOSTNAME. For each article the metadata files object is retrieved
and any PDFs listed are then harvested and stored in a directory structure forms from the
CLPID, the RDM record id and the PDF name. 

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

This results in a directory tree under Grubbs-R-H of RDM record ids that hold
PDFs associated with the record.

`
)