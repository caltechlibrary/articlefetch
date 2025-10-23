Installation for development of **articlefetch**
===========================================

**articlefetch** A client that will submit an search to an RDM instance and retrieve the articles in the results.

Quick install with curl or irm
------------------------------

There is an experimental installer.sh script that can be run with the following command to install latest table release. This may work for macOS, Linux and if you’re using Windows with the Unix subsystem. This would be run from your shell (e.g. Terminal on macOS).

~~~shell
curl https://caltechlibrary.github.io/articlefetch/installer.sh | sh
~~~

This will install the programs included in articlefetch in your `$HOME/bin` directory.

If you are running Windows 10 or 11 use the Powershell command below.

~~~ps1
irm https://caltechlibrary.github.io/articlefetch/installer.ps1 | iex
~~~

### If your are running macOS or Windows

You may get security warnings if you are using macOS or Windows. See the notes for the specific operating system you're using to fix issues.

- [INSTALL_NOTES_macOS.md](INSTALL_NOTES_macOS.md)
- [INSTALL_NOTES_Windows.md](INSTALL_NOTES_Windows.md)

Installing from source
----------------------

### Required software


### Steps

1. git clone https://github.com/caltechlibrary/articlefetch
2. Change directory into the `articlefetch` directory
3. Make to build, test and install

~~~shell
git clone https://github.com/caltechlibrary/articlefetch
cd articlefetch
make
make test
make install
~~~

