// Copyright 2014 Codehack. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strarr

import (
	"math/rand"
	"strings"
	"time"
)

func streq(s1, s2 string) bool { return s1 == s2 }

func indexFunc(a []string, s string, f func(string, string) bool) int {
	for k, v := range a {
		if f(v, s) {
			return k
		}
	}
	return -1
}

// Index returns the index of the first instance of 's' in 'a', or -1 if not found
func Index(a []string, s string) int {
	return indexFunc(a, s, streq)
}

// IndexPrefix returns the index of the first entry in 'a' with prefix 'prefix', or -1 if not found
func IndexPrefix(a []string, prefix string) int {
	return indexFunc(a, prefix, strings.HasPrefix)
}

// IndexSuffix returns the index of the first entry in 'a' with suffix 'suffix', or -1 if not found
func IndexSuffix(a []string, suffix string) int {
	return indexFunc(a, suffix, strings.HasSuffix)
}

// Search returns the index of the first entry containing the substring 's' in 'a', or -1 if not found
func Search(a []string, s string) int {
	return indexFunc(a, s, strings.Contains)
}

// Contains returns true if 's' is in 'a', false otherwise
func Contains(a []string, s string) bool {
	return Index(a, s) >= 0
}

// ContainsPrefix returns true if any entry in 'a' has prefix 'prefix', false otherwise
func ContainsPrefix(a []string, prefix string) bool {
	return IndexPrefix(a, prefix) >= 0
}

// ContainsSuffix returns true if any entry in 'a' has suffix 'suffix', false otherwise
func ContainsSuffix(a []string, suffix string) bool {
	return IndexSuffix(a, suffix) >= 0
}

// Count returns the number of occurrances of 's' in 'a'
func Count(a []string, s string) int {
	n := 0
	for _, v := range a {
		if s == v {
			n++
		}
	}
	return n
}

func lastIndexFunc(a []string, s string, f func(string, string) bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if f(a[i], s) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last instance of 's' in 'a', or -1 if not found
func LastIndex(a []string, s string) int {
	return lastIndexFunc(a, s, streq)
}

// LastIndexPrefix returns the index of the last entry in 'a' with prefix 'prefix', or -1 if not found
func LastIndexPrefix(a []string, prefix string) int {
	return lastIndexFunc(a, prefix, strings.HasPrefix)
}

// LastIndexSuffix returns the index of the last entry in 'a' with suffix 'suffix', or -1 if not found
func LastIndexSuffix(a []string, suffix string) int {
	return lastIndexFunc(a, suffix, strings.HasSuffix)
}

// LastSearch returns the index of the last entry containing the substring 's' in 'a', or -1 if not found
func LastSearch(a []string, s string) int {
	return lastIndexFunc(a, s, strings.Contains)
}

// Map returns an array copy of 'a' with the function 'mapping' applied to each element
func Map(mapping func(string) string, a []string) []string {
	maxlen := len(a)
	var b []string
	for k, v := range a {
		s := mapping(v)
		if b == nil {
			b = make([]string, maxlen)
			copy(b, a)
		}
		if s == v {
			continue
		}
		b[k] = s
	}
	if b == nil {
		return a // nothing changed
	}
	return b
}

// ToUpper returns an array with all entries of 'a' changed to upper case
func ToUpper(a []string) []string {
	return Map(strings.ToUpper, a)
}

// ToLower returns an array with all entries of 'a' changed to lower case
func ToLower(a []string) []string {
	return Map(strings.ToLower, a)
}

// ToTitle returns an array with all entries of 'a' changed to title case
func ToTitle(a []string) []string {
	return Map(strings.Title, a)
}

func trimFunc(a []string, s string, f func(string, string) bool) []string {
	b := make([]string, 0)
	for _, v := range a {
		if f(v, s) {
			continue
		}
		b = append(b, v)
	}
	return b
}

// Trim returns an array with all the entries of 'a' that don't match string 's'
func Trim(a []string, s string) []string {
	return trimFunc(a, s, streq)
}

// TrimFunc returns an array with all the entries of 'a' that don't match string 's' using a callback function 'f'
// Callback is f(value, key string) where value is in 'a' and is checked for key. If true, value will be trimmed.
func TrimFunc(a []string, s string, f func(string, string) bool) []string {
	return trimFunc(a, s, f)
}

// TrimPrefix returns an array with all the entries of 'a' that don't have prefix 'prefix'
func TrimPrefix(a []string, prefix string) []string {
	return trimFunc(a, prefix, strings.HasPrefix)
}

// TrimSuffix returns an array with all the entries of 'a' that don't have suffix 'suffix'
func TrimSuffix(a []string, suffix string) []string {
	return trimFunc(a, suffix, strings.HasSuffix)
}

func filterFunc(a []string, s string, f func(string, string) bool) []string {
	g := func(v string, s string) bool { return !f(v, s) }
	return trimFunc(a, s, g)
}

// Filter returns an array with all the entries of 'a' that match string 's'
func Filter(a []string, s string) []string {
	return filterFunc(a, s, streq)
}

// FilterFunc returns an array with all the entries of 'a' that match string 's' using a callback function 'f'
// Callback is f(value, key string) where value is in 'a' and is checked for key. If true, value will be filtered.
func FilterFunc(a []string, s string, f func(string, string) bool) []string {
	return filterFunc(a, s, f)
}

// FilterPrefix returns an array with all the entries of 'a' that have prefix 'prefix'
func FilterPrefix(a []string, prefix string) []string {
	return filterFunc(a, prefix, strings.HasPrefix)
}

// FilterSuffix returns an array with all the entries of 'a' that have suffix 'suffix'
func FilterSuffix(a []string, suffix string) []string {
	return filterFunc(a, suffix, strings.HasSuffix)
}

func diffFunc(a, b []string, f func(a []string, s string) bool) []string {
	c := make([]string, 0)
	for _, v := range a {
		if f(b, v) {
			c = append(c, v)
		}
	}
	return c
}

// Diff returns an array with all the entries of 'a' that are not found in 'b'
func Diff(a, b []string) []string {
	f := func(a []string, s string) bool { return !Contains(a, s) }
	return diffFunc(a, b, f)
}

// Intersect returns an array with all the entries of 'a' that are found in 'b'
func Intersect(a, b []string) []string {
	f := func(a []string, s string) bool { return Contains(a, s) }
	return diffFunc(a, b, f)
}

// Repeat returns a new array consisting of 'count' copies of 's'
func Repeat(s string, count int) []string {
	a := make([]string, 0)
	for i := 0; i < count; i++ {
		a = append(a, s)
	}
	return a
}

// Fill is an alias of Repeat (for PHP converts)
func Fill(n int, s string) []string {
	return Repeat(s, n)
}

// Replace returns an array with the values of 'a' replaced with the index-matching
// values of 'b'. If 'b' has more entries than 'a' they will be appended.
func Replace(a, b []string) []string {
	m, n := len(a), len(b)
	c := make([]string, m)
	copy(c, a)
	for i := 0; i < m; i++ {
		if i >= n {
			break
		}
		c[i] = b[i]
	}
	for i := m; i < n; i++ {
		c = append(c, b[i])
	}
	return c
}

// Rand returns a slice with 'n' number of random entries of 'a'
func Rand(a []string, n int) []string {
	b := make([]string, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, m := 0, len(a); i < n; i++ {
		b = append(b, a[r.Intn(m)])
	}
	return b
}

// Shuffle returns a slice with randomized order of elements in 'a'
func Shuffle(a []string) []string {
	n := len(a)
	b := make([]string, n)
	p := rand.New(rand.NewSource(time.Now().UnixNano())).Perm(n)
	for i := 0; i < n; i++ {
		b[i] = a[p[i]]
	}
	return b
}

// Reverse returns an array copy of 'a' in reverse index order
func Reverse(a []string) []string {
	b := make([]string, 0)
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return b
}

// Shift shifts the first element of '*a' and returns it, shortening the array by one.
// If '*a' is empty returns empty string ""
func Shift(a *[]string) string {
	if m := len(*a); m > 0 {
		s := (*a)[0]
		if m > 1 {
			*a = (*a)[1:]
		} else {
			*a = []string{}
		}
		return s
	}
	return ""
}

// Unshift prepends one or more elements to the beginning of '*a' and returns the number of entries
func Unshift(a *[]string, s ...string) int {
	if s != nil {
		b := Reverse(*a)
		for _, v := range s {
			b = append(b, v)
		}
		*a = Reverse(b)
	}
	return len(*a)
}

// Pop removes the last element in '*a' and returns it, shortening the array by one.
// If '*a' is empty returns empty string ""
func Pop(a *[]string) string {
	if m := len(*a); m > 0 {
		s := (*a)[m-1]
		*a = (*a)[0 : m-1]
		return s
	}
	return ""
}

// Push prepends one or more elements to the end of '*a' and returns the number of entries
func Push(a *[]string, s ...string) int {
	if s != nil {
		for _, v := range s {
			*a = append(*a, v)
		}
	}
	return len(*a)
}
