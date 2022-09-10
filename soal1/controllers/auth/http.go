package auth

import (
	"encoding/json"
	"net/http"

	"majoo/soal1/controllers"
	"majoo/soal1/usecases/auth"
)

type authHTTPController struct {
	authUC auth.AuthUsecase
}

func NewAuthHTTPController(authUC auth.AuthUsecase) *authHTTPController {
	return &authHTTPController{
		authUC: authUC,
	}
}

func (ctrl *authHTTPController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(LoginRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		controllers.WriteResponse(w, http.StatusBadRequest, err.Error(), nil, nil)
	}

	result, err := ctrl.authUC.Login(ctx, auth.AuthEntity{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		controllers.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	res := LoginResponse{
		Token: result,
	}
	controllers.WriteResponse(w, http.StatusOK, "", res, nil)
}
