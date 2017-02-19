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

Item {
    id: colorSelector
    property color selectedColor

    height: ((width - 4) * colorSelectorModel.count) + 4

    Rectangle {
        anchors.fill: parent
        color: "black"
        radius: width / 10
        opacity: 0.3
    }

    ListModel {
        id: colorSelectorModel
        ListElement { caseColor: "#eeeeee" }
        ListElement { caseColor: "#111111" }
        ListElement { caseColor: "#ffe400" }
        ListElement { caseColor: "#469835" }
        ListElement { caseColor: "#fa0000" }
    }

    GridView {
        id: colorSelectorGrid
        anchors.fill: colorSelector
        anchors.margins: 4
        model: colorSelectorModel
        interactive: false
        cellWidth: width
        cellHeight: cellWidth + 4
        delegate: Component {
            Rectangle {
                id: colorDelegate
                width: colorSelectorGrid.cellWidth
                height: colorSelectorGrid.cellWidth
                color: caseColor
                border.color: "lightgray"
                border.width: 2
                radius: width / 10
                MouseArea {
                    anchors.fill: parent
                    onClicked: {
                        selectedColor = parent.color
                        colorDelegateAnimation.start()
                    }
                }
                NumberAnimation {
                    id: colorDelegateAnimation
                    running: false
                    loops: 1
                    target: colorDelegate
                    property: "opacity"
                    from: 0.1
                    to: 1.0
                    duration: 500
                }
            }
        }
    }
}

