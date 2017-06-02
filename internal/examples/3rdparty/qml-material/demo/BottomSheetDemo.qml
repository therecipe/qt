/*
 * QML Material - An application framework implementing Material Design.
 * Copyright (C) 2015 Steve Coffey <scoffey@barracuda.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as
 * published by the Free Software Foundation, either version 2.1 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */
import QtQuick 2.4
import Material 0.2
import Material.ListItems 0.1

Item {
    anchors.fill: parent
    id: bottomSheetDemo

    Button {
        anchors.centerIn: parent
        elevation: 1
        text: "Open Bottom Sheet"
        onClicked: {
            actionSheet.open()
        }
    }

    BottomActionSheet {
        id: actionSheet

        title: "Demo!"

        actions: [
            Action {
                iconName: "social/share"
                name: "Share"
            },

            Action {
                iconName: "file/file_download"
                name: "Download (Disabled)"
                enabled: false
            },

            Action {
                iconName: "action/autorenew"
                name: "THIS SHOULD BE HIDDEN"
                visible: false
            },

            Action {
                iconName: "action/settings"
                name: "Details"
                hasDividerAfter: true
            },

            Action {
                iconName: "content/forward"
                name: "Move"
            },

            Action {
                iconName: "action/delete"
                name: "Delete"
            },

            Action {
                iconName: "content/create"
                name: "Rename"
            }
        ]    
    }
}

