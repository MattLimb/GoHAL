package gohal

import "fmt"

type ErrHalConfused string

func (e ErrHalConfused) Error() string {
	return  "I'm sorry Dave, I'm afraid I can't do that."
}

func (e ErrHalConfused) DetailedError() string {
	return fmt.Sprintf("I'm sorry Dave, I'm afraid I can't do that. (ERROR - %s)", e)
}

// MUST END Errors
type ErrHalReallyConfused string

func (e ErrHalReallyConfused) Error() string {
	return  "I'm sorry Dave, I'm afraid I can't do that."
}

func (e ErrHalReallyConfused) DetailedError() string {
	return fmt.Sprintf("I'm sorry Dave, I'm afraid I can't do that. (ERROR - %s)", e)
}