BUILD_VERSION   := v0.0.1
BUILD_TIME      := $(shell date "+%F %T")
BUILD_NAME      := app_$(shell date "+%Y%m%d%H" )
SOURCE          := ./main.go
TARGET_DIR      := ./bin/
COMMIT_SHA1     := $(shell git rev-parse HEAD )
#export GOOS=linux
all:
	go build -ldflags                           \
    "                                           \
    -X 'gserver/pkg/version.buildVersion=${BUILD_VERSION}'     \
    -X 'gserver/pkg/version.buildGitRevision=${COMMIT_SHA1}'     \
    "                                           \
    -o ${TARGET_DIR}/${BUILD_NAME}

linux:
	export GOOS=linux
	make all

clean:
	rm ${BUILD_NAME} -f

install:
	mkdir -p ${TARGET_DIR}
	cp ${BUILD_NAME} ${TARGET_DIR} -f

.PHONY :
	all clean install ${BUILD_NAME}