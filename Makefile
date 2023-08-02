BUILDPATH=$(CURDIR)
BINARY=book-cron

test: # run all test in project
	@echo "Executing tests..."
	@go test ./... -v

coverage: # creates a coverage file
	@echo "Making coverage evaluation ..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -func coverfile_out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' > coverage_tmp.txt
	@while read line; do \
		COVERAGE=$$(echo $${line%.*}) ; \
		echo $$COVERAGE > coverage.txt ; \
	done <coverage_tmp.txt

race: # Run data race detector
	@echo "Running race detector ..."
	go test ./... -race -short 

mod:
	@echo "Vendoring..."
	@go mod vendor

build: 
	@echo "Building binary file ..."
	@go build -mod vendor -ldflags "-s -w" -o $(BUILDPATH)/build/${BINARY} cmd/main.go
	@echo "Build file generated in folder build/"${BINARY}

build_report:
	@echo "Compilando con reporte..."
	@go build -gcflags="-m -l"
