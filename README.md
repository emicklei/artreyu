# Typhoon - artifact assembly tool

### copy an artifact

```
typhoon archive target/app --version=$BUILD_ID --group=com.company target\app
```
in the repo
```
$TYPHOON_REPO/com/company/app/2015-01-29_09-04-55/app
```

### fetch latest artifact

```
typhoon fetch --group=com.company --artifact=app
```

### fetch specific artifact version

```
typhoon fetch --group=com.company --artifact=app --version=v3
```

### list artifacts by group, version or its name

```
typhoon list --group=com.company
```





### Nexus directory layout

	$groupId/$artifactId/$version/$os-arch/$artifactId-$version.$extension
	
	com.ubanita/firespark-web/1.0-SNAPSHOT/Linux/firespark-web-1.0-SNAPSHOT.tgz


The build process will create the `firespark-web.tgz`

	typhoon archive firespark-web.tgz

will upload firespark-web.tgz to com/ubanita/firespark-web/1.0-SNAPSHOT/Linux/firespark-web-1.0-SNAPSHOT.tgz using `typhoon.yaml`