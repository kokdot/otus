package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		// t.Skip()
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("Unlimited errors", func(t *testing.T) {
		// t.Skip()
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		
		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10

		require.NoError(t, Run(tasks, workersCount, 0))
		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		// t.Skip()
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		// workersCount := 60
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)
		fmt.Println("runTasksCount :  --", runTasksCount)
		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
	t.Run("invalid n", func(t *testing.T) {
		err := Run(nil, 0, 0)
		require.Truef(t, errors.Is(err, ErrInvalidN), "actual err - %v", err)
	})
}

func TestRunConcurrency(t *testing.T) {
	const workersCount = 5
	tasks := make([]Task, workersCount)
	waitch := make(chan struct{})
	var runTaskCount int32
	for i := 0; i < len(tasks); i++ {
		tasks[i] = func() error {
			atomic.AddInt32(&runTaskCount, 1)
			<- waitch
			return nil
		}
	}

	runErrCh := make(chan error, 1)
	go func() {
		runErrCh <- Run(tasks, workersCount, 0)
	}()

	require.Eventually(t, func() bool {
		return atomic.LoadInt32(&runTaskCount) == workersCount
	}, 5 * time.Second, time.Millisecond)

	close(waitch)

	var runErr error
	require.Eventually(t, func() bool {
		select {
		case runErr = <- runErrCh:
			return true
		default:
			return false
		}
	}, 5 * time.Second, time.Millisecond)

	require.NoError(t, runErr)
}
