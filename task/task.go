package task

import (
	"math/rand"
	"time"
)

type Task struct {
	ID int `json:"id"`
	Data string `json:"data"`
}

func (t Task) Process() (bool,string) {
	time.Sleep(1*time.Second)
	randInteger := rand.Intn(3)

	if randInteger == 0 {
		return false,string(rune(t.ID))
	}
	return true, "Processed Task ID: " + string(rune(t.ID))
}