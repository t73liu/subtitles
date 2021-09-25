package srt

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadSRTFile(path string) ([]*Subtitle, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var subtitles []*Subtitle
	var prevPosition int

	for scanner.Scan() {
		currLine := strings.TrimSpace(scanner.Text())
		if len(currLine) == 0 {
			continue
		}
		position, err := parsePositionLine(scanner)
		if position != prevPosition+1 {
			log.Fatalf("unexpected position: %s", currLine)
		}
		start, end, err := parseTimestampsLine(scanner)
		if err != nil {
			log.Fatalln(err)
		}
		textLines, err := parseTextLines(scanner)
		if err != nil {
			log.Fatalln(err)
		}
		subtitles = append(subtitles, &Subtitle{
			Position:  position,
			Start:     start,
			End:       end,
			TextLines: textLines,
		})
		prevPosition = position
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return subtitles, nil
}

func parsePositionLine(scanner *bufio.Scanner) (position int, err error) {
	if !scanner.Scan() {
		return position, errors.New("unexpected EOF for position line")
	}
	line := strings.TrimSpace(scanner.Text())
	if len(line) == 0 {
		return position, errors.New("missing position line")
	}
	position, err = strconv.Atoi(line)
	if err != nil {
		return position, err
	}
	return position, nil
}

func parseTimestampsLine(scanner *bufio.Scanner) (start time.Time, end time.Time, err error) {
	if !scanner.Scan() {
		return start, end, errors.New("unexpected EOF for timestamps line")
	}
	line := strings.TrimSpace(scanner.Text())
	timestampsArr := strings.Split(line, " ")
	if len(timestampsArr) != 3 {
		return start, end, errors.New("improperly formatted timestamps line")
	}
	start, err = time.Parse(TimestampFormat, timestampsArr[0])
	if err != nil {
		return start, end, err
	}
	end, err = time.Parse(TimestampFormat, timestampsArr[2])
	if err != nil {
		return start, end, err
	}
	return start, end, nil
}

func parseTextLines(scanner *bufio.Scanner) (text []string, err error) {
	if !scanner.Scan() {
		return text, errors.New("unexpected EOF for text line(s)")
	}
	currLine := strings.TrimSpace(scanner.Text())
	if len(currLine) == 0 {
		return text, errors.New("missing text line(s)")
	}
	text = append(text, currLine)
	// Read lines until an empty line is reached.
	for scanner.Scan() {
		currLine = strings.TrimSpace(scanner.Text())
		if len(currLine) == 0 {
			break
		}
		text = append(text, currLine)
	}
	return text, nil
}

func WriteSRTFile(subtitles []*Subtitle, outputPath string, append bool) error {
	return nil
}
