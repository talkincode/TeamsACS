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
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
)


type DateRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// WEB 参数
type WebForm struct {
	FormItem interface{}
	Posts    url.Values        `json:"-" form:"-" query:"-"`
	Gets     url.Values        `json:"-" form:"-" query:"-"`
	Params   map[string]string `json:"-" form:"-" query:"-"`
}
func EmptyWebForm() *WebForm {
	v := &WebForm{}
	v.Params = make(map[string]string,0)
	v.Posts = make(url.Values,0)
	v.Gets = make(url.Values,0)
	return v
}

func NewWebForm(c echo.Context) *WebForm {
	v := &WebForm{}
	v.Params = make(map[string]string)
	v.Posts, _ = c.FormParams()
	v.Gets = c.QueryParams()
	for _, p := range c.ParamNames() {
		v.Params[p] = c.Param(p)
	}
	return v
}

func (f *WebForm) Set(name string, value string) {
	f.Gets.Set(name, value)
}

func (f *WebForm) Param(name string) string {
	return f.Param(name)
}

func (f *WebForm) Param2(name string, defval string) string {
	if val, ok := f.Params[name]; ok {
		return val
	}
	return defval
}

func (f *WebForm) GetDateRange(name string) (DateRange, error) {
	var dr = DateRange{Start: "", End: ""}
	val := f.GetVal(name)
	if val == "" {
		return dr, nil
	}
	err := json.Unmarshal([]byte(val), &dr)
	if err != nil {
		return dr, err
	}
	return dr, nil
}

func (f *WebForm) GetVal(name string) string {
	val := f.Posts.Get(name)
	if val != "" {
		return val
	}
	val = f.Gets.Get(name)
	if val != "" {
		return val
	}
	return ""
}

func (f *WebForm) GetMustVal(name string) (string, error) {
	val := f.Posts.Get(name)
	if val != "" {
		return val, nil
	}
	val = f.Gets.Get(name)
	if val != "" {
		return val, nil
	}
	return "", errors.New(name + " 不能为空")
}

func (f *WebForm) GetVal2(name string, defval string) string {
	val := f.Posts.Get(name)
	if val != "" {
		return val
	}
	val = f.Gets.Get(name)
	if val != "" {
		return val
	}
	return defval
}

func (f *WebForm) GetIntVal(name string, defval int) int {
	val := f.GetVal(name)
	if val == "" {
		return defval
	}
	v, _ := strconv.Atoi(val)
	return v
}

func (f *WebForm) GetInt64Val(name string, defval int64) int64 {
	val := f.GetVal(name)
	if val == "" {
		return defval
	}
	v, _ := strconv.ParseInt(val, 10, 64)
	return v
}


type PageResult struct {
	TotalCount int64       `json:"total_count,omitempty"`
	Pos        int64       `json:"pos"`
	Data       interface{} `json:"data"`
}

var EmptyPageResult = &PageResult{
	TotalCount: 0,
	Pos:        0,
	Data:       common.EmptyList,
}
