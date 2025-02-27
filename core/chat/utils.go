package chat

import (
	"syscall"

	log "github.com/sirupsen/logrus"
)

func getMaximumConcurrentConnectionLimit() int64 {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		log.Fatalln(err)
	}

	// Return the limit to 70% of max so the machine doesn't die even if it's maxed out for some reason.
	proposedLimit := int64(float32(rLimit.Max) * 0.7)

	return proposedLimit
}
