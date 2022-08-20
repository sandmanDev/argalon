package controllers

import (
	"errors"
	"net/http"

	"github.com/argalon/config"
	"github.com/argalon/requestDM"
	"github.com/argalon/responseDM"
	"github.com/argalon/services"
	"github.com/argalon/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// DivideTwoController is the handler of divide_two endpoint
// @Summary Divide two request parameters and response back two results in json object
// @Description
// @Tags card Detail
// @Accept  json
// @Param order body requestDM.DivideRequestDM
// @Success 200 {object} responseDM.RequestResponse
// @Router /argalon/divide_two [post]
func DivideTwoController(ctx *gin.Context) {

	var logger = config.ZapLogger
	var divideRequestDM requestDM.DivideRequestDM
	var divideResponseDM responseDM.RequestResponse

	if err := ctx.ShouldBindJSON(&divideRequestDM); err != nil {
		var ve validator.ValidationErrors
		divideResponseDM.Msg = "required paramters not passed in request, please pass a and b in json request"
		divideResponseDM.Success = false
		if errors.As(err, &ve) {
			out := make(map[string]interface{})
			for _, fe := range ve {
				out[fe.Namespace()] = utils.GetErrorMsg(fe)
			}

			divideResponseDM.Errors = out
			ctx.JSON(http.StatusBadRequest, divideResponseDM)
			return
		}
		ctx.JSON(http.StatusBadRequest, divideResponseDM)
		return
	}

	result1, result2, errSer := services.DivideTwoByEachOther(divideRequestDM.ParamterA, divideRequestDM.ParamterB)
	if errSer != nil {
		logger.Error("error in division ", zap.Error(errSer), zap.Any("Provided Request DivideRequestDM", divideRequestDM))
		divideResponseDM.Msg = errSer.Error()
		divideResponseDM.Success = false
		ctx.JSON(http.StatusBadRequest, divideResponseDM)
		return
	}
	data := make(map[string]interface{})
	data["result1"] = result1
	data["result2"] = result2
	divideResponseDM.Msg = "division completed"
	divideResponseDM.Success = true
	divideResponseDM.Data = data

	ctx.JSON(http.StatusOK, divideResponseDM)
	return
}
