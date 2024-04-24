package model

type Profile struct {
	Id                     int       `gorm:"primary_key;auto_increment"`
	FirstName              string    `gorm:"default:null"`
	LastName               string    `gorm:"default:null"`
	ProfilePicture         string    `gorm:"default:null"`
	Biography              string    `gorm:"default:null"`
	Motto                  string    `gorm:"default:null"`
	UserID                 int64     `gorm:"primaryKey"`
	IsActive               bool      `gorm:"default:false"`
	Follows                []*Follow `gorm:"many2many:profiles_follows;"`
	TourPreference         TourPreference
	XP                     int  `gorm:"default:0"`
	IsFirstPurchased       bool `gorm:"default:false"`
	QuestionnaireDone      bool `gorm:"default:false"`
	NumberOfCompletedTours int  `gorm:"default:0"`
	RequestSent            bool `gorm:"default:false"`
}
