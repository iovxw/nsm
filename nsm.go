/*
 Copyright 2015 Bluek404

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/conformal/gotk3/gtk"
)

var net string

func readIntFromFile(path string) (int, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	b = b[:len(b)-1] // delete "\n"
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func getNetIn() int {
	i, err := readIntFromFile("/sys/class/net/" + net + "/statistics/rx_bytes")
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func getNetOut() int {
	i, err := readIntFromFile("/sys/class/net/" + net + "/statistics/tx_bytes")
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func parseSize(src int) string {
	if src == 0 {
		return "0 B/s"
	}
	var unit string

	for i := 0; ; i++ {
		size := 1 << uint(i*10)
		if src < size {
			size = 1 << uint((i-1)*10)
			src = src / size
			switch i - 1 {
			case 0:
				unit = "B"
			case 1:
				unit = "KB"
			case 2:
				unit = "MB"
			case 3:
				unit = "GB"
			}
			break
		}
	}
	return strconv.Itoa(src) + " " + unit + "/s"
}

func main() {
	if len(os.Args) == 1 {
		net = "wlan0"
	} else {
		net = os.Args[1]
	}

	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetSkipTaskbarHint(true)
	win.SetDecorated(false)
	win.SetKeepAbove(true)

	l, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	l2, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		log.Fatal(err)
	}

	box.PackStart(l, false, true, 0)
	box.PackStart(l2, false, true, 0)

	win.Add(box)
	win.ShowAll()

	ticker := time.NewTicker(time.Second * 1)
	go func() {
		oldNetIn := getNetIn()
		oldNetOut := getNetOut()
		for _ = range ticker.C {
			netIn := getNetIn()
			netOut := getNetOut()
			l.SetMarkup("<span foreground='red'>⇩</span> " + parseSize(netIn-oldNetIn))
			l2.SetMarkup("<span foreground='green'>⇧</span> " + parseSize(netOut-oldNetOut))
			oldNetIn = netIn
			oldNetOut = netOut
		}
	}()

	gtk.Main()
}
