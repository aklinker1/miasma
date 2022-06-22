package utils

import (
	"encoding/json"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/samber/lo"
)

func ToEnvMap(input map[string]any) internal.EnvMap {
	return lo.Reduce(
		lo.Entries(input),
		func(res internal.EnvMap, e lo.Entry[string, interface{}], i int) internal.EnvMap {
			switch v := e.Value.(type) {
			case string:
				res[e.Key] = v
			case json.Number:
				res[e.Key] = v.String()
			case bool:
				res[e.Key] = fmt.Sprint(v)
			}
			return res
		},
		internal.EnvMap{},
	)
}

func ToAnyMap(input internal.EnvMap) map[string]any {
	return lo.Reduce(
		lo.Entries(input),
		func(res map[string]any, e lo.Entry[string, string], i int) map[string]any {
			res[e.Key] = e.Value
			return res
		},
		map[string]any{},
	)
}
