/*
A tool for handling versioned, platform dependent artifacts.
Its primary purpose it to create assembly artifacts from build artifacts archived in a (remote) repository.

Currently, it supports a local (filesystem) and Sonatype Nexus repository.

Usage:
  artreyu [flags]
  artreyu [command]
Available Commands:
  archive     upload an artifact to the repository
  fetch       download an artifact from the repository
  assemble    upload a new artifact by assembling fetched parts as specified in the descriptor
  help        Help about any command

Flags:
  -c, --config="/Users/emicklei/.artreyu": location of the artreyu repositories configuration
  -d, --descriptor="artreyu.yaml": overwrite if the artifact descriptor has a different name or location
  -h, --help=false: help for artreyu
  -o, --os="darwin": overwrite if assembling for different OS
  -v, --verbose=false: set to true for more execution details
*/
package main
