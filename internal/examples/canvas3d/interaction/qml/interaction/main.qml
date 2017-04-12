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
//! [0]
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.0
//! [0]

import "interaction.js" as GLCode

Item {
    id: mainview
    width: 1280
    height: 768
    visible: true

    Canvas3D {
        id: canvas3d
        anchors.fill: parent
        focus: true
        //! [3]
        property double xRotSlider: 0
        property double yRotSlider: 0
        property double zRotSlider: 0
        //! [3]

        // Emitted when one time initializations should happen
        onInitializeGL: {
            GLCode.initializeGL(canvas3d);
        }

        // Emitted each time Canvas3D is ready for a new frame
        onPaintGL: {
            GLCode.paintGL(canvas3d);
        }

        onResizeGL: {
            GLCode.resizeGL(canvas3d);
        }
    }

    //! [1]
    RowLayout {
        id: controlLayout
        spacing: 5
        x: 12
        y: parent.height - 100
        width: parent.width - (2 * x)
        height: 100
        visible: true
        //! [1]

        Label {
            id: xRotLabel
            Layout.alignment: Qt.AlignRight
            Layout.fillWidth: false
            text: "X-axis:"
        }

        //! [2]
        Slider {
            id: xSlider
            Layout.alignment: Qt.AlignLeft
            Layout.fillWidth: true
            minimumValue: 0;
            maximumValue: 360;
            //! [4]
            onValueChanged: canvas3d.xRotSlider = value;
            //! [4]
        }
        //! [2]

        Label {
            id: yRotLabel
            Layout.alignment: Qt.AlignRight
            Layout.fillWidth: false
            text: "Y-axis:"
        }

        Slider {
            id: ySlider
            Layout.alignment: Qt.AlignLeft
            Layout.fillWidth: true
            minimumValue: 0;
            maximumValue: 360;
            onValueChanged: canvas3d.yRotSlider = value;
        }

        Label {
            id: zRotLabel
            Layout.alignment: Qt.AlignRight
            Layout.fillWidth: false
            text: "Z-axis:"
        }

        Slider {
            id: zSlider
            Layout.alignment: Qt.AlignLeft
            Layout.fillWidth: true
            minimumValue: 0;
            maximumValue: 360;
            onValueChanged: canvas3d.zRotSlider = value;
        }
    }
}
