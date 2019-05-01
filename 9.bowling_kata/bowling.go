package __bowling_kata

import (
	"fmt"
	"strconv"
)

func Score(rolls string) int {
	var score int = 0
	scores := convertScoreArray(rolls)
	frame := 1
	throw := 0
	fmt.Println(scores)

	for i, roll := range scores {
		if frame < 10 {
			switch string(rolls[i]) {
			case "X":
				score += int(roll) + scores[i+1] + scores[i+2]
				frame += 1
			case "/":
				score += int(roll) + scores[i+1]
				frame += 1
				throw = 0
			default:
				score += int(roll)
				if throw += 1; throw >= 2 {
					frame += 1
					throw = 0
				}
			}
		} else {
			score += int(roll)
		}
	}

	return score
}

func convertScoreArray(rolls string) []int {
	scores := make([]int, 0)
	for i, r := range rolls {
		roll := string(r)
		switch roll {
		case "-":
			scores = append(scores, 0)
		case "X":
			scores = append(scores, 10)
		case "/":
			scores = append(scores, 10-scores[i-1])
		default:
			scores = append(scores, StringToInt(roll))
		}
	}
	return scores
}

func StringToInt(val string) int {
	parsed, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return parsed
}
