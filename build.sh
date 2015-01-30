rm -rf target \
&& mkdir target \
&& docker build -t typhoon-builder . \
&& docker run --rm -v -e VERSION=$GIT_COMMIT `pwd`/target:/target -t $(docker images -q | head -1) \
&& ls -l target