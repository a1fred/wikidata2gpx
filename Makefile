BRANCH=$(shell git branch --show-current)
GITREV=$(shell git describe --abbrev=7 --always --tags)
REV=$(GITREV)-$(BRANCH)-$(shell date +%Y%m%d-%H:%M:%S)

BUILD_DIR=build
BIN=wikidata2gpx
TARGET=$(BUILD_DIR)/$(BIN)

all: $(TARGET)

.PHONY: test
test:
	go test ./wikidata2gpx/...

.PHONY: $(TARGET)
$(TARGET): test
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build -tags=NOCOMPLETION -ldflags "-X main.revision=$(REV) -s -w" -o $(TARGET) ./
	strip $(TARGET)

	mkdir -p completions
	go run main.go completion bash > completions/completion.bash
	go run main.go completion fish > completions/completion.fish
	go run main.go completion powershell > completions/completion.powershell
	go run main.go completion zsh > completions/completion.zsh

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

latest-all.json.bz2:
	wget -c https://dumps.wikimedia.org/wikidatawiki/entities/latest-all.json.bz2

gpxallfiles: latest-all.json.bz2 $(TARGET)
	mkdir -p gpxallfiles
	$(TARGET) wikidata ./latest-all.json.bz2 --gpx-metadata-name="Wiki все" --lang=ru -o gpxallfiles/ru_ALL.gpx
	$(TARGET) wikidata ./latest-all.json.bz2 --gpx-metadata-name="Wiki all" --lang=en -o gpxallfiles/en_ALL.gpx

gpxfiles:
	mkdir -p gpxfiles
	$(TARGET) filter gpxallfiles/ru_ALL.gpx --gpx-metadata-name="Wiki Россия"     --country-code=RU -o gpxfiles/ru_RU.gpx
	$(TARGET) filter gpxallfiles/en_ALL.gpx --gpx-metadata-name="Wiki USA"   --country-code=US -o gpxfiles/en_US.gpx
	$(TARGET) filter gpxallfiles/en_ALL.gpx --gpx-metadata-name="Wiki GB"    --country-code=GB -o gpxfiles/en_GB.gpx
