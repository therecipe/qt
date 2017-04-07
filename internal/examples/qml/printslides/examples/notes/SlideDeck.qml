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
import Qt.labs.presentation 1.0

Presentation {

    width: 1280
    height: 720

    Rectangle {
        anchors.fill: parent
        gradient: Gradient {
            GradientStop { position: 0.16; color: "black" }
            GradientStop { position: 0.17; color: "white" }
            GradientStop { position: 0.92; color: "white" }
            GradientStop { position: 0.93; color: "black" }
        }
    }

    Clock { textColor: "white" }
    SlideCounter { textColor: "white" }

    titleColor: "white"

    showNotes: true

    Slide {
        id: firstSlide;
        centeredText: "Using Notes..."
        fontScale: 2
        notes: "In this window you will see the notes for the very first slide..."
    }

    CodeSlide {
        title: "Enable using Presentation.showNotes"
        notes: "Here you can see the code required to enable the slides view.

It is as simple as setting 'showNotes: true' in the Presentation {} element and then add a text to the 'notes' property of the Slide.

The text will then update automatically as you go from slide to slide."
        code:
"Presentation {
    showNotes: true;

    Slide {
        title: 'Slide One'
        content: ['bullet point', 'bullet point'];
        notes: 'Here cometh the notes...'
    }
"
    }

    Slide {
        title: "Example Slide"
        content: [
            "There should be a second window on your desktop",
            "It should read 'once upon a time'",
            "",
            "Try changing the slide..."
        ]
        notes: "once upon a time..."
    }

    Slide {
        title: "Second Example Slide"
        content: [
            "The text on the second window should now change",
            "The notes window shows plain text with plain line brakes only"
        ]
        notes:
"This is the second example slide...

This notes system uses the QtQuick.Window element to pop up a second window. Quite simple and quite convenient..."
    }

}
