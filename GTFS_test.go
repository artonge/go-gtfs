package GTFS

import (
	"fmt"
	"path"
	"testing"
)

func TestLoad(t *testing.T) {
	gtfs, err := Load(path.Join("gtfs_test", "ratp"))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(gtfs.Calendars) == 0 ||
		len(gtfs.Routes) == 0 ||
		len(gtfs.Stops) == 0 ||
		len(gtfs.StopsTimes) == 0 ||
		len(gtfs.Transfers) == 0 ||
		len(gtfs.Trips) == 0 {
		t.Fail()
	}
}

func TestLoadBad(t *testing.T) {
	_, err := Load(path.Join("gtfs_test", "ratp_bad"))
	if err == nil {
		t.Fail()
	}
}

func TestLoadNoDir(t *testing.T) {
	_, err := Load(path.Join("gtfs_test", "no_dir"))
	if err == nil {
		t.Fail()
	}
}

func TestLoadSplitted(t *testing.T) {
	gtfss, err := LoadSplitted(path.Join("gtfs_test", "ratp_splitted"))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(gtfss) != 2 {
		t.Fail()
	}
}

func TestLoadSplittedBad(t *testing.T) {
	_, err := LoadSplitted(path.Join("gtfs_test", "ratp_splitted_bad"))
	if err == nil {
		t.Fail()
	}
}

func TestLoadSplittedNoDir(t *testing.T) {
	_, err := LoadSplitted(path.Join("gtfs_test", "no_dir"))
	if err == nil {
		t.Fail()
	}
}
