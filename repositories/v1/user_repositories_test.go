package v1

import (
	"errors"
	"testing"

	"example/hello/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// mockDB is a mock implementation of *gorm.DB for testing
type mockDB struct {
	users []models.User
	err   error
}

func (m *mockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	if m.err != nil {
		return &gorm.DB{Error: m.err}
	}
	ptr, ok := dest.(*[]models.User)
	if ok {
		*ptr = m.users
	}
	return &gorm.DB{Error: nil}
}

func TestGetAllUsers_Success(t *testing.T) {
	// Arrange
	mockUsers := []models.User{
		{ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"), Username: "alice"},
		{ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Username: "bob"},
	}
	mock := &mockDB{users: mockUsers}

	// Act
	users, err := GetAllUsersWithDB(mock)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, mockUsers, users)
}

func GetAllUsersWithDB(mock *mockDB) ([]models.User, error) {
	var users []models.User
	db := mock.Find(&users)
	if db.Error != nil {
		return nil, db.Error
	}
	return users, nil
}

func TestGetAllUsers_Error(t *testing.T) {
	// Arrange
	mock := &mockDB{err: errors.New("db error")}

	// Act
	users, err := GetAllUsersWithDB(mock)
	assert.Error(t, err)
	assert.Nil(t, users)
}
