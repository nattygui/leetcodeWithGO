package subdomain

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domainMap := make(map[string]int)
	for _, domain := range cpdomains {
		domainSplit := strings.Split(domain, " ")
		times, mainDomain := domainSplit[0], domainSplit[1]
		// 1. 最高级domain
		timesInt, err := strconv.Atoi(times)
		if err != nil {
			continue
		}
		if _, ok := domainMap[mainDomain]; ok {
			domainMap[mainDomain] += timesInt
		} else {
			domainMap[mainDomain] += timesInt
		}
		// 2. 分离高级domain
		domains := strings.Split(mainDomain, ".")
		if len(domains) < 2 {
			continue
		} else {
			domains = domains[1:]
			for len(domains) != 0 {
				subDomain := strings.Join(domains, ".")
				if _, ok := domainMap[subDomain]; ok {
					domainMap[subDomain] += timesInt
				} else {
					domainMap[subDomain] += timesInt
				}
				domains = domains[1:]
			}
		}
	}
	// 最后再整合
	domainTemplate := "%d %s"
	res := make([]string, 0)
	for domain, times := range domainMap {
		res = append(res, fmt.Sprintf(domainTemplate, times, domain))
	}
	return res
}
