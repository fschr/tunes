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

package api

import (
	"fmt"
	"net/http"

	"github.com/fschr/tunes/core"
	log "github.com/sirupsen/logrus"
)

func reply(w http.ResponseWriter, r *http.Request, status int, msg string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "%v: %v\n", status, msg)
}

func addHandler(w http.ResponseWriter, r *http.Request, q *core.TuneQueue) {
	log.Info("hit /add")
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	url := r.Form.Get("url")
	if url == "" {

		log.Info("no url provided")
		reply(w, r, 400, "url not provided")
	}
	q.Add(url)
	log.Infof("added %v to the queue", url)
	reply(w, r, 200, "success")
}

func nextHandler(w http.ResponseWriter, r *http.Request, q *core.TuneQueue) {
	q.Next()
	reply(w, r, 200, "success")
}

func RunServer(dir, player, port string) {
	q := core.NewTuneQueue(dir, player)
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) { addHandler(w, r, q) })
	http.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) { nextHandler(w, r, q) })
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
