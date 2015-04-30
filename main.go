package main

//go:generate genqrc assets

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/qml.v1"
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

	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()
	component, err := engine.LoadFile("qrc:///assets/main.qml")
	if err != nil {
		return err
	}

	window := component.CreateWindow(nil)
	root := window.Root()
	upText:=root.ObjectByName("upText")
	downText:=root.ObjectByName("downText")
	window.Show()

	go func() {
		var (
			oldNetIn  = getNetIn()
			oldNetOut = getNetOut()
			ticker    = time.NewTicker(time.Second * 1)
		)
		for _ = range ticker.C {
			netIn := getNetIn()
			netOut := getNetOut()
			upText.Set("text", `<font color="green">⇧</font> ` + parseSize(netIn-oldNetIn))
			downText.Set("text", `<font color="red">⇩</font> ` + parseSize(netOut-oldNetOut))
			oldNetIn = netIn
			oldNetOut = netOut
		}
	}()

	window.Wait()

	return nil
}
