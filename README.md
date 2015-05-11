# Artreyu - artifact assembly tool

Artreyu is a command line tool for build pipelines that need to create products that are composed of multiple versioned build parts.
An example of such a product is a Photo Editor that is composed of a platform specific binary, a folder with examples,
a folder containing a HTML documentation site, a folder with sample textures, another binary for image format conversions.
Each of the components or parts may have been build by a separate continuous build job, 
have their own version lifecycle and may be platform specific. 

Artreyu can be used to realize a continuous product build job that assembles artifacts which are stored in an artifact repository.
The tool uses descriptor files that contain meta data about the artifact (e.g. version,name,type). To support assembly, a descriptor can also list the part descriptors needed for creating the container artifact.

The design of this tool is inspired by the Apache Maven project which provides assembly support for Java projects. Compared to the Maven repository layout, Artreyu uses the OS name in the container path to support platform specific artifacts.

### Archive a build result

Given the artifact descriptor artreyu.yaml

	api: 1
		
	artifact: 	my-app
	version: 	1.0-SNAPSHOT
	group: 		com.company
	type: 		tgz
	
When running the "archive" command with a build result in target.
	
	artreyu archive target/my-app.tgz	

Then the artifact is uploaded to the repo under:

	<SOME_REPO>/com/company/my-app/1.0-SNAPSHOT/darwin/my-app-1.0-SNAPSHOT.tgz	

In the above example, $osname is set to "darwin" when running from an OS X machine.
It can be overriden using the command flag `--os`. 
$osname can by `any` when the artifact is not operating system dependent (e.g texts, scripts, Java). 
Such artifacts will have the descriptor field `anyos` set to true.

#### Directory layout

	$group/$artifact/$version/$osname/$artifact-$version.$type


### Assemble a new artifact

Given the artifact descriptor artreyu.yaml which references parts that are already archived.

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
	  anyos:    true

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
	
#### MIT License, ernestmicklei.com