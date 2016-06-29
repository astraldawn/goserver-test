#!/bin/sh
# To set up the environment
#
# gvm pkgset create env_name
# gvm pkgset use env_name
#
# Following that, run this script . ./export.sh (otherwise it runs in another shell)
#
# Then run gpm install (Godeps file required), followed by go install
# To run the server, just type server into any command window
#
PROJECTDIR="$HOME/projects/gotest"
export GOPATH="$GOPATH:$PROJECTDIR"
echo $GOPATH
export PATH="$PATH:$PROJECTDIR/bin"
echo $PATH
