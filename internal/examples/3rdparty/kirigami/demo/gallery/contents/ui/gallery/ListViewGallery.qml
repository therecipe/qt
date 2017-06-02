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

import QtQuick 2.4
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.1 as Kirigami

Kirigami.ScrollablePage {
    id: page
    Layout.fillWidth: true
    title: "Long List view"

    actions {
        main: Kirigami.Action {
            iconName: sheet.sheetOpen ? "dialog-cancel" : "document-edit"
            text: "Main Action Text"
            checkable: true
            onCheckedChanged: sheet.sheetOpen = checked;
        }
    }

    //Close the drawer with the back button
    onBackRequested: {
        if (sheet.sheetOpen) {
            event.accepted = true;
            sheet.close();
        }
    }

    supportsRefreshing: true
    onRefreshingChanged: {
        if (refreshing) {
            refreshRequestTimer.running = true;
        } else {
            showPassiveNotification("Example refreshing completed")
        }
    }

    background: Rectangle {
        color: Kirigami.Theme.viewBackgroundColor
    }
    Kirigami.OverlaySheet {
        id: sheet
        onSheetOpenChanged: page.actions.main.checked = sheetOpen;
        parent: applicationWindow().overlay
        ListView {
            model: 100
            implicitWidth: Kirigami.Units.gridUnit * 30
            delegate: Kirigami.BasicListItem {
                label: "Item in sheet" + modelData
            }
        }
    }

    ListView {
        Timer {
            id: refreshRequestTimer
            interval: 3000
            onTriggered: page.refreshing = false
        }
        model: 200
        delegate: Kirigami.SwipeListItem {
            id: listItem
            contentItem: Kirigami.Label {
                height: Math.max(implicitHeight, Kirigami.Units.iconSizes.smallMedium)
                anchors.verticalCenter: parent.verticalCenter
                text: "Item " + modelData
                color: listItem.checked || (listItem.pressed && !listItem.checked && !listItem.sectionDelegate) ? listItem.activeTextColor : listItem.textColor
            }
            actions: [
                Kirigami.Action {
                    iconName: "document-decrypt"
                    onTriggered: showPassiveNotification(model.text + " Action 1 clicked")
                },
                Kirigami.Action {
                    iconName: "mail-reply-sender"
                    onTriggered: showPassiveNotification(model.text + " Action 2 clicked")
                }]
        }
    }
}
