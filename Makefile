TARGET := goawk-srcget

help:
	@echo "Work in Process"

build: fmt
	@docker build -t $(TARGET) .

fmt:
	@cd src && go fmt

run:
	@docker run -it --rm $(TARGET)

shell:
	@docker run -it --rm $(TARGET) sh

brun: build run
