# Typhoon - artifact assembly tool

### Archive a build result

Given the artifact descriptor typhoon.yaml
	```
	typhoon-api: 1
	
	group: com.company
	artifact: my-app
	version: 1.0-SNAPSHOT
	extension: tgz
	```

When running the "archive" command with a file location
	```
	typhoon archive target/my-app.tgz
	```

Then the artifact is stored (uploaded) in the repo under
	```
	<SOME_REPO>/com/company/my-app/1.0-SNAPSHOT/Darwin/my-app-1.0-SNAPSHOT.tgz
	```

### Directory layout

	$group/$artifact/$version/$osname/$artifact-$version.$type


$osname can by `any` when the artifact is not operating system dependent (e.g texts,scripts,Java,...)

### Assemble a new artifact

Given the artifact descriptor typhoon.yaml

	```
	typhoon-api: 1
	
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
	```

When running the "assemble" command with a directory

	```
	typhoon assemble target
	```
	
Then the parts are downloaded to directory `target`

	```
	/target
		rest-service-1.9.tgz
		ui-app-2.1.tgz
	```
	
### Sample .typhoon

	repository: nexus
	url:		https://yours.com/nexus/content/repositories
	user: 		admin
	password:	admin
	osname:		Darwin