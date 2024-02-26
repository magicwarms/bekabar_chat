package user

import (
	"bekabar_chat/apps/user/model"
	"bekabar_chat/apps/utils"
	"bekabar_chat/structs"
	"net/http"

	"gorm.io/gorm"
)

func RegisterUser(connection *structs.MyConnectionService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user = model.UserModel{
			Username: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		// Handle errors directly from the transaction
		if err := connection.DB.Transaction(func(tx *gorm.DB) error {
			return tx.Create(&user).Error
		}); err != nil {
			utils.AppResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Continue with normal response if no errors
		utils.AppResponse(w, http.StatusOK, user)
	})
}
