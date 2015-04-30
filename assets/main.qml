import QtQuick 2.2
import QtQuick.Window 2.0
import QtQuick.Controls 1.1

ApplicationWindow {
    id: mainwindow
    flags: Qt.FramelessWindowHint | Qt.WindowStaysOnTopHint
    visible: true
    title: qsTr("Network Speed Monitor")
    width: 80
    height: 37
    x: (Screen.width - width)
    y: (Screen.height - height)
    color: "transparent"

    Rectangle {
        anchors.fill: parent
        radius: 2
        opacity: 0.7
        color: "black"
    }

    MouseArea {
        id: mouseRegion
        anchors.fill: parent;

        // 用于记录鼠标开始拖动的位置
        property variant clickPos: "0, 0"

        onPressed: {
            clickPos  = Qt.point(mouse.x, mouse.y)
        }

        onPositionChanged: {
            mainwindow.x = mainwindow.x + mouse.x - clickPos.x
            mainwindow.y = mainwindow.y + mouse.y - clickPos.y
        }

        Text {
            id: up
            objectName: "upText"
            x: 3
            color: "white"
            text: '<font color="green">⇧</font> 0 B/s'
            font.pixelSize: 13
        }
        Text {
            id: down
            objectName: "downText"
            x: 3
            y: up.y + font.pixelSize + 3
            color: "white"
            text: '<font color="red">⇩</font> 0 B/s'
            font.pixelSize: 13
        }
    }
}

