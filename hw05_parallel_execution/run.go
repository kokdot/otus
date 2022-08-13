package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
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

/*
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	tasksCount := len(tasks)
	workersCount := n
	maxNumber := tasksCount
	maxErrorsCount := m
	wg := sync.WaitGroup{}
	var errSum int
	var noErrSum int
	var number int
	var stopNumber int
	checkTaskError := false
	checkTaskErrorDone := false
	checkMaxNumber := false
	checkMaxNumberDone := false
	checkTaskStop := false
	checkTaskStopDone := false
	checkTaskEnd := false
	checkTaskEndDone := false
	chTask := make(chan Task, tasksCount)
	chStop := make(chan struct{}, n)
	chStopGoRutine := make(chan struct{}, n)
	chManageGoRutine := make(chan struct{}, n)
	chGetMistake := make(chan struct{}, n)
	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func(numberOfGoRutine int) {
			defer wg.Done()
			fmt.Println("Запущена горутина: ", numberOfGoRutine)
			for {
				select {
				case task, ok := <-chTask:
					if ok {
						err := task()
						if err != nil {
							chGetMistake <- struct{}{}
							fmt.Println("Выполнена задача с ошибкой", "numberOfGoRutine -- ", numberOfGoRutine)
						} else {
							chManageGoRutine <- struct{}{}
							fmt.Println("Выполнена задача ", "numberOfGoRutine -- ", numberOfGoRutine)
						}
					}
				case <-chStop:
					fmt.Println("---------------Получено задане на остановку работы", " numberOfGoRutine -- ", numberOfGoRutine)
					chStopGoRutine <- struct{}{}
					return
				}
			}
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fistCount := workersCount
		if fistCount > maxNumber {
			fistCount = maxNumber
		}
		for i := 0; i < fistCount; i++ {
			number++
			chTask <- tasks[i]
		}
		// fmt.Println("number = ", number, "; maxNumber = ", maxNumber)
		for {
			if ((number == maxNumber) || (errSum == maxErrorsCount)) && !checkTaskStopDone {
				checkTaskStopDone = true
				checkTaskStop = true
			}
			// fmt.Println("(((((((((((((((((((((((((((((((((((maxNumber ", maxNumber)

			select {
			case <-chGetMistake:
				errSum++
				if !checkTaskErrorDone {
					checkTaskError = true
					checkTaskErrorDone = true
					if (workersCount + maxErrorsCount) < maxNumber {
						maxNumber = workersCount + m
					}
				}
				if !checkTaskStop {
					number++
					chTask <- tasks[number-1]
				}
				fmt.Println("  |   errSum = ", errSum, "; noErrSum = ",
					noErrSum, "; maxNumber = ", maxNumber, "; number", number, "++++++++++")
				fmt.Println("получ ошиб", "Общее кол ош: ", errSum, "Общее кол отпр заданий: ",
					number, " checkTaskStop: ", checkTaskStop)

			case <-chManageGoRutine:
				noErrSum++
				if !checkTaskStop {
					number++
					chTask <- tasks[number-1]
				}
				fmt.Println("Выполнено задание", "Общее кол заданий без ошибок: ",
					noErrSum, "Общее кол отпр заданий: ", number)

			case <-chStopGoRutine:
				stopNumber++
				fmt.Println("____________Остановлено горутин: ", stopNumber)
			default:
			}
			if errSum + noErrSum == maxNumber && !checkMaxNumberDone {
				fmt.Println("  |   errSum = ", errSum, "; noErrSum = ", noErrSum,
					"; maxNumber = ", maxNumber, "; number", number, "++++++++++")
				checkMaxNumber = true
				checkMaxNumberDone = true
			}
			if checkMaxNumber {
				checkMaxNumber = false
				for i := 0; i < workersCount; i++ {
					chStop <- struct{}{}
				}
				fmt.Println("  |   errSum = ", errSum, "; noErrSum = ", noErrSum,
					"; maxNumber = ", maxNumber, "; number", number, "++++++++++")
				fmt.Println("Выполнены все задания, флаг checkTaskStop установлен в true, отпрвлены задачи о завершении работы.")
			}
			if stopNumber == workersCount && !checkTaskEndDone {
				checkTaskEndDone = true
				checkTaskEnd = true
			}

			if checkTaskEnd {
				return
			}
		}
	}()
	wg.Wait()
	if checkTaskError {
		return ErrErrorsLimitExceeded
	}
	return nil
}
*/
