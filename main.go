package main

import (
	"fmt"
	"os"

	"github.com/a1fred/wikidata2gpx.git/wikidata2gpx/cmd"
)

var revision = "unknown"

func main() {
	os.Stderr.WriteString(fmt.Sprintf("wikidata2gpx version %s\n", revision))
	cmd.Execute()
	os.Exit(0)
}
