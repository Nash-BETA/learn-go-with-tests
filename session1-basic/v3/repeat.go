package iteration

const repeatCount = 5

func Repeat(character string) (repeatText string) {
	for i := 0; i <= repeatCount; i++ {
		repeatText += character
	}
	return repeatText
}
