// Copyright 2018 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package recover

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vicanso/hes"

	"github.com/vicanso/elton"
)

const (
	// ErrCategory recover error category
	ErrCategory = "elton-recover"
)

// New new recover
func New() elton.Handler {
	return func(c *elton.Context) error {
		defer func() {
			// 可针对实际需求调整，如对于每个recover增加邮件通知等
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				he := hes.Wrap(err)
				he.Category = ErrCategory
				he.StatusCode = http.StatusInternalServerError
				err = he
				c.Elton().EmitError(c, err)
				// 出错时清除部分响应头
				for _, key := range []string{
					elton.HeaderETag,
					elton.HeaderLastModified,
					elton.HeaderContentEncoding,
					elton.HeaderContentLength,
				} {
					c.SetHeader(key, "")
				}
				// 如果已直接对Response写入数据，则将 Committed设置为 true
				c.Committed = true
				resp := c.Response
				buf := []byte(err.Error())
				if strings.Contains(c.GetRequestHeader("Accept"), "application/json") {
					c.SetHeader(elton.HeaderContentType, elton.MIMEApplicationJSON)
					buf = he.ToJSON()
				}
				resp.WriteHeader(he.StatusCode)
				resp.Write(buf)
			}
		}()
		return c.Next()
	}
}
