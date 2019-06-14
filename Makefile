GO = go
TARGET = yo

.PHONY: all
all: $(TARGET)

$(TARGET):
	$(GO) build

.PHONY:	test
test: yo
	./test.sh

.PHONY: clean
clean:
	rm -rf yo temp