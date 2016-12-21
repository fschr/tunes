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
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type tune struct {
	path    string
	url     string
	loadCmd *exec.Cmd
	playCmd *exec.Cmd
}

func NewTune(url, dir, player string) *tune {
	id, err := idFromURL(url)
	if err != nil {
		log.Warn(err)
		id = "invalid"
	}
	path := fmt.Sprintf("%v.mp3", filepath.Join(dir, id))
	outputString := fmt.Sprintf("--output=\"%v\"", path)
	urlString := fmt.Sprintf("\"%v\"", url)
	loadArgs := []string{urlString, "--format=bestaudio/best", "-x", "--audio-format=mp3", outputString}
	return &tune{
		path:    path,
		url:     url,
		loadCmd: exec.Command("youtube-dl", loadArgs...),
		playCmd: exec.Command(player, []string{path}...),
	}
}

func (t *tune) AsyncLoad() error {
	if err := t.loadCmd.Start(); err != nil {
		return err
	}
	return nil
}

func (t *tune) WaitOnLoad() error {
	if err := t.loadCmd.Wait(); err != nil {
		return err
	}
	return nil
}

func (t *tune) AsyncPlay() error {
	if err := t.playCmd.Start(); err != nil {
		return err
	}
	return nil
}

func (t *tune) Stop() error {
	if err := t.playCmd.Process.Kill(); err != nil {
		return err
	}
	return nil
}

type tuneNode struct {
	t          *tune
	next, last *tuneNode
}

type TuneQueue struct {
	dir, player string
	first, last *tuneNode
}

func NewTuneQueue(dir, player string) *TuneQueue {
	return &TuneQueue{
		dir:    dir,
		player: player,
		first:  nil,
		last:   nil,
	}
}

func (tq *TuneQueue) Add(url string) {

	tq.last = &tuneNode{
		t:    NewTune(url, tq.dir, tq.player),
		last: tq.last,
		next: nil,
	}
	if tq.last.last != nil {
		tq.last.last.next = tq.last
	}
	if err := tq.last.t.AsyncLoad(); err != nil {
		log.Warn(err)
	}
	if tq.first == nil {
		t := tq.last
		for t.next != nil {
			t = t.next
		}
		tq.first = t
		if err := tq.first.t.WaitOnLoad(); err != nil {
			log.Warn(err)
		}
		if err := tq.first.t.AsyncPlay(); err != nil {
			log.Warn(err)
		}
	}
}

func (tq *TuneQueue) Next() {
	if tq.first == nil {
		log.Warn("tq.first is nil or the queue is empty!")
		return
	}
	if err := tq.first.t.Stop(); err != nil {
		log.Warn(err)
	}
	tq.first = tq.first.next
	tq.first.last = nil
	if err := tq.first.t.AsyncPlay(); err != nil {
		log.Warn(err)
	}
}
