package middlewares

import (
	"Auth_Api_Gateway/dto"
	"Auth_Api_Gateway/utils"
	"context"
	"net/http"
)

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.SignupUserRequestDTO
		Jsonerr := utils.ReadJsonBody(r, &payload)
		if Jsonerr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Error in reading the payload", Jsonerr)
			return
		}

		utils.Validator.Struct(payload)
		if ValidationErr := utils.Validator.Struct(payload); ValidationErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input type", ValidationErr)
			return
		}



		req_context := r.Context()   // original context comming with req
		
		ctx := context.WithValue(req_context,"signupPaylaod", &payload) // creating a new context and appending our payload
		// key -> "loginPayload" value -> our payload

		next.ServeHTTP(w, r.WithContext(ctx)) 	// call next with the new context (which having our payload)
	})
}


func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDTO
		Jsonerr := utils.ReadJsonBody(r, &payload)
		if Jsonerr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Error in reading the payload", Jsonerr)
			return
		}

		utils.Validator.Struct(payload)
		if ValidationErr := utils.Validator.Struct(payload); ValidationErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input type", ValidationErr)
			return
		}
		ctx := context.WithValue(r.Context(),"loginPayload",&payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}