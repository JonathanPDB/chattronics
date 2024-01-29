PROJECT=chattronics
model=gpt-4-turbo
temperature=0.2
iterations=1
testCase=themometry

.PHONY: build
build:
	@make clean
	@go build -o $(PROJECT) ./cmd/chattronics/main.go

run:
	clear
	@./$(PROJECT) -model=$(model) -temperature=$(temperature)


.PHONY: build-direct
build-direct:
	@make clean
	@go build -o $(PROJECT)-direct ./cmd/tests/direct/direct.go

run-direct:
	clear
	@./$(PROJECT)-direct -model=$(model) -temperature=$(temperature) -iterations=$(iterations) -testCase=$(testCase)

.PHONY: build-investigative
build-investigative:
	@make clean
	@go build -o $(PROJECT)-investigative ./cmd/tests/investigative/investigative.go

run-investigative:
	clear
	@./$(PROJECT)-investigative -model=$(model) -temperature=$(temperature) -iterations=$(iterations) -testCase=$(testCase)

unit-test:
	go test ./pkg/*** -cover

clean:
	@rm -f $(PROJECT)
	@rm -f $(PROJECT)-direct
	@rm -f $(PROJECT)-investigative
	@find ./runs -type d -name "*_TEST" -exec rm -rf {} +

SHELL=bash
delete-all-runs:
	@read -p "Are you sure you want delete all run folders? [y/n] " -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		rm -rf ./runs; \
	fi
	@echo "  Logs deleted."

