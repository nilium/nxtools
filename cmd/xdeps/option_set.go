package main

import "strings"

type optionSet map[string]struct{}

var Options = optionSet{}

func newStringSet(vals ...string) optionSet {
	o := optionSet{}
	for _, v := range vals {
		o.Add(v)
	}
	return o
}

func (s optionSet) Has(k string) bool {
	_, ok := s[k]
	return ok
}

func (s optionSet) Add(k string) {
	s[k] = struct{}{}
}

func (s optionSet) Remove(k string) {
	delete(s, k)
}

func (s optionSet) String() string {
	return ""
}

func (s optionSet) Set(k string) error {
	for _, opt := range strings.Split(k, ",") {
		if strings.HasPrefix(opt, "~") {
			s.Remove(opt[1:])
		} else {
			s.Add(opt)
		}
	}
	return nil
}
