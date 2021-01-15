package test

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTimeToStr(t *testing.T) {
	nt := time.Now()
	fmt.Println(nt)
	const base_format = "2006-01-02 15:04:05"
	str_time := nt.Format(base_format)
	fmt.Println(str_time)
	parse_str_time, _ := time.Parse(base_format, str_time)
	fmt.Println(parse_str_time)
	fmt.Println(parse_str_time.Format(base_format))

	dd, _ := time.ParseDuration("24h")
	dd1 := parse_str_time.Add(dd)
	fmt.Println(dd1)
}

func TestCurrentTimeToStr(t *testing.T) {
	nt := time.Now()
	fmt.Println(nt)
	const base_format = "2006-01-02 15:04:05"
	str_time := nt.Format(base_format)
	fmt.Println(str_time)
	cur_time := strings.Split(str_time, ":")[0] + ":00:00"
	fmt.Println(cur_time)
}
