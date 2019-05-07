package dist

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestJankenpoProxy_Play(t *testing.T) {
	type args struct {
		player1Move string
		player2Move string
	}
	jp := *NewJankenpoProxy("127.0.0.1", 1234, 100)

	tests := []struct {
		name string
		jp   JankenpoProxy
		args args
		want int
	}{
		{"Teste 1",
			jp,
			args{"P", "T"},
			1,
		},
		{"Teste 2",
			jp,
			args{"P", "P"},
			0,
		},
		{"Teste 3",
			jp,
			args{"T", "P"},
			2,
		},
	}

	inv := InvokerImpl{}
	go inv.Invoke()
	defer inv.Stop()

	time.Sleep(1 * time.Second)

	//var wg sync.WaitGroup
	for _, tt := range tests {
		//wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.jp.Play(tt.args.player1Move, tt.args.player2Move)

			if err != nil {
				//wg.Done()
				t.Errorf("RequestorImpl.Invoke() = Error %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				//wg.Done()
				t.Errorf("RequestorImpl.Invoke() = %v, want %v", got, tt.want)
			}
			//wg.Done()
			fmt.Println("Teste finalizado")
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
