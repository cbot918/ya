INSTALL_BIN_PATH=$(HOME)/go/bin
BIN_NAME=ya
install: bin
	rm -f $(INSTALL_BIN_PATH)/$(BIN_NAME)
	sudo cp bin/$(BIN_NAME) $(INSTALL_BIN_PATH)
	echo "command: ya"

bin:
	mkdir -p bin
	go build -o bin/$(BIN_NAME) cmd/main.go

clean:
	rm *.test

test-install:
	go run test/test_install/main.go

run:
	go run .

.PHONY: clean run install bin
.SILENT: clean run install bin