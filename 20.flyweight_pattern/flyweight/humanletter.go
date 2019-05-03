package flyweight

import "fmt"

var Factory humanLetterFactory = humanLetterFactory{make(map[string]HumanLetter)}

type humanLetterFactory struct {
	letterMap map[string]HumanLetter
}

func (f humanLetterFactory) GetHumanLetter(letter string) HumanLetter {
	humanLetter, ok := f.letterMap[letter]
	if !ok {
		fmt.Println("generate letter: ", letter)
		humanLetter = HumanLetter(letter)
		f.letterMap[letter] = humanLetter
	}
	return humanLetter
}

type HumanLetter string

func (l HumanLetter) GetLetter() string {
	return string(l)
}
