package api

import (
	"net/http"

	db "github.com/KaungthuKhant/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// no balance because when we first create the account, the balanc should be zero
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

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

	// if there is not error in creating the account, send statusok response and the account that is created
	ctx.JSON(http.StatusOK, account)
}
