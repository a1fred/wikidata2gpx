# wikidata2gpx

 * [Cli docs](CLI.md)

[![build](https://github.com/a1fred/wikidata2gpx/actions/workflows/main.yml/badge.svg)](https://github.com/a1fred/wikidata2gpx/actions/workflows/main.yml)

Generate gpx file from wikidata's structured data dumps.

* https://www.topografix.com/GPX/1/1/
* https://dumps.wikimedia.org/wikidatawiki/entities/
* https://www.wikidata.org/wiki/Property:P625

Install [pbzip2](http://compression.ca/pbzip2/) recomended for `.bz2` dumps.

## Examples
| GPX file                        |  Pois language |  Pois country |
|---------------------------------|----------------|---------------|
| [en_US.gpx](gpxfiles/en_US.gpx) | en             | US            |
| [en_GB.gpx](gpxfiles/en_GB.gpx) | en             | GB            |
| [ru_RU.gpx](gpxfiles/ru_RU.gpx) | ru             | RU            |

See all files in [gpxfiles](gpxfiles) folder, or generate your-own.

# Build
```sh
$ make
```

# Usage
```sh
$ ./build/wikidata2gpx --help
wikidata2gpx version 7d1d0be-master-20210825-23:57:08
Wikidata pois exporter

Usage:
  wikidata2gpx [flags]
  wikidata2gpx [command]

Available Commands:
  wikidata    Generate gpx from wikidata dumps
  filter      Filter gpx file points
  help        Help about any command

Flags:
  -h, --help   help for wikidata2gpx

Use "wikidata2gpx [command] --help" for more information about a command.
```

* Download bz2 dump
    ```sh
    $ wget -c https://dumps.wikimedia.org/wikidatawiki/entities/latest-all.json.bz2
    ```

* Generate localized GPX file from dump
    ```sh
    $ ./bin/wikidata2gpx wikidata ./latest-all.json.bz2 --lang=ru -o=all.gpx
    ```
    `all.gpx` will contain all wikipedia pois localized in russian language.

* Optional filter by some params
    ```sh
    # Get only Russian points
    $ ./bin/wikidata2gpx filter ./all.gpx --country-code=RU -o=result.gpx
    ```
    `result.gpx` will contain only Russian pois localized in russian language.

## Completions
 * [bash completion](completions/completion.bash)
 * [fish completion](completions/completion.fish)
 * [powershell completion](completions/completion.powershell)
 * [zsh completion](completions/completion.zsh)
