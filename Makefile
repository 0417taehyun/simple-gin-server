NAME=simple-gin-server

.PHONY: build
build:
	@go build -o $(NAME)

.PHONY: run
run: build
	@./$(NAME)
	
.PHONY: clean
clean:
	@rm -f $(NAME)

.PHONY: fmt
fmt:
	@go fmt

.PHONY: test
test:
	@go test -v ./tests/*
