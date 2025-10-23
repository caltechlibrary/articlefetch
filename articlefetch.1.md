%articlefetch(1) user manual | version 0.0.0 0fd37b8
% R. S. Doiel
% 2025-10-23

# NAME

articlefetch

# SYNOPSIS

articlefetch [OPTIONS] HOSTNAME QUERY_STRING

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
articlefetch authors.library.caltech.edu "Grubbs, Robert"
~~~


