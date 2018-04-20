/*
 *   Copyright 2018 Eike Hein <mart@kde.org>
 *   Copyright 2018 Marco Martin <mart@kde.org>
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

import QtQuick 2.7
import QtQuick.Layouts 1.2
import QtQuick.Controls 2.0 as Controls
import org.kde.kirigami 2.4 as Kirigami

Kirigami.ScrollablePage {
    id: page

    title: "Inline Messages"

    actions.main: Kirigami.Action {
        iconName: "documentinfo"
        text: "Info"
        checkable: true
        onCheckedChanged: sheet.sheetOpen = checked;
        shortcut: "Alt+I"
    }

    Kirigami.OverlaySheet {
        id: sheet
        onSheetOpenChanged: page.actions.main.checked = sheetOpen
        header: RowLayout {
            Kirigami.Heading {
                Layout.fillWidth: true
                text: "Inline Messages"
            }
            Controls.ToolButton {
                text: "HIG..."
                enabled: false
                onClicked: Qt.openUrlExternally("")
            }
            Controls.ToolButton {
                text: "Source code..."
                onClicked: Qt.openUrlExternally("https://cgit.kde.org/kirigami.git/tree/examples/gallery/contents/ui/gallery/InlineMessagesGallery.qml")
            }
        }

        Controls.Label {
            property int implicitWidth: Kirigami.Units.gridUnit * 25
            wrapMode: Text.WordWrap
            text: "Inline messages allow you to show various types of messages placed the same layout as content they relate to, instead of showing a message in an overlay. They are invisible by default and need to be explicitly set visible to be revealed."
        }
    }

    ColumnLayout {
        spacing: Kirigami.Units.largeSpacing

        Kirigami.InlineMessage {
            Layout.fillWidth: true

            visible: true

            text: "This is an informational inline message (the default type). Use it for example to announce a result or provide commentary."
        }

        Kirigami.InlineMessage {
            Layout.fillWidth: true

            visible: true

            type: Kirigami.MessageType.Positive

            text: "This is a positive inline message. Use it for example to announce a successful result or the succesful completion of a procedure."
        }

        Kirigami.InlineMessage {
            Layout.fillWidth: true

            visible: true

            type: Kirigami.MessageType.Warning

            text: "This is a warning inline message. Use it for example to provide critical guidance or warn about something that is not going to work."
        }

        Kirigami.InlineMessage {
            Layout.fillWidth: true

            visible: true

            type: Kirigami.MessageType.Error

            text: "This is an error inline message. Use it for example to announce something has gone wrong or that input will not be accepted."
        }

        Kirigami.Separator {
            Layout.fillWidth: true
        }

        Kirigami.InlineMessage {
            Layout.fillWidth: true

            visible: true

            icon.source: "system-run"

            text: "Inline messages can optionally have a custom icon set."
        }

        Kirigami.InlineMessage {
            Layout.fillWidth: true

            visible: true

            text: "You can use rich text in inline messages and optionally handle clicks on links (opens in browser): <a href=\"https://www.kde.org\">https://www.kde.org/<a/>"

            onLinkActivated: Qt.openUrlExternally(link)
        }

        Kirigami.InlineMessage {
            Layout.fillWidth: true

            visible: true

            showCloseButton: true

            text: "Inline messages can have an optional close button."
        }

        Kirigami.InlineMessage {
            id: actionsMessage

            Layout.fillWidth: true

            visible: true

            readonly property string initialText: "Inline messages can have optional actions."

            text: initialText

            actions: [
                Kirigami.Action {
                    enabled: actionsMessage.text == actionsMessage.initialText

                    text: "Add text"
                    icon.name: "list-add"

                    onTriggered: {
                        actionsMessage.text = actionsMessage.initialText + " Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.";
                    }
                },
                Kirigami.Action {
                    enabled: actionsMessage.text != actionsMessage.initialText

                    text: "Reset text"
                    icon.name: "list-remove"

                    onTriggered: actionsMessage.text = actionsMessage.initialText
                }
            ]
        }

        Kirigami.Separator {
            Layout.fillWidth: true
        }

        ColumnLayout {
            Layout.alignment: Qt.AlignHCenter

            spacing: Kirigami.Units.smallSpacing

            Controls.Label {
                Layout.alignment: Qt.AlignHCenter

                text: "Inline messages are initially hidden and animate when revealed. Try it!"
            }

            Controls.Button {
                Layout.alignment: Qt.AlignHCenter

                enabled: !toggleMessage.visible

                text: "Show additional message"
                //TODO: enable when we can depend from Qt 5.10
               // icon.name: "list-add"

                onClicked: toggleMessage.visible = true
            }
        }

        Kirigami.InlineMessage {
            id: toggleMessage

            Layout.fillWidth: true

            visible: false

            type: Kirigami.MessageType.Warning

            text: "Boo!"

            actions: [
                Kirigami.Action {
                    text: "Shush"
                    icon.name: "dialog-cancel"

                    onTriggered: toggleMessage.visible = false
                }
            ]
        }
    }
}
