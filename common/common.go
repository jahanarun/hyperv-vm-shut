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
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	return s
}

func GroupMap(s string, expr string) map[string]string{
    r := regexp.MustCompile(expr)
    values := r.FindStringSubmatch(s)
    keys := r.SubexpNames()

	fmt.Println("keys", keys)
	fmt.Println("values", values)
	fmt.Println("values len", len(values))
	fmt.Println("keys[0]", keys[0])
	fmt.Println("keys[1]", keys[1])
	fmt.Println("keys[2]", keys[2])
	fmt.Println("values[0]", values[0])
	fmt.Println("values[1]", values[1])
	fmt.Println("values[2]", values[2])
	fmt.Println("*********************")

    // create map
    d := make(map[string]string)
    for i := 1; i < len(keys); i++ {
        d[keys[i]] = values[i]
    }
	fmt.Println(d)
	fmt.Println("*********************")

	return d

// mi := make(map[string]interface{}, len(d))
// 	for k, v := range d {
// 		mi[k] = v
// 	}

//     return mi
}

func GetValue(s string, expr string) string{
    r := regexp.MustCompile(expr)
    values := r.FindStringSubmatch(s)
    keys := r.SubexpNames()

	fmt.Println("keys", keys)
	fmt.Println("values", values)
	fmt.Println("values len", len(values))
	fmt.Println("keys[0]", keys[0])
	fmt.Println("keys[1]", keys[1])
	fmt.Println("values[0]", values[0])
	fmt.Println("values[1]", values[1])
	fmt.Println("*********************")

	return values[1]
}