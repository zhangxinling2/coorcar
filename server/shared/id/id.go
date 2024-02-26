package id

type AccountId string

func (a AccountId) String() string {
	return string(a)
}

type TripId string

func (a TripId) String() string {
	return string(a)
}
