package profiling

// When we wish to look carefully at the speed of our programs, the best technique for
// identifying  the  critical  code  is  profiling.

//* A CPU profile identifies the functions whose execution requires the most CPU time.
//* The currently running thread on each CPU is interrupted periodically by the operating
//* system  every  few  milliseconds,  with  each  interruption  recording  one  profile  event
//* before normal execution resumes.
//
//* A heap profile identifies the statements responsible for allocating the most memory.
//* The profiling library samples calls to the internal memory allocation routines so that
//* on average, one profile event is recorded per 512KB of allocated memory.
//
//* A  blocking  profile  identifies  the  operations  responsible  for  blocking  goroutines  the
//* longest, such as system calls, channel sends and receives, and acquisitions of locks.
//* The profiling library records an event every time a goroutine is blocked by one of
//* these operations.

//? pprof needs the executable in order to make sense of the log.

// Although go test usually discards  the  test  executable  once  the  test  is  complete,
// when  profiling  is  enabled  it saves the executable as foo.test, where foo is the name of the tested package.

//* go test -run=NONE -bench=ClientServerParallelTLS64 -cpuprofile=cpu.log net/http
