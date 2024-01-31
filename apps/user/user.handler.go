package user

import (
	"bekabar_chat/apps/user/model"
	"bekabar_chat/apps/utils"
	"net/http"

	"gorm.io/gorm"
)

func RegisterUser(DB *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		var user = model.UserModel{
			Username: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		// Handle errors directly from the transaction
		if err := DB.Transaction(func(tx *gorm.DB) error {
			return tx.Create(&user).Error
		}); err != nil {
			utils.AppResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Continue with normal response if no errors
		utils.AppResponse(w, http.StatusOK, user)
	})
}
