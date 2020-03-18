package throttling

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func newTestHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		<-r.Context().Done()
	})
}

func setup(ctx context.Context) (*httptest.ResponseRecorder, *http.Request) {
	r := httptest.NewRequest("GET", "/health", nil)
	r = r.WithContext(ctx)
	return httptest.NewRecorder(), r
}

func TestCallsNextWhenConnectionsOK(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	handler := NewLimitHandler(2, newTestHandler(ctx))
	rw, r := setup(ctx)

	go handler.ServeHTTP(rw, r)
	cancel()
	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, http.StatusOK, rw.Code)
}

func TestReturnsBusyWhen0Connection(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	handler := NewLimitHandler(0, newTestHandler(ctx))
	rw, r := setup(ctx)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
	})
	handler.ServeHTTP(rw, r)

	assert.Equal(t, http.StatusTooManyRequests, rw.Code)
}

func TestReturnsOKWith2ConnectionsAndConnectionLimit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	handler := NewLimitHandler(2, newTestHandler(ctx))
	rw, r := setup(ctx)
	rw2, r2 := setup(ctx2)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
		cancel2()
	})

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		handler.ServeHTTP(rw, r)
		waitGroup.Done()
	}()

	go func() {
		handler.ServeHTTP(rw2, r2)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	if rw.Code != http.StatusOK || rw2.Code != http.StatusOK {
		t.Fatalf("Both requests should be OK, request 1 : %v, request 2 : %v", rw.Code, rw2.Code)
	}
}

func TestReturnBusyWhenConnectionsExhausted(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	// LimitHandler를 생성하고 이 핸들러에 모의 핸들러를 전달
	handler := NewLimitHandler(1, newTestHandler(ctx))
	rw, r := setup(ctx)
	rw2, r2 := setup(ctx2)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
		cancel2()
	})

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		handler.ServeHTTP(rw, r)
		waitGroup.Done()
	}()

	go func() {
		handler.ServeHTTP(rw2, r2)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	if rw.Code == http.StatusOK && rw2.Code == http.StatusOK {
		t.Fatalf("One request should have bean busy, request 1 : %v, request 2 : %v", rw.Code, rw2.Code)
	}
}

func TestReleasesConnectionLockWhenFinished(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	handler := NewLimitHandler(1, newTestHandler(ctx))
	rw, r := setup(ctx)
	rw2, r2 := setup(ctx2)

	cancel()
	cancel2()

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		handler.ServeHTTP(rw, r)
		waitGroup.Done()
		handler.ServeHTTP(rw2, r2)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	if rw.Code != http.StatusOK || rw2.Code != http.StatusOK {
		t.Fatalf("One request should have been busy, request 1: %v, request 2: %v", rw.Code, rw2.Code)
	}

}
