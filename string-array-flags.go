package main

import "strings"

type stringArrayFlags []string

func (a *stringArrayFlags) String() string {
	return strings.Join(*a, ",")
}

func (a *stringArrayFlags) Set(value string) error {
	*a = append(*a, value)
	return nil
}
