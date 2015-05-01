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
	$SOME_REPO/com/company/my-app/1.0-SNAPSHOT/Darwin/my-app-1.0-SNAPSHOT.tgz
	```

### Directory layout

	$groupId/$artifactId/$version/$os-name/$artifactId-$version.$extension


$os-name can by `any` when the artifact is not operating system dependent (e.g texts,scripts,Java,...)

### Assemble a new artifact

Given the artifact descriptor typhoon.yaml

	```
	typhoon-api: 1
	
	group: com.company
	artifact: company-linux-sdk
	version: 2.1
	extension: tgz
	
	parts:
	- group: com.company
	  artifact: rest-service
	  version: 1.9
	  extension: tgz
	- group: com.company
	  artifact: ui-app
	  version: 2.1
	  extension: tgz
	```

When running the "fetch" command with a directory

	```
	typhoon fetch target
	```
	
Then the artifacts are unpacked in directory `target`

	```
	/target
		rest-service.exe
		rest-service.properties
		ui-app.js
		ui-app.html
	```