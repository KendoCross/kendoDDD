#!/bin/bash

set -e

if [ $# -lt 2 ];then
    echo "use: build_docker.sh versionNum dev[test][prod] "
    echo "example: build_docker.sh v0.0.1 dev"
    exit 1
fi

./build_app.sh

#commitid="$1.$2."`date +%Y%m%d%H%M%S`
commitid="$1_$2"

echo "commitid = ${commitid}"

app="auth_srv"

docker build -t ${app}:${commitid} .

# docker tag ${app}:${commitid} harbor.nilinside.com/stepsflow/${app}:${commitid}
# docker rmi ${app}:${commitid}
# docker push harbor.nilinside.com/stepsflow/${app}:${commitid}

# docker run --name auth_srv -p 19195:18185 -p 19196:18186 467
