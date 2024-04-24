package model

type TourPreference struct {
	Id            int `gorm:"primary_key;auto_increment"`
	Difficulty    int
	WalkingRating int
	BicycleRating int
	CarRating     int
	BoatRating    int
	Tags          []string
}
