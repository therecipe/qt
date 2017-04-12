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

Item {
    id: fpsDisplayControl
    property bool hidden: true
    property real fps: 0.0

    onHiddenChanged: {
        if (fpsDisplayControl.hidden)
            fpsDisplay.color = "transparent";
        else
            fpsDisplay.color = "#000000FF";
    }

    onFpsChanged: {
        fpsDisplay.updateFps();
    }

    Rectangle {
        anchors.fill: parent
        id: fpsDisplay
        color: "transparent"

        property real maxFps: 60.0
        property color maxFpsColor: "#00FF00"
        property color minFpsColor: "#FF0000"

        function updateFps() {
            var scale = (fps > maxFps)?1.0:(fps/maxFps);
            var r = (1 - scale) * minFpsColor.r + scale * maxFpsColor.r;
            var g = (1 - scale) * minFpsColor.g + scale * maxFpsColor.g;
            var b = (1 - scale) * minFpsColor.b + scale * maxFpsColor.b;
            var a = (1 - scale) * minFpsColor.a + scale * maxFpsColor.a;
            fpsCauge.height = scale * fpsDisplay.height;
            fpsCauge.color = Qt.rgba(r,g,b,a);
        }

        Rectangle {
            id: fpsCauge
            width: parent.width
            anchors.bottom: parent.bottom
            visible: !fpsDisplayControl.hidden
        }

        Text {
            id: fpsText
            text: ""+(fps | 0)
            font.family: "Helvetica"
            font.pixelSize: 16
            font.weight: Font.Light
            color: "white"
            anchors.fill: parent
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            visible: !fpsDisplayControl.hidden
        }
    }
    MouseArea {
        anchors.fill: parent
        onClicked: {
            fpsDisplayControl.hidden = !fpsDisplayControl.hidden;
        }
    }
}

