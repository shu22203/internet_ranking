package ir

import (
	"time"
)

type Section struct {
	SectionChallenges
	startAt    time.Time
	endAt      time.Time
}
