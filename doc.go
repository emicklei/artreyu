/*
artreyu - artifact assembly tool

	Usage:
	  artreyu [flags]
	  artreyu [command]
	Available Commands:
	  archive     upload an artifact to the repository
	  fetch       download an artifact from the repository to [destination]
	  assemble    create a new artifact [destination] by assembling parts from the descriptor
	  help        Help about any command

	Flags:
	  -h, --help=false: help for artreyu
	      --os="": overwrite if assembling for different OS
*/
package main
