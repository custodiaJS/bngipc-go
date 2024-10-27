package tests

import (
	"fmt"
	"sync"
	"testing"

	bngipcgo "github.com/custodiaJS/bngipc-go"
)

func TestIPC(m *testing.T) {
	_, err := bngipcgo.SetupNewIpcServer("custodiajs")
	if err != nil {
		fmt.Println(err)
		m.Fail()
	}
	t := new(sync.WaitGroup)
	t.Add(1)
	t.Wait()
}
