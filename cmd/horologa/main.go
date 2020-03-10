package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"

  "github.com/jt-l/horologa/internal/horologa"
)


func main() {

  // tasks to schedule
  tasks := horologa.Tasks{}

  // parse tasks file
  file, _ := ioutil.ReadFile("tasks.json")
  err := json.Unmarshal([]byte(file), &tasks)

  if err != nil {
    fmt.Println(err)
    os.Exit(horologa.Errors.FailedToParseTasks.Code)
  }

  // determine frame length
  frame_length, err := horologa.DetermineFrameLength(tasks)
  if err != nil {
    fmt.Println(err)
    os.Exit(horologa.Errors.FailedToDetermineFrameLength.Code)
  }

  hyper_period, err := horologa.DetermineHyperPeriod(tasks)
  if err != nil {
    fmt.Println(err)
    os.Exit(horologa.Errors.FailedToDetermineHyperPeriod.Code)
  }

  utilzation := horologa.DetermineUtilization(tasks)

  fmt.Println("Hyper Period:", hyper_period)
  fmt.Println("Frame Length:", frame_length)
  fmt.Println("Utilization:", utilzation)
}
