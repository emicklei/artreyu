rm -rf target \
&& mkdir target \
&& docker build --no-cache=true -t artreyu-builder . \
&& echo `pwd` \
&& docker run --rm -e VERSION=$GIT_COMMIT -v `pwd`/target:/target -t $(docker images -q | head -1) \
&& ls -l target