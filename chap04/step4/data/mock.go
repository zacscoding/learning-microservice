package data

import "github.com/stretchr/testify/mock"

// MockStore는 테스트용 데이터 저장소의 모의 구현
type MockStore struct {
	mock.Mock
}

// Search 메소드는 초기 설정에서 매개 변수로 전달된 객체를 리턴
func (m *MockStore) Search(name string) []Kitten {
	args := m.Mock.Called(name)
	return args.Get(0).([]Kitten)
}
