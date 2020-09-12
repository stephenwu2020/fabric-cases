package v3

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stephenwu2020/fabric-cases/web-app/server/pkg/app"
	"github.com/stephenwu2020/fabric-cases/web-app/server/sdk"
)

func QueryCar(ctx *gin.Context) {
	appGin := app.Gin{C: ctx}

	type Body struct {
		Key string `json:"key"`
	}
	var body Body

	if err := ctx.ShouldBind(&body); err != nil {
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	rsp, err := sdk.ChannelQuery("queryCar", body.Key)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(rsp, &data)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	appGin.Response(http.StatusOK, "success", data)
}

func QueryAllCars(ctx *gin.Context) {
	appGin := app.Gin{C: ctx}

	rsp, err := sdk.ChannelQuery("queryAllCars")
	if err != nil {
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	var data interface{}
	err = json.Unmarshal(rsp, &data)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	appGin.Response(http.StatusOK, "success", data)

}

func CreateCar(ctx *gin.Context) {
	appGin := app.Gin{C: ctx}

	type Body struct {
		Make   string `json:"make"`
		Model  string `json:"model"`
		Colour string `json:"colour"`
		Owner  string `json:"owner"`
	}
	var body Body
	if err := ctx.ShouldBind(&body); err != nil {
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	_, err := sdk.ChannelExecute("createCar", body.Make, body.Model, body.Colour, body.Owner)
	if err != nil {
		log.Println(errors.WithMessage(err, "create car failed."))
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	appGin.Response(http.StatusOK, "success", "")
}

func ChangeCarOwner(ctx *gin.Context) {
	appGin := app.Gin{C: ctx}

	type Body struct {
		Key   string `json:"key"`
		Owner string `json:"owner"`
	}
	var body Body
	if err := ctx.ShouldBind(&body); err != nil {
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	_, err := sdk.ChannelExecute("changeCarOwner", body.Key, body.Owner)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, "fail", err.Error())
		return
	}

	appGin.Response(http.StatusOK, "success", "")
}
