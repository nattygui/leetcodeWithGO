package _482LicenseKeyFormatting

func licenseKeyFormatting(S string, K int) string {
	res := make([]byte, 0)

	for i, v := range S {
		if v != '-' {
			if v >= 'a' && v <= 'z' {
				res = append(res, S[i]-32)
			} else {
				res = append(res, S[i])
			}
		}
	}

	temp := ""
	l := len(res)
	start := 0
	end := 0
	for l == 0 {
		end = l % 4 + start
		temp += "-"
		l -= end - start
		start = end
	}

	return temp
}
