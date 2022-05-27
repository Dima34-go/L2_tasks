package codeUnderstanding


type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test1() *customError {
	{
		// do something
	}
	return nil
}

func MainList5() {
	var err error
	err = test1()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
