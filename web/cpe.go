/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/web"
)

// QueryCpe
func (h *HttpHandler) QueryCpe(c echo.Context) error {
	var result = make(map[string]interface{})
	data, err := h.GetManager().GetCpeManager().QueryCpes(web.NewWebForm(c))
	if err != nil {
		return h.GetInternalError(err)
	}
	result["data"] = data
	return c.JSON(http.StatusOK, result)
}


func (h *HttpHandler) AddCpe(c echo.Context) error {
	err := h.GetManager().GetCpeManager().AddCpe(web.NewWebForm(c))
	if err != nil {
		return h.GetInternalError(err)
	}
	h.AddOpsLog(c, fmt.Sprintf("Add Cpe sn=%s", c.FormValue("sn")))
	return c.JSON(200, h.RestSucc("Success"))
}


func (h *HttpHandler) UpdateCpeAttrs(c echo.Context) error {
	form := web.NewWebForm(c)
	err := h.GetManager().GetCpeManager().UpdateCpeAttrs(form)
	common.Must(err)
	h.AddOpsLog(c, fmt.Sprintf("Update CPE sn=%s", form.GetVal("sn")))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}


func (h *HttpHandler) DeleteCpe(c echo.Context) error {
	sn := c.QueryParam("sn")
	common.Must(h.GetManager().GetCpeManager().DeleteCpe(sn))
	h.AddOpsLog(c, fmt.Sprintf("Delete CPE sn=%s", sn))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}


