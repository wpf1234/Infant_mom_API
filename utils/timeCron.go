package utils

import "time"

// 定时器，启动的时候执行一次，以后每天零点执行一次
func StartTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

func StartTicker(f func(), t *time.Ticker) {
	go func() {
		for {
			// 程序一执行就执行一次该函数
			f()

			// 根据传入的 ticker 定时执行
			<-t.C
		}
	}()
}

// 每天 8 点执行
func StartTimerAt8(f func()) {
	go func() {
		for {
			now := time.Now()
			//f()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 8, 0, 0, 0, next.Location())
			//h,_:=time.ParseDuration("1h")
			//next=next.Add(8*h)
			t := time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}
