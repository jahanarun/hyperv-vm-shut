package common

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func MapToStruct(m map[string]string, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}

func ResponseCleanup(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	return s
}

func GroupMap(s string, expr string) map[string]string{
    r := regexp.MustCompile(expr)
    values := r.FindStringSubmatch(s)
    keys := r.SubexpNames()

    // create map
    d := make(map[string]string)
    for i := 1; i < len(keys); i++ {
        d[keys[i]] = values[i]
    }

	return d

// mi := make(map[string]interface{}, len(d))
// 	for k, v := range d {
// 		mi[k] = v
// 	}

//     return mi
}

func GetValue(s string, key string) string{
	cleanedResult := ResponseCleanup(s)
	expr := fmt.Sprintf(`%s:(?P<%s>.*)`, key, key)
	r := regexp.MustCompile(expr)
    values := r.FindStringSubmatch(cleanedResult)
    // keys := r.SubexpNames()
	
	return strings.TrimSpace(values[1])
}