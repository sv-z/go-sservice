package errors

type StoreError struct {
}

type DuplicateRecord struct {
	Err error
	StoreError
}

func (e *StoreError) Error() string {
	return "An infrastructure error occurred."
}
