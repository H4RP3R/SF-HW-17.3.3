// Перепишите приведённый выше пример со счётчиком из основного текста,но вместо
// примитивов из пакета atomic используйте условную переменную и попробуйте
// реализовать динамическую проверку достижения конечного значения счётчиком.

package main

import (
	"fmt"
	"sync"
)

const (
	step            int = 1
	endCounterValue int = 1000
)

var (
	cond *sync.Cond
	lock sync.Mutex

	counter int = 0
)

func init() {
	cond = sync.NewCond(&lock)
}

func main() {
	increment := func() {
		cond.L.Lock()
		defer cond.L.Unlock()
		counter += step
		if counter == endCounterValue {
			cond.Signal()
		}
	}

	iterationCount := endCounterValue / step
	for i := 0; i < iterationCount; i++ {
		go increment()
	}

	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
	fmt.Println(counter)
}
