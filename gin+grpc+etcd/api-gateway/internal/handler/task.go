package handler

import (
	"api-gateway/internal/service/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTask(c *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(c.Bind(&tReq))
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserId)
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskShow(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Status: uint(taskResp.Code),
		Data:   taskResp,
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	c.JSON(http.StatusOK, r)
}

func CreateTask(c *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(c.Bind(&tReq))
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserId)
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskCreate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Status: uint(taskResp.Code),
		Data:   taskResp,
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	c.JSON(http.StatusOK, r)
}

func UpdataTask(c *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(c.Bind(&tReq))
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserId)
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskUpdate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Status: uint(taskResp.Code),
		Data:   taskResp,
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	c.JSON(http.StatusOK, r)
}

func DeleteTask(c *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(c.Bind(&tReq))
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserId)
	taskService := c.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskDelete(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Status: uint(taskResp.Code),
		Data:   taskResp,
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	c.JSON(http.StatusOK, r)
}
