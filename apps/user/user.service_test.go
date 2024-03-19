package user

import (
	"testing"
)

func TestStoreUser(t *testing.T) {
	// db, mock := dbmock.NewMockDB()
	// userService := &UserService{DB: db}

	// t.Run("Test storing a user successfully", func(t *testing.T) {
	// user := &entity.AddUserRequestDTO{
	// 	ID:       "018e51ee-5473-7aca-9984-0b62f9555f73",
	// 	Username: "testuser",
	// 	Email:    "test@example.com",
	// 	Password: "password123",
	// }

	// mock.ExpectBegin()
	// mock.ExpectExec("INSERT INTO \"users\" (.+) VALUES (.+)").WithArgs(user.ID, user.Username, user.Email, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg(), nil)
	// mock.ExpectCommit()

	// err := userService.StoreUser(user)
	// if err != nil {
	// 	t.Errorf("unexpected error: %v", err)
	// }

	// err = mock.ExpectationsWereMet()
	// if err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
	// })

	// t.Run("Test handling an error from the transaction", func(t *testing.T) {
	// 	user := &entity.AddUserRequestDTO{
	// 		ID:       "018e51ee-5473-7aca-9984-0b62f9555f73",
	// 		Username: "testuser",
	// 		Email:    "test@example.com",

	// 		Password: "password123",
	// 	}

	// 	mock.ExpectBegin()
	// 	mock.ExpectExec("INSERT INTO users").WithArgs(user.ID, user.Email, user.Username, user.Password).WillReturnError(errors.New("error inserting user"))
	// 	mock.ExpectRollback()

	// 	err := userService.StoreUser(user)
	// 	if err == nil {
	// 		t.Error("expected an error but got nil")
	// 	}

	// 	err = mock.ExpectationsWereMet()
	// 	if err != nil {
	// 		t.Errorf("there were unfulfilled expectations: %s", err)
	// 	}
	// })
}
