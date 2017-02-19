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

import QtQuick 2.0
import QtQuick.Layouts 1.1

Item {
    id: infoSheet
    anchors.topMargin: 59
    anchors.leftMargin: 48
    property alias headingText1: heading1.text
    property alias headingText2: heading2.text
    property alias text: description.text
    state: "hidden"

    states: [
        State {
            name: "hidden"
            when: !visible
            PropertyChanges { target: heading1; color: "#00000000"; }
            PropertyChanges { target: heading2; color: "#005caa15"; }
            PropertyChanges { target: description; color: "#00000000"; }
            PropertyChanges { target: infoSheet; visible: false; }
        },
        State {
            name: "visible"
            when: visible
            PropertyChanges { target: heading1; color: "#FF000000"; }
            PropertyChanges { target: heading2; color: "#FF5caa15"; }
            PropertyChanges { target: description; color: "#FF000000"; }
            PropertyChanges { target: infoSheet; visible: true; }
        }
    ]

    transitions: [
        Transition {
            from: "visible"
            to: "hidden"
            ColorAnimation {
                properties: "color"
                easing.type: Easing.InOutCubic
                duration: 300
            }
            PropertyAnimation {
                properties: "visible"
                duration: 300
            }
        },
        Transition {
            from: "hidden"
            to: "visible"
            ColorAnimation {
                properties: "color"
                easing.type: Easing.InOutCubic
                duration: 300
            }
        }
    ]

    ColumnLayout {
        Text {
            id: heading1
            text: ""
            font.family: "Helvetica"
            font.pointSize: 36
            font.weight: Font.Light
            color: "black"

            Text {
                id: heading2
                anchors.top: heading1.top
                anchors.left: heading1.right
                text: ""
                font.family: "Helvetica"
                font.pointSize: 36
                font.weight: Font.Light
                color: "#5caa15"
            }
        }

        Item {
            id: layoutHelper
            Layout.minimumHeight: 37
            Layout.preferredHeight: 37
            Layout.maximumHeight: 37
        }

        Text {
            id: description
            text: ""
            width: (infoSheet.width - infoSheet.anchors.leftMargin) * 0.3
            font.family: "Helvetica"
            font.pointSize: 16
            font.weight: Font.Light
        }
    }
}
