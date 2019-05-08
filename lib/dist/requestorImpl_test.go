package dist

import (
	"reflect"
	"testing"
	"time"
)

func TestRequestorImpl_Invoke(t *testing.T) {
	type args struct {
		inv Invocation
	}
	tests := []struct {
		name string
		r    RequestorImpl
		args args
		want Termination
	}{
		{"Teste 1",
			RequestorImpl{},
			args{Invocation{1000, "127.0.0.1", 1234, "jankenpo.play", []interface{}{"P", "T"}}},
			Termination{float64(1)},
		},
		{"Teste 2",
			RequestorImpl{},
			args{Invocation{1000, "127.0.0.1", 1234, "jankenpo.play", []interface{}{"P", "P"}}},
			Termination{float64(0)},
		},
		{"Teste 3",
			RequestorImpl{},
			args{Invocation{1000, "127.0.0.1", 1234, "jankenpo.play", []interface{}{"T", "P"}}},
			Termination{float64(2)},
		},
	}

	inv := InvokerImpl{}
	go inv.Invoke(1234)
	defer inv.Stop()

	time.Sleep(1 * time.Second)

	//var wg sync.WaitGroup
	for _, tt := range tests {
		//wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.r.Invoke(tt.args.inv)

			if err != nil {
				//wg.Done()
				t.Errorf("RequestorImpl.Invoke() = Error %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				//wg.Done()
				t.Errorf("RequestorImpl.Invoke() = %v, want %v", got, tt.want)
			}
			//wg.Done()
		})
	}
	//wg.Wait()
}

/*func TestInvoker(t *testing.T) {
	r := new(RequestorImpl)

	var parameters [2]string
	parameters[0] = "Pedra"
	parameters[1] = "Tesoura"

	i := InvocationImpl{ObjectId: 1000, IpAddress: "127.0.0.1:1234", PortNumber: 1234, OperationName: "jankenpo", Parameters: parameters}

	r.Invoke(i)
}*/
