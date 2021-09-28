package srt

import (
	"fmt"
	"strings"
	"time"
)

const timestampFormat = "15:04:05,000"

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

func (s *Subtitle) ToSRT() string {
	return fmt.Sprintf(`%d
%s --> %s
%s

`,
		s.Position,
		s.Start.Format(timestampFormat),
		s.End.Format(timestampFormat),
		strings.Join(s.TextLines, "\n"),
	)
}
