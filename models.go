package main

// Source ::
type Source struct {
	ID   *int    `gorm:"id"`
	UUID *string `gorm:"uuid"`
	Name *string `gorm:"name"`
}

func (Source) TableName() string {
	return "source"
}

// Target ::
type Target struct {
	ID   *int    `gorm:"id"`
	UUID *string `gorm:"uuid"`
	Name *string `gorm:"name"`
}

func (Target) TableName() string {
	return "target"
}
