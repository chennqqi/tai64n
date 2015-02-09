package tai64n

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

const TAI64OriginalBase = 4611686018427387904

var (
	nextLS       = time.Date(2015, time.July, 1, 0, 0, 0, 0, time.UTC)
	nextLSOffset = 36
	curLS        = time.Date(2012, time.July, 1, 0, 0, 0, 0, time.UTC)
	curLSOffset  = 35
)

func nowBase(now time.Time) int64 {
	sec := now.Unix()

	// perf bias: most users set their server time to the current
	// time on earth, so we bias this to check that we're in that
	// time region before checking the complete leap second table.

	switch {
	case sec >= nextLS.Unix():
		return int64(TAI64OriginalBase + nextLSOffset)
	case sec >= curLS.Unix():
		return int64(TAI64OriginalBase + curLSOffset)
	default:
		return int64(TAI64OriginalBase + LeapSecondsInvolved(now))
	}
}

type TimeComparison int

const (
	Before TimeComparison = 0
	Equal                 = iota
	After                 = iota
)

func Now() *TAI64N {
	t := time.Now()

	return &TAI64N{
		Seconds:     uint64(t.Unix() + nowBase(t)),
		Nanoseconds: uint32(t.Nanosecond()),
	}
}

func FromTime(t time.Time) *TAI64N {
	return &TAI64N{
		Seconds:     uint64(t.Unix() + int64(TAI64OriginalBase+LeapSecondsInvolved(t))),
		Nanoseconds: uint32(t.Nanosecond()),
	}
}

func (tai *TAI64N) Time() time.Time {
	t := time.Unix(int64(tai.Seconds-TAI64OriginalBase), int64(tai.Nanoseconds)).UTC()

	return t.Add(-time.Duration(LeapSecondsInvolved(t)) * time.Second)
}

func (tai *TAI64N) WriteStorage(buf []byte) {
	binary.BigEndian.PutUint64(buf[:], tai.Seconds)
	binary.BigEndian.PutUint32(buf[8:], tai.Nanoseconds)
}

func (tai *TAI64N) ReadStorage(buf []byte) {
	tai.Seconds = binary.BigEndian.Uint64(buf[:])
	tai.Nanoseconds = binary.BigEndian.Uint32(buf[8:])
}

func (tai *TAI64N) Label() string {
	var buf [12]byte

	tai.WriteStorage(buf[:])

	s := fmt.Sprintf("@%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X%02X",
		buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6],
		buf[7], buf[8], buf[9], buf[10], buf[11])

	return s
}

func ParseTAI64NLabel(label string) *TAI64N {
	if label[0] != '@' {
		return nil
	}

	buf, err := hex.DecodeString(label[1:])

	if len(buf) != 12 || err != nil {
		return nil
	}

	ts := &TAI64N{}

	ts.ReadStorage(buf[:])

	return ts
}

func (tai TAI64N) MarshalJSON() ([]byte, error) {
	return tai.Time().MarshalJSON()
}

func (tai *TAI64N) UnmarshalJSON(data []byte) (err error) {
	var t time.Time
	err = t.UnmarshalJSON(data)

	*tai = *FromTime(t)

	return err
}

func (tai *TAI64N) Before(other *TAI64N) bool {
	return tai.Compare(other) == Before
}

func (tai *TAI64N) After(other *TAI64N) bool {
	return tai.Compare(other) == After
}

func (tai *TAI64N) Compare(other *TAI64N) TimeComparison {
	if tai.Seconds < other.Seconds {
		return Before
	}

	if tai.Seconds > other.Seconds {
		return After
	}

	if tai.Nanoseconds < other.Nanoseconds {
		return Before
	}

	if tai.Nanoseconds > other.Nanoseconds {
		return After
	}

	return Equal
}

func (tai *TAI64N) Add(dur time.Duration) *TAI64N {
	return FromTime(tai.Time().Add(dur))
}
