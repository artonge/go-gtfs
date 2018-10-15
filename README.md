# go-gtfs
Load GTFS files in golang

[![godoc for artonge/go-gtfs](https://godoc.org/github.com/artonge/go-gtfs?status.svg)](http://godoc.org/github.com/artonge/go-gtfs)

[![Build Status](https://travis-ci.org/artonge/go-gtfs.svg?branch=master)](https://travis-ci.org/artonge/go-gtfs)
![cover.run go](https://cover.run/go/github.com/artonge/go-gtfs.svg)
[![goreportcard for artonge/go-gtfs](https://goreportcard.com/badge/github.com/artonge/go-gtfs)](https://goreportcard.com/report/artonge/go-gtfs)

[![Sourcegraph for artonge/go-gtfs](https://sourcegraph.com/github.com/artonge/go-gtfs/-/badge.svg)](https://sourcegraph.com/github.com/artonge/go-gtfs?badge)

# Install
`go get github.com/artonge/go-gtfs`

# Examples
## Load one directory containing GTFS files:
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
g, err := gtfs.Load("path/to/gtfs_files", nil)
```

## Load a directory containing sub directories containing GTFS files:
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
gs, err := gtfs.LoadSplitted("path/to/gtfs_directories", nil)
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
