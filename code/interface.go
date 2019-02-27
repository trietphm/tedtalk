package main

import "fmt"

type Locker interface {
	Lock()
	Unlock()
}

type A struct{}

func (A) Lock()           { fmt.Println("A Lock") }
func (A) Unlock()         { fmt.Println("A Unlock") }
func LockF(locker Locker) { locker.Lock() }

func main() {
	var lock Locker
	lock = A{}
	lock.Lock()
	LockF(lock)
}
