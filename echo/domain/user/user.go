package user

import (
	ulid "echo/domain/shared"
	"errors"
	"unicode/utf8"
)

type User struct {
	id   string
	name string
}

func newUser(name string) (*User, error) {
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errors.New("nameが不正")
	}
	return &User{
		id:   ulid.NewULID(), // idはここでUlidを生成
		name: name,
	}, nil
}

// DBからデータ取ってきて再構成するとき用
func reconstructUser(id string, name string) (*User, error) {
	return &User{
		id:   id,
		name: name,
	}, nil
}

const (
	// 名前の長さ
	nameLengthMin = 1
	nameLengthMax = 100
)
