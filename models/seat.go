package models

type Seat struct {
	Pid        string
	Id         string
	Price      int
	Reservable bool
	// ..
}

type SeatCache struct {
	Pid        string
	Id         string
	Reservable bool
}
