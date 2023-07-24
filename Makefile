PROJECT=chattronics

run-mock:
	@go build -o $(PROJECT) ./cmd/$(PROJECT)/mock.go
	@./$(PROJECT) --mock --model=GPT3Dot5Turbo
	@make clean

run-real:
	@go build -o $(PROJECT) ./cmd/$(PROJECT)/chattronics.go
	@./$(PROJECT) --mcok=false --model=GPT3Dot5Turbo --temperature=0.1
	@make clean

clean:
	@rm -f $(PROJECT)
	@find ./runs -type d -empty -delete

SHELL=bash
delete-run-folder:
	@read -p "Are you sure you want delete all run folders? [y/n]" -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		rm -rf ./runs; \
	fi
	@echo " Logs deleted."
