package cronline

import (
	"regexp"
	"strings"
)

// SafeSplit split a command into an array
func SafeSplit(s string) []string {
	split := strings.Split(s, " ")

	var result []string
	var inquote string
	var block string
	for _, i := range split {
		if inquote == "" {
			if strings.HasPrefix(i, "'") || strings.HasPrefix(i, "\"") {
				inquote = string(i[0])
				block = strings.TrimPrefix(i, inquote) + " "
			} else {
				result = append(result, i)
			}
		} else {
			if !strings.HasSuffix(i, inquote) {
				block += i + " "
			} else {
				block += strings.TrimSuffix(i, inquote)
				inquote = ""
				result = append(result, block)
				block = ""
			}
		}
	}
	return result
}

// GetCommand return the command of a cron line
func GetCommand(line string) string {
	r, _ := regexp.Compile("@(every|yearly,monthly|weekly|daily|hourly){1} ")
	if r.MatchString(line) {
		r2, _ := regexp.Compile("^@every (\\w*) ")
		if r2.MatchString(line) {
			return GetCommandOfEvery(line)
		} else {
			str := r.FindAllString(line, -1)
			return line[len(str[0]):]
		}
	}
	r, _ = regexp.Compile("^(([0-9]+-[0-9]+|\\*/?[0-9]+|\\*|[0-9]+) )+")
	if len(r.FindAllStringIndex(line, -1)) != 0 {
		res := r.FindAllString(line, -1)
		return line[len(res[0]):]
	}

	return ""
}

// GetCron return the cron of a cron line
func GetCron(line string) string {
	r, _ := regexp.Compile("@(every|yearly,monthly|weekly|daily|hourly){1} ")
	if r.MatchString(line) {
		r, _ := regexp.Compile("^@every (\\w*) ")
		if r.MatchString(line) {
			str := r.FindAllString(line, 1)
			return line[:len(str[0])]
		} else {
			r, _ := regexp.Compile("@(yearly,monthly|weekly|daily|hourly){1} ")
			if r.MatchString(line) {
				str := r.FindAllString(line, 1)
				return line[:len(str[0])]
			}
		}
	}

	r, _ = regexp.Compile("^(([0-9]+-[0-9]+|\\*/?[0-9]+|\\*|[0-9]+) )+")
	if len(r.FindAllStringIndex(line, -1)) != 0 {
		res := r.FindAllString(line, -1)
		return res[0]
	}
	return ""
}

// GetCommandOfEvery return the command if the line is something like @every ...
func GetCommandOfEvery(line string) string {

	r, _ := regexp.Compile("^@every (\\w*) ")
	if r.MatchString(line) {
		str := r.FindAllString(line, 1)
		return line[len(str[0]):]
	}
	return ""
}
