package repository

import "core-api-example/entity"

type UserRepository struct {
	// 데이터베이스 연결을 위한 필드 예시
}

func (ur *UserRepository) Save(user entity.User) (*entity.User, error) {
	// 실제 데이터베이스 저장 로직 구현
	// 예시: return nil, errors.New("Database not connected")
	return &user, nil
}
