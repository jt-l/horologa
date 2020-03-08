package horologa

import (
	"errors"
	"fmt"
)

var Errors = NewHorologaErrors()

// custom error type
type HorologaError struct {
	Code int
	Err  error
}

// implement the error interface
func (r *HorologaError) Error() string {
	return fmt.Sprintf("code: %d: err %v", r.Code, r.Err)
}

// create a new application error
func NewHorologaError(err error, code int) *HorologaError {
	return &HorologaError{Err: err, Code: code}
}

// struct to hold all error types
type HorologaErrors struct {
	FailedToParseTasks           *HorologaError
	FailedToDetermineHyperPeriod *HorologaError
	FailedToDetermineFrameLength *HorologaError
}

// create an instance of horologa errors
func NewHorologaErrors() *HorologaErrors {
	failed_to_parse_tasks := &HorologaError{1, errors.New("failed to parse tasks.")}
	failed_to_determine_hyper_period := &HorologaError{2, errors.New("failed to determine hyper period.")}
	failed_to_determine_frame_length := &HorologaError{3, errors.New("failed to determine frame length.")}

	return &HorologaErrors{
		FailedToParseTasks:           failed_to_parse_tasks,
		FailedToDetermineHyperPeriod: failed_to_determine_hyper_period,
		FailedToDetermineFrameLength: failed_to_determine_frame_length,
	}
}
