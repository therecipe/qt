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
import QtCanvas3D 1.0

import "imagecube.js" as GLCode

Canvas3D {
    id: cube
    //! [0]
    state: "image6"
    property color backgroundColor: "#FCFCFC"
    property real angleOffset: -180 / 8.0
    property string image1: ""
    //! [0]
    property string image2: ""
    property string image3: ""
    property string image4: ""
    property string image5: ""
    property string image6: ""
    property real xRotation: 0
    property real yRotation: 0
    property real zRotation: 0

    onBackgroundColorChanged: { GLCode.setBackgroundColor(cube.backgroundColor); }

    //! [1]
    states: [
        State {
            name: "image1"
            PropertyChanges { target: cube; xRotation: 0; }
            PropertyChanges { target: cube; yRotation: 180 * 1.5 + angleOffset; }
            PropertyChanges { target: cube; zRotation: 0 }
        },
        //! [1]
        State {
            name: "image2"
            PropertyChanges { target: cube; xRotation: 0; }
            PropertyChanges { target: cube; yRotation: 180 * 1.0 + angleOffset; }
            PropertyChanges { target: cube; zRotation: 0 }
        },
        State {
            name: "image3"
            PropertyChanges { target: cube; xRotation: 0; }
            PropertyChanges { target: cube; yRotation: 180 * 0.5 + angleOffset; }
            PropertyChanges { target: cube; zRotation: 0 }
        },
        State {
            name: "image4"
            PropertyChanges { target: cube; xRotation: 0; }
            PropertyChanges { target: cube; yRotation: 0 + angleOffset; }
            PropertyChanges { target: cube; zRotation: 0 }
        },
        State {
            name: "image5"
            PropertyChanges { target: cube; xRotation: 180 / 2.0; }
            PropertyChanges { target: cube; yRotation: 0; }
            PropertyChanges { target: cube; zRotation: -angleOffset; }
        },
        State {
            name: "image6"
            PropertyChanges { target: cube; xRotation: -180 / 2.0; }
            PropertyChanges { target: cube; yRotation: 0; }
            PropertyChanges { target: cube; zRotation: angleOffset; }
        }
    ]

    //! [2]
    transitions: [
        Transition {
            id: turnTransition
            from: "*"
            to: "*"
            RotationAnimation {
                properties: "xRotation,yRotation,zRotation"
                easing.type: Easing.InOutCubic
                direction: RotationAnimation.Shortest
                duration: 450
            }
        }
    ]
    //! [2]

    //! [3]
    onInitializeGL: {
        GLCode.initializeGL(cube);
    }

    onPaintGL: {
        GLCode.paintGL(cube);
    }

    onResizeGL: {
        GLCode.resizeGL(cube);
    }
    //! [3]

    SwipeArea {
        id: swipeArea
        anchors.fill: parent

        onSwipeRight: {
            if (cube.state === "image1")
                cube.state = "image4";
            else if (cube.state == "image2")
                cube.state = "image1";
            else if (cube.state == "image3")
                cube.state = "image2";
            else if (cube.state == "image4")
                cube.state = "image3";
            else if (cube.state == "image5")
                cube.state = "image3";
            else if (cube.state == "image6")
                cube.state = "image3";
        }

        onSwipeLeft: {
            if (cube.state === "image1")
                cube.state = "image2";
            else if (cube.state == "image2")
                cube.state = "image3";
            else if (cube.state == "image3")
                cube.state = "image4";
            else if (cube.state == "image4")
                cube.state = "image1";
            else if (cube.state == "image5")
                cube.state = "image1";
            else if (cube.state == "image6")
                cube.state = "image1";
        }

        onSwipeUp: {
            if (cube.state === "image1")
                cube.state = "image6";
            else if (cube.state == "image2")
                cube.state = "image6";
            else if (cube.state == "image3")
                cube.state = "image6";
            else if (cube.state == "image4")
                cube.state = "image6";
            else if (cube.state == "image5")
                cube.state = "image4";
            else if (cube.state == "image6")
                cube.state = "image2";
        }

        onSwipeDown: {
            if (cube.state === "image1")
                cube.state = "image5";
            else if (cube.state == "image2")
                cube.state = "image5";
            else if (cube.state == "image3")
                cube.state = "image5";
            else if (cube.state == "image4")
                cube.state = "image5";
            else if (cube.state == "image5")
                cube.state = "image2";
            else if (cube.state == "image6")
                cube.state = "image4";
        }
    }

}

