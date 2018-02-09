
all: testapp

testapp:
	go build -v -i
dev: all
	spk dev
clean:
	rm -f testapp

.PHONY: testapp all dev clean
