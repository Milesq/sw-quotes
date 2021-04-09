package utils

// Expect .
func Expect(err error) {
	if err != nil {
		panic(err)
	}
}
