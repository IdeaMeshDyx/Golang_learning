package interface_chan_routine

import (
	"fmt"
	"testing"
	"time"
)

/*
demo: interface/channel/routine配合使用
channel以参数的形式传入routine即可
*/

type Person interface {
	Say() string
}

type Worker struct {
	WorkID int
	Name   string
}

type Student struct {
	StuID int
	Name  string
}

func (w Worker) Say() string {
	return fmt.Sprintf("ID: %d, Name: %s", w.WorkID, w.Name)
}

func (s Student) Say() string {
	return fmt.Sprintf("ID: %d, Name: %s", s.StuID, s.Name)
}
func Send(ch chan Person) {
	for i := 1; ; i++ {
		ch <- Worker{
			WorkID: i,
			Name:   fmt.Sprintf("Worker %d", i),
		}
		ch <- Student{
			StuID: i,
			Name:  fmt.Sprintf("Student %d", i),
		}

		time.Sleep(1 * time.Second)
	}
}

func Recv(ch chan Person) {
	for p := range ch {
		fmt.Println(p)
	}

}

func Test_main(t *testing.T) {
	ch := make(chan Person)
	go Send(ch)
	go Recv(ch)
	time.Sleep(5 * time.Second)
}
