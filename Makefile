.PHONY: run
run:
	@echo "create windows app"
	GOOS=windows go build -o file_manipulation.exe main.go