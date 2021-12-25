package jobqueue

import "github.com/EricChiou/jobqueue"

type user struct {
	list jobqueue.Queue
}

func (u *user) NewUserListJob(run func() interface{}) error {
	wait := make(chan interface{})
	u.list.Add(worker{run: run, wait: &wait})

	result := <-wait
	switch r := result.(type) {
	case error:
		return r
	default:
		return nil
	}
}

var User user = user{
	list: jobqueue.New(1024),
}
