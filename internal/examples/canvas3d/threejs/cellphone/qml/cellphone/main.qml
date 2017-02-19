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
import QtQuick.Window 2.2

Window {
    property int initialWidth: 800
    property int initialHeight: 600
    id: mainView
    width: initialWidth
    height: initialHeight
    visible: true
    title: "Interactive Mobile Phone Demo"
    color: "black"

    //! [0]
    Item {
        id: textureSource
        layer.enabled: true
        layer.smooth: true
        layer.textureMirroring: ShaderEffectSource.NoMirroring
        //! [0]
        anchors.centerIn: parent
        // The dimensions of the texture source item determine the actual OpenGL texture dimensions.
        // We want a fairly large texture, so that even the smallest text we use renders nicely.
        width: 512
        height: 1024
        // Hide texture source behind the canvas normally, so you can't interact with it.
        // It would also be possible to set visible property to false instead, but if the item
        // is hidden, some things don't update correctly. For example, the Flickable doesn't update
        // its content if it is not visible.
        z: -2
        // Specify transform to make the visual representation of the item match the size and
        // orientation of the ui presented in Canvas. If the window initial dimensions,
        // textureSource dimensions, or phone mesh dimensions or position are changed,
        // scaling needs to be adjusted accordingly.
        //! [2]
        transform: [
            Scale {
                origin.y: textureSource.height / 2
                origin.x: textureSource.width / 2
                yScale: 0.5 * mainView.height / mainView.initialHeight
                xScale: 0.5 * mainView.height / mainView.initialHeight
            }
        ]
        opacity: 0.0
        //! [2]

        // CellphoneApp implements the UI of the phone
        CellphoneApp {
            id: cellphoneApp
            anchors.fill: parent
            backgroundColor: "black"
            textColor: "white"
            onLocked: canvas3d.state = "running"
        }
    }

    // CellphoneCanvas displays the rotating phone and the color selector
    CellphoneCanvas {
        id: canvas3d
        anchors.fill:parent
        textureSource: textureSource
        onRotationStopped: {
            cellphoneApp.resetLockTimer()
            // Bring the UI to the foreground so that it can be interacted with
            textureSource.z = 1
        }
        onRotationStarted: {
            // Hide the texture source behind canvas to ensure UI cannot be
            // interacted while the phone is rotating.
            textureSource.z = -1
        }
    }

    ColorSelector {
        anchors.right: parent.right
        anchors.top: parent.top
        width: parent.width / 16

        onSelectedColorChanged: canvas3d.caseColor = selectedColor
    }

    // FPS display, initially hidden, clicking will show it
    FpsDisplay {
        anchors.left: parent.left
        anchors.top: parent.top
        width: 32
        height: 64
        fps: canvas3d.fps
    }
}
