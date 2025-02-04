from alpine:3.21

env GO_VERS=1.23.5
env GO_TARBALL=go${GO_VERS}.linux-amd64.tar.gz
run mkdir /wa/
run mkdir /wa/bin

workdir /wa/

run apk add wget bash shadow
run chsh -s /bin/bash root

workdir /wa/bin/
run wget https://go.dev/dl/$GO_TARBALL
run tar -xf ${GO_TARBALL}


workdir /wa/bin/
env TASK_VERS=3.41.0
run wget https://github.com/go-task/task/releases/download/v${TASK_VERS}/task_linux_386.tar.gz
run tar -xf task_linux_386.tar.gz

run echo "PATH=$PATH:/wa/bin/go/bin:/wa/bin/" >> ~/.bashrc

workdir /wa/gtfs_tracker
