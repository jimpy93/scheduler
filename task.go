package scheduler

import "time"

type Task struct {
    Time time.Time
    Task interface{}
}
