### copy an artifact

```
typhoon archive --artifact=target/torch --version=$BUILD_ID --group=com.ubanita
```
in the repo
```
$TYPHOON_REPO/com/ubanita/torch/2015-01-29_09-04-55/torch
```

### fetch latest artifact

```
typhoon fetch --group=com.ubanita --artifact=torch
```

### fetch specific artifact version

```
typhoon fetch --group=com.ubanita --artifact=torch --version=v3
```