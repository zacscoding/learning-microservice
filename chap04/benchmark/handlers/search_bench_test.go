package handlers

import (
	"bytes"
	"github.com/zacscoding/learning-microservice-with-go/chap04/benchmark/data"
	"net/http/httptest"
	"testing"
)

// go test -bench=. -benchmem
// go test -bench=. -cpuprofile=cpu.prof -blockprofile=block.prof -memprofile=mem.prof
func BenchmarkSearchHandler(b *testing.B) {
	var mockStore = &data.MockStore{}
	mockStore.On("Search", "Fat Freddy's Cat").Return([]data.Kitten{
		{
			Name: "Fat Freddy's Cat",
		},
	})

	search := Search{DataStore: mockStore}

	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest("POST", "/search", bytes.NewReader([]byte(`{"query", "Fat Freddy's Cat'"}`)))
		rr := httptest.NewRecorder()

		search.ServeHTTP(rr, r)
	}
}
