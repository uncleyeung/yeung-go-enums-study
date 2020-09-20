package base

import "time"

//父类
type Base struct {
	Id         int64
	CreamTime  time.Time
	UpdateTime time.Time
}

func (b *Base) InitData() {
	b.Id = 123456
	b.CreamTime = time.Now()
	b.UpdateTime = time.Now()
}
