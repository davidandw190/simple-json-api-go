package main

type apiError struct {
	Err    string
	Status int
}

// returns the error message associated with the apiError instance.
func (e apiError) Error() string {
	return e.Err
}
