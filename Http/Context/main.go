package main

import (
	"ctx/examples"
	"ctx/management"
)

/*
Context
Creation
Send value
Cancellation from request
Internal cancellation
timeout
Deadline

*/
func main() {
	// Creation
	examples.SendValueExample()

	//External Cancellation
	management.Run()

	// Internal Cancellation
	examples.InternalCancellationExample()

}
