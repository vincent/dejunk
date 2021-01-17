# Dejunk 

Dejunk is a file sorter, to arrange movies, TV shows, and music files in a nice structure.

> This is a learning project, do not use it for production !

## Install
- with Go 
```
go get github.com/vincent/dejunk
```
- with Docker
```
tbd
```

## Options
```shell
$ dejunk --help
Application Options:
  -l, --level=      Log level (fatal, error, info, debug) (default: error)
  -d, --dryrun      Print result without actually executing any tasks
  -i, --in=         Directories to scan
  -o, --out=        Directoriy to write files
  -r, --rules=      Rules file (default: ./samples/rules.yml)
  -b, --rollback=   Rollback file
  -x, --dorollback  Do rollback last run
```

## Rollback
When provided a rollback file, you can revert the last run if anything went wrong.<br>
Rollback file will be cleaned up so you can revert multiple runs sequentially.
```
$ dejunk -i /input -o /output -b rollback.log
$ ... oops something went wrong !
$ dejunk -x -b rollback.log
```

## Example
```
$ ls /downloads # a pile of mixed files
├ u2
│ ├ sunday_bloody_sunday.flac
│ └ new_year's_day.wav
├ TheCurrent_war_2020.mkv
├ dirk-gently-s1e03.avi
├ man_in the_high_castle-s2e01.avi
└ back.to.the.future_1990.mp4

gives

$ dejunk --in /downloads --out /library
├ Movies
│ ├ The Current War (2020)
│ │ └ The Current War.mkv
│ └ Back To The Future (1990)
│   └ Back To The Future.mp4
├ Music
│ ├ U2
│ │ └ War (1983)
│ │   ├ Sunday Bloody Sunday.flac
│ │   └ New Year's Day.wav
└ TV Shows
  ├ Dirk Gently
  │ └ Season 1
  │   └ 01 - Horizons.avi
  │   └ 02 - Lost and Found.avi
  │   └ 03 - Rogue Wall Enthusiasts.avi
  └ The Man In The High Castle
    └ Season 2
      └ 01 - The Tiger's Cave.avi
      └ 02 - The Road Less Traveled.avi
      └ 03 - Travelers.avi
```

Sorting rules are simply described by YAML files

```yaml
- name: Music                             # The rule name. This will be used as the first output directory 
  match: "ext(:audio)"                    # Matching rules. :audio will match all audio files
  type: Music                             # Internal category
  store: ":artist/:album (:year)/:title"  # The final storage path with all dynamic parts replaced
  with: [dummy, artwork]                  # Additional features.
                                          #     dummy: try to guess some tag values from file name
                                          #      artwork: fetch cover artwork
```
