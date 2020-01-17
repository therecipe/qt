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

import QtQuick 2.1
import QtQuick.Window 2.1

Effect {
    id: root
    divider: false
    parameters: ListModel {
        ListElement {
            name: "Radius"
            value: 0.5
        }
        ListElement {
            name: "Diffraction"
            value: 0.5
        }
        onDataChanged: updateParameters()
    }

    function updateParameters()
    {
        radius   = parameters.get(0).value * 100;
        diffractionIndex = parameters.get(1).value;
    }

    property real posX: -1
    property real posY: -1
    property real pixDens: Screen.pixelDensity

    QtObject {
        id: d
        property real oldTargetWidth: root.targetWidth
        property real oldTargetHeight: root.targetHeight
    }

    // Transform slider values, and bind result to shader uniforms
    property real radius: 0.5 * 100
    property real diffractionIndex: 0.5

    onTargetWidthChanged: {
        if (posX == -1)
            posX = targetWidth / 2
        else if (d.oldTargetWidth != 0)
            posX *= (targetWidth / d.oldTargetWidth)
        d.oldTargetWidth = targetWidth
    }

    onTargetHeightChanged: {
        if (posY == -1)
            posY = targetHeight / 2
        else if (d.oldTargetHeight != 0)
            posY *= (targetHeight / d.oldTargetHeight)
        d.oldTargetHeight = targetHeight
    }

    fragmentShaderFilename: "magnify.fsh"

    MouseArea {
        anchors.fill: parent
        onPositionChanged: { root.posX = mouse.x; root.posY = root.targetHeight - mouse.y }
    }
}
