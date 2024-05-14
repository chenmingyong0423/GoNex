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

package domain

type Page struct {
	Size    int64
	Skip    int64
	Keyword string
	Field   string
	Order   string

	CategoryFilter []string
	TagFilter      []string
}

type PostsQueryCondition struct {
	Size int64
	Skip int64

	Keyword *string

	Sorting

	Categories []string
	Tags       []string
}

type DetailPostVO struct {
	PrimaryPost
	ExtraPost
	IsLiked bool `json:"is_liked"`
}

type Post struct {
	PrimaryPost
	ExtraPost
	Likes []string `json:"-"`
}

type ExtraPost struct {
	Content          string `json:"content"`
	MetaDescription  string `json:"meta_description"`
	MetaKeywords     string `json:"meta_keywords"`
	WordCount        int    `json:"word_count"`
	UpdatedAt        int64  `json:"updated_at"`
	IsDisplayed      bool   `json:"is_displayed"`
	IsCommentAllowed bool   `json:"is_comment_allowed"`
}

type PrimaryPost struct {
	Id           string          `json:"_id"`
	Author       string          `json:"author"`
	Title        string          `json:"title"`
	Summary      string          `json:"summary"`
	CoverImg     string          `json:"cover_img"`
	Categories   []Category4Post `json:"category"`
	Tags         []Tag4Post      `json:"tags"`
	LikeCount    int             `json:"like_count"`
	CommentCount int             `json:"comment_count"`
	VisitCount   int             `json:"visit_count"`
	StickyWeight int             `json:"sticky_weight"`
	CreatedAt    int64           `json:"created_at"`
}

type Category4Post struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tag4Post struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PostRequest struct {
	PageRequest
	Categories []string `form:"categories"`
	Tags       []string `form:"tags"`
}

type PageRequest struct {
	Page2
	// 排序字段
	Sorting
	// 搜索内容
	Keyword *string `form:"keyword,omitempty"`
}

func (p *PageRequest) ValidateAndSetDefault() {
	if p.PageNo <= 0 {
		p.PageNo = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

type Sorting struct {
	Field *string `form:"sortField,omitempty"`
	Order *string `form:"sortOrder,omitempty"`
}

type Page2 struct {
	// 当前页
	PageNo int64 `form:"pageNo" binding:"required"`
	// 每页数量
	PageSize int64 `form:"pageSize" binding:"required"`
}
