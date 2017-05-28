/*
 *   Copyright 2015 Marco Martin <mart@kde.org>
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

import QtQuick 2.0
import QtQuick.Controls 2.0 as Controls
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.0

ScrollablePage {
    id: page
    Layout.fillWidth: true
    title: "Switches"

    ColumnLayout {
        width: page.width

        Item {
            Layout.fillWidth: true
            Layout.minimumHeight: Units.gridUnit * 10
            GridLayout {
                anchors.centerIn: parent
                columns: 3
                rows: 3
                rowSpacing: Units.smallSpacing

                Item {
                    width: 1
                    height: 1
                }
                Label {
                    text: "Normal"
                }
                Label {
                    text: "Disabled"
                    enabled: false
                }
                Label {
                    text: "On"
                }
                Controls.Switch {
                    checked: true
                    text: "On"
                }
                Controls.Switch {
                    checked: true
                    enabled: false
                    text: "On"
                }
                Label {
                    text: "Off"
                }
                Controls.Switch {
                    checked: false
                    text: "Off"
                }
                Controls.Switch {
                    checked: false
                    enabled: false
                    text: "Off"
                }
            }
        }
    }
}
