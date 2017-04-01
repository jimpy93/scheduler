# scheduler
Basic scheduler implementation in Go.
Very simple to use.

Create the scheduler with:

```
sch := scheduler.NewScheduler();
```

Start it using:

```
sch.Start()
```

Add Tasks:

```
t := time.Date(year, time.Month(month), date, hour, min, 0, 0, time.Now().Location())

task := "some task. This can be of any instance type"

sch.Add(scheduler.Task{Time: t, Task: task})
```

You will be notified whenever any scheduled task needs to be triggered

```
 task := <- sch.TriggerChan()
```
