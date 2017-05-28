/*
 *   Copyright 2016 Marco Martin <mart@kde.org>
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
import org.kde.kirigami 2.0 as Kirigami

Kirigami.ApplicationWindow {
    id: root

    header: Kirigami.ApplicationHeader {
    }

    globalDrawer: Kirigami.GlobalDrawer {
        title: "Hello App"
        titleIcon: "applications-graphics"

        actions: [
            Kirigami.Action {
                text: "push"
                onTriggered: pageStack.push(secondPageComponent)
            },
            Kirigami.Action {
                text: "pop"
                onTriggered: pageStack.pop()
            },
            Kirigami.Action {
                text: "clear"
                onTriggered: pageStack.clear()
            },
            Kirigami.Action {
                text: "replace"
                onTriggered: pageStack.replace(secondPageComponent)
            }
        ]
    }

    pageStack.initialPage: mainPageComponent

    Component {
        id: mainPageComponent
        Kirigami.Page {
            title: "First Page"
            Rectangle {
                anchors.fill: parent
                Kirigami.Label {
                    text: "First Page"
                }
            }
        }
    }

    Component {
        id: secondPageComponent
        Kirigami.Page {
            title: "Second Page"
            Rectangle {
                color: "red"
                anchors.fill: parent
                Kirigami.Label {
                    text: "Second Page"
                }
            }
        }
    }

}
