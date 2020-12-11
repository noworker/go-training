package model

import "time"

type UnixTime int64

func (ut UnixTime) Before(other UnixTime) bool {
	return ut < other
}

func CurrentUnixTime() UnixTime {
	return UnixTime(time.Now().Unix())
}

func (ut UnixTime) AddMinutes(min int64) UnixTime {
	return ut + UnixTime(min*60)
}

func (ut UnixTime) AddHours(hour int64) UnixTime {
	return ut.AddMinutes(hour * 60)
}
