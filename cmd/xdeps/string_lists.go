package main

import "sort"

type stringLists map[string][]string

func (d stringLists) Add(k, v string) {
	d[k] = append(d[k], v)
}

func (d stringLists) Keys() []string {
	keys := make([]string, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
