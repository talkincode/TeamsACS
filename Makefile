BUILD_VERSION   := latest
BUILD_TIME      := $(shell date "+%F %T")
BUILD_NAME      := teamsacs
RELEASE_VERSION := v1.0.1
SOURCE          := main.go
RELEASE_DIR     := ./release
COMMIT_SHA1     := $(shell git show -s --format=%H )
COMMIT_DATE     := $(shell git show -s --format=%cD )
COMMIT_USER     := $(shell git show -s --format=%ce )
COMMIT_SUBJECT     := $(shell git show -s --format=%s )

clean:
	rm -f teamsacs

gen:
	go generate

build:
	go generate
	CGO_ENABLED=0 go build -a -ldflags \
	'\
	-X "main.BuildVersion=${BUILD_VERSION}"\
	-X "main.ReleaseVersion=${RELEASE_VERSION}"\
	-X "main.BuildTime=${BUILD_TIME}"\
	-X "main.BuildName=${BUILD_NAME}"\
	-X "main.CommitID=${COMMIT_SHA1}"\
	-X "main.CommitDate=${COMMIT_DATE}"\
	-X "main.CommitUser=${COMMIT_USER}"\
	-X "main.CommitSubject=${COMMIT_SUBJECT}"\
	-s -w -extldflags "-static"\
	' \
    -o ${BUILD_NAME} ${SOURCE}

build-linux:
	go generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags \
	'\
	-X "main.BuildVersion=${BUILD_VERSION}"\
	-X "main.ReleaseVersion=${RELEASE_VERSION}"\
	-X "main.BuildTime=${BUILD_TIME}"\
	-X "main.BuildName=${BUILD_NAME}"\
	-X "main.CommitID=${COMMIT_SHA1}"\
	-X "main.CommitDate=${COMMIT_DATE}"\
	-X "main.CommitUser=${COMMIT_USER}"\
	-X "main.CommitSubject=${COMMIT_SUBJECT}"\
	-s -w -extldflags "-static"\
	' \
    -o ${RELEASE_DIR}/${BUILD_NAME} ${SOURCE}

fastpub:
	make build-linux
	make upx
	echo 'FROM alpine' > .build
	echo 'ARG CACHEBUST="$(shell date "+%F %T")"' >> .build
	echo 'COPY ./release/teamsacs /teamsacs' >> .build
	echo 'RUN chmod +x /teamsacs' >> .build
	echo 'EXPOSE 20991 20992 20993 1812/udp 1813/udp 20914/udp 20924/udp 20914/udp' >> .build
	echo 'ENTRYPOINT ["/teamsacs"]' >> .build
	docker build -t teamsacs . -f .build
	rm -f .build
	docker tag teamsacs docker.pkg.github.com/ca17/teamsacs/teamsacs:latest
	docker push docker.pkg.github.com/ca17/teamsacs/teamsacs:latest
	docker tag teamsacs alab.189csp.cn:5000/teamsacs:latest
	docker push alab.189csp.cn:5000/teamsacs:latest

upx:
	upx ${RELEASE_DIR}/${BUILD_NAME}

ci:
	@read -p "type commit message: " cimsg; \
	git ci -am "$(shell date "+%F %T") $${cimsg}"

push:
	@read -p "type commit message: " cimsg; \
	git ci -am "$(shell date "+%F %T") $${cimsg}"
	git push origin main

.PHONY: clean build rpccert webcert


