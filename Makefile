ifeq ($(OS),Windows_NT)
    BINARY_NAME = tamagotchi.exe
else
    BINARY_NAME = tamagotchi
endif

SRC_DIR = ./cmd/console

.PHONY: all build run clean

all: build

build:
	go build -ldflags "-s -w" -o $(BINARY_NAME) $(SRC_DIR)/main.go

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
