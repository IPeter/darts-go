package helper

import (
	"reflect"
	"sort"
)

func GetMapPosition(mp interface{}, ky string) (pos int, prev string, next string) {
	k := reflect.ValueOf(mp)
	if k.Kind() != reflect.Map {
		panic("Invalid MAP argument!")
	}
	t := k.Type()
	if t.Key().Kind() != reflect.String {
		panic("Invalid key type!")
	}
	var keys []string
	for _, kv := range k.MapKeys() {
		keys = append(keys, kv.String())
	}

	sort.Strings(keys)

	for i, k := range keys {
		if ky == k {
			pos = i

			if (i - 1) < 0 {
				prev = keys[(len(keys) - 1)]
			} else {
				prev = keys[i-1]
			}

			if (i + 1) >= len(keys) {
				next = keys[0]
			} else {
				next = keys[i+1]
			}

			break
		}
	}
	return
}
