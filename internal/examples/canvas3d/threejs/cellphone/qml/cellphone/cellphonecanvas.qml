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

import QtQuick 2.6
import QtCanvas3D 1.1

import "cellphone.js" as GLCode

Canvas3D {
    id: canvas3d
    property double xRotAnim: 0
    property double yRotAnim: 0
    property double zRotAnim: 0
    property double rotateDistance: 2
    property double uiDistance: 1.51 // This distance is selected so that UI on model matches the size of QML UI
    property double distance: rotateDistance
    property double cameraAngle: 0
    property string caseColor: "#eeeeee"
    property Item textureSource
    signal rotationStopped
    signal rotationStarted

    Component.onCompleted: {
        GLCode.setSphereTexture("qrc:/plutomap1k.jpg")
        GLCode.setIconTexture("qrc:/qtlogo_with_alpha.png")
        GLCode.setMeshFiles("qrc:/cellphone_case.json",
                            "qrc:/cellphone_front.json",
                            "qrc:/cellphone_icon.json")
        GLCode.setCaseColor(caseColor);
    }
    //! [1]
    onInitializeGL: GLCode.initializeGL(canvas3d, textureSource)
    //! [1]
    onPaintGL: GLCode.paintGL(canvas3d)
    onResizeGL: GLCode.resizeGL(canvas3d)
    onCaseColorChanged: GLCode.setCaseColor(caseColor)

    state: "running"
    states: [
        State {
            name: "running"
            StateChangeScript {
                script: {
                    resetAnimationX.stop()
                    resetAnimationY.stop()
                    resetAnimationZ.stop()
                    resetAnimationDistance.stop()
                    resetAnimationDistance.duration = 2000
                    resetAnimationDistance.from = canvas3d.distance
                    resetAnimationDistance.to = canvas3d.rotateDistance
                    resetAnimationDistance.start()
                    objAnimationX.start()
                    objAnimationY.start()
                    objAnimationZ.start()
                    rotationStarted()
                }
            }
        },
        State {
            name: "stopped"
            StateChangeScript {
                script: {
                    objAnimationX.stop()
                    objAnimationY.stop()
                    objAnimationZ.stop()
                    resetAnimationDistance.stop()
                    resetAnimationX.from = canvas3d.xRotAnim
                    resetAnimationY.from = canvas3d.yRotAnim
                    resetAnimationZ.from = canvas3d.zRotAnim
                    resetAnimationDistance.duration = resetAnimationX.duration
                    resetAnimationDistance.from = canvas3d.distance
                    resetAnimationDistance.to = canvas3d.uiDistance
                    resetAnimationDistance.start()
                    resetAnimationX.start()
                    resetAnimationY.start()
                    resetAnimationZ.start()
                    rotationStopped()
                }
            }
        }
    ]

    MouseArea {
        anchors.fill: parent
        onClicked: {
            if (canvas3d.state === "stopped")
                canvas3d.state = "running"
            else
                canvas3d.state = "stopped"
        }
    }

    SequentialAnimation {
        id: objAnimationX
        loops: Animation.Infinite
        running: true
        NumberAnimation {
            target: canvas3d
            property: "xRotAnim"
            from: -90.0
            to: 270.0
            duration: 9200
            easing.type: Easing.InOutQuad
        }
        NumberAnimation {
            target: canvas3d
            property: "xRotAnim"
            from: 270.0
            to: -90.0
            duration: 9200
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
            to: 360.0
            duration: 8500
            easing.type: Easing.InOutCubic
        }
        NumberAnimation {
            target: canvas3d
            property: "yRotAnim"
            from: 360.0
            to: 0.0
            duration: 8500
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
            from: 180.0
            to: -180.0
            duration: 7600
            easing.type: Easing.InOutSine
        }
        NumberAnimation {
            target: canvas3d
            property: "zRotAnim"
            from: -180.0
            to: 180.0
            duration: 7600
            easing.type: Easing.InOutSine
        }
    }

    NumberAnimation {
        id: resetAnimationX
        running: false
        loops: 1
        target: canvas3d
        property: "xRotAnim"
        from: 0.0
        to: -90.0
        duration: 600
        easing.type: Easing.InOutSine
    }
    NumberAnimation {
        id: resetAnimationY
        running: false
        loops: 1
        target: canvas3d
        property: "yRotAnim"
        from: 0.0
        to: 0.0
        duration: resetAnimationX.duration
        easing.type: Easing.InOutSine
    }
    NumberAnimation {
        id: resetAnimationZ
        running: false
        loops: 1
        target: canvas3d
        property: "zRotAnim"
        from: 0.0
        to: 180.0
        duration: resetAnimationX.duration
        easing.type: Easing.InOutSine
    }
    NumberAnimation {
        id: resetAnimationDistance
        running: false
        loops: 1
        target: canvas3d
        property: "distance"
        from: canvas3d.rotateDistance
        to: canvas3d.uiDistance
        duration: resetAnimationX.duration
        easing.type: Easing.InOutSine
    }
    RotationAnimation {
        id: cameraRotationAnimation
        running: true
        target: canvas3d
        property: "cameraAngle"
        loops: Animation.Infinite
        from: 0
        to: 360
        duration: 60000
        direction: RotationAnimation.Clockwise
    }
}
