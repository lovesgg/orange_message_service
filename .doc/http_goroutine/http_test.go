package http_goroutine

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func BenchmarkHttp1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			httpSleep(1)
			wg.Done()
		}()
		go func() {
			httpSleep(1)
			wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkHttp2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		httpSleep(1)
		httpSleep(1)
	}
}

func httpSleep(second int) {
	http.Get(fmt.Sprintf("http://127.0.0.1:1234/sleep?second=%d", second))
}
