TARGET=gameboy-advance
NAME=game-of-life
SRCS=main.go
GO=tinygo

build:
	$(GO) build -target=$(TARGET) -o $(NAME).gba $(SRCS)

run:
	$(GO) run -target=$(TARGET) $(SRCS)

clean:
	$(GO) clean
	rm -f $(NAME).gba
