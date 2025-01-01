package entities_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ardwiinoo/micro-music/authentications/internal/domains/users/entities"
)

func TestDetailUser_Validate(t *testing.T) {
	t.Run("Should return error when missing properties", func(t *testing.T) {
		user := &entities.DetailUser{
			ID:        0,
			Email:     "",
			FullName:  "",
			Password:  "",
			PublicId:  uuid.Nil,
			Role:      0,
			CreatedAt: "",
			UpdatedAt: "",
		}

		err := user.Validate()

		assert.NotNil(t, err)
		assert.EqualError(t, err, "DETAIL_USER.NOT_CONTAIN_NEEDED_PROPERTY")
	})

	t.Run("Should return nil when all properties are valid", func(t *testing.T) {
		user := &entities.DetailUser{
			ID:        1,
			Email:     "ardwiinoo@gmail.com",
			FullName:  "John Doe",
			Password:  "securepassword",
			PublicId:  uuid.New(),
			Role:      1,
			CreatedAt: "2024-12-26T00:00:00Z",
			UpdatedAt: "2024-12-26T00:00:00Z",
		}

		err := user.Validate()

		assert.Nil(t, err)
	})
}