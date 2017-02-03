package utils

func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfFalse(expr bool, msg string) {
	if !expr {
		panic(msg)
	}
}
