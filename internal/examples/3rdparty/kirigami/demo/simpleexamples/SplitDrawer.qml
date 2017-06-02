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
import org.kde.kirigami 2.0 as Kirigami

Kirigami.ApplicationWindow {
    id: root
    width: Kirigami.Units.gridUnit * 60
    height: Kirigami.Units.gridUnit * 40


    pageStack.initialPage: mainPageComponent
    globalDrawer: Kirigami.SplitDrawer {
        id: drawer
        drawerOpen: true
        //modal: false
        contentItem: Column {
            //implicitWidth: Kirigami.Units.gridUnit * 10

            Kirigami.Label {
                anchors.horizontalCenter: parent.horizontalCenter
                text: "This is a sidebar"
                width: parent.width - Kirigami.Units.smallSpacing * 2
                wrapMode: Text.WordWrap
            }
            Controls.Button {
                anchors.horizontalCenter: parent.horizontalCenter
                text: "Modal Drawer"
                checkable: true
                checked: true
                onCheckedChanged: drawer.modal = checked
            }
        }
    }

    //Main app content
    Component {
        id: mainPageComponent
        MultipleColumnsGallery {}
    }
}
