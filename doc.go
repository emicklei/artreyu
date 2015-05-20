/*
A tool for handling versioned, platform dependent artifacts.
Its primary purpose is to create assembly artifacts from build artifacts archived in a repository.

Currently, it supports a local (filesystem) and Sonatype Nexus repository.

See https://github.com/emicklei/artreyu for more details.

(c)2015 ernestmicklei.com, MIT license

Usage:
  artreyu [flags]
  artreyu [command]
Available Commands:
  archive     upload an artifact to the repository
  fetch       download an artifact from the repository
  assemble    upload a new artifact by assembling fetched parts as specified in the descriptor
  format      process the artifact descriptor with the given template and write to stdout
  help        Help about any command

Flags:
  -c, --config="$HOME/.artreyu": location of the artreyu repositories configuration
  -d, --descriptor="artreyu.yaml": overwrite if the artifact descriptor has a different name or location
  -h, --help=false: help for artreyu
  -o, --os="$(uname)": overwrite if assembling for different OS
  -r, --repository="local": name of the repository as defined in the artreyu repositories configuration
  -v, --verbose=false: set to true to log more execution details
*/
package main
