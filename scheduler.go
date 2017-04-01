package scheduler

import "time"

type Scheduler struct{
    tasks []Task
    trig chan Task
    changed chan bool
    running bool
}

func (q *Scheduler) Start(){
    if !q.running {
        q.running = true
        go func(){
            for q.running{
                h, hasHead := q.Head()
                for hasHead && h.Time.Before(time.Now()){
                    q.trig <- h
                    q.RemoveHead()
                    h, hasHead = q.Head()
                }
                var tchan <- chan time.Time
                if !hasHead{
                    tchan = time.After(5 * time.Minute)
                } else {
                    tchan = time.After(time.Now().Sub(h.Time))
                }
                select{
                case <- q.changed:
                case <- tchan:
                }
            }
        }()
    }
}

func (q *Scheduler) Stop(){
    if q.running{
        q.running = false
        q.changed <- true
    }
}

func (q *Scheduler) IsRunning() bool{
    return q.running
}

func (q *Scheduler) HasTasks() bool{
    return len(q.tasks) > 0
}

func (q *Scheduler) Tasks() []Task{
    return q.tasks
}

func (q *Scheduler) Head() (Task, bool){
    var t Task
    ok := true
    if !q.HasTasks(){
        ok = false
    } else {
        t = q.tasks[len(q.tasks)-1]
    }
    return t, ok
}

func (q *Scheduler) RemoveHead(){
    q.tasks = q.tasks[:len(q.tasks)-1]
}

func (q *Scheduler) findNewPos(t Task) int{
    for i,v := range q.tasks{
        if v.Time.Before(t.Time){
            return i
        }
    }
    return len(q.tasks)
}

func (q *Scheduler) findElem(t Task) int{
    for i,v := range q.tasks{
        if v == t{
            return i
        }
    }
    return -1
}

func (q *Scheduler) TriggerChan() chan Task{
    return q.trig
}

func (q *Scheduler) Add(t Task) {
    pos := q.findNewPos(t)
    if pos == len(q.tasks){
        q.tasks = append(q.tasks, t)
    } else {
        q.tasks = append(q.tasks, Task{})
        copy(q.tasks[pos+1:], q.tasks[pos:])
        q.tasks[pos] = t
    }
    if q.running{
        q.changed <- true
    }
}

func (q *Scheduler) Remove(t Task) bool{
    pos := q.findElem(t)
    if pos == -1 {
        return false
    }
    return q.RemoveFromPosition(pos)
}

func (q *Scheduler) RemoveFromPosition(pos int) bool{
    if pos < 0 || pos >= len(q.tasks){
        return false
    }
    q.tasks = append(q.tasks[:pos], q.tasks[pos+1:]...)
    if q.running{
        q.changed <- true
    }
    return true
}
