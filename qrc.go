package main

// This file is automatically generated by gopkg.in/qml.v1/cmd/genqrc

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/qml.v1"
)

func init() {
	var r *qml.Resources
	var err error
	if os.Getenv("QRC_REPACK") == "1" {
		err = qrcRepackResources()
		if err != nil {
			panic("cannot repack qrc resources: " + err.Error())
		}
		r, err = qml.ParseResources(qrcResourcesRepacked)
	} else {
		r, err = qml.ParseResourcesString(qrcResourcesData)
	}
	if err != nil {
		panic("cannot parse bundled resources data: " + err.Error())
	}
	qml.LoadResources(r)
}

func qrcRepackResources() error {
	subdirs := []string{"assets"}
	var rp qml.ResourcesPacker
	for _, subdir := range subdirs {
		err := filepath.Walk(subdir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			rp.Add(filepath.ToSlash(path), data)
			return nil
		})
		if err != nil {
			return err
		}
	}
	qrcResourcesRepacked = rp.Pack().Bytes()
	return nil
}

var qrcResourcesRepacked []byte
var qrcResourcesData = "qres\x00\x00\x00\x01\x00\x00\x05\a\x00\x00\x00\x14\x00\x00\x04\xdf\x00\x00\x04\xc7import QtQuick 2.2\nimport QtQuick.Window 2.0\nimport QtQuick.Controls 1.1\n\nApplicationWindow {\n    id: mainwindow\n    flags: Qt.FramelessWindowHint | Qt.WindowStaysOnTopHint\n    visible: true\n    title: qsTr(\"Network Speed Monitor\")\n    width: 80\n    height: 37\n    x: (Screen.width - width)\n    y: (Screen.height - height)\n    color: \"transparent\"\n\n    Rectangle {\n        anchors.fill: parent\n        radius: 2\n        opacity: 0.7\n        color: \"black\"\n    }\n\n    MouseArea {\n        id: mouseRegion\n        anchors.fill: parent;\n\n        // 用于记录鼠标开始拖动的位置\n        property variant clickPos: \"0, 0\"\n\n        onPressed: {\n            clickPos  = Qt.point(mouse.x, mouse.y)\n        }\n\n        onPositionChanged: {\n            mainwindow.x = mainwindow.x + mouse.x - clickPos.x\n            mainwindow.y = mainwindow.y + mouse.y - clickPos.y\n        }\n\n        Text {\n            id: up\n            x: 3\n            color: \"white\"\n            text: speed.up\n            font.pixelSize: 13\n        }\n        Text {\n            id: down\n            x: 3\n            y: up.y + font.pixelSize + 3\n            color: \"white\"\n            text: speed.down\n            font.pixelSize: 13\n        }\n    }\n}\n\n\x00\x06\x06\x8a\x9c\xb3\x00a\x00s\x00s\x00e\x00t\x00s\x00\b\b\x01Z\\\x00m\x00a\x00i\x00n\x00.\x00q\x00m\x00l\x00\x00\x00\x00\x00\x02\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x02\x00\x00\x00\x01\x00\x00\x00\x02\x00\x00\x00\x12\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00"