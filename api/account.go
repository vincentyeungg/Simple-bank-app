package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vincentyeungg/Simple-bank-app/db/sqlc"
)

type createAccountRequest struct {
	// default Balance will be 0 when creating a new acocunt
	// we can do input validation from the client since Gin uses a valiadator under the hood
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// no errors in the incoming request parameters, can create new account
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// account created successfully, can return to client
	ctx.JSON(http.StatusOK, account)
}
