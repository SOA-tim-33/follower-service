package model

type Follow struct {
	Id         int `gorm:"primary_key;auto_increment"`
	ProfileId  int64
	FollowerId int64
}
