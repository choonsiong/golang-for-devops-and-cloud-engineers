package main

type RequestError struct {
	HTTPCode int
	Body     string
	Err      string
}

func (e RequestError) Error() string {
	return e.Err
}
