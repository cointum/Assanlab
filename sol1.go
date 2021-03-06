package main

import (
	"fmt"
	"strconv"
	"strings"
)

var cpdomains = [...]string{"900 google.com", "60 mail.yahoo.com", "10 mobile.sports.yahoo.com",
	"40 sports.yahoo.com", "300 yahoo.com", "10 stackoverflow.com", "20 overflow.com",
	"5 com.com", "2 en.wikipedia.org", "1 m.wikipedia.org", "1 mobile.sports", "1 google.co.uk"}

func subdomainVisits(cpdomains []string) []string {
	result := make([]string, 0)
	if len(cpdomains) == 0 {
		return result
	}
	domainCountMap := make(map[string]int, 0)
	for _, domain := range cpdomains {
		countDomain := strings.Split(domain, " ")
		allDomains := strings.Split(countDomain[1], ".")
		temp := make([]string, 0)
		for i := len(allDomains) - 1; i >= 0; i-- {
			temp = append([]string{allDomains[i]}, temp...)
			ld := strings.Join(temp, ".")
			count, _ := strconv.Atoi(countDomain[0])
			if val, ok := domainCountMap[ld]; !ok {
				domainCountMap[ld] = count
			} else {
				domainCountMap[ld] = count + val
			}
		}
	}
	for k, v := range domainCountMap {
		t := strings.Join([]string{strconv.Itoa(v), k}, " ")
		result = append(result, t)
	}
	return result
}

func subdomainVisits1(cpdomains []string) []string {
	out := make([]string, 0)
	var b strings.Builder
	domains := make(map[string]int, 0)
	for _, v := range cpdomains {
		splitDomain(v, domains)
	}
	for k, v := range domains {
		b.WriteString(strconv.Itoa(v))
		b.WriteString(" ")
		b.WriteString(k)
		out = append(out, b.String())
		b.Reset()
	}
	return out
}

func splitDomain(domain string, domains map[string]int) {
	visits := 0
	var e error
	subdomains := make([]string, 0)
	for i, v := range domain {
		if v == ' ' {
			visits, e = strconv.Atoi(domain[0:i])
			if e != nil {
				panic(e)
			}
			break
		}
	}
	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			subdomains = append(subdomains, domain[i+1:])
		} else if domain[i] == ' ' {
			subdomains = append(subdomains, domain[i+1:])
			break
		}
	}
	for _, v := range subdomains {
		count, ok := domains[v]
		if ok {
			domains[v] = count + visits
		} else {
			domains[v] = visits
		}
	}
}
func main() {
	fmt.Println("\n", "Total no of  Count + Subdomain  = ", subdomainVisits(cpdomains[:]))

}
