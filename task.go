package main

import (
	"fmt"
	"log"
	"time"
)

func taskObserveChannelsConnections() {
	ticker := time.NewTicker(time.Duration(observeInterval) * time.Second)
	for range ticker.C {
		for _, ch := range observedChannels {
			total, err := selectActiveConnectionsCount(ch.ID)
			if err != nil {
				log.Println(err)
			}
			log.Println(fmt.Sprintf("current active channel connections from %s is: %v", ch.Name, total))
			gauge.WithLabelValues(fmt.Sprint(ch.ID), ch.Name, ch.UUID).Set(float64(total))
		}
	}
}
