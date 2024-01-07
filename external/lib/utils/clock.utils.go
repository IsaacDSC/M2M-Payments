package utils

import (
	"strconv"
	"strings"
	"time"
)

func FromNumbExpiredToTimer(expired string) (output time.Time, err error) {
	input := strings.Split(expired, "/")
	if len(input) != 2 {
		return
	}
	year, err := strconv.Atoi("20" + input[0])
	if err != nil {
		return
	}
	month, err := strconv.Atoi(input[1])
	if err != nil {
		return
	}
	output = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
	return
}
