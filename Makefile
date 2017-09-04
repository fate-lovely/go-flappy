install:
	go install
.PHONY: install

run: install
	go-flappy
.PHONY: run
