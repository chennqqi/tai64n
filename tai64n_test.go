package tai64n

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNowBase(t *testing.T) {
	var n time.Time

	n = time.Date(2016, time.May, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, TAI64OriginalBase+36, nowBase(n))

	n = time.Date(2014, time.May, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, TAI64OriginalBase+35, nowBase(n))

	n = time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, TAI64OriginalBase+34, nowBase(n))
}

func TestNow(t *testing.T) {
	t1 := Now()
	t2 := Now()

	assert.True(t, t1.Seconds <= t2.Seconds)
	assert.True(t, t1.Nanoseconds < t2.Nanoseconds)
}

func TestFromTime(t *testing.T) {
	var n time.Time

	n = time.Date(2016, time.May, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, uint64(4611686019889448740), FromTime(n).Seconds)

	n = time.Date(2014, time.May, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, uint64(4611686019826290339), FromTime(n).Seconds)

	n = time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, uint64(4611686019731595938), FromTime(n).Seconds)
}

func TestTime(t *testing.T) {
	m1 := time.Date(2011, time.May, 1, 0, 0, 0, 0, time.UTC)

	t1 := FromTime(m1)

	m2 := t1.Time()

	assert.True(t, m1.Equal(m2))
}

func TestLabel(t *testing.T) {
	m1 := FromTime(time.Date(2014, time.May, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, "@4000000053618EA300000000", m1.Label())
}

func TestParseLabel(t *testing.T) {
	m1 := FromTime(time.Date(2014, time.May, 1, 0, 0, 0, 0, time.UTC))
	m2 := ParseTAI64NLabel("@4000000053618EA300000000")

	assert.Equal(t, m1.Seconds, m2.Seconds)
}

type cont struct {
	Time *TAI64N
}

func TestJSON(t *testing.T) {
	m1 := Now()

	c1 := cont{Time: m1}

	bytes, err := json.Marshal(&c1)
	require.NoError(t, err)

	var c2 cont

	err = json.Unmarshal(bytes, &c2)
	require.NoError(t, err)

	assert.True(t, c2.Time.Seconds == m1.Seconds)
	assert.True(t, c2.Time.Nanoseconds == m1.Nanoseconds)
}

func TestCompare(t *testing.T) {
	m1 := Now()
	m2 := Now()

	assert.True(t, m1.Before(m2))
	assert.Equal(t, m1.Compare(m2), Before)

	assert.True(t, m2.After(m1))
	assert.Equal(t, m2.Compare(m1), After)

	assert.True(t, m1.Equal(m1))
	assert.Equal(t, m1.Compare(m1), Equal)
}

func TestAdd(t *testing.T) {
	m1 := Now()

	m2 := m1.Add(1 * time.Second)

	assert.Equal(t, m1.Seconds+1, m2.Seconds)
}