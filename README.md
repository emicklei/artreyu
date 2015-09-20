# Artreyu - artifact assembly tool

Artreyu is a command line tool for build pipelines that need to create artifacts that are composed of multiple versioned build parts.
An example of such an artifact is an SDK that is composed of a platform specific binaries, a folder with examples,
a folder containing a HTML documentation site, a folder with samples. 
Each of the components or parts may have been build by a separate continuous build job, 
have their own version lifecycle and may be platform specific. 

Artreyu can be used to realize a continuous product build job that assembles artifacts which are stored in an artifact repository.
The tool uses descriptor files that contain meta data about the artifact (e.g. version,name,type). To support assembly, a descriptor will have to list the part descriptors needed for creating the container artifact.

The design of this tool is inspired by the Apache Maven project which provides assembly support for Java projects. Compared to the Maven repository layout, Artreyu uses the OS name in the container path to support platform specific artifacts.

### Archive a build result

Given the artifact descriptor artreyu.yaml

	api: 1
		
	artifact: 	my-app
	version: 	1.0-SNAPSHOT
	group: 		com.company
	type: 		tgz
	
When running the `archive` command with a build result located in target.
	
	artreyu archive target/my-app.tgz	

Then the artifact is uploaded to the repo under:

	<SOME_REPO>/com/company/my-app/1.0-SNAPSHOT/darwin/my-app-1.0-SNAPSHOT.tgz	

#### Directory layout

	$group/$artifact/$version/$osname/$artifact-$version.$type

In the above example, $osname is set to `darwin` when running from an OSX machine. It can be overriden using the command flag `--os`. 
$osname can by `any-os` when the artifact is not operating system dependent (e.g texts, scripts, Java). 
Such artifacts will have the descriptor field `any-os` set to true.

Unless run with the flag `--repository` or `-r`, the artifacts are stored only on the local filesystem at a location specified in the artreyu configuration (see below). 

### Plugin commands

Artreyu uses a simple plugin architecture to support other repository types. For example, the `artreyu-nexus` program is called to store and fetch artifacts from a Sonatype Nexus repository.  See [artreyu-nexus](https://github.com/emicklei/artreyu-nexus). To store an artifact in Nexus, you run:

	artreyu archive -r nexus target/my-app.tgz


### Assemble a new artifact

Given the artifact descriptor `artreyu.yaml` which references parts that are already archived.

	api: 1
		
	artifact: 	my-app
	version: 	2.1
	group: 		com.company
	type: 		tgz
	
	parts:
	- artifact:	rest-service
	  version: 	1.9
	  group: 	com.company
	  type: 	tgz

	- artifact: ui-app
	  version: 	2.1
	  group: 	com.company
	  type:		tgz
	  any-os:   true

When running the "assemble" command.

	artreyu assemble
	
Then the parts are downloaded to a temporary directory, the parts are extracted,
all content is compressed again into a new artifact and then the new artifact is stored. 
You can override the temporary directory explicitly by appending its name to the command line, e.g. `artreyu assemble target`

	target/
		my-app-2.1.tgz
		rest-service.bin
		rest-service.properties
		ui-app.html
		ui-app.js
	
### Print descriptor details

The format subcommand can be used to print information about the descriptor (artifact or assembly). The format command requires a template using the Go syntax. See [http://golang.org/pkg/text/template/](http://golang.org/pkg/text/template/)

	artreyu format "{{.Name}}-{{.Version}}.{{.Type}}"
	
	doc-2.0.tgz
	
### Tree of components	
	
The tree subcommand will recursively retrieve the artreyu descriptors to construct and print a composition hierarchy.

	artreyu tree
	
	root-commponent-2.0
	|-- child-1.0.tgz
	|-- another-2.9.tgz
	|  |-- sub.1.0.tgz
	
### Local caching artifacts

Versioned artifacts are cached using the local repository (filesystem).
Archiving a version of an artifact will put it in the local repository
after storing it on a remote (using a plugin).
Fetching the version of an artifact will first try to get it from the local repository.
If that fails then the remote repository is used. If that succeeds, a copy of the artifact is put 
in the local repository.
If the target repository is set to `local` then both versions and snapshots are store locally.
An artifact is called a snapshot if the Version property has the substring "SNAPSHOT".
	
### Sample configuration file .artreyu
Default location for this configuration file is $HOME. You can override the location using `--config`. 

	api: 1
	
	repositories:
	- name:		local
	  path:     /Users/you/artreyu	

	- name:		nexus
	  url:		https://yours.com/nexus
	  path:     /content/repositories
	  user: 	admin
	  password:	****  
	
### Installation from source

	VERSION=latest make local	
	
(c)2015, MIT License, [http://ernestmicklei.com](http://ernestmicklei.com)