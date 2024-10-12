BIN_NAME=go-restapi
CMD_DIR=cmd
VERSION=v1.0.2
BUILD_DIR=build
CONFIG_PATH=./config/config.yml

build:
	go build -C ${CMD_DIR}/${BIN_NAME}  -o ../../${BUILD_DIR}/${BIN_NAME}
	tar -czf ${BUILD_DIR}/${BIN_NAME}-${VERSION}.tar.gz ${BUILD_DIR}/${BIN_NAME}

run: build
	./${BUILD_DIR}/${BIN_NAME} --config ${CONFIG_PATH}

clean:
	rm -rf ${BUILD_DIR}
