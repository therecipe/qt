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
    //implicitWidth: Units.gridUnit * (Math.floor(Math.random() * 35) + 10)

    title: "Buttons"

    actions {
        main: Action {
            iconName: sheet.sheetOpen ? "dialog-cancel" : "document-edit"
            text: "Main Action Text"
            checkable: true
            onCheckedChanged: sheet.sheetOpen = checked;
            shortcut: "Alt+S"
        }
        left: Action {
            iconName: "go-previous"
            text: "Left Action Text"
            onTriggered: {
                showPassiveNotification("Left action triggered")
            }
        }
        right: Action {
            iconName: "go-next"
            text: "Right Action Text"
            onTriggered: {
                showPassiveNotification("Right action triggered")
            }
        }
        contextualActions: [
            Action {
                text:"Action for buttons"
                iconName: "bookmarks"
                onTriggered: showPassiveNotification("Action 1 clicked")
            },
            Action {
                text:"Disabled Action"
                iconName: "folder"
                enabled: false
            },
            Action {
                text: "Action for Sheet"
                visible: sheet.sheetOpen
            }
        ]
    }


    //Close the drawer with the back button
    onBackRequested: {
        if (bottomDrawer.drawerOpen) {
            event.accepted = true;
            bottomDrawer.close();
        }
        if (sheet.sheetOpen) {
            event.accepted = true;
            sheet.close();
        }
    }

    OverlayDrawer {
        id: bottomDrawer
        edge: Qt.BottomEdge
        contentItem: Item {
            implicitHeight: childrenRect.height + Units.gridUnit
            ColumnLayout {
                anchors.centerIn: parent
                Controls.Button {
                    text: "Button1"
                    onClicked: showPassiveNotification("Button 1 clicked")
                }
                Controls.Button {
                    text: "Button2"
                    onClicked: showPassiveNotification("Button 2 clicked")
                }
                Item {
                    Layout.minimumHeight: Units.gridUnit * 4
                }
            }
        }
    }
    

    OverlaySheet {
        id: sheet
        onSheetOpenChanged: page.actions.main.checked = sheetOpen
        ColumnLayout {
            Label {
                Layout.fillWidth: true
                wrapMode: Text.WordWrap
                text: "
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id risus id augue euismod accumsan. Nunc vestibulum placerat bibendum. Morbi commodo auctor varius. Donec molestie euismod ultrices. Sed facilisis augue nec eros auctor, vitae mattis quam rhoncus. Nam ut erat diam. Curabitur iaculis accumsan magna, eget fermentum massa scelerisque eu. Cras elementum erat non erat euismod accumsan. Vestibulum ac mi sed dui finibus pulvinar. Vivamus dictum, leo sed lobortis porttitor, nisl magna faucibus orci, sit amet euismod arcu elit eget est. Duis et vehicula nibh. In arcu sapien, laoreet sit amet porttitor non, rhoncus vel magna. Suspendisse imperdiet consectetur est nec ornare. Pellentesque bibendum sapien at erat efficitur vehicula. Morbi sed porta nibh. Vestibulum ut urna ut dolor sagittis mattis."
            }

            Controls.TextField {
                Layout.alignment: Qt.AlignHCenter
            }

            Label {
                Layout.fillWidth: true
                wrapMode: Text.WordWrap
                text: "
    Morbi dictum, sapien at maximus pulvinar, sapien metus condimentum magna, quis lobortis nisi dui mollis turpis. Aliquam sit amet scelerisque dui. In sit amet tellus placerat, condimentum enim sed, hendrerit quam. Integer dapibus lobortis finibus. Suspendisse faucibus eros vitae ante posuere blandit. Nullam volutpat quam id diam hendrerit aliquam. Donec non sem at diam posuere convallis. Vivamus ut congue quam. Ut dictum fermentum sapien, eu ultricies est ornare ut.

    Nullam fringilla a libero vehicula faucibus. Donec euismod sodales nulla, in vehicula lectus posuere a. Donec nisi nulla, pulvinar eu porttitor vitae, varius eget ante. Nam rutrum eleifend elit, quis facilisis leo sodales vitae. Aenean accumsan a nulla at sagittis. Integer placerat tristique magna, vitae iaculis ante cursus sit amet. Sed facilisis mollis turpis nec tristique. Etiam quis feugiat odio. Vivamus sagittis at purus nec aliquam.

    Morbi neque dolor, elementum ac fermentum ac, auctor ut erat. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vivamus non nibh sit amet quam luctus congue. Donec in eros varius, porta metus sed, sagittis lacus. Mauris dapibus lorem nisi, non eleifend massa tristique egestas. Curabitur nec blandit urna. Mauris rhoncus libero felis, commodo viverra ante consectetur vel. Donec dictum tincidunt orci, quis tristique urna. Quisque egestas, dui ac mollis dictum, purus velit elementum est, at pellentesque erat est fermentum purus. Nulla a quam tellus. Vestibulum a congue ligula. Quisque feugiat nulla et tortor sodales viverra. Maecenas dolor leo, elementum sed urna vel, posuere hendrerit metus. Mauris pellentesque, mi non luctus aliquam, leo nulla varius arcu, vel pulvinar enim enim nec nisl.

    Etiam sapien leo, venenatis eget justo at, pellentesque mollis tellus. Fusce consequat ullamcorper vulputate. Duis tellus nisi, dictum ut augue non, elementum congue ligula. Fusce in vehicula arcu. Nulla facilisi. Quisque a convallis sapien. Aenean pellentesque convallis egestas. Phasellus rhoncus, nulla in tempor maximus, arcu ex venenatis diam, sit amet egestas mi dolor non ante. "
            }
            Controls.Button {
                text: "Close"
                anchors.horizontalCenter: parent.horizontalCenter
                onClicked: sheet.close()
            }
        }
    }
    ColumnLayout {
        width: page.width
        spacing: Units.smallSpacing

        Controls.Button {
            text: "Open Bottom drawer"
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: bottomDrawer.open()
        }
        Controls.Button {
            text: "Open Sheet"
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: sheet.open()
        }
        Controls.Button {
            text: "Toggle Action Button"
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: mainAction.visible = !mainAction.visible;
        }
        Controls.Button {
            text: "Show Passive Notification"
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: showPassiveNotification("This is a passive message", 3000);
        }
        Controls.Button {
            text: "Passive Notification Action"
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: showPassiveNotification("This is a passive message", "long", "Action", function() {showPassiveNotification("Passive notification action clicked")});
        }
        Controls.ToolButton {
            text: "Toggle controls"
            checkable: true
            checked: true
            anchors.horizontalCenter: parent.horizontalCenter
            onCheckedChanged: applicationWindow().controlsVisible = checked
        }
        Controls.Button {
            text: "Disabled Button"
            enabled: false
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: showPassiveNotification("clicked")
        }
        Controls.ToolButton {
            text: "Tool Button"
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: showPassiveNotification(text + " clicked")
        }
        Controls.ToolButton {
            text: "Tool Button non flat"
            flat: false
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: showPassiveNotification(text + " clicked")
        }
    }
}
