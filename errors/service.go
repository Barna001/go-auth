package errors

// CriticalHandling checks for error and panics if present
func CriticalHandling(err error) {
	if err != nil {
		panic(err)
	}
}
