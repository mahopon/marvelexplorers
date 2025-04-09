package utils

import (
	"log"
	"sync/atomic"
)

var maintenanceMode atomic.Bool

func EnableMaintenance(mode bool) {
	maintenanceMode.Store(mode)
	if maintenanceMode.Load() {
		log.Println("Maintenance mode activated")
	} else {
		log.Println("Maintenance mode deactivated")
	}
}

func CheckMaintenance() bool {
	return maintenanceMode.Load()
}
