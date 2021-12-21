package jobqueue

import (
	"errors"
	"otter-v2/api/http/response"

	"github.com/EricChiou/jobqueue"
)

type worker struct {
	run  func() interface{}
	wait *chan interface{}
}

func Init() {
	run(&User.list)
}

func Wait() {
	User.list.Wait()
}

func run(queue *jobqueue.Queue) {
	queue.SetWorker(func(w interface{}) {
		if w, ok := w.(worker); ok {
			*w.wait <- w.run()
		} else {
			*w.wait <- errors.New(string(response.ServerError))
		}
	})
	queue.Run()
}
