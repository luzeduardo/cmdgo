package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task       string
	Done       bool
	CreateAt   time.Time
	CompleteAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:       task,
		Done:       false,
		CreateAt:   time.Now(),
		CompleteAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompleteAt = time.Now()
	return nil
}
