// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

type CreateTagRequest struct {
	Name    string `json:"name" binding:"required"`
	Route   string `json:"route" binding:"required"`
	Enabled bool   `json:"enabled"`
}

type TagEnabledRequest struct {
	Enabled *bool `json:"enabled" binding:"required"`
}

type PageRequest struct {
	// 当前页
	PageNo int64 `form:"pageNo" binding:"required"`
	// 每页数量
	PageSize int64 `form:"pageSize" binding:"required"`
	// 排序字段
	Field string `form:"sortField,omitempty"`
	// 排序规则
	Order string `form:"sortOrder,omitempty"`
	// 搜索内容
	Keyword string `form:"keyword,omitempty"`
}
