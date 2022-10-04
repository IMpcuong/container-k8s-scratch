# ARTIFACT_DIR=bootstrap
ZIP_FILE = bootstrap

clean:
	@rm -rf dist
	@mkdir -p dist

# zip -q -Z bzip2 ./go-lambda.zip ./go-lambda
# cp ./go-lambda.zip $(ARTIFACT_DIR)
build: clean
	GOOS=linux go build -v -o dist/handler/go-lambda ./...
	zip -jq ./$(ZIP_FILE) dist/handler/go-lambda
	cp ./$(ZIP_FILE) $(ARTIFACT_DIR)
	chmod +x $(ARTIFACT_DIR)/$(ZIP_FILE)