package mysqlModule

import "time"

const (
	TableTest = "test"
)

type Test struct {
	ID    uint      `gorm:"primarykey"`
	Time  time.Time `gorm:"column:time"`
	Key   string    `gorm:"column:key;index:idx_key"`
	Value string    `gorm:"column:value"`
}

func (Test) TableName() string {
	return "test"
}
