package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

/* A periodic task is defined to have:
phase: release time of the task
period: task needs to be scheduled every period and within the frame
execution_time: cycles that the task takes to execute
relative_deadline: time from the phase in which in the task must complete
*/
type Task struct {
	Phase            int `json:"phase"`
	Period           int `json:"period"`
	ExecutionTime    int `json:"executionTime"`
	RelativeDeadline int `json:"relativeDeadline"`
}

type Tasks struct {
	T []Task `json:"tasks"`
}

var Errors = newHorologaErrors()

// custom error type 
type HorologaError struct {
  code int
  err error
}

// implement the error interface
func (r *HorologaError) Error() string {
  return fmt.Sprintf("code: %d: err %v", r.code, r.err)
}

// create a new application error
func NewHorologaError(err error, code int) *HorologaError {
  return &HorologaError{err: err, code: code}
}

// struct to hold all error types
type HorologaErrors struct {
  FailedToParseTasks *HorologaError
  FailedToDetermineHyperPeriod *HorologaError
  FailedToDetermineFrameLength *HorologaError
}

// create an instance of horologa errors
func newHorologaErrors() *HorologaErrors {
  failed_to_parse_tasks := &HorologaError{1, errors.New("failed to parse tasks.")}
  failed_to_determine_hyper_period := &HorologaError{2, errors.New("failed to determine hyper period.")}
  failed_to_determine_frame_length := &HorologaError{3, errors.New("failed to determine frame length.")}

  return &HorologaErrors{
    FailedToParseTasks: failed_to_parse_tasks, 
    FailedToDetermineHyperPeriod: failed_to_determine_hyper_period, 
    FailedToDetermineFrameLength: failed_to_determine_frame_length,
  }
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
		if task.Period%canidate_frame == 0 {
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

func determineFrameLength(tasks Tasks) (int, error) {

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
    return -1, Errors.FailedToDetermineFrameLength
	}

	return canidate_frame, nil
}

func lcm(a int, b int) (int, error) {
	var lcm int = 1

	if a > b {
		lcm = a
	} else {
		lcm = b
	}

	for {
		if lcm%a == 0 && lcm%b == 0 {
			return lcm, nil
		}
		lcm++
	}

	// shouldn't reach here since the loop above continues until the lcm is found
	return -1, errors.New("lcm failed")
}

func determineHyperPeriod(tasks Tasks) (int, error) {

	// inital val
	hyper_period := tasks.T[0].Period
	var err error = nil

	for _, task := range tasks.T[2:] {
		b := task.Period
		hyper_period, err = lcm(hyper_period, b)

    if err != nil {
      return -1, Errors.FailedToDetermineHyperPeriod
    }
	}

	return hyper_period, nil
}

func main() {

	// tasks to schedule
	tasks := Tasks{}

  // parse tasks file
	file, _ := ioutil.ReadFile("tasks.json")
	err := json.Unmarshal([]byte(file), &tasks)

  if err != nil {
    fmt.Println(err)
    os.Exit(Errors.FailedToParseTasks.code)
  }

  // determine frame length
	frame_length, err := determineFrameLength(tasks)
  if err != nil {
    fmt.Println(err)
    os.Exit(Errors.FailedToDetermineFrameLength.code)
  }

  // determine hyper period
	hyper_period, err := determineHyperPeriod(tasks)
  if err != nil {
    fmt.Println(err)
    os.Exit(Errors.FailedToDetermineHyperPeriod.code)
  }

  fmt.Println("Hyper Period:", hyper_period)
  fmt.Println("Frame Length:",frame_length)
}
