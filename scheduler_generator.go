package scheduler

func NewScheduler() Scheduler{
    trig := make(chan Task, 100)
    return Scheduler{ tasks: make([]Task, 0), trig: trig, changed: make(chan bool) }
}
