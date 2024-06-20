CONCURRENCY : Design pattern and way to write code that has possibility to run in parallel (depends of cpu)
PARALLELISM : Code that runs on parallel (multiprocessor cpu is assumed)
SYNCHRONIZATION VARS: Are variables used for synchronization of thread execution (wait-group,atomic,mutex)
SHARED VALUES : IN golang are passed around on channels (Don't communicate by sharing memory,share memory by communicating)
METHOD SETS : Determine what methods attach to a type.Also method set of a type determines the INTERFACES that a type implements
RACE CONDITION : State when we have shared variables and threads compete for read-write,and we get incorrect data
GO-ROUTINE: Function executing concurrently with other goroutines in same address space.GO-ROUTINES are multiplexed onto multiple OS threads.Function or method can be run on goroutine.