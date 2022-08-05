package hw06pipelineexecution

// import "github.com/sqs/goreturns/returns"

// "fmt"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func wrap(in In, done In) Out {
	out := make(Bi)
	go func() () {
		defer close(out)
		for {
			select {
			case <- done:
				return
			case v, ok := <- in:
				if !ok {
					return
				}
				out <- v
				// return
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// out := make(chan Out)
	out := in
	for i := 0; i< len(stages); i++ {
		// out = stages[i](out)
		out = stages[i](wrap(out, done))

	}
	return out
}
