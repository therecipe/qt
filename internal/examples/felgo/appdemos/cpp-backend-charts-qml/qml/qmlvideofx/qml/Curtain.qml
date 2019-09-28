/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the Qt Mobility Components.
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

Rectangle {
    id: root
    color: "transparent"
    radius: 5
    property alias value: grip.value
    property color gripColor: "transparent"
    property real gripSize: 20
    property real gripTolerance: 3.0
    property real increment: 0.1
    property bool enabled: true
    property string imageSource: "../../../assets/qmlvideofx/Triangle_Top.png"

    Rectangle {
        id: grip
        property real value: 0.5
        x: (value * parent.width) - width/2
        anchors.verticalCenter: parent.verticalCenter
        width: root.gripTolerance * root.gripSize
        height: width
        radius: width/2
        color: "transparent"

        Image {
            id: sliderhandleimage
            source: imageSource
            anchors.centerIn: parent
            width: root.gripSize
            height: width / sourceSize.width * sourceSize.height
        }

        MouseArea {
            id: mouseArea
            enabled: root.enabled
            anchors.fill:  parent
            drag {
                target: grip
                axis: Drag.XAxis
                minimumX: -parent.width/2
                maximumX: root.width - parent.width/2
            }
            onPositionChanged:  {
                if (drag.active)
                    updatePosition()
            }
            onReleased: {
                updatePosition()
            }
            function updatePosition() {
                value = (grip.x + grip.width/2) / grip.parent.width
            }
        }
    }
}
