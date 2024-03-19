package user

import (
	"bekabar_chat/apps/user/entity"
	"bekabar_chat/apps/utils"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(
	userService *UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (user *UserHandler) RegisterUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userDTO = &entity.AddUserRequestDTO{
			Username: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		validationErr := utils.ValidateFields(*userDTO)
		if validationErr != nil {
			utils.AppResponse(w, http.StatusUnprocessableEntity, validationErr.Message)
			return
		}

		generateUuid, errGenerateUuid := uuid.NewV7()
		if errGenerateUuid != nil {
			fmt.Println("error generate uuid", errGenerateUuid)
			utils.AppResponse(w, http.StatusInternalServerError, errGenerateUuid)
			return
		}

		userDTO.ID = generateUuid.String()

		storeUserData := user.userService.StoreUser(userDTO)

		// Continue with normal response if no errors
		utils.AppResponse(w, http.StatusOK, storeUserData)
	})
}

func (user *UserHandler) GetAllUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		users, err := user.userService.FetchAllUser()
		if err != nil {
			utils.AppResponse(w, http.StatusInternalServerError, err)
		}

		// Continue with normal response if no errors
		utils.AppResponse(w, http.StatusOK, users)
	})
}
