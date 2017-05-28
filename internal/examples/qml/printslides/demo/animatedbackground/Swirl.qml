/****************************************************************************
**
** Copyright (C) 2012 Digia Plc and/or its subsidiary(-ies).
** Contact: http://www.qt-project.org/legal
**
** This file is part of the QML Presentation System.
**
** $QT_BEGIN_LICENSE:LGPL$
** Commercial License Usage
** Licensees holding valid commercial Qt licenses may use this file in
** accordance with the commercial license agreement provided with the
** Software or, alternatively, in accordance with the terms contained in
** a written agreement between you and Digia.  For licensing terms and
** conditions see http://qt.digia.com/licensing.  For further information
** use the contact form at http://qt.digia.com/contact-us.
**
** GNU Lesser General Public License Usage
** Alternatively, this file may be used under the terms of the GNU Lesser
** General Public License version 2.1 as published by the Free Software
** Foundation and appearing in the file LICENSE.LGPL included in the
** packaging of this file.  Please review the following information to
** ensure the GNU Lesser General Public License version 2.1 requirements
** will be met: http://www.gnu.org/licenses/old-licenses/lgpl-2.1.html.
**
** In addition, as a special exception, Digia gives you certain additional
** rights.  These rights are described in the Digia Qt LGPL Exception
** version 1.1, included in the file LGPL_EXCEPTION.txt in this package.
**
** GNU General Public License Usage
** Alternatively, this file may be used under the terms of the GNU
** General Public License version 3.0 as published by the Free Software
** Foundation and appearing in the file LICENSE.GPL included in the
** packaging of this file.  Please review the following information to
** ensure the GNU General Public License version 3.0 requirements will be
** met: http://www.gnu.org/copyleft/gpl.html.
**
**
** $QT_END_LICENSE$
**
****************************************************************************/


import QtQuick 2.0

ShaderEffect {
    id: shader

    width: 400
    height: 300

    property real speed: 1

    property color d: Qt.rgba(Math.random() * 0.7,
                              Math.random() * 0.5,
                              Math.random() * 0.7,
                              Math.random() * 0.5)
    property real tx
    NumberAnimation on tx { from: 0; to: Math.PI * 2; duration: (Math.random() * 30 + 30) * 1000 / speed; loops: Animation.Infinite }
    property real ty
    NumberAnimation on ty { from: 0; to: Math.PI * 2; duration: (Math.random() * 30 + 30) * 1000 / speed; loops: Animation.Infinite }
    property real tz
    NumberAnimation on tz { from: 0; to: Math.PI * 2; duration: (Math.random() * 30 + 30) * 1000 / speed; loops: Animation.Infinite }
    property real tw
    NumberAnimation on tw { from: 0; to: Math.PI * 2; duration: (Math.random() * 30 + 30) * 1000 / speed; loops: Animation.Infinite }

    property real amplitude: height / 2

    property variant colorTable: ShaderEffectSource { sourceItem: Rectangle { width: 4; height: 4; color: "steelblue" } }

    fragmentShader: "
    uniform lowp float qt_Opacity;
    uniform lowp sampler2D colorTable;
    varying highp vec2 qt_TexCoord0;

    void main() {
        gl_FragColor = texture2D(colorTable, qt_TexCoord0);
        gl_FragColor.w *= qt_Opacity;
    }
    "

    vertexShader: "
    uniform lowp vec4 d;
    uniform highp float tx;
    uniform highp float ty;
    uniform highp float tz;
    uniform highp float tw;
    uniform highp float amplitude;
    uniform highp mat4 qt_Matrix;
    attribute highp vec4 qt_Vertex;
    attribute highp vec2 qt_MultiTexCoord0;
    varying highp vec2 qt_TexCoord0;

    void main() {
        highp vec4 pos = qt_Vertex;

        highp float y1 = sin(tx + d.x * qt_MultiTexCoord0.x * 17. + 2. * d.y) + sin(ty + d.z * qt_MultiTexCoord0.x * 11. + 5. * d.w);
        highp float y2 = sin(tz + d.w * qt_MultiTexCoord0.x * 7. + 3. * d.z) + sin(tw + d.y * qt_MultiTexCoord0.x * 19. + 3. * d.x);

        pos.y += mix(y1, y2, qt_MultiTexCoord0.y) * amplitude * 0.5;

        gl_Position = qt_Matrix * pos;
        qt_TexCoord0 = qt_MultiTexCoord0;
    }
    "

    mesh: GridMesh { resolution: Qt.size(width / 10, 4) }

}
