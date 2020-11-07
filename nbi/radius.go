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
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
)

func (h *HttpHandler) QueryRadiusAccounting(c echo.Context) error {
	params := h.RequestParse(c)
	data, err := h.GetManager().GetRadiusManager().QueryAccountings(params)
	common.Must(err)
	return c.JSON(http.StatusOK, data)
}

func (h *HttpHandler) QueryRadiusAuthlog(c echo.Context) error {
	params := h.RequestParse(c)
	data, err := h.GetManager().GetRadiusManager().QueryAuthlogs(params)
	common.Must(err)
	return c.JSON(http.StatusOK, data)
}

func (h *HttpHandler) QueryRadiusOnline(c echo.Context) error {
	params := h.RequestParse(c)
	data, err := h.GetManager().GetRadiusManager().QueryOnlines(params)
	common.Must(err)
	return c.JSON(http.StatusOK, data)
}

