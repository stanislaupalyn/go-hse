package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func removeDuplicates(a []int) (b []int) {
	sorted := a
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	for i := 0; i < len(sorted); i++ {
		if i == 0 || sorted[i-1] != sorted[i] {
			b = append(b, sorted[i])
		}
	}
	return
}

func bubbleSort(a []int) (b []int) {
	b = a
	for i := 0; i < len(b); i++ {
		for j := 0; j+1 < len(b); j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
	return
}

func fibonacciSequence(n int) (a []int) {
	if n > 0 {
		a = append(a, 1)
	}
	if n > 1 {
		a = append(a, 1)
	}
	for i := 2; i <= n; i++ {
		a = append(a, a[i-1]+a[i-2])
	}
	return
}

func countOccurences(a []int, x int) int {
	ans := 0
	for _, val := range a {
		if val == x {
			ans++
		}
	}
	return ans
}

func arrayIntersection(a []int, b []int) (c []int) {
	var haveInA map[int]bool
	haveInA = make(map[int]bool)
	for _, val := range a {
		haveInA[val] = true
	}

	for _, val := range b {
		_, ok := haveInA[val]
		if ok {
			c = append(c, val)
			haveInA[val] = false
		}
	}
	return c
}

func areStringsAnagrams(a string, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	ra := []rune(a)
	rb := []rune(b)
	sort.Slice(ra, func(i int, j int) bool {
		return ra[i] < ra[j]
	})
	sort.Slice(rb, func(i int, j int) bool {
		return rb[i] < rb[j]
	})

	return reflect.DeepEqual(ra, rb)
}

func mergeSorted(a []int, b []int) (result []int) {
	ptr1 := 0
	ptr2 := 0
	for ptr1 < len(a) || ptr2 < len(b) {
		if ptr1 < len(a) && (ptr2 == len(b) || b[ptr2] > a[ptr1]) {
			result = append(result, a[ptr1])
			ptr1++
		} else {
			result = append(result, b[ptr2])
			ptr2++
		}
	}
	return result
}

type HashTable struct {
	length  int
	buckets [][]entry
}

const initialSize = 512

type entry struct {
	key   []byte
	value interface{}
}

func NewHashTableWithSize(size int) *HashTable {
	return &HashTable{
		buckets: make([][]entry, size),
	}
}

func NewHashTable() *HashTable {
	return NewHashTableWithSize(initialSize)
}

// https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
func hashValue(str string, limit int) int {
	fnvOffSetBasis := uint64(14695981039346656037)
	fnvPrime := uint64(1099511628211)
	hash := fnvOffSetBasis
	bytesSlice := []byte(str)

	for _, b := range bytesSlice {
		hash = hash ^ uint64(b)
		hash = hash * fnvPrime
	}
	return int(hash % uint64(limit))
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	bytesSlice := []byte(key)
	hash := hashValue(key, len(ht.buckets))
	for _, v := range ht.buckets[hash] {
		if bytes.Equal(v.key, bytesSlice) {
			return v.value, true
		}
	}
	return nil, false
}

func (ht *HashTable) Set(key string, value interface{}) {
	bytesSlice := []byte(key)
	hash := hashValue(key, len(ht.buckets))
	pos := -1

	for idx, v := range ht.buckets[hash] {
		if bytes.Equal(v.key, bytesSlice) {
			pos = idx
			break
		}
	}
	if pos == -1 {
		ht.buckets[hash] = append(ht.buckets[hash], entry{bytesSlice, value})
	} else {
		ht.buckets[hash][pos].value = value
	}
}

func (ht *HashTable) Remove(key string) error {
	bytesSlice := []byte(key)
	hash := hashValue(key, len(ht.buckets))
	pos := -1

	for idx, v := range ht.buckets[hash] {
		if bytes.Equal(v.key, bytesSlice) {
			pos = idx
			break
		}
	}
	if pos == -1 {
		return fmt.Errorf("Key is not present")
	} else {
		bucket := ht.buckets[hash]
		bucket[pos], bucket[len(bucket)-1] = bucket[len(bucket)-1], bucket[pos]
		ht.buckets[hash] = bucket[:len(bucket)-1]

		return nil
	}
}

func BinarySearch(a []int, x int) int {
	l := 0
	r := len(a) - 1

	for l < r {
		mid := l + (r-l)/2

		if a[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if l < len(a) && a[l] == x {
		return l
	} else {
		return -1
	}
}

type Queue struct {
	st1, st2 []interface{}
}

func NewQueue() *Queue {
	return &Queue{make([]interface{}, 0), make([]interface{}, 0)}
}

func (q *Queue) Push(x interface{}) {
	q.st1 = append(q.st1, x)
}

func (q *Queue) GetLength() int {
	return len(q.st1) + len(q.st2)
}

func (q *Queue) IsEmpty() bool {
	return q.GetLength() == 0
}

func (q *Queue) Pop() error {
	if q.IsEmpty() {
		return errors.New("popping from empty queue")
	}

	if len(q.st2) == 0 {
		for len(q.st1) > 0 {
			q.st2 = append(q.st2, q.st1[len(q.st1)-1])
			q.st1 = q.st1[0 : len(q.st1)-1]
		}
	}
	q.st2 = q.st2[0 : len(q.st2)-1]

	return nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errors.New("peeking in empty queue")
	}
	var val interface{}
	if len(q.st2) > 0 {
		val = q.st2[len(q.st2)-1]
	} else {
		val = q.st1[0]
	}

	return val, nil
}

func main() {
	fmt.Println("Array before removing duplicates [2, 1, 2, 1, 5, 1, 4, 0, 5]")
	fmt.Println("After:", removeDuplicates([]int{2, 1, 2, 1, 5, 1, 4, 0, 5}))

	fmt.Println("Array [4, 1, 3, 12, 10] after bubbleSort():", bubbleSort([]int{4, 1, 3, 12, 10}))

	fmt.Println("First 10 fibonacci numbers are ", fibonacciSequence(10))

	fmt.Println("Number 4 has", countOccurences([]int{4, 1, 3, 4, 4, 5, 4, 4}, 4), "occurences in the array [4, 1, 3, 4, 4, 5, 4, 4]")

	fmt.Println("Array intersection of [4, 1, 10, 3, 5] and [10, 3, 8, 9, 5] is", arrayIntersection([]int{4, 1, 10, 3, 5}, []int{10, 3, 8, 9, 5}))

	fmt.Println("Are \"abbaBc\" and \"BAbcAB\" anagrams:", areStringsAnagrams("abbaBc", "BAbcAb"))

	fmt.Println("Result of merging arrays [1, 3, 4, 12] and [2, 4, 10, 12, 12] is", mergeSorted([]int{1, 3, 4, 12}, []int{2, 4, 10, 12, 12}))

	ht := NewHashTable()
	ht.Set("aba", 3)
	ht.Set("abc", 11)
	ht.Set("x", 4)

	fmt.Println(ht.Get("aba"))
	fmt.Println(ht.Get("abc"))
	fmt.Println(ht.Get("aboba"))

	ht.Set("abc", 12)

	fmt.Println(ht.Get("abc"))

	ht.Remove("abc")
	fmt.Println(ht.Get("abc"))

	fmt.Println("Binary search of 3 on array [-2, -1, 0, 1, 1, 2, 3, 3, 3, 5]:", BinarySearch([]int{-2, -1, 0, 1, 1, 2, 3, 3, 3, 5}, 3))
	fmt.Println("Binary search of 9 on array [-2, -1, 0, 1, 1, 2, 3, 3, 3, 5]:", BinarySearch([]int{-2, -1, 0, 1, 1, 2, 3, 3, 3, 5}, 9))
	fmt.Println("Binary search of 4 on array [-2, -1, 0, 1, 1, 2, 3, 3, 3, 5]:", BinarySearch([]int{-2, -1, 0, 1, 1, 2, 3, 3, 3, 5}, 4))

	var q = NewQueue()
	q.Push(12)
	q.Push(17)
	q.Push(3)
	fmt.Println(q.Peek())
	q.Pop()
	fmt.Println(q.Peek())
	q.Pop()
	fmt.Println(q.Peek())
	q.Pop()
	q.Push(13)
	fmt.Println(q.Peek())
}
