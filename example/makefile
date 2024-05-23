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
	go get github.com/DennisTheodoreNedry/ArgumentParser
	go get github.com/DennisTheodoreNedry/notify_handler
	go get github.com/DennisTheodoreNedry/Go-tools