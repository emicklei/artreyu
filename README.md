# Artreyu - artifact assembly tool

### Archive a build result

Given the artifact descriptor artreyu.yaml

	api: 1
		
	artifact: 	my-app
	version: 	1.0-SNAPSHOT
	group: 		com.company
	type: 		tgz
	
When running the "archive" command with a file location
	
	artreyu archive target/my-app.tgz	

Then the artifact is uploaded to the repo under

	<SOME_REPO>/com/company/my-app/1.0-SNAPSHOT/darwin/my-app-1.0-SNAPSHOT.tgz	

### Directory layout

	$group/$artifact/$version/$osname/$artifact-$version.$type


$osname can by `any` when the artifact is not operating system dependent (e.g texts,scripts,Java,...). Such artifacts will have the descriptor field `anyos` set to true. In the above example, osname is set to "darwin" when running from an OS X machine. It can be overriden using the command flag `--os`

### Assemble a new artifact

Given the artifact descriptor artreyu.yaml

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

When running the "assemble"

	artreyu assemble
	
Then the parts are downloaded to a temporary directory, the parts are extracted, all content is compressed again into a new artifact and the artifact is stored. You can set the temporary director using the command flag `--target`.

	target/
		my-app-2.1.tgz
		rest-service.bin
		rest-service.properties
		ui-app.html
		ui-app.js
	
### Sample configuration file .artreyu, stored in $HOME

	api: 1
	
	repositories:
	- name:		local
	  path:     /Users/you/artreyu	

	- name:		nexus
	  url:		https://yours.com/nexus
	  path:     /content/repositories
	  user: 	admin
	  password:	****  	  