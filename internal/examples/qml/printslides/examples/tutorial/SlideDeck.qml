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


import Qt.labs.presentation 1.0
import QtQuick 2.0

Presentation
{
    id: presentation

    width: 1280
    height: 720

    SlideCounter {}
    Clock {}


    Slide {
        centeredText: "Use [space] or [keypad] to see intro"
    }

    Slide {
        title: "Presentation {} Element"
        content: [
            "Toplevel element",
            "Defines background",
            "Foreground color",
        ]

        CodeSection {
            text: "Presentation
{
    width: 640
    height: 360

    // Define a background color...
    Rectangle {
        anchors.fill: parent
        color: \"white\"
    }

    property color textColor: \"black\"

    // Then define slide elements
    Slide { ... }
    Slide { ... }
    Slide { ... }
    ...
}"
        }
    }


    Slide {
        title: "Slide {} Element"
        content: [
            "Bullet points",
            "Should be short",
            "And to the point",
            " Sub point",
            "  Sub Sub point",
            " Sub point"
        ]

        CodeSection {

            text: "Slide {\n" +
                  "    id: areaSlide\n" +
                  "    title: \"Slide {} Element\"\n" +
                  "    content: [\n" +
                  "              \"Bullet points\",\n" +
                  "              \"Should be short\",\n" +
                  "              \"And to the point\",\n" +
                  "              \" Sub point\",\n" +
                  "              \"  Sub Sub point\",\n" +
                  "              \" Sub point\"\n" +
                  "             ]\n" +
                  "}\n"
        }
    }

    Slide {
        title: "Slide {}, continued"
        Rectangle {
            anchors.fill: parent
            color: "lightGray"
            Text {
                text: "Slide fills this area..."
                anchors.centerIn: parent
            }
        }
    }

    Slide {
        id: fillAreaSlide
        title: "Slide {}, continued"
        content: ["The built-in property \"contentWidth\" can be used to let the bullet points fill a smaller area of the slide..."]

        SequentialAnimation on contentWidth {
            PropertyAction { value: fillAreaSlide.width }
            PauseAnimation { duration: 2500 }
            NumberAnimation { to: fillAreaSlide.width / 2; duration: 5000; easing.type: Easing.InOutCubic }
            running: fillAreaSlide.visible
        }

        Rectangle {
            height: parent.height
            width: parent.contentWidth

            color: "lightGray"
            z: -1
        }
    }

    Slide {
        title: "Slide {}, continued"
        centeredText: "Use the predefined <i><b><code>centeredText</code></b></i> property to put a single block of text at the center of the Slide{}"
    }

    Slide {
        title: "Slide {}, continued"
        centeredText: '<font color="red"><i>Use</i> rich text, <font color="blue">if <b>you</b> like...'
    }

    Slide {
        title: "Slide {}, continued"
        writeInText: "You can also use the 'writeInText' property for text that should appear character-by-character, like this...\n\n\nProin vulputate pretium tortor, ut bibendum nisi ultricies et. Nullam pharetra tincidunt lorem eu consequat. Sed placerat sem non lacus dictum at lobortis tellus molestie. Fusce sit amet iaculis odio. Ut dictum nibh quis justo lacinia pulvinar. In hac habitasse platea dictumst. Aliquam erat volutpat."
    }


    CodeSlide {
        title: "CodeSlide {} Element"
        code:
"CodeSlide {
    title: \"CodeSlide {} Element\"
    code:
\"
// Whitespaces are preserved,
// so we start at the beginning of the line...

// You can mouse click on any line

// Navigate with keypad when the code has focus

int main(int argc, char **argv) {
    QGuiApplication app;
    QWindow window;
    window.show();
    return app.exec();
}
\" "
    }

    Slide {
        title: "Font size relative to screen size"
        content: [
            "Which means you don't need to worry about it",
            "Bullet points wrap around on the edges, regardless of how long they are, like this. Even if you should choose to use a very long bullet point (which would distract your audience) it would still look ok'ish",
            "If you run out of height, you're out of luck though..."
        ]
    }



    Slide {
        id: interactiveSlide

        title: "Embed Interactive Content"

        Rectangle {
            id: box
            width: parent.fontSize * 10
            height: width
            color: mouse.pressed ? "lightsteelblue" : "steelblue"

            NumberAnimation on rotation { from: 0; to: 360; duration: 10000; loops: Animation.Infinite; running: visible }

            Text {
                text: "Click Me!"
                anchors.centerIn: parent
            }

            MouseArea {
                id: mouse
                anchors.fill: parent
                drag.target: box
            }
        }
    }


    Slide {
        title: "Features"
        centeredText: 'Hit [C] to fade out the current page if there are questions from the audience'
    }

    Slide {
        title: "Features"
        centeredText: 'Navigate back and forth using [left] and [right]\n[space] or [click] takes you to the next slide.'
    }

    CodeSlide {
        title: "Slide Counter"
        code:
"Presentation {

    SlideCounter {
        // Defaults:
        // anchors.right: parent.right
        // anchors.bottom: parent.bottom
        // anchors.margins: fontSize;
        // textColor: 'black'
        // fontFamily: 'Helvetica' (from presentation)
        // fontScale: 0.5;
        }
    }

    Slide {
        ...
    }
}"
    }

    CodeSlide {
        title: "Clock"
        code:
"Presentation {

    Clock {
        // Defaults:
        // anchors.let: parent.left
        // anchors.bottom: parent.bottom
        // anchors.margins: fontSize;
        // textColor: 'black'
        // fontFamily: 'Helvetica' (from presentation)
        // fontScale: 0.5;
        }
    }

    Slide {
        ...
    }
}"
    }

    Slide {
        title: "Customizations"
        titleColor: "white"

        textColor: "green"
        fontFamily: "Times New Roman"

        Rectangle {
            x: -parent.x
            y: -parent.y
            width: presentation.width
            height: parent.y * 0.9
            color: "black"
        }

        content: [
            "Bullet points, centered text, write-in-text or code listings, can be changed using 'textColor'",
            "Title can be changed with 'textColor'",
            "Font can be changed using 'fontFamily'",
            "Change 'fontScale' for bigger or smaller fonts",
            "All can be set globally on the presentation (and then inherited in all Slide {} elements), or set directly on a specific Slide {} element."
        ]
    }

    CodeSlide {
        title: "Slide Notes in Another Window"
        code:
"Presentation {
    showNotes: true;

    Slide {
        title: 'Slide One'
        content: ['bullet point', 'bullet point'];
        notes: 'Here cometh the notes...'
    }

    ...

    // Check out examples/notes for a running example
}
"
    }


    Slide {
        centeredText: "Now go make our own presentations\n\nEnjoy!"
    }


}
