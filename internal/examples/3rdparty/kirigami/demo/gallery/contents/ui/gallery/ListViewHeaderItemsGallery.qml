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
        contextualActions: [
            Kirigami.Action {
                id: overlayHeaderAction
                checkable: true
                checked: true
                text:"Overlay Header"
                onTriggered: {
                    mainList.headerPositioning = ListView.OverlayHeader;
                    overlayHeaderAction.checked = pullBackHeaderAction.checked = inlineHeaderAction.checked = false;
                }
            },
            Kirigami.Action {
                id: pullBackHeaderAction
                checkable: true
                text:"PullBack Header"
                onTriggered: {
                    mainList.headerPositioning = ListView.PullBackHeader;
                    overlayHeaderAction.checked = pullBackHeaderAction.checked = inlineHeaderAction.checked = false;
                }
            },
            Kirigami.Action {
                id: inlineHeaderAction
                checkable: true
                text:"Inline Header"
                onTriggered: {
                    mainList.headerPositioning = ListView.InlineHeader;
                    overlayHeaderAction.checked = pullBackHeaderAction.checked = inlineHeaderAction.checked = false;
                }
            },
            Kirigami.Action {
                checkable: true
                checked: true
                text:"Header overlaps list"
                onCheckedChanged: {
                    mainList.headerItem.z = checked ? 10 : 0
                }
            }
        ]
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

    ListView {
        id: mainList
        Timer {
            id: refreshRequestTimer
            interval: 3000
            onTriggered: page.refreshing = false
        }
        headerPositioning: ListView.OverlayHeader
        //headerPositioning: ListView.PullBackHeader
        header: Kirigami.ItemViewHeader {
            backgroundImage.source: "../banner.jpg"
            title: page.title
        }

        model: 200
        delegate: Kirigami.BasicListItem {
            id: listItem
            label: "Item " + modelData
        }
    }
}
