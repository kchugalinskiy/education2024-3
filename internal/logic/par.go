package logic

import (
	"fmt"
	"github.com/kchugalinskiy/education2024-3/pkg/mutex"
)

func ParGo() {
	m := mutex.New()
	go func() {
		defer m.Done()
		fmt.Println("hello")
	}()
	m.Wait()
}
