package singleton

import (
	"fmt"
	"sync"
	"testing"
)

const parCount = 100

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	fmt.Println(instance1)
	if instance1 != instance2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T)  {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instance := [parCount]singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instance[index] = GetInstance()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < parCount ; i++ {
		if instance[i] != instance[i - 1] {
			t.Fatal("instance is not equal")
		}
	}
}
