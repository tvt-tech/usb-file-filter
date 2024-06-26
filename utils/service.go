package utils

import (
	"time"
)

var connected_drives []string

func upd_drives() {
	if drives, err := Detect(); err != nil {
		upd_drives()
	} else {
		connected_drives = drives
	}
}

func init() {
	upd_drives()
}

func isInList(search string, list []string) bool {
	for _, item := range list {
		if item == search {
			return true // If found, return true
		}
	}
	return false // If not found, return false
}

func update() {
	// beepFunc.Call(0xffffffff)
	prev_drives := make([]string, len(connected_drives))
	copy(prev_drives, connected_drives)

	upd_drives()
	for _, new_drive := range connected_drives {
		if !isInList(new_drive, prev_drives) {
			Logger.Infof("New connected: %s", new_drive)
			go runArchiveInRoutine(new_drive)
		}
	}
}

func runArchiveInRoutine(drive string) {
	if err := ArchiveNotAllowed(drive); err != nil {
		Logger.Error(err)
	}
}

func StartService() {
	for {
		time.Sleep(100 * time.Millisecond)
		update()
	}
}
