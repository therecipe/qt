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

    title: "Radio buttons"
    actions {
        main: Action {
            iconName: "document-edit"
            text: "Main Action Text"
            onTriggered: {
                showPassiveNotification("Action button in buttons page clicked");
            }
        }
        left: Action {
            iconName: "folder-sync"
            text: "Left Action Text"
            onTriggered: {
                showPassiveNotification("Left action triggered")
            }
        }
    }

    Controls.ButtonGroup {
        buttons: column1.children
    }

    ColumnLayout {
        width: page.width

        Item {
            Layout.fillWidth: true
            Layout.minimumHeight: Units.gridUnit * 10
            RowLayout {
                anchors.centerIn: parent
                ColumnLayout {
                    Layout.fillHeight: true
                    Item {
                        width: 1
                        height: 1
                        Layout.fillHeight: true
                    }
                    Label {
                        text: "On"
                        Layout.preferredHeight: radio1.height
                    }
                    Label {
                        text: "Off"
                        Layout.preferredHeight: radio1.height
                    }
                }
                ColumnLayout {
                    id: column1
                    Label {
                        text: "Normal"
                    }
                    Controls.RadioButton {
                        id: radio1
                        text: "On"
                        checked: true
                    }
                    Controls.RadioButton {
                        text: "Off"
                        checked: false
                    }
                }
                ColumnLayout {
                    id: column2
                    enabled: false
                    Label {
                        text: "Disabled"
                    }
                    Controls.RadioButton {
                        text: "On"
                        checked: true
                    }
                    Controls.RadioButton {
                        text: "Off"
                        checked: false
                    }
                }
            }
        }
    }
}
