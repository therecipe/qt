/****************************************************************************
**
** Copyright (C) 2013 BlackBerry Limited. All rights reserved.
** Copyright (C) 2017 The Qt Company Ltd.
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
import Qt.labs.folderlistmodel 1.0

Item {
    Rectangle {
        id: title
        opacity: 0.7
        height: titleLabel.height + 90
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.top: parent.top
        anchors.margins: 70

        color: "#5c9fba"
        Text {
            id: titleLabel
            text: "Select picture"
            color: "white"
            font.bold: true
            font.pointSize: 15
            anchors.centerIn: parent
        }
    }

    ListView {
        id: listView
        clip: true
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        anchors.top: title.bottom
        anchors.topMargin: 0
        anchors.margins: 70
        spacing: 5
        add: Transition {
            NumberAnimation { properties: "x"; from: 1000; duration: 500 }
        }

//! [FileSelect-1]
        model: FolderListModel {
            folder: "file://"+SystemPictureFolder
            showDirs: false
        }

        delegate: Rectangle {
//! [FileSelect-1]
            opacity: 0.7
            height: label.height + 130
            width: listView.width + 2
            x: -1
            border.color: Qt.lighter("#67b0d1")
            border.width: 1
            color: mArea.pressed ? "#5c9fba" : "#67b0d1"
            Image {
                id: picture
                anchors.verticalCenter: parent.verticalCenter
                height: parent.height-10
                 fillMode: Image.PreserveAspectFit
                asynchronous: true
                source: "file://" + model.filePath
            }

//! [FileSelect-2]
            Text {
//! [FileSelect-2]
                id: label
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right
                anchors.margins: 20
                anchors.left: picture.right
//! [FileSelect-3]
                text: model.fileName
//! [FileSelect-3]
                font.bold: true
                font.pointSize: 10
                color: "white"
                wrapMode: Text.WordWrap
//! [FileSelect-4]
            }
            MouseArea {
                id: mArea
                anchors.fill: parent
                onClicked: {
                    print ("path: " + model.filePath + " " + model.fileName)
                    root.fileName = model.filePath
                }
            }
        }
//! [FileSelect-4]
    }
}
