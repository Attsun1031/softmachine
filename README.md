# jobnetes
kubernetes workflow engine

## Setup dev env
1. set GOROOT
1. set GOPATH
1. clone this repository in $GOPATH/src/github.com/Attsun1031/jobnetes
1. `go get -u github.com/golang/dep/cmd/dep`
1. cd to jobnetes dir
1. `dep ensure`
1. add config.yaml and kube-config to $HOME/.jobnetes
1. start local mysql container
  `docker run --name jobnetes-db -p3333:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -v ~/jobnetes-db:/var/lib/mysql -d mysql --character-set-server=utf8 --collation-server=utf8_unicode_ci`
1. execute cmd/dbmigration/dbmigration.go
  
## Setup local k8s env
1. start local kubernetes
1. apply `setting/k8s/deploy-mysql.yaml`
