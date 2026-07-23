package helper

var StringHelper = stringHelper{}

type stringHelper struct {
}

func (t *stringHelper) SubString(source string, start int, end int) string {
	var substring = ""
	var pos = 0
	for _, c := range source {
		if pos < start {
			pos++
			continue
		}
		if pos >= end {
			break
		}
		pos++
		substring += string(c)
	}

	return substring
}
