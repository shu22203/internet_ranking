package main

import (
	"errors"
	"time"
)

type IrSections struct {
	sections []Section
}

func (irs *IrSections) FirstStartAt() (time.Time, error) {
	if len(irs.sections) == 0 {
		return time.Time{}, errors.New("sections is empty")
	}

	firstStartAt := irs.sections[0].startAt
	for _, section := range irs.sections {
		if section.startAt.Before(firstStartAt) {
			firstStartAt = section.startAt
		}
	}

	return firstStartAt, nil
}

func (irs *IrSections) LastEndAt() (time.Time, error) {
	if len(irs.sections) == 0 {
		return time.Time{}, errors.New("sections is empty")
	}

	lastEndAt := irs.sections[0].endAt
	for _, section := range irs.sections {
		if section.endAt.After(lastEndAt) {
			lastEndAt = section.endAt
		}
	}

	return lastEndAt, nil
}
