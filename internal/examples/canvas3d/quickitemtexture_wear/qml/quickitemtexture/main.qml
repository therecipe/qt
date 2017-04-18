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

import QtQuick 2.2
import QtCanvas3D 1.1
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.1
import QtQuick.Window 2.2

import "quickitemtexture.js" as GLCode

Window {
    id: mainview
    width: 800
    height: 600
    visible: true
    title: "Qt Quick Item as Texture"
    color: "#f9f9f9"

    ColumnLayout {
        Layout.fillWidth: true
        x: 4
        y: 4
        //! [0]
        Rectangle {
            id: textureSource
            color: "lightgreen"
            width: 256
            height: 256
            border.color: "blue"
            border.width: 4
            layer.enabled: true
            layer.smooth: true
            Label {
                anchors.fill: parent
                anchors.margins: 16
                text: "X Rot:" + (canvas3d.xRotAnim | 0) + "\n"
                    + "Y Rot:" + (canvas3d.yRotAnim | 0) + "\n"
                    + "Z Rot:" + (canvas3d.zRotAnim | 0) + "\n"
                    + "FPS:" + canvas3d.fps
                color: "red"
                font.pointSize: 26
                horizontalAlignment: Text.AlignLeft
                verticalAlignment: Text.AlignVCenter
            }
        }
        //! [0]
        Button {
            Layout.fillWidth: true
            Layout.minimumWidth: 256
            text: textureSource.visible ? "Hide texture source" : "Show texture source"
            onClicked: textureSource.visible = !textureSource.visible
        }
        Button {
            Layout.fillWidth: true
            Layout.minimumWidth: 256
            text: "Quit"
            onClicked: Qt.quit()
        }
    }

    Canvas3D {
        id: canvas3d
        anchors.fill:parent
        focus: true
        property double xRotAnim: 0
        property double yRotAnim: 0
        property double zRotAnim: 0
        property bool isRunning: true

        // Emitted when one time initializations should happen
        onInitializeGL: {
            GLCode.initializeGL(canvas3d, textureSource);
        }

        // Emitted each time Canvas3D is ready for a new frame
        onPaintGL: {
            GLCode.paintGL(canvas3d);
        }

        onResizeGL: {
            GLCode.resizeGL(canvas3d);
        }

        Keys.onSpacePressed: {
            canvas3d.isRunning = !canvas3d.isRunning
            if (canvas3d.isRunning) {
                objAnimationX.pause();
                objAnimationY.pause();
                objAnimationZ.pause();
            } else {
                objAnimationX.resume();
                objAnimationY.resume();
                objAnimationZ.resume();
            }
        }

        SequentialAnimation {
            id: objAnimationX
            loops: Animation.Infinite
            running: true
            NumberAnimation {
                target: canvas3d
                property: "xRotAnim"
                from: 0.0
                to: 120.0
                duration: 7000
                easing.type: Easing.InOutQuad
            }
            NumberAnimation {
                target: canvas3d
                property: "xRotAnim"
                from: 120.0
                to: 0.0
                duration: 7000
                easing.type: Easing.InOutQuad
            }
        }

        SequentialAnimation {
            id: objAnimationY
            loops: Animation.Infinite
            running: true
            NumberAnimation {
                target: canvas3d
                property: "yRotAnim"
                from: 0.0
                to: 240.0
                duration: 5000
                easing.type: Easing.InOutCubic
            }
            NumberAnimation {
                target: canvas3d
                property: "yRotAnim"
                from: 240.0
                to: 0.0
                duration: 5000
                easing.type: Easing.InOutCubic
            }
        }

        SequentialAnimation {
            id: objAnimationZ
            loops: Animation.Infinite
            running: true
            NumberAnimation {
                target: canvas3d
                property: "zRotAnim"
                from: -100.0
                to: 100.0
                duration: 3000
                easing.type: Easing.InOutSine
            }
            NumberAnimation {
                target: canvas3d
                property: "zRotAnim"
                from: 100.0
                to: -100.0
                duration: 3000
                easing.type: Easing.InOutSine
            }
        }
    }
}
