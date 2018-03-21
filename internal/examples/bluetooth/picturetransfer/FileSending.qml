/****************************************************************************
**
** Copyright (C) 2013 BlackBerry Limited. All rights reserved.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the QtBluetooth module of the Qt Toolkit.
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

import QtQuick 2.1

Item {
    Button {
        x: 40
        y: 40
        text: "Back"
        onClicked: loader.source = "DeviceDiscovery.qml"
    }

    Text {
        anchors.bottom: theProgress.top
        anchors.left: theProgress.left
        anchors.right: parent.right
        anchors.rightMargin: 100
        anchors.bottomMargin: 40
        wrapMode: Text.Wrap
        color: "white"
        text: "Sending file: " + root.fileName + "<br>" + "Remote-Device:" + root.remoteDevice
    }

    Rectangle {
        id: theProgress
        anchors.left: parent.left
        anchors.verticalCenter: parent.verticalCenter
        anchors.leftMargin: 100
        width: (parent.width - 200) * fileTransfer.progress
        height: 50
        color: "#5c9fba"
        opacity: 0.7
        Behavior on width {
            NumberAnimation {duration: 300}
        }
    }
    Text {
        text: "0%"
        anchors.horizontalCenter: theProgress.left
        anchors.top: theProgress.bottom
        anchors.topMargin: 10
        color: "white"
    }

    Text {
        text: "100%"
        anchors.right: parent.right
        anchors.rightMargin: 100 - width/2
        anchors.top: theProgress.bottom
        anchors.topMargin: 10
        color: "white"
    }

    Image {
        anchors.top: theProgress.bottom
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        anchors.margins: 60
        fillMode: Image.PreserveAspectFit
        source: "file://" + root.fileName
    }
}
