package ir

import (
	"errors"
	"time"
)

type IrSections []Section

func (irs *IrSections) FirstStartAt() (time.Time, error) {
	if len(*irs) == 0 {
		return time.Time{}, errors.New("sections is empty")
	}

	firstStartAt := (*irs)[0].startAt
	for _, section := range *irs {
		if section.startAt.Before(firstStartAt) {
			firstStartAt = section.startAt
		}
	}

	return firstStartAt, nil
}

func (irs *IrSections) LastEndAt() (time.Time, error) {
	if len(*irs) == 0 {
		return time.Time{}, errors.New("sections is empty")
	}

	lastEndAt := (*irs)[0].endAt
	for _, section := range *irs {
		if section.endAt.After(lastEndAt) {
			lastEndAt = section.endAt
		}
	}

	return lastEndAt, nil
}
