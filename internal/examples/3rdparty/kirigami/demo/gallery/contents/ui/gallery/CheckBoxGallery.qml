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
    actions {
        main: Action {
            iconName: sheet.sheetOpen ? "dialog-cancel" : "document-edit"
            text: "Main Action Text"
            checked: sheet.sheetOpen
            checkable: true
            onCheckedChanged: sheet.sheetOpen = checked;
        }
        contextualActions: [
            Action {
                text:"Action 1"
                onTriggered: showPassiveNotification("Action 1 clicked")
            },
            Action {
                text:"Action 2"
                onTriggered: showPassiveNotification("Action 2 clicked")
            }
        ]
    }

    Layout.fillWidth: true
    title: "Checkboxes"

    OverlaySheet {
        id: sheet
        Label {
            property int implicitWidth: Units.gridUnit * 45
            wrapMode: Text.WordWrap
            text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id risus id augue euismod accumsan. Nunc vestibulum placerat bibendum. Morbi commodo auctor varius. Donec molestie euismod ultrices. Sed facilisis augue nec eros auctor, vitae mattis quam rhoncus. Nam ut erat diam. Curabitur iaculis accumsan magna, eget fermentum massa scelerisque eu. Cras elementum erat non erat euismod accumsan. Vestibulum ac mi sed dui finibus pulvinar. Vivamus dictum, leo sed lobortis porttitor, nisl magna faucibus orci, sit amet euismod arcu elit eget est. Duis et vehicula nibh. In arcu sapien, laoreet sit amet porttitor non, rhoncus vel magna. Suspendisse imperdiet consectetur est nec ornare. Pellentesque bibendum sapien at erat efficitur vehicula. Morbi sed porta nibh. Vestibulum ut urna ut dolor sagittis mattis."
        }
    }

    ColumnLayout {
        //This OverlaySheet is put in the "wrong place", but will be automatically reparented
        // to "page"

        Item {
            Layout.fillWidth: true
            Layout.minimumHeight: Units.gridUnit * 10
            GridLayout {
                id: grid
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
                Controls.CheckBox {
                    text: "On"
                    checked: true
                }
                Controls.CheckBox {
                    text: "On"
                    checked: true
                    enabled: false
                }
                Label {
                    text: "Off"
                }
                Controls.CheckBox {
                    text: "Off"
                    checked: false
                }
                Controls.CheckBox {
                    text: "Off"
                    checked: false
                    enabled: false
                }
            }
            Controls.CheckBox {
                anchors {
                    top: grid.bottom
                    left: grid.left
                }
                text: "Tristate"
                checked: true
                tristate: true
            }
        }
    }
}
