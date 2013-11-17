package gorinju

import (
	"fmt"
	"time"
	"sync"
	"net/http"
)

func bang() {
	fmt.Print(".")
}

func Bullet(url string) {
	bang()
	http.Get(url)
}

func MultiBullet(url string, wg *sync.WaitGroup) {
	bang()
	http.Get(url)
	wg.Done()
}

func HandGun(url string, bullet int) {
	start := time.Now()
	for i := 0;i < bullet; i++ {
		Bullet(url)
	}
	end := time.Now()
	fmt.Println()
	result := end.Sub(start)

	rep := Report{url, result}
	rep.Print()
}

func MachineGun (url string, bullet int) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < bullet; i++ {
		wg.Add(1)
		go MultiBullet(url, &wg)
    }
	wg.Wait()

	end := time.Now()
	result := end.Sub(start)

	rep := Report{url, result}
	fmt.Println()
	rep.Print()
}
