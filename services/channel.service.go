package services

func AsyncFunctionWithoutParam[T any](function func() (T, error)) (chan T, chan error) {
	resultChannel := make(chan T)
	errorChannel := make(chan error)
	go func() {
		defer close(resultChannel)
		defer close(errorChannel)

		result, err := function()
		if err != nil {
			errorChannel <- err
		} else {
			resultChannel <- result
		}
	}()

	return resultChannel, errorChannel
}

func AsyncFunctionWith1Param[T any, Arg any](function func(Arg) (T, error), arg Arg) (chan T, chan error) {
	resultChannel := make(chan T)
	errorChannel := make(chan error)
	go func() {
		defer close(resultChannel)
		defer close(errorChannel)

		result, err := function(arg)
		if err != nil {
			errorChannel <- err
		} else {
			resultChannel <- result
		}
	}()

	return resultChannel, errorChannel
}
