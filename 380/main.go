package main

import "math/rand"

type Void struct{}

type RandomizedSet map[int]Void

func Constructor() RandomizedSet {
	return RandomizedSet{}
}

func (this RandomizedSet) Insert(val int) bool {
	if _, ok := this[val]; ok {
		return false
	}
	this[val] = Void{}
	return true
}

func (this RandomizedSet) Remove(val int) bool {
	if _, ok := this[val]; !ok {
		return false
	}
	delete(this, val)
	return true
}

func (this RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(len(this))
	i := 0
	for k := range this {
		if i == randIdx {
			return k
		}
		i++
	}
	return -1
}
