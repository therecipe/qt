/***************************************************************************
**
** Copyright (C) 2014 BlackBerry Limited. All rights reserved.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the examples of the QtBluetooth module of the Qt Toolkit.
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

import QtQuick 2.0
import QtQuick.Window 2.1

Rectangle {
    id: board
    width: 600
    height: 300

    // 1 - server role; left pedal
    // 2 - client role; right pedal
    property int roleSide: pingPong.role
    onRoleSideChanged: {
        if (pingPong.role == 1) {
            rightMouse.opacity = 0.7
            rightMouse.enabled = false
        }
        else if (pingPong.role == 2) {
            leftMouse.opacity = 0.7
            leftMouse.enabled = false
        }
    }

    property bool deviceMessage: pingPong.showDialog
    onDeviceMessageChanged: {
        if (pingPong.showDialog) {
            info.visible = true;
            board.opacity = 0.5;
        } else {
            info.visible = false;
            board.opacity = 1;
        }
    }

    // Left pedal - server role
    Rectangle {
        id: leftblock
        y: (parent.height/2)
        width: (parent.width/27)
        height: (parent.height/5)
        anchors.left: parent.left
        color: "#363636"
        radius: 10

        MouseArea {
            id: leftMouse
            anchors.fill: parent
            acceptedButtons: Qt.LeftButton
            drag.target: leftblock
            drag.axis: Drag.YAxis
            drag.minimumY: 0
            drag.maximumY: (board.height - leftblock.height)
        }
    }

    // Right pedal - client role
    Rectangle {
        id: rightblock
        y: (parent.height/2)
        width: (parent.width/27)
        height: (parent.height/5)
        anchors.right: parent.right
        color: "#363636"
        radius: 10

        MouseArea {
            id: rightMouse
            anchors.fill: parent
            acceptedButtons: Qt.LeftButton
            drag.target: rightblock
            drag.axis: Drag.YAxis
            drag.minimumY: 0
            drag.maximumY: (board.height - rightblock.height)
        }
    }

    property double leftBlockY: leftblock.y
    onLeftBlockYChanged: pingPong.updateLeftBlock(leftblock.y)

    property double leftBlockUpdate: pingPong.leftBlockY
    onLeftBlockUpdateChanged: leftblock.y = pingPong.leftBlockY

    property double rightBlockY: rightblock.y
    onRightBlockYChanged: pingPong.updateRightBlock(rightblock.y)

    property double rightBlockUpdate: pingPong.rightBlockY
    onRightBlockUpdateChanged: rightblock.y = pingPong.rightBlockY

    Rectangle {
        id: splitter
        color: "#363636"
        anchors.horizontalCenter: parent.horizontalCenter
        height: parent.height
        width: parent.width/100
    }

    Text {
        id: leftResult
        text: pingPong.leftResult
        font.bold: true
        font.pixelSize: 30
        anchors.right: splitter.left
        anchors.top: parent.top
        anchors.margins: 15
    }

    Text {
        id: rightResult
        text: pingPong.rightResult
        font.bold: true
        font.pixelSize: 30
        anchors.left: splitter.right
        anchors.top: parent.top
        anchors.margins: 15
    }

    Rectangle {
        id: ball
        width: leftblock.width/2
        height: leftblock.width/2
        radius: width
        color: "#363636"
        x: pingPong.ballX
        y: pingPong.ballY

        SequentialAnimation {
            running: true
            NumberAnimation { target: ball; property: "x"; duration: 50 }
            NumberAnimation { target: ball; property: "y"; duration: 50 }
        }
    }

    Component.onCompleted: {
        if (menulist.height == Screen.height && menulist.width == Screen.width)
            pingPong.setSize(Screen.width, Screen.height)
        else
            pingPong.setSize(board.width, board.height)
        pingPong.updateLeftBlock(leftblock.y)
        pingPong.updateRightBlock(rightblock.y)
    }
}
