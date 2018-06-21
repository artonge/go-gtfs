package gtfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/artonge/go-csv-tag"
)

// Load - load GTFS files
// @param dirPath: the directory containing the GTFS
// @param filter:
// It is possible to partialy load the gtfs files
// If you don't want to load all the files, add an param to the Load function
// containing only the files that must be loaded
// Example:
// Load("path/to/gtfs", map[string]bool{"routes": true})
// @return a filled GTFS or an error
func Load(dirPath string, filter map[string]bool) (*GTFS, error) {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("Error loading GTFS: directory does not existe")
	}
	g := &GTFS{Path: dirPath}
	err = loadGTFS(g, filter)
	if err != nil {
		return nil, fmt.Errorf("Error loading GTFS: '%v'\n	==> %v", g.Path, err)
	}
	return g, nil
}

// LoadSplitted - load splitted GTFS files
// ==> When GTFS are splitted into sub directories
// @param dirPath: the directory containing the sub GTFSs
// @param filter: see Load function
// @return an array of filled GTFS or an error
func LoadSplitted(dirPath string, filter map[string]bool) ([]*GTFS, error) {
	// Get directory list
	subDirs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	// Prepare the array of GTFSs
	GTFSs := make([]*GTFS, len(subDirs))
	i := 0
	// Load each sub directory into a new GTFS struct
	for _, dir := range subDirs {
		if !dir.IsDir() {
			continue
		}
		GTFSs[i] = &GTFS{Path: path.Join(dirPath, dir.Name())}
		err := loadGTFS(GTFSs[i], filter)
		if err != nil {
			return nil, fmt.Errorf("Error loading GTFS: '%v'\n	==> %v", GTFSs[i].Path, err)
		}
		i++
	}
	// Return a slice of GTFSs, because we skipped non directory files
	return GTFSs[:i], nil
}

// Load a directory containing:
// 	- routes.txt
// 	- stops.txt
// @param g: the GTFS struct that will receive the data
// @return an error
func loadGTFS(g *GTFS, filter map[string]bool) error {
	// Create a slice of agency to load agency.txt
	var agencySlice []Agency
	// List all files that will be loaded and there dest
	filesToLoad := map[string]interface{}{
		"agency.txt":         &agencySlice,
		"calendar.txt":       &g.Calendars,
		"calendar_dates.txt": &g.CalendarDates,
		"routes.txt":         &g.Routes,
		"stops.txt":          &g.Stops,
		"stop_times.txt":     &g.StopsTimes,
		"transfers.txt":      &g.Transfers,
		"trips.txt":          &g.Trips,
	}
	// Load the files
	for file, dest := range filesToLoad {
		// If filter not nil check if me need to load the current file
		if filter != nil && !filter[file[:len(file)-4]] {
			continue
		}
		filePath := path.Join(g.Path, file)
		// If the file does not existe, skip it
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			continue
		}
		err = csvtag.Load(csvtag.Config{
			Path: filePath,
			Dest: dest,
		})
		if err != nil {
			return fmt.Errorf("Error loading file (%v)\n	==> %v", file, err)
		}
	}
	// Put the loaded agency in g.Agency
	if len(agencySlice) > 0 {
		g.Agency = agencySlice[0]
	}
	return nil
}

// Dump GTFS data to an already existing directory
// @param g: GTFS struct to dump
// @param dirPath: Target directory
// @param filter: same as for load function
// @return error
func Dump(g *GTFS, dirPath string, filter map[string]bool) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, os.ModeDir)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	files := map[string]interface{}{
		"agency.txt":			[]Agency{g.Agency},
		"calendar.txt":			g.Calendars,
		"calendar_dates.txt": 	g.CalendarDates,
		"routes.txt":			g.Routes,
		"stops.txt":			g.Stops,
		"stop_times.txt":		g.StopsTimes,
		"transfers.txt":		g.Transfers,
		"trips.txt":			g.Trips,
	}
	for file, src := range files {
		if filter != nil && !filter[file[:len(file)-4]] {
			continue
		}
		if src == nil {
			continue
		}
		filePath := path.Join(dirPath, file)

		err := csvtag.DumpToFile(src, filePath)
		if err != nil {
			return fmt.Errorf("Error dumping file %v: %v", file, err)
		}
	}
	return err
}
