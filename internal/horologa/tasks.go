package horologa

/*
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
