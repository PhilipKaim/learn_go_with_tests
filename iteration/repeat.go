package iteration

func Repeat(character string, cycles int) string {
	var repeated string
	for i := 0; i < cycles; i++ {
		repeated += character
	}
	return repeated
}
