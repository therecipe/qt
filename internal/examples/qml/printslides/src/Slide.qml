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

Item {
    /*
      Slides can only be instantiated as a direct child of a Presentation {} as they rely on
      several properties there.
     */

    id: slide

    property bool isSlide: true;

    property bool delayPoints: false;
    property int _pointCounter: 0;
    function _advance() {
        if (!parent.allowDelay)
            return false;

        _pointCounter = _pointCounter + 1;
        if (_pointCounter < content.length)
            return true;
        _pointCounter = 0;
        return false;
    }

    property string title;
    property variant content: []
    property string centeredText
    property string writeInText;
    property string notes;

    property real fontSize: parent.height * 0.05
    property real fontScale: 1

    property real baseFontSize: fontSize * fontScale
    property real titleFontSize: fontSize * 1.2 * fontScale
    property real bulletSpacing: 1

    property real contentWidth: width

    // Define the slide to be the "content area"
    x: parent.width * 0.05
    y: parent.height * 0.2
    width: parent.width * 0.9
    height: parent.height * 0.7

    property real masterWidth: parent.width
    property real masterHeight: parent.height

    property color titleColor: parent.titleColor;
    property color textColor: parent.textColor;
    property string fontFamily: parent.fontFamily;
    property int textFormat: Text.PlainText

    visible: false

    Text {
        id: titleText
        font.pixelSize: titleFontSize
        text: title;
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.bottom: parent.top
        anchors.bottomMargin: parent.fontSize * 1.5
        font.bold: true;
        font.family: slide.fontFamily
        color: slide.titleColor
        horizontalAlignment: Text.Center
        z: 1
    }

    Text {
        id: centeredId
        width: parent.width
        anchors.centerIn: parent
        anchors.verticalCenterOffset: - parent.y / 3
        text: centeredText
        horizontalAlignment: Text.Center
        font.pixelSize: baseFontSize
        font.family: slide.fontFamily
        color: slide.textColor
        wrapMode: Text.Wrap
    }

    Text {
        id: writeInTextId
        property int length;
        font.family: slide.fontFamily
        font.pixelSize: baseFontSize
        color: slide.textColor

        anchors.fill: parent;
        wrapMode: Text.Wrap

        text: slide.writeInText.substring(0, length);

        NumberAnimation on length {
            from: 0;
            to: slide.writeInText.length;
            duration: slide.writeInText.length * 30;
            running: slide.visible && parent.visible && slide.writeInText.length > 0
        }

        visible: slide.writeInText != undefined;
    }


    Column {
        id: contentId
        anchors.fill: parent

        Repeater {
            model: content.length

            Row {
                id: row

                function decideIndentLevel(s) { return s.charAt(0) == " " ? 1 + decideIndentLevel(s.substring(1)) : 0 }
                property int indentLevel: decideIndentLevel(content[index])
                property int nextIndentLevel: index < content.length - 1 ? decideIndentLevel(content[index+1]) : 0
                property real indentFactor: (10 - row.indentLevel * 2) / 10;

                height: text.height + (nextIndentLevel == 0 ? 1 : 0.3) * slide.baseFontSize * slide.bulletSpacing
                x: slide.baseFontSize * indentLevel
                visible: (!slide.parent.allowDelay || !delayPoints) || index <= _pointCounter

                Rectangle {
                    id: dot
                    anchors.baseline: text.baseline
                    anchors.baselineOffset: -text.font.pixelSize / 2
                    width: text.font.pixelSize / 3
                    height: text.font.pixelSize / 3
                    color: slide.textColor
                    radius: width / 2
                    opacity: text.text.length == 0 ? 0 : 1
                }

                Item {
                    id: space
                    width: dot.width * 1.5
                    height: 1
                }

                Text {
                    id: text
                    width: slide.contentWidth - parent.x - dot.width - space.width
                    font.pixelSize: baseFontSize * row.indentFactor
                    text: content[index]
                    textFormat: slide.textFormat
                    wrapMode: Text.WordWrap
                    color: slide.textColor
                    horizontalAlignment: Text.AlignLeft
                    font.family: slide.fontFamily
                }
            }
        }
    }

}
