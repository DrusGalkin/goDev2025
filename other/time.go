package other

import "fmt"

type Time struct {
	hour,
	min,
	sec int
}

func NewTime(sec uint) Time {
	return Time{int(sec / 3600), int(sec%3600) / 60, int(sec % 60)}
}

func (t Time) Print() {
	fmt.Printf("%d:%d:%d\n", t.hour, t.min, t.sec)
}
