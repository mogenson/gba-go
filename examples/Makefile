GO := tinygo
NAME := hello-world
SRCS := main.go
TARGET := gameboy-advance

build:
	$(GO) build -target=$(TARGET) -o $(NAME).gba $(SRCS)

run:
	$(GO) run -target=$(TARGET) $(SRCS)

debug:
	$(GO) gdb -target=$(TARGET) -opt=1 $(SRCS)

clean:
	$(GO) clean
	rm -f $(NAME).gba
