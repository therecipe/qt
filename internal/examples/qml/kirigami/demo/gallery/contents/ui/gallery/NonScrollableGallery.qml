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
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.0

Page {
    id: page
    Layout.fillWidth: true
    title: "Simple Page"

    actions {
        main: Action {
            iconName: sheet.sheetOpen ? "dialog-cancel" : "document-edit"
            text: "Main Action Text"
            checkable: true
            onCheckedChanged: sheet.sheetOpen = checked;
        }
    }

    Rectangle {
        anchors.fill: parent
        color: "red"
        Label {
            anchors.centerIn: parent
            text: "Rectangle with automatic margins"
        }
    }

    OverlaySheet {
        id: sheet
        onSheetOpenChanged: page.actions.main.checked = sheetOpen;
        Label {
            property int implicitWidth: Units.gridUnit * 30
            wrapMode: Text.WordWrap
            text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id risus id augue euismod accumsan. Nunc vestibulum placerat bibendum. Morbi commodo auctor varius. Donec molestie euismod ultrices. Sed facilisis augue nec eros auctor, vitae mattis quam rhoncus. Nam ut erat diam. Curabitur iaculis accumsan magna, eget fermentum massa scelerisque eu. Cras elementum erat non erat euismod accumsan. Vestibulum ac mi sed dui finibus pulvinar. Vivamus dictum, leo sed lobortis porttitor, nisl magna faucibus orci, sit amet euismod arcu elit eget est. Duis et vehicula nibh. In arcu sapien, laoreet sit amet porttitor non, rhoncus vel magna. Suspendisse imperdiet consectetur est nec ornare. Pellentesque bibendum sapien at erat efficitur vehicula. Morbi sed porta nibh. Vestibulum ut urna ut dolor sagittis mattis."
        }
    }
}
