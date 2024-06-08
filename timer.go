package main

import (
	"fmt"
	"time"
)

func timeTrack(start time.Time, name string) {
	if DEBUG == 1 {
		elapsed := time.Since(start)
		fmt.Println(name+":", "Finished in", elapsed)
	}
}
