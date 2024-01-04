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

package vo

type FriendVO struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
}

type AdminFriendVO struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Show        bool   `json:"show"`
	Accepted    bool   `json:"accepted"`
	CreateTime  int64  `json:"create_time"`
}
