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
import QtQuick.Controls 2.0 as Controls
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.0

Page {
    id: page

    title: "Tab Bar"

    header: Controls.TabBar {
        id: tabBar
        currentIndex: swipeView.currentIndex

        Controls.TabButton {
            text: "First"
        }
        Controls.TabButton {
            text: "Second"
        }
        Controls.TabButton {
            text: "Third"
        }
    }
 
    Controls.SwipeView {
        id: swipeView
        anchors.fill: parent
        currentIndex: tabBar.currentIndex
        clip: true

        Repeater {
            model: 3

            Item {
                Label {
                    width: parent.width
                    wrapMode: Label.Wrap
                    horizontalAlignment: Qt.AlignHCenter
                    text: "Page " + modelData
                }
                Controls.TabBar {
                    position: Controls.TabBar.Footer
                    anchors {
                        bottom: parent.bottom
                        horizontalCenter: parent.horizontalCenter
                    }
                    Controls.TabButton {
                        text: "First"
                    }
                    Controls.TabButton {
                        text: "Second"
                    }
                }
            }
        }
    }
}
