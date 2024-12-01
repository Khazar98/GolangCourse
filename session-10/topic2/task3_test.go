package topic2

import "testing"

func TestSendingReceiving(t *testing.T) {



	ch :=make(chan int)
	go SendingReceiving(ch,42)
	select{
	case value:=<-ch:
		if value!=42{
			t.Errorf("Excepted value is 42 ,but got %d ",value)
		}
	}

}
