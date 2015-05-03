# Artreyu - artifact assembly tool

### Archive a build result

Given the artifact descriptor artreyu.yaml

	api: 1
	
	group: com.company
	artifact: my-app
	version: 1.0-SNAPSHOT
	extension: tgz
	
When running the "archive" command with a file location
	
	artreyu archive target/my-app.tgz	

Then the artifact is uploaded to the repo under

	<SOME_REPO>/com/company/my-app/1.0-SNAPSHOT/Darwin/my-app-1.0-SNAPSHOT.tgz	

### Directory layout

	$group/$artifact/$version/$osname/$artifact-$version.$type


$osname can by `any` when the artifact is not operating system dependent (e.g texts,scripts,Java,...). Such artifacts will have the property `anyos` set to true. In the above example, osname is set to "Darwin" because the artreyu configuration (see below) specifies that. It can be overriden using the command flag `-osname`

### Assemble a new artifact

Given the artifact descriptor artreyu.yaml

	api: 1
	
	group: 		com.company
	artifact: 	my-app
	version: 	2.1
	extension: 	tgz
	
	parts:
	- group: 	com.company
	  artifact:	rest-service
	  version: 	1.9
	  type: 	tgz
	- group: 	com.company
	  artifact: ui-app
	  version: 	2.1
	  type:		tgz
	  anyos:    true

When running the "assemble" command with a directory

	artreyu assemble target
	
Then the parts are downloaded to directory `target`, the parts are extracted and all content is compressed again into a new artifact.

	/target
		my-app-2.1.tgz
		rest-service.bin
		rest-service.properties
		ui-app.html
		ui-app.js
	
### Sample configuration file .artreyu, stored in $HOME

	repository: nexus
	url:		https://yours.com/nexus/content/repositories
	user: 		admin
	password:	admin
	osname:		Darwin