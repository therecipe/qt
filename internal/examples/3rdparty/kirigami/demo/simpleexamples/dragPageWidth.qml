/*
 *   Copyright 2017 Eike Hein <hein@kde.org>
 *
 *   This program is free software; you can redistribute it and/or modify
 *   it under the terms of the GNU Library General Public License as
 *   published by the Free Software Foundation; either version 2 or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU Library General Public License for more details
 *
 *   You should have received a copy of the GNU Library General Public
 *   License along with this program; if not, write to the
 *   Free Software Foundation, Inc.,
 *   51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */

import QtQuick 2.1

import org.kde.kirigami 2.4 as Kirigami

Kirigami.ApplicationWindow {
    id: root

    property int defaultColumnWidth: Kirigami.Units.gridUnit * 13
    property int columnWidth: defaultColumnWidth

    pageStack.defaultColumnWidth: columnWidth
    pageStack.initialPage: [firstPageComponent, secondPageComponent]

    MouseArea {
        id: dragHandle

        visible: pageStack.wideMode

        anchors.top: parent.top
        anchors.bottom: parent.bottom

        x: columnWidth - (width / 2)
        width: Kirigami.Units.devicePixelRatio * 2

        property int dragRange: (Kirigami.Units.gridUnit * 5)
        property int _lastX: -1

        cursorShape: Qt.SplitHCursor

        onPressed: _lastX = mouseX

        onPositionChanged: {
            if (mouse.x > _lastX) {
                columnWidth = Math.min((defaultColumnWidth + dragRange),
                    columnWidth + (mouse.x - _lastX));
            } else if (mouse.x < _lastX) {
                columnWidth = Math.max((defaultColumnWidth - dragRange),
                    columnWidth - (_lastX - mouse.x));
            }
        }

        Rectangle {
            anchors.fill: parent

            color: "blue"
        }
    }

    Component {
        id: firstPageComponent

        Kirigami.Page {
            id: firstPage

            background: Rectangle { color: "red" }
        }
    }

    Component {
        id: secondPageComponent

        Kirigami.Page {
            id: secondPage

            background: Rectangle { color: "green" }
        }
    }
}
