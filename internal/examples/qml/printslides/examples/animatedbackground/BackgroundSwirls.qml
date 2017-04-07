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
import QtQuick.Particles 2.0

Item {
    anchors.fill: parent

    Rectangle {
        anchors.fill: parent
        gradient: Gradient {
            GradientStop { position: 0; color: "lightsteelblue" }
            GradientStop { position: 1; color: "black" }
        }
    }

    Rectangle {
        id: colorTableItem
        width: 16
        height: 16
        anchors.fill: parent

        property color color1: Qt.rgba(0.8, 0.8, 1, 0.3)
        property color color2: Qt.rgba(0.8, 0.8, 1, 0.3)

        gradient: Gradient {
            GradientStop { position: 0.0; color: "transparent"}
            GradientStop { position: 0.05; color: colorTableItem.color1 }
            GradientStop { position: 0.3; color: "transparent" }
            GradientStop { position: 0.7; color: "transparent" }
            GradientStop { position: 0.95; color: colorTableItem.color2 }
            GradientStop { position: 1.0; color: "transparent" }
        }

        visible: false
    }

    ShaderEffectSource {
        id: colorTableSource
        sourceItem: colorTableItem
        smooth: true
    }

    Repeater {
        model: 4
        Swirl {

            width: parent.width
            anchors.bottom: parent.bottom
            height: parent.height / (2 + index)
            opacity: 0.3
            speed: (index + 1) / 5
            colorTable: colorTableSource
        }
    }

    ParticleSystem{
        id: particles
    }
    ImageParticle{
        anchors.fill: parent
        system: particles
        source: "particle.png"
        alpha: 0
        colorVariation: 0.3
    }

    Emitter{
        anchors.fill: parent
        system: particles
        emitRate: Math.sqrt(parent.width * parent.height) / 30
        lifeSpan: 2000
        size: 4
        sizeVariation: 2

        acceleration: AngleDirection { angle: 90; angleVariation: 360; magnitude: 20; }
        velocity: AngleDirection { angle: -90; angleVariation: 360; magnitude: 10; }
    }

}
