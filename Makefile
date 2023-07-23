PROJECT=chattronics

.PHONY: build
build:
	go build -o chattronics .

clean:
	rm chattronics
	find ./logs -type d -empty -delete

run-mock: build
	./chattronics --mock --model=GPT3Dot5Turbo
	make clean

run-real: build
	./chattronics --mcok=false --model=GPT3Dot5Turbo --temperature=0.1
	make clean

SHELL=bash
delete-logs:
	@read -p "Are you sure you want delete all logs? [y/n]" -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		rm -rf ./logs; \
	fi
	@echo " Logs deleted."
