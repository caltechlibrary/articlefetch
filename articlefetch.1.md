%articlefetch(1) user manual | version 0.0.1 7afe26c
% R. S. Doiel
% 2025-10-24

# NAME

articlefetch

# SYNOPSIS

articlefetch [OPTIONS] RDM_HOSTNAME CLPID

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
articlefetch authors.library.caltech.edu Grubbs-R-H
~~~


