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


.PHONY: build-limited
build-limited:
	@make clean
	@go build -o $(PROJECT)-limited ./cmd/tests/limited/limited.go

run-limited:
	clear
	@./$(PROJECT)-limited -model=$(model) -temperature=$(temperature) -iterations=$(iterations) -testCase=$(testCase)

.PHONY: build-informative
build-informative:
	@make clean
	@go build -o $(PROJECT)-informative ./cmd/tests/informative/informative.go

run-informative:
	clear
	@./$(PROJECT)-informative -model=$(model) -temperature=$(temperature) -iterations=$(iterations) -testCase=$(testCase)

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

