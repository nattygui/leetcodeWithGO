package uniqueemail

import "strings"

func numUniqueEmails(emails []string) int {
	emailMap := make(map[string]bool)
	for _, email := range emails {
		strs := strings.Split(email, "@")
		if len(strs) != 2 {
			continue
		}
		firstName, domainName := strs[0], strs[1]
		index := strings.IndexByte(firstName, '+')
		if index != -1 {
			firstName = firstName[:index]
		}
		for {
			index1 := strings.IndexByte(firstName, '.')
			if index1 == -1 {
				break
			}
			firstName = firstName[0:index1] + firstName[index1+1:]
		}
		trimEmail := firstName + "@" + domainName
		if !emailMap[trimEmail] {
			emailMap[trimEmail] = true
		}
	}
	return len(emailMap)
}
