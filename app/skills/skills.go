package skills

import (
	"i-moscow-backend/app/db"
	"strings"
)

func GenerateSkills(text string) (skills []string) {
	words := strings.Split(text, " ")
	for _, v := range words {
		res := db.FullTextSearch(v, 2)
		if len(res) > 0 {
			skills = append(skills, res[0])
		}
	}
	return
}
