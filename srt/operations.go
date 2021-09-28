package srt

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const byteOrderMark = string('\uFEFF')

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
		currLine := strings.TrimSpace(strings.TrimPrefix(scanner.Text(), byteOrderMark))
		if len(currLine) == 0 {
			continue
		}
		position, err := parsePositionLine(currLine)
		if err != nil {
			log.Fatalln(err)
		}
		if position != prevPosition+1 {
			log.Fatalf("unexpected position: %d to %d, %t", prevPosition, position, position != prevPosition+1)
		}
		start, end, err := readTimestamps(scanner)
		if err != nil {
			log.Fatalln(err)
		}
		textLines, err := readSubtitleText(scanner)
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

func parsePositionLine(line string) (position int, err error) {
	position, err = strconv.Atoi(line)
	if err != nil {
		return position, err
	}
	return position, nil
}

func readTimestamps(scanner *bufio.Scanner) (start time.Time, end time.Time, err error) {
	if !scanner.Scan() {
		return start, end, errors.New("unexpected EOF for timestamps line")
	}
	line := strings.TrimSpace(scanner.Text())
	timestampsArr := strings.Split(line, " ")
	if len(timestampsArr) != 3 {
		return start, end, errors.New("improperly formatted timestamps line")
	}
	start, err = time.Parse(timestampFormat, timestampsArr[0])
	if err != nil {
		return start, end, err
	}
	end, err = time.Parse(timestampFormat, timestampsArr[2])
	if err != nil {
		return start, end, err
	}
	return start, end, nil
}

func readSubtitleText(scanner *bufio.Scanner) (text []string, err error) {
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

func WriteSRTFile(subtitles []*Subtitle, outputPath string) error {
	file, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}

	w := bufio.NewWriter(file)

	for _, sub := range subtitles {
		if _, err := w.WriteString(sub.ToSRT()); err != nil {
			return fmt.Errorf("failed to write subtitle: %+v", sub)
		}
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %s", err)
	}

	return file.Close()
}
