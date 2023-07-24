PROJECT=chattronics

.PHONY: build-mock
build-mock:
	@go build -o $(PROJECT) ./cmd/$(PROJECT)/mock.go

.PHONY: build-real
build-real:
	@go build -o $(PROJECT) ./cmd/$(PROJECT)/chattronics.go

run-mock:
	@go build -o $(PROJECT) ./cmd/$(PROJECT)/mock.go
	@./$(PROJECT)
	@make clean

run-real:
	@go build -o $(PROJECT) ./cmd/$(PROJECT)/chattronics.go
	@./$(PROJECT) -model=$(model) -temperature=$(temperature)
	@make clean

clean:
	@rm -f $(PROJECT)
	@find ./runs -type d -empty -delete

SHELL=bash
delete-run-folder:
	@read -p "Are you sure you want delete all run folders? [y/n]\n" -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		rm -rf ./runs; \
	fi
	@echo " Logs deleted."
