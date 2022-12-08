# wikidata2gpx cli
## Commands
```sh
$ ./bin/wikidata2gpx --help
Wikidata pois exporter

Usage:
  wikidata2gpx [flags]
  wikidata2gpx [command]

Available Commands:
  wikidata    Generate gpx from wikidata dumps
  filter      Filter gpx file points
  help        Help about any command
  completion  generate the autocompletion script for the specified shell

Flags:
  -h, --help   help for wikidata2gpx

Use "wikidata2gpx [command] --help" for more information about a command.
```
## wikidata
```sh
$ ./bin/wikidata2gpx wikidata --help
Generate gpx from wikidata dumps

Usage:
  wikidata2gpx wikidata <wikidata-dump-file> [flags]

Flags:
      --cache string               Cache file path, used for cache wikidata and nominatim api responses. (default ".wikidata2gpx_cache.bolt")
      --gpx-metadata-name string   Gpx metadata name (default "wikidata2gpx")
  -h, --help                       help for wikidata
  -l, --lang string                Poi language (default "en")
  -o, --output string              result .gpx file (default "result.gpx")
```
## filter
```sh
$ ./bin/wikidata2gpx filter --help
Filter gpx file points

Usage:
  wikidata2gpx filter <gpx-file.gpx> [flags]

Flags:
      --cache string               Cache file path, used for cache wikidata and nominatim api responses. (default ".wikidata2gpx_cache.bolt")
      --country-code string        Generate only for country code (ISO_3166-1), all countries if empty (default)
      --gpx-metadata-name string   Gpx metadata name (default "wikidata2gpx")
  -h, --help                       help for filter
  -o, --output string              result .gpx file (default "result.gpx")
```
