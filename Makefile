.PHONY: version build build_linux docker_login docker_build docker_push_dev docker_push_pro
.PHONY: rm_stop

Version = v0.0.1
VersionFile = version/version.go
CommitVersion = $(shell git rev-parse --short=8 HEAD)
BuildVersion = $(shell date "+%F %T")
GOBIN = $(shell pwd)

_ProjectPath = github.com/kooksee/ktask
ImagesPrefix = registry.cn-hangzhou.aliyuncs.com/ybase/
ImageName = ktask
ImageLatestName = $(ImagesPrefix)$(ImageName)
ImageTestName = $(ImagesPrefix)$(ImageName):test
ImageVersionName = $(ImagesPrefix)$(ImageName):$(Version)

version:
	@echo "项目版本处理"
	@echo "package version" > $(VersionFile)
	@echo "const Version = "\"$(Version)\" >> $(VersionFile)
	@echo "const BuildVersion = "\"$(BuildVersion)\" >> $(VersionFile)
	@echo "const CommitVersion = "\"$(CommitVersion)\" >> $(VersionFile)

build:
	@echo "开始编译"
	GOBIN=$(GOBIN) go install main.go

build_linux: version
	@echo "交叉编译成linux应用"
	docker run --rm -v $(GOPATH):/go golang:latest go build -o /go/src/$(_ProjectPath)/main /go/src/$(_ProjectPath)/main.go

rm_stop:
	@echo "删除所有的的容器"
	sudo docker rm -f $(sudo docker ps -qa)
	sudo docker ps -a

rm_none:
	@echo "删除所为none的image"
	sudo docker images  | grep none | awk '{print $3}' | xargs docker rmi

docker_push_pro: docker_build
	@echo "docker push pro"
	sudo docker tag $(ImageLatestName) $(ImageVersionName)
	sudo docker push $(ImageVersionName)
	sudo docker push $(ImageVersionName)

docker_push_dev: docker_build
	@echo "docker push test"
	sudo docker tag $(ImageLatestName) $(ImageTestName)
	sudo docker push $(ImageTestName)

docker_build: build_linux
	@echo "构建docker镜像"
	sudo docker build -t $(ImageLatestName) .