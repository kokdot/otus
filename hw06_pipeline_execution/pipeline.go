package hw06pipelineexecution

// import "github.com/sqs/goreturns/returns"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func wrap(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer func() {
			close(out)
			for range in {
			}
		}()
		for {
			select {
			case <-done:
				return
			default:
			}

			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

// ExecutePipeline builds pipline from stages
// if stage is nil it will be skipped.
func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := wrap(in, done)

	for i := 0; i < len(stages); i++ {
		if stages[i] != nil {
			out = stages[i](wrap(out, done))
		}
	}
	return out
}
