package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidN            = errors.New("n should be positive")
)

type Task func() error

// Run start tasks in n gorutines and stop its work when receiveing M errors from task
// M <= 0 means no limits failed tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrInvalidN
	}

	if m <= 0 {
		m = len(tasks) + 1
	}
	var errCnt int32
	taskChen := make(chan Task)

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				fmt.Println("Завершена гоурутина")
			}()
			defer wg.Done()
			fmt.Println("Запущена гоурутина")
			for task := range taskChen {
				if err := task(); err != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		fmt.Println("Запущена задача в канал")
		if atomic.LoadInt32(&errCnt) >= int32(m) {
			break
		}
		taskChen <- task
	}
	close(taskChen)
	wg.Wait()
	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
