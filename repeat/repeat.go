package iteration

// Repeat iterates the giving character for giving nums times
func Repeat(character string, nums int) string {
	var str string
	for i := 0; i < nums; i++ {
		str += character
	}
	return str
}
