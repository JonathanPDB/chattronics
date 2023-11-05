PROJECT=chattronics
model=gpt-3.5-turbo
temperature=0.2
iterations=1
testCase=themometry

.PHONY: build
build:
	@go build -o $(PROJECT) ./cmd/chattronics/main.go

run:
	clear
	@./$(PROJECT) -model=$(model) -temperature=$(temperature)
	@make clean


.PHONY: build-limited
build-limited:
	@go build -o $(PROJECT)-limited ./cmd/tests/limited/limited.go

run-limited:
	clear
	@./$(PROJECT)-limited -model=$(model) -temperature=$(temperature) -iterations=$(iterations) -testCase=$(testCase)
	@make clean

.PHONY: build-informative
build-informative:
	@go build -o $(PROJECT)-informative ./cmd/tests/informative/informative.go

run-informative:
	clear
	@./$(PROJECT)-informative -model=$(model) -temperature=$(temperature) -iterations=$(iterations) -testCase=$(testCase)
	@make clean

unit-test:
	go test ./pkg/*** -cover

clean:
	@rm -f $(PROJECT)
	@rm -f $(PROJECT)-limited
	@rm -f $(PROJECT)-informative
	@find ./runs -type d -empty -delete
	@find ./runs -type d -name "*_TEST" -exec rm -rf {} +

SHELL=bash
delete-all-runs:
	@read -p "Are you sure you want delete all run folders? [y/n] " -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		rm -rf ./runs; \
	fi
	@echo "  Logs deleted."

