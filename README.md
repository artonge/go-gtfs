# go-gtfs
Load GTFS files in golang

![cover.run go](https://cover.run/go/github.com/artonge/go-gtfs.svg)
[![godoc for artonge/go-gtfs](https://godoc.org/github.com/artonge/go-gtfs?status.svg)](http://godoc.org/github.com/artonge/go-gtfs)
[![goreportcard for artonge/go-gtfs](https://goreportcard.com/badge/github.com/artonge/go-gtfs)](https://goreportcard.com/report/artonge/go-gtfs)
[![Sourcegraph for artonge/go-gtfs](https://sourcegraph.com/github.com/artonge/go-gtfs/-/badge.svg)](https://sourcegraph.com/github.com/artonge/go-gtfs?badge)


# Install
`go get github.com/artonge/go-gtfs`

# Examples
To load one directory containing GTFS files:
```bash
path/to/gtfs_files
├── agency.txt
├── calendar_dates.txt
├── calendar.txt
├── routes.txt
├── stops.txt
├── stop_times.txt
├── transfers.txt
└── trips.txt
```
```go
gtfs, err := Load("path/to/gtfs_files")
```

To load a directory containing sub directories containing GTFS files:
```bash
path/to/gtfs_directories
├── gtfs1
│   ├── agency.txt
│   ├── calendar_dates.txt
│   ├── routes.txt
│   ├── stops.txt
│   ├── stop_times.txt
│   ├── transfers.txt
│   └── trips.txt
└── gtfs2
    ├── agency.txt
    ├── calendar_dates.txt
    ├── calendar.txt
    ├── routes.txt
    ├── stops.txt
    ├── stop_times.txt
    ├── transfers.txt
    └── trips.txt

```
```go
gtfss, err := LoadSplitted("path/to/gtfs_directories")
```

You can then access the data through the GTFS structure.
That structure contains arrays of approriate structures for each files.
```go
type GTFS struct {
	Path       string // The path to the containing directory
	Agency     Agency
	Routes     []Route
	Stops      []Stop
	StopsTimes []Stop
	Trips      []Trip
	Calendars  []Calendar
	Transfers  []Transfer
}

type Route struct {
	ID        string `csv:"route_id"`
	AgencyID  string `csv:"agency_id"`
	ShortName string `csv:"route_short_name"`
	LongName  string `csv:"route_long_name"`
	Type      int    `csv:"route_type"`
	Desc      string `csv:"route_url"`
	URL       string `csv:"route_desc"`
	Color     string `csv:"route_color"`
	TextColor string `csv:"route_text_color"`
}

...
```

# Contributions
Pull requests are welcome ! :)

## TODO
- [ ] Add `Dump(g *GTFS, file string) error` function to write a GTFS struct to the disk (need go-csv-tag implementation of Dump)   
