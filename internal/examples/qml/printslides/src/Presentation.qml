/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: https://www.qt.io/licensing/
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


import QtQuick 2.5
import QtQuick.Window 2.0

Item {
    id: root

    property variant slides: []
    property int currentSlide: 0

    property bool showNotes: false;
    property bool allowDelay: true;
    property alias mouseNavigation: mouseArea.enabled
    property bool arrowNavigation: true
    property bool keyShortcutsEnabled: true

    property color titleColor: textColor;
    property color textColor: "black"
    property string fontFamily: "Helvetica"
    property string codeFontFamily: "Courier New"

    // Private API
    property bool _faded: false
    property int _userNum;
    property int _lastShownSlide: 0

    Component.onCompleted: {
        var slideCount = 0;
        var slides = [];
        for (var i=0; i<root.children.length; ++i) {
            var r = root.children[i];
            if (r.isSlide) {
                slides.push(r);
            }
        }

        root.slides = slides;
        root._userNum = 0;

        // Make first slide visible...
        if (root.slides.length > 0)
            root.slides[root.currentSlide].visible = true;
    }

    function switchSlides(from, to, forward) {
        from.visible = false
        to.visible = true
        return true
    }

    onCurrentSlideChanged: {
        switchSlides(root.slides[_lastShownSlide], root.slides[currentSlide], currentSlide > _lastShownSlide)
        _lastShownSlide = currentSlide
    }

    function goToNextSlide() {
        root._userNum = 0
        if (_faded)
            return
        if (root.slides[currentSlide].delayPoints) {
            if (root.slides[currentSlide]._advance())
                return;
        }
        if (currentSlide + 1 < root.slides.length)
            ++currentSlide;
    }

    function goToPreviousSlide() {
        root._userNum = 0
        if (root._faded)
            return
        if (currentSlide - 1 >= 0)
            --currentSlide;
    }

    function goToUserSlide() {
        --_userNum;
        if (root._faded || _userNum >= root.slides.length)
            return
        if (_userNum < 0)
            goToNextSlide()
        else {
            currentSlide = _userNum;
            root.focus = true;
        }
    }

    // directly type in the slide number: depends on root having focus
    Keys.onPressed: {
        if (event.key >= Qt.Key_0 && event.key <= Qt.Key_9)
            _userNum = 10 * _userNum + (event.key - Qt.Key_0)
        else {
            if (event.key == Qt.Key_Return || event.key == Qt.Key_Enter)
                goToUserSlide();
            _userNum = 0;
        }
    }

    // navigate with arrow keys
    Shortcut { sequence: StandardKey.MoveToNextLine; enabled: root.arrowNavigation; onActivated: goToNextSlide() }
    Shortcut { sequence: StandardKey.MoveToPreviousLine; enabled: root.arrowNavigation; onActivated: goToPreviousSlide() }
    Shortcut { sequence: StandardKey.MoveToNextChar; enabled: root.arrowNavigation; onActivated: goToNextSlide() }
    Shortcut { sequence: StandardKey.MoveToPreviousChar; enabled: root.arrowNavigation; onActivated: goToPreviousSlide() }

    // presentation-specific single-key shortcuts (which interfere with normal typing)
    Shortcut { sequence: " "; enabled: root.keyShortcutsEnabled; onActivated: goToNextSlide() }
    Shortcut { sequence: "c"; enabled: root.keyShortcutsEnabled; onActivated: root._faded = !root._faded }

    // standard shortcuts
    Shortcut { sequence: StandardKey.MoveToNextPage; onActivated: goToNextSlide() }
    Shortcut { sequence: StandardKey.MoveToPreviousPage; onActivated: goToPreviousSlide() }
    Shortcut { sequence: StandardKey.Quit; onActivated: Qt.quit() }

    Rectangle {
        z: 1000
        color: "black"
        anchors.fill: parent
        opacity: root._faded ? 1 : 0
        Behavior on opacity { NumberAnimation { duration: 250 } }
    }

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        acceptedButtons: Qt.LeftButton | Qt.RightButton
        onClicked: {
            if (mouse.button == Qt.RightButton)
                goToPreviousSlide()
            else
                goToNextSlide()
        }
        onPressAndHold: goToPreviousSlide(); //A back mechanism for touch only devices
    }

    Window {
        id: notesWindow;
        width: 400
        height: 300

        title: "QML Presentation: Notes"
        visible: root.showNotes

        Flickable {
            anchors.fill: parent
            contentWidth: parent.width
            contentHeight: textContainer.height

            Item {
                id: textContainer
                width: parent.width
                height: notesText.height + 2 * notesText.padding

                Text {
                    id: notesText

                    property real padding: 16;

                    x: padding
                    y: padding
                    width: parent.width - 2 * padding


                    font.pixelSize: 16
                    wrapMode: Text.WordWrap

                    property string notes: root.slides[root.currentSlide].notes;

                    onNotesChanged: {
                        var result = "";

                        var lines = notes.split("\n");
                        var beginNewLine = false
                        for (var i=0; i<lines.length; ++i) {
                            var line = lines[i].trim();
                            if (line.length == 0) {
                                beginNewLine = true;
                            } else {
                                if (beginNewLine && result.length) {
                                    result += "\n\n"
                                    beginNewLine = false
                                }
                                if (result.length > 0)
                                    result += " ";
                                result += line;
                            }
                        }

                        if (result.length == 0) {
                            font.italic = true;
                            text = "no notes.."
                        } else {
                            font.italic = false;
                            text = result;
                        }
                    }
                }
            }
        }
    }
}
