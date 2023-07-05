CC := go
OPTION := build
BIN := example
SRC := main.go

compile:
	$(CC) $(OPTION) -o $(BIN) $(SRC)

clean: clean_binary 

clean_binary:
	-rm $(BIN)

dependencies:
	go get github.com/s9rA16Bf4/ArgumentParser@nightly
	go get github.com/s9rA16Bf4/notify_handler
	go get github.com/s9rA16Bf4/Go-tools/tools