package skills

import (
	"i-moscow-backend/app/db"
	"strings"
)

func GenerateSkills(text string) (skills []string) {
	words := strings.Split(text, " ")
	for _, v := range words {
		res := db.FullTextSearch(v, 2)
		var isRepeat bool
		if len(res) > 0 {
			for _, v2 := range skills {
				if v2 == res[0] {
					isRepeat = true
					break
				}
			}
			if !isRepeat {
				skills = append(skills, res[0])
			}
		}
	}
	return
}
