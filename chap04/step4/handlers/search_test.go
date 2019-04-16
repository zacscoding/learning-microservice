package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zacscoding/learning-microservice-with-go/chap04/step4/data"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
# go test
go test -v -race
-v : 상세 출력 형식으로 출력
-race 동시성 문제가 있는 버그를 발견하는 Go의 레이스 탐지(race detector)활성화

# 코드 커버리지
go test -cover ./...
*/

var mockStore *data.MockStore

func TestSearchHandlerReturnBadRequestWhenNoSearchCriteriaIsSend(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{})

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSend(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{})

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest god %v", rw.Code)
	}
}

func TestSearchHandlerCallsDataStoreWithValidQuery(t *testing.T) {
	query := "Fat Freddy's Cat"
	r, rw, handler := setupTest(&searchRequest{Query: query})
	mockStore.On("Search", query).Return(make([]data.Kitten, 0))

	handler.ServeHTTP(rw, r)

	decoder := json.NewDecoder(rw.Body)
	searchResponse := new(searchResponse)
	err := decoder.Decode(searchResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(searchResponse)

	mockStore.AssertExpectations(t)
}

func setupTest(d interface{}) (*http.Request, *httptest.ResponseRecorder, Search) {
	mockStore = &data.MockStore{}

	h := Search{
		Store: mockStore,
	}
	rw := httptest.NewRecorder()

	if d == nil {
		return httptest.NewRequest("POST", "/search", nil), rw, h
	}

	body, _ := json.Marshal(d)
	return httptest.NewRequest("POST", "/search", bytes.NewReader(body)), rw, h
}
