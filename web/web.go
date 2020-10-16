package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/config"
	"github.com/ca17/teamsacs/models"
)

type RestResult struct {
	Code    int         `json:"code"`
	Msgtype string      `json:"msgtype"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}



type WebContext struct {
	Manager *models.ModelManager
	Config  *config.AppConfig
}

type WebHandler interface {
	InitRouter(group *echo.Group)
}

type HttpHandler struct {
	Ctx *WebContext
}

func NewHttpHandler(ctx *WebContext) HttpHandler {
	return HttpHandler{Ctx: ctx}
}

func (h *HttpHandler) InitRouter(group *echo.Group) {

}


func (h *HttpHandler) GetConfig() *config.AppConfig {
	return h.Ctx.Config
}

func (h *HttpHandler) GetManager() *models.ModelManager {
	return h.Ctx.Manager
}

func (h *HttpHandler) GetInternalError(err interface{}) *echo.HTTPError {
	switch err.(type) {
	case error:
		return echo.NewHTTPError(http.StatusInternalServerError, err.(error).Error())
	case string:
		return echo.NewHTTPError(http.StatusInternalServerError, err.(string))
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}

func (h *HttpHandler) GoInternalErrPage(c echo.Context, err interface{}) error {
	switch err.(type) {
	case error:
		return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message":err.(error).Error()})
	case string:
		return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message":err.(string)})
	}
	return c.Render(http.StatusInternalServerError, "err500", map[string]string{"message":err.(string)})
}


func (h *HttpHandler) RestResult(data interface{}) *RestResult {
	return &RestResult{
		Code:    0,
		Msgtype: "info",
		Msg:     "Operation Success",
		Data:    data,
	}
}

func (h *HttpHandler) RestSucc(msg string) *RestResult {
	return &RestResult{
		Code:    0,
		Msgtype: "info",
		Msg:     msg,
	}
}

func (h *HttpHandler) RestError(msg string) *RestResult {
	return &RestResult{
		Code:    9999,
		Msgtype: "error",
		Msg:     msg,
	}
}

func (h *HttpHandler) ParseFormInt64(c echo.Context, name string) (int64, error) {
	return strconv.ParseInt(c.FormValue("id"), 10, 64)

}



func (h *HttpHandler) FetchExcelData(c echo.Context, sheet string) ([]map[string]string, error) {

	file, err := c.FormFile("upload")
	if err != nil {
		return nil, err
	}
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		return nil, err
	}
	// 获取 Sheet1 上所有单元格
	rows := f.GetRows(sheet)
	head := make(map[int]string)
	var data []map[string]string
	for i, row := range rows {
		item := make(map[string]string)
		for k, colCell := range row {
			if i == 0 {
				head[k] = colCell
			} else {
				item[common.ToCamelCase(head[k])] = colCell
			}
		}
		if i == 0 {
			continue
		}
		data = append(data, item)
	}

	return data, nil
}

type HTTPError struct {
	Code     int         `json:"-"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"` // Stores the error returned by an external dependency
}

func NewHTTPError(code int, message ...interface{}) *HTTPError {
	he := &HTTPError{Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d:%s", e.Code, e.Message)
}
