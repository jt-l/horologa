package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
  "os"
  "errors"
)

/* A periodic task is defined to have:
phase: release time of the task
period: task needs to be scheduled every period and within the frame
execution_time: cycles that the task takes to execute
relative_deadline: time from the phase in which in the task must complete
*/
type Task struct {
  Phase int `json:"phase"`
  Period int `json:"period"`
  ExecutionTime int `json:"executionTime"`
  RelativeDeadline int `json:"relativeDeadline"`
}

type Tasks struct {
  T []Task `json:"tasks"`
}

// helper function to determine the gcd
func gcd(a int, b int) int {
  for a != b {

    if a > b {
      a -= b
    } else {
      b -= a
    }
  }

  return a
}

// determine the max execution time from all the tasks
func maxExecutionTime(tasks Tasks) int {
  max_execution_time := 0

  for _, task := range tasks.T {
    if task.ExecutionTime > max_execution_time {
      max_execution_time = task.ExecutionTime
    }
  }

  return max_execution_time
}


// the frame must divde evenly into the period of at least one task
func dividesEvenly(canidate_frame int, tasks Tasks) bool {

  divides_evenly := false

  for _, task := range tasks.T {
    if task.Period % canidate_frame == 0 { 
      divides_evenly = true
    }
  }

  return divides_evenly 
}


/* scheduler needs to check that jobs complete by their deadline, this means
that there should be at least one frame boundary between the release time of a 
job and its deadline
*/
func jobsCompleteByDeadline(canidate_frame int, tasks Tasks) bool {

  jobs_complete := true

  for _, task := range tasks.T {
    if (2*canidate_frame - gcd(task.Period, canidate_frame)) > task.Period {
      jobs_complete = false
    }
  }

  return jobs_complete
}

func determineFrameLength(tasks Tasks) int {

  // initalize canidate_frame to be the max of the execution times
  canidate_frame := maxExecutionTime(tasks)

  frame_divides_evenly := dividesEvenly(canidate_frame, tasks)
  jobs_complete := jobsCompleteByDeadline(canidate_frame, tasks)

  /* frame_divides_evenly is true, check if jobsCompleteByDeadline
  if not then try another canidate up until 49 */  
  for !(frame_divides_evenly && jobs_complete) && canidate_frame < 50 {
    canidate_frame = canidate_frame + 1
    frame_divides_evenly = dividesEvenly(canidate_frame, tasks)
    jobs_complete = jobsCompleteByDeadline(canidate_frame, tasks)
  }

  if !(frame_divides_evenly && jobs_complete) {
    // failed to find a valid frame size
    fmt.Println("Failed to find a valid frame length. Try modifying tasks to have different execution/deadline times.")
    os.Exit(0)
  }

  return canidate_frame
}

func lcm(a int, b int) (int, error) {
  var lcm int = 1

  if (a > b) {
    lcm = a
  } else {
    lcm = b
  }

  for {
    if(lcm % a == 0 && lcm % b == 0) {
      return lcm, nil
    }
    lcm++
  }

  // shouldn't reach here since the loop above continues until the lcm is found
  return -1, errors.New("lcm failed")
}


func determineHyperPeriod(tasks Tasks) int {

  // inital val
  hyper_period := tasks.T[0].Period
  var err error = nil

  for _, task := range tasks.T[2:] {
    b := task.Period
    hyper_period, err = lcm(hyper_period,b)
    check(err)
  }

  return hyper_period
}


// helper function to check for errors
func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  
  // tasks to schedule
  tasks := Tasks{}

  // read task data from json
  file, _ := ioutil.ReadFile("tasks.json")
  err := json.Unmarshal([]byte(file), &tasks)
  check(err)

  frame_length := determineFrameLength(tasks)
  check(err)
  hyper_period := determineHyperPeriod(tasks)
 
  fmt.Println(hyper_period)
  fmt.Println(lcm)
  fmt.Println(frame_length)
}
