export BIN_DIR=.
export BIN_NAME=ae86

build:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BIN_NAME)

run: build
	$(BIN_DIR)/$(BIN_NAME)