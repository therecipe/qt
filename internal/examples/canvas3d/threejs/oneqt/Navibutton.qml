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

Text {
    id: menubarItem
    text: ""
    font.family: "Helvetica"
    font.pointSize: 20
    font.weight: Font.Light
    color: "#404244"
    Layout.alignment: Qt.AlignHCenter
    state: "default"
    verticalAlignment: Text.AlignVCenter
    property bool selected: false
    property string stateSelect: ""
    property Item stateTarget

    states: [
        State {
            name: "default"
            when: !menubarItem.selected
            PropertyChanges { target: menubarItem; color: "#404244"; }
            PropertyChanges { target: selectedHighlight; height: 0; }
        },
        State {
            name: "mouseover"
            PropertyChanges { target: menubarItem; color: "#5caa15"; }
            PropertyChanges { target: selectedHighlight; height: 4; }
        },
        State {
            name: "selected"
            when: menubarItem.selected
            PropertyChanges { target: menubarItem; color: "#5caa15"; }
            PropertyChanges { target: selectedHighlight; height: 4; }
        }
    ]

    onStateChanged: {
        if (state == "selected") {
            menubarItem.stateTarget.state = menubarItem.stateSelect;
        }
    }

    transitions: [
        Transition {
            from: "*"
            to: "*"
            ColorAnimation {
                properties: "color"
                easing.type: Easing.InOutCubic
                duration: 250
            }
        }
    ]

    Rectangle {
        id: selectedHighlight
        color: "#7cc54d"
        width: parent.width
        anchors.bottom: parent.bottom
        height: 0

        Behavior on height {

                    NumberAnimation {
                        easing.type: Easing.InOutCubic
                        duration: 250
                    }
                }
    }

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        hoverEnabled: true

        onPressed: {
            menubarItem.state = "mouseover";
        }

        onReleased: {
            if (menubarItem.state == "mouseover") {
                menubarItem.selected = true;
            }
        }

        onEntered: {
            menubarItem.state = "mouseover";
        }

        onExited: {
            if (menubarItem.selected) {
                menubarItem.state = "selected";
            } else {
                menubarItem.state = "default";
            }
        }
    }
}

