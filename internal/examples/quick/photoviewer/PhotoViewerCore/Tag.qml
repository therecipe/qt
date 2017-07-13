/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the QtQml module of the Qt Toolkit.
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

Flipable {
    id: flipable

    property alias frontLabel: frontButton.label
    property alias backLabel: backButton.label

    property int angle: 0
    property int randomAngle: Math.random() * (2 * 6 + 1) - 6
    property bool flipped: false

    signal frontClicked
    signal backClicked
    signal tagChanged(string tag)

    front: EditableButton {
        id: frontButton; rotation: flipable.randomAngle
        anchors { centerIn: parent; verticalCenterOffset: -20 }
        onClicked: flipable.frontClicked()
        onLabelChanged: flipable.tagChanged(label)
    }

    back: Button {
        id: backButton; tint: "red"; rotation: flipable.randomAngle
        anchors { centerIn: parent; verticalCenterOffset: -20 }
        onClicked: flipable.backClicked()
    }

    transform: Rotation {
        origin.x: flipable.width / 2; origin.y: flipable.height / 2
        axis.x: 0; axis.y: 1; axis.z: 0
        angle: flipable.angle
    }

    states: State {
        name: "back"; when: flipable.flipped
        PropertyChanges { target: flipable; angle: 180 }
    }

    transitions: Transition {
        ParallelAnimation {
            NumberAnimation { properties: "angle"; duration: 400 }
            SequentialAnimation {
                NumberAnimation { target: flipable; property: "scale"; to: 0.8; duration: 200 }
                NumberAnimation { target: flipable; property: "scale"; to: 1.0; duration: 200 }
            }
        }
    }
}
