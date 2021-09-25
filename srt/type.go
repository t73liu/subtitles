package srt

import (
	"fmt"
	"time"
)

const TimestampFormat = "15:04:05,000"

const TimestampSeparator = "-->"

type Subtitle struct {
	Position  int       `json:"position"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	TextLines []string  `json:"textLines"`
}

func (s *Subtitle) AddDuration(duration time.Duration) {
	s.Start = s.Start.Add(duration)
	s.End = s.End.Add(duration)
}

func (s *Subtitle) write() string {
	return fmt.Sprintf("%d", s.Position)
}
