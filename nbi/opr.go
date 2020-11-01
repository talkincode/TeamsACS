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

package nbi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/models"
)

// QueryOperator
func (h *HttpHandler) QueryOperator(c echo.Context) error {
	var result = make(map[string]interface{})
	params := h.RequestParse(c)
	data, err := h.GetManager().GetOpsManager().QueryOperators(params)
	if err != nil {
		return h.GetInternalError(err)
	}
	result["data"] = data
	return c.JSON(http.StatusOK, result)
}

// AddOperator
func (h *HttpHandler) AddOperator(c echo.Context) error {
	item := new(models.Operator)
	common.Must(c.Bind(item))
	_id, err := h.GetManager().GetOpsManager().AddOperator(item)
	if err != nil {
		return h.GetInternalError(err)
	}
	h.AddOpsLog(c, fmt.Sprintf("Add Operator id=%s", _id))
	return c.JSON(200, h.RestSucc("Success"))
}


// UpdateOperator
func (h *HttpHandler) UpdateOperator(c echo.Context) error {
	item := new(models.Operator)
	common.Must(c.Bind(item))
	err := h.GetManager().GetOpsManager().UpdateOperator(item)
	common.Must(err)
	h.AddOpsLog(c, fmt.Sprintf("Update Operator id=%s",item.ID))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}

// DeleteOperator
func (h *HttpHandler) DeleteOperator(c echo.Context) error {
	params := h.RequestParse(c)
	username := params.GetMustString("username")
	common.Must(h.GetManager().GetOpsManager().DeleteOperator(username))
	h.AddOpsLog(c, fmt.Sprintf("Delete Operator username=%s", username))
	return c.JSON(http.StatusOK, h.RestSucc("Success"))
}



