package overrustlelogs

import (
	"bufio"
	"net/http"
	"strings"
	"time"
)

const TimestampFormat = "2006-01-02 15:04:05"

type OverrustleLog struct {
	Message   string
	Timestamp time.Time
	User      string
}

type OverrustleLogs struct {
	scanner *bufio.Scanner
}

func (logs *OverrustleLogs) Scan() bool {
	return logs.scanner.Scan()
}

func (logs *OverrustleLogs) Log() (OverrustleLog, error) {
	line := logs.scanner.Text()
	message := line[strings.Index(line, ": ")+2:]

	timestamp, err := time.Parse(TimestampFormat, line[1:20])
	if err != nil {
		return OverrustleLog{}, err
	}

	user := line[26:strings.Index(line, ": ")]
	return OverrustleLog{
		Message:   message,
		Timestamp: timestamp,
		User:      user,
	}, nil
}

func New(url string) (*OverrustleLogs, error) {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := netClient.Get(url)
	if err != nil {
		return nil, err
	}

	return &OverrustleLogs{
		scanner: bufio.NewScanner(res.Body),
	}, nil
}
