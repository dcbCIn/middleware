package main

import (
	"middleware/lib/dist"
	"reflect"
	"sync"
	"testing"
)

func main() {
	//teste := dist.InvokerImpl{}
	//teste.Invoke()
	type args struct {
		inv dist.Invocation
	}

	requestor := dist.RequestorImpl{}
	tests := []struct {
		name string
		r    dist.RequestorImpl
		args args
		want dist.Termination
	}{
		{"Teste 1",
			requestor,
			args{dist.Invocation{1000, "127.0.0.1", 1234, "play", []interface{}{"P", "T"}}},
			dist.Termination{1},
		},
		{"Teste 2",
			requestor,
			args{dist.Invocation{1000, "127.0.0.1", 1234, "play", []interface{}{"P", "P"}}},
			dist.Termination{2},
		},
		{"Teste 3",
			requestor,
			args{dist.Invocation{1000, "127.0.0.1", 1234, "play", []interface{}{"T", "P"}}},
			dist.Termination{0},
		},
	}

	/*inv := dist.InvokerImpl{}
	go inv.Invoke()
	defer inv.Stop()

	time.Sleep(1 * time.Second)*/

	var t *testing.T
	var wg sync.WaitGroup

	for _, tt := range tests {
		wg.Add(1)
		//t.Run(tt.name, func(t *testing.T) {

		got, err := tt.r.Invoke(tt.args.inv)

		if err != nil {
			t.Errorf("RequestorImpl.Invoke() = Error %v", err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RequestorImpl.Invoke() = %v, want %v", got, tt.want)
		}
		wg.Done()
		//})

	}
	wg.Wait()
}
