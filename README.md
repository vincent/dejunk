# Dejunk 

Dejunk is a file sorter, to arrange movies, TV shows, and music files in a nice structure.

```shell
dejunk --out /library --in /my-music --in /my-movies
├ Movies
│ ├ The Current War (2020)
│ │ └ The Current War.mkv
│ └ Back To The Future (1990)
│   └ Back To Th Future.mp4
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
- name: Music
  match: "ext(:audio)"
  type: Music
  store: ":artist/:album (:year)/:title"
  with: [dummy, tags, artwork]

- name: Movies
  match: "ext(:video)not(:episode)"
  type: Movie
  store: ":title (:year)/:title"
  with: [dummy, tags, artwork]

- name: TV Shows
  match: "ext(:video)is(:episode)"
  type: TVShow
  store: ":title/Season :season/:episode - :title"
  with: [dummy, tags, artwork]
```
