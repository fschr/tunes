// Copyright © 2016 Thomas Fischer <tdf.tomfischer@gmail.com>
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

package core

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func idFromURL(url string) (id string, err error) {
	for i := len(url) - 1; i > 0; i-- {
		if url[i] == '=' {
			return url[i+1:], nil
		}
	}
	log.WithFields(log.Fields{"url": url}).Warn("invalid url")
	return "", fmt.Errorf("invalid url")
}
