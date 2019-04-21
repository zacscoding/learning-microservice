package throttling

import "net/http"

type LimitHandler struct {
	// 버퍼링 된 채널로서 연결의 수를 저장
	connections chan struct{}
	// 시스템이 정상이라는 것을 확인 후 호출하는 핸들러
	handler http.Handler
}

func NewLimitHandler(connections int, next http.Handler) *LimitHandler {
	cons := make(chan struct{}, connections)
	for i := 0; i < connections; i++ {
		cons <- struct{}{}
	}

	return &LimitHandler{
		connections: cons,
		handler:     next,
	}
}

func (l *LimitHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	select {
	case <-l.connections:
		l.handler.ServeHTTP(rw, r)
		l.connections <- struct{}{} // release the lock
	default:
		http.Error(rw, "Busy", http.StatusTooManyRequests)
	}
}
