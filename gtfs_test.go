package gtfs

import (
	"fmt"
	"path"
	"testing"
)

func TestLoad(t *testing.T) {
	gtfs, err := Load(path.Join("gtfs_test", "ratp"), nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(gtfs.Agencies) == 0 ||
		len(gtfs.Calendars) == 0 ||
		len(gtfs.CalendarDates) == 0 ||
		len(gtfs.Routes) == 0 ||
		len(gtfs.Stops) == 0 ||
		len(gtfs.StopsTimes) == 0 ||
		len(gtfs.Transfers) == 0 ||
		len(gtfs.Trips) == 0 {
		t.Fail()
	}

	if gtfs.Agencies[0] != gtfs.Agency {
		t.Fail()
	}
}

func TestLoadBad(t *testing.T) {
	_, err := Load(path.Join("gtfs_test", "ratp_bad"), nil)
	if err == nil {
		t.Fail()
	}
}

func TestLoadNoDir(t *testing.T) {
	_, err := Load(path.Join("gtfs_test", "no_dir"), nil)
	if err == nil {
		t.Fail()
	}
}

func TestLoadPartial(t *testing.T) {
	gtfs, err := Load(path.Join("gtfs_test", "ratp"), map[string]bool{"routes": true})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(gtfs.Agencies) != 0 ||
		len(gtfs.Calendars) != 0 ||
		len(gtfs.CalendarDates) != 0 ||
		len(gtfs.Routes) == 0 ||
		len(gtfs.Stops) != 0 ||
		len(gtfs.StopsTimes) != 0 ||
		len(gtfs.Transfers) != 0 ||
		len(gtfs.Trips) != 0 {
		t.Fail()
	}
}

func TestLoadSplitted(t *testing.T) {
	gtfss, err := LoadSplitted(path.Join("gtfs_test", "ratp_splitted"), nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(gtfss) != 2 {
		t.Fail()
	}
}

func TestLoadSplittedBad(t *testing.T) {
	_, err := LoadSplitted(path.Join("gtfs_test", "ratp_splitted_bad"), nil)
	if err == nil {
		t.Fail()
	}
}

func TestLoadSplittedNoDir(t *testing.T) {
	_, err := LoadSplitted(path.Join("gtfs_test", "no_dir"), nil)
	if err == nil {
		t.Fail()
	}
}
