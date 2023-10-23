PROJECT=chattronics
model=gpt-3.5-turbo
temperature=0.2

.PHONY: build
build:
	@go build -o $(PROJECT) ./main.go

run:
	@go build -o $(PROJECT) ./main.go
	clear
	@./$(PROJECT) -model=$(model) -temperature=$(temperature)
	@make clean

unit-test:
	go test ./pkg/*** -cover

clean:
	@rm -f $(PROJECT)
	@find ./runs -type d -empty -delete

SHELL=bash
delete-all-runs:
	@read -p "Are you sure you want delete all run folders? [y/n] " -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		rm -rf ./runs; \
	fi
	@echo "  Logs deleted."

delete-mock-runs:
	@find ./runs -type d -name "*_MOCK" -exec rm -rf {} +
	@echo "Mock run folders deleted."
