/*
 *   Copyright 2016 Aleix Pol Gonzalez <aleixpol@kde.org>
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
    title: "Misc controls"

    header: Controls.ToolBar {
        RowLayout {
            anchors.verticalCenter: parent.verticalCenter
            Controls.ToolButton {
                text: "ToolButton"
            }
            Controls.ToolButton {
                text: "Menu"
                onClicked: menu.open();
                Controls.Menu {
                    id: menu
                    y: parent.height

                    Controls.MenuItem {
                        checkable: true
                        text: "Item1"
                    }
                    Controls.MenuItem {
                        text: "Item2"
                    }
                }
            }
        }
    }

    Controls.Popup {
        id: dialog
        modal: true
        focus: true
        x: (page.width - width) / 2
        y: page.height / 2 - height
        width: Math.min(page.width - Units.gridUnit * 4, Units.gridUnit * 20)
        contentHeight: popupColumn.height

        Column {
            id: popupColumn
            spacing: 20

            Controls.Label {
                width: dialog.availableWidth
                text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id risus id augue euismod accumsan. Nunc vestibulum placerat bibendum. Morbi commodo auctor varius. Donec molestie euismod ultrices. Sed facilisis augue nec eros auctor."
                wrapMode: Label.Wrap
            }
            Controls.Button {
                anchors.right: parent.right
                text: "Ok"
                onClicked: dialog.close()
            }
        }
    }

    ColumnLayout {
        anchors.centerIn: parent
        Controls.Button {
            Layout.alignment: Qt.AlignHCenter
            text: "Dialog"
            onClicked: dialog.open()
        }
        Controls.Dial {
            Layout.alignment: Qt.AlignHCenter
        }
        Controls.SpinBox {
            editable: true
            Layout.alignment: Qt.AlignHCenter
        }
        Controls.ComboBox {
            model: ["First", "Second", "Third"]
            Layout.alignment: Qt.AlignHCenter
        }
        Controls.GroupBox {
            title: "Title"
            Layout.alignment: Qt.AlignHCenter

            ColumnLayout {
                id: options

                Controls.RadioButton {
                    text: "First"
                    checked: true
                }
                Controls.RadioButton {
                    text: "Second"
                    checked: false
                }
                Controls.RadioButton {
                    text: "Third"
                    checked: false
                }
            }
        }
        Column {
            Layout.alignment: Qt.AlignHCenter
            Controls.ItemDelegate {
                width: 300
                text: "Delegate1"
            }
            Controls.ItemDelegate {
                width: 300
                text: "Delegate2"
            }
            Controls.CheckDelegate {
                width: 300
                text: "Delegate3"
            }
            Controls.SwitchDelegate {
                width: 300
                text: "Delegate4"
            }
            Controls.RadioDelegate {
                width: 300
                text: "Delegate5"
            }
        }
    }
}
