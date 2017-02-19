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
import QtQuick.Window 2.1
import QtCanvas3D 1.0
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.1

import "jsonmodels.js" as GLCode

Window {
    visible: true
    width: 1200
    height: 900
    color: "#e6e6e6"

    //! [3]
    property int previousY: 0
    property int previousX: 0
    //! [3]

    Rectangle {
        id: valuePanel
        width: 150
        height: 100
        anchors.left: parent.left
        anchors.top: parent.top
        opacity: 0.3
        border.color: "black"
        border.width: 2
        radius: 5
        z: 1
    }
    ColumnLayout {
        width: valuePanel.width
        height: valuePanel.height
        x: 10
        z: 2
        Label {
            font.pixelSize: 20
            text: "x angle: " + angle
            readonly property int angle: canvas3d.xRot
        }
        Label {
            font.pixelSize: 20
            text: "y angle: " + angle
            readonly property int angle: canvas3d.yRot
        }
        Label {
            font.pixelSize: 20
            text: "distance: " + distance
            readonly property int distance: canvas3d.distance * 10
        }
    }

    ColumnLayout {
        anchors.fill: parent
        RowLayout {
            Layout.fillHeight: true
            Layout.fillWidth: true
            Canvas3D {
                id: canvas3d
                Layout.fillHeight: true
                Layout.fillWidth: true
                //! [1]
                property double xRot: 0.0
                property double yRot: 45.0
                property double distance: 2.0
                //! [1]
                property double itemSize: 1.0
                property double lightX: 0.0
                property double lightY: 45.0
                property double lightDistance: 2.0
                property bool animatingLight: false
                property bool animatingCamera: false
                property bool drawWireframe: false

                onInitializeGL: {
                    GLCode.initializeGL(canvas3d);
                }

                onPaintGL: {
                    GLCode.paintGL(canvas3d);
                }

                onResizeGL: {
                    GLCode.resizeGL(canvas3d);
                }

                //! [0]
                MouseArea {
                    anchors.fill: parent
                    //! [0]
                    //! [2]
                    onMouseXChanged: {
                        // Do not rotate if we don't have previous value
                        if (previousY !== 0)
                            canvas3d.yRot += mouseY - previousY
                        previousY = mouseY
                        // Limit the rotation to -90...90 degrees
                        if (canvas3d.yRot > 90)
                            canvas3d.yRot = 90
                        if (canvas3d.yRot < -90)
                            canvas3d.yRot = -90
                    }
                    onMouseYChanged: {
                        // Do not rotate if we don't have previous value
                        if (previousX !== 0)
                            canvas3d.xRot += mouseX - previousX
                        previousX = mouseX
                        // Wrap the rotation around
                        if (canvas3d.xRot > 180)
                            canvas3d.xRot -= 360
                        if (canvas3d.xRot < -180)
                            canvas3d.xRot += 360
                    }
                    onReleased: {
                        // Reset previous mouse positions to avoid rotation jumping
                        previousX = 0
                        previousY = 0
                    }
                    //! [2]
                    //! [4]
                    onWheel: {
                        canvas3d.distance -= wheel.angleDelta.y / 1000.0
                        // Limit the distance to 0.5...10
                        if (canvas3d.distance < 0.5)
                            canvas3d.distance = 0.5
                        if (canvas3d.distance > 10)
                            canvas3d.distance = 10
                    }
                    //! [4]
                }
            }
        }
        RowLayout {
            Layout.fillWidth: true
            Button {
                id: lightButton
                Layout.fillWidth: true
                text: "Animate Light"
                checkable: true
                onCheckedChanged: canvas3d.animatingLight = checked
            }
            Button {
                id: cameraButton
                Layout.fillWidth: true
                text: "Animate Camera"
                checkable: true
                onCheckedChanged: canvas3d.animatingCamera = checked
            }
            Button {
                id: drawButton
                Layout.fillWidth: true
                text: "Wireframe"
                checkable: true
                onCheckedChanged: canvas3d.drawWireframe = checked
            }
            Button {
                Layout.fillWidth: true
                text: "Reset"
                onClicked: {
                    canvas3d.xRot = 0.0
                    canvas3d.yRot = 45.0
                    canvas3d.distance = 2.0
                    canvas3d.itemSize = 1.0
                    canvas3d.lightX = 0.0
                    canvas3d.lightY = 45.0
                    canvas3d.lightDistance = 2.0
                    lightButton.checked = false
                    cameraButton.checked = false
                    drawButton.checked = false
                }
            }
            Button {
                Layout.fillWidth: true
                text: "Quit"
                onClicked: Qt.quit()
            }
        }
    }

    SequentialAnimation {
        loops: Animation.Infinite
        running: canvas3d.animatingLight
        NumberAnimation {
            target: canvas3d
            property: "lightX"
            from: 0.0
            to: 360.0
            duration: 5000
        }
    }

    SequentialAnimation {
        loops: Animation.Infinite
        running: canvas3d.animatingLight
        NumberAnimation {
            target: canvas3d
            property: "lightY"
            from: 0.0
            to: 90.0
            duration: 10000
        }
        NumberAnimation {
            target: canvas3d
            property: "lightY"
            from: 90.0
            to: 0.0
            duration: 10000
        }
    }

    SequentialAnimation {
        loops: Animation.Infinite
        running: canvas3d.animatingLight
        NumberAnimation {
            target: canvas3d
            property: "lightDistance"
            from: 10.0
            to: 0.5
            duration: 30000
        }
        NumberAnimation {
            target: canvas3d
            property: "lightDistance"
            from: 0.5
            to: 10.0
            duration: 30000
        }
    }

    SequentialAnimation {
        loops: Animation.Infinite
        running: canvas3d.animatingCamera
        NumberAnimation {
            target: canvas3d
            property: "xRot"
            from: -180.0
            to: 180.0
            duration: 5000
            easing.type: Easing.InOutSine
        }
        NumberAnimation {
            target: canvas3d
            property: "xRot"
            from: 180.0
            to: -180.0
            duration: 5000
            easing.type: Easing.InOutSine
        }
    }

    SequentialAnimation {
        loops: Animation.Infinite
        running: canvas3d.animatingCamera
        NumberAnimation {
            target: canvas3d
            property: "yRot"
            from: 0.0
            to: 90.0
            duration: 10000
            easing.type: Easing.InOutSine
        }
        NumberAnimation {
            target: canvas3d
            property: "yRot"
            from: 90.0
            to: 0.0
            duration: 10000
            easing.type: Easing.InOutSine
        }
    }

    SequentialAnimation {
        loops: Animation.Infinite
        running: canvas3d.animatingCamera
        NumberAnimation {
            target: canvas3d
            property: "distance"
            from: 10.0
            to: 0.5
            duration: 30000
            easing.type: Easing.InOutSine
        }
        NumberAnimation {
            target: canvas3d
            property: "distance"
            from: 0.5
            to: 10.0
            duration: 30000
            easing.type: Easing.InOutSine
        }
    }
}
