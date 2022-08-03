package hw06pipelineexecution
import(
	"fmt"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	fmt.Println("stages")
	fmt.Println(len(stages))
	// if in == nil {
	// 	return nil
	// }
	select {
	case <- done:
		terminate := make(Bi)
		defer close(terminate)
		return terminate
	default:
		// return in 
	}

	if len(stages) == 0 {
		return in
	}

	// if len(stages) == 0 {
	// 	return in 
	// }
	// out := ExecutePipeline(in, done, stages...)
	// stage := stages[0]
	// stages = stages[1:]
	
	select {
	case <- done:
		fmt.Println("PPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPP")
		return nil
	default:
	// case <- in:
	}
	
	// go func(done In) () {
		// 	done
		// }(done)
		
	stage := stages[0]
	stages = stages[1:]
	return ExecutePipeline(stage(in), done, stages...)
	// return ExecutePipeline(stage(in), done, stages...)

}
