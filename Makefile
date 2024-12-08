BIN_NAME=go-restapi
CMD_DIR=cmd
VERSION=v1.0.7
BUILD_DIR=build
CONFIG_PATH=./config/config.yml
TEST_DB=db
TEST_DB_HOST=localhost
TEST_DB_PORT=5432
TEST_DB_TIMEOUT=5
MAX_RETRIES=3
RETRY_COUNT=0

debug:
	cd ${CMD_DIR}/${BIN_NAME} && go run main.go

build:
	go build -C ${CMD_DIR}/${BIN_NAME}  -o ../../${BUILD_DIR}/${BIN_NAME}
	tar -czf ${BUILD_DIR}/${BIN_NAME}-${VERSION}.tar.gz ${BUILD_DIR}/${BIN_NAME}

run: build
	./${BUILD_DIR}/${BIN_NAME} --config ${CONFIG_PATH}

clean-test:
	@docker-compose down
	@go clean -testcache

test:
	docker-compose up -d db
	@while true; do \
		DB_NAME=$$(docker exec -i $$(docker-compose ps -q db) pg_isready -h localhost -p $(TEST_DB_PORT)); \
		if [[ $$DB_NAME == *"accepting connections"* ]]; then \
			break; \
		fi; \
		$(eval RETRY_COUNT=$(RETRY_COUNT+1)) \
		if [ $(RETRY_COUNT) -ge $(MAX_RETRIES) ]; then \
			echo "Maximum retries reached. Exiting..."; \
			exit 1; \
		fi; \
		echo "Postgres connection failed"; \
		sleep 2; \
	done
	go test ./...

clean:
	rm -rf ${BUILD_DIR}
