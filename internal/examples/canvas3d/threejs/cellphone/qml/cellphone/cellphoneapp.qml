/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the QtCanvas3D module of the Qt Toolkit.
**
** $QT_BEGIN_LICENSE:BSD$
** Commercial License Usage
** Licensees holding valid commercial Qt licenses may use this file in
** accordance with the commercial license agreement provided with the
** Software or, alternatively, in accordance with the terms contained in
** a written agreement between you and The Qt Company. For licensing terms
** and conditions see https://www.qt.io/terms-conditions. For further
** information use the contact form at https://www.qt.io/contact-us.
**
** BSD License Usage
** Alternatively, you may use this file under the terms of the BSD license
** as follows:
**
** "Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions are
** met:
**   * Redistributions of source code must retain the above copyright
**     notice, this list of conditions and the following disclaimer.
**   * Redistributions in binary form must reproduce the above copyright
**     notice, this list of conditions and the following disclaimer in
**     the documentation and/or other materials provided with the
**     distribution.
**   * Neither the name of The Qt Company Ltd nor the names of its
**     contributors may be used to endorse or promote products derived
**     from this software without specific prior written permission.
**
**
** THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
** "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
** LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
** A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
** OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
** SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
** LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
** OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE."
**
** $QT_END_LICENSE$
**
****************************************************************************/

import QtQuick 2.6
import QtQuick.Controls 1.2

Rectangle {
    id: mainScreen
    property color backgroundColor
    property color textColor
    signal locked
    color: backgroundColor
    // The state indicates which app is to be shown on the screen.
    state: "menu"
    states: [
        State {
            name: "menu"
            PropertyChanges {
                target: menuApp; visible: true
            }
        },
        State {
            name: "about"
            PropertyChanges {
                target: aboutApp; visible: true
            }
        },
        State {
            name: "dummy"
            PropertyChanges {
                target: dummyApp; visible: true
            }
        }
    ]

    // This model contains names, icons, and corresponsing mainScreen states of phone apps
    ListModel {
        id: appGridModel
        ListElement {
            name: "Calendar"
            image: "qrc:/calendar.png"
            toState: "dummy"
        }
        ListElement {
            name: "Camera"
            image: "qrc:/camera.png"
            toState: "dummy"
        }
        ListElement {
            name: "Clock"
            image: "qrc:/clock.png"
            toState: "dummy"
        }
        ListElement {
            name: "Contacts"
            image: "qrc:/contacts.png"
            toState: "dummy"
        }
        ListElement {
            name: "Gallery"
            image: "qrc:/gallery.png"
            toState: "dummy"
        }
        ListElement {
            name: "Games"
            image: "qrc:/games.png"
            toState: "dummy"
        }
        ListElement {
            name: "Mail"
            image: "qrc:/mail.png"
            toState: "dummy"
        }
        ListElement {
            name: "Maps"
            image: "qrc:/maps.png"
            toState: "dummy"
        }
        ListElement {
            name: "Music"
            image: "qrc:/music.png"
            toState: "dummy"
        }
        ListElement {
            name: "Settings"
            image: "qrc:/settings.png"
            toState: "dummy"
        }
        ListElement {
            name: "Todo"
            image: "qrc:/todo.png"
            toState: "dummy"
        }
        ListElement {
            name: "Videos"
            image: "qrc:/videos.png"
            toState: "dummy"
        }
        ListElement {
            name: "About"
            image: "qrc:/qtlogo_with_alpha.png"
            toState: "about"
        }
        ListElement {
            name: "Lock"
            image: "qrc:/lock.png"
            toState: "lock"
        }
    }

    Timer {
        id: clockTimer
        interval: 1000
        repeat: true
        running: lockScreen.visible || menuApp.visible
        triggeredOnStart: true
        onTriggered: clockLabel.text = Qt.formatDateTime(new Date(), "hh:mm:ss")
    }

    // This grid shows the app menu on the phone
    Rectangle {
        id: menuApp
        visible: false
        anchors.fill: parent
        color: mainScreen.color
        Image {
            anchors.fill: parent
            source: "qrc:/menu_background.jpg"
        }
        Rectangle {
            id: menuHeader
            anchors.top: parent.top
            anchors.left: parent.left
            anchors.right: parent.right
            anchors.margins: 16
            height: parent.height / 24
            color: "transparent"
            Label {
                id: menuClockLabel
                anchors.top: parent.top
                anchors.left: parent.left
                anchors.bottom: parent.bottom
                width: menuHeader.width
                text: clockLabel.text
                color: mainScreen.textColor
                font.pixelSize: 24
                horizontalAlignment: Text.AlignLeft
                verticalAlignment: Text.AlignVCenter
            }
        }
        GridView {
            id: menu
            anchors.top: menuHeader.bottom
            anchors.bottom: parent.bottom
            anchors.left: parent.left
            anchors.right: parent.right
            anchors.margins: 16
            model: appGridModel
            cellWidth: width / 3
            cellHeight: cellWidth
            boundsBehavior: Flickable.StopAtBounds
            delegate: Component {
                id: appGridDelegate
                Item {
                    id: wrapper
                    width: menu.cellWidth
                    height: menu.cellHeight
                    Column {
                        anchors.fill: parent
                        height: width
                        Image {
                            width: (parent.width * 3) / 4
                            height : width
                            source: image
                            anchors.horizontalCenter: parent.horizontalCenter
                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    // Lock option is a special case, as it doesn't launch
                                    // an app. Instead it simply slides the lock screen over
                                    // the current app.
                                    if (toState === "lock") {
                                        mainScreen.lock()
                                    } else if (toState === "dummy"){
                                        dummyApp.dummyImage = image
                                        mainScreen.state = toState
                                    } else {
                                        mainScreen.state = toState
                                        mainScreen.resetLockTimer()
                                    }
                                }
                            }
                        }
                        Text {
                            width: parent.width
                            height: parent.width / 4
                            text: name
                            font.pixelSize: 20
                            color: mainScreen.textColor
                            horizontalAlignment: Text.AlignHCenter
                            anchors.horizontalCenter: parent.horizontalCenter
                        }
                    }
                }
            }
        }
    }

    Rectangle {
        id: aboutApp
        anchors.fill: parent
        anchors.margins: 16
        color: mainScreen.backgroundColor
        visible: false
        Label {
            anchors.fill: parent
            text: "This example demonstrates how to use Qt Quick Item as a texture source"
                  + " for Canvas3D in conjunction with three.js."
                  + "\n\nClick outside the phone to restart the rotation animation."
            color: mainScreen.textColor
            font.pixelSize: 24
            wrapMode: Text.WordWrap
            horizontalAlignment: Text.AlignHCenter
            verticalAlignment: Text.AlignVCenter
        }
        MouseArea {
            anchors.fill: parent
            onClicked: {
                mainScreen.state = "menu"
                mainScreen.resetLockTimer()
            }
        }
    }

    Rectangle {
        id: dummyApp
        property string dummyImage
        anchors.fill: parent
        anchors.margins: 16
        color: mainScreen.backgroundColor
        visible: false
        Image {
            anchors.centerIn: parent
            width: parent.width / 2
            height: width
            source: parent.dummyImage
        }
        MouseArea {
            anchors.fill: parent
            onClicked: {
                mainScreen.state = "menu"
                mainScreen.resetLockTimer()
            }
        }
    }

    Flickable {
        id: lockScreen
        property int flickSpeed: 1500

        anchors.fill: parent
        z: mainScreen.z + 0.1
        contentWidth: clock.width
        contentHeight: clock.height
        contentX: 0
        boundsBehavior: Flickable.StopAtBounds
        flickableDirection: Flickable.HorizontalFlick
        flickDeceleration: 0.1
        onMovementEnded: {
            if (contentX == width) {
                visible = false
                mainScreen.resetLockTimer()
            } else if (contentX != 0) {
                flick(-flickSpeed, 0)
            }
        }
        onFlickEnded: {
            if (contentX == width) {
                visible = false
                mainScreen.resetLockTimer()
            }
        }

        Timer {
            id: lockTimer
            interval: 60000
            repeat: false
            onTriggered: mainScreen.lock()
        }

        Item {
            id: clock
            width: lockScreen.width * 2
            height: lockScreen.height
            Rectangle {
                anchors.top: parent.top
                anchors.bottom: parent.bottom
                x: 0
                width: lockScreen.width
                color: mainScreen.backgroundColor
                Label {
                    id: clockLabel
                    anchors.top: parent.top
                    anchors.left: parent.left
                    anchors.bottom: parent.bottom
                    width: lockScreen.width
                    text: Qt.formatDateTime(new Date(), "hh:mm:ss")
                    color: mainScreen.textColor
                    font.pixelSize: 60
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }
                Label {
                    anchors.left: parent.left
                    anchors.bottom: parent.bottom
                    height: 48
                    width: lockScreen.width
                    text: "Swipe left to unlock"
                    color: mainScreen.textColor
                    font.pixelSize: 24
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                }
            }
        }
    }

    function lock() {
        lockTimer.stop()
        lockScreen.visible = true
        lockScreen.flick(lockScreen.flickSpeed, 0)
        locked() // emit
    }

    function resetLockTimer() {
        lockTimer.restart()
    }
}
