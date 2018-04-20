/*
 * This file is part of Fluid.
 *
 * Copyright (C) 2018 Pier Luigi Fiorini <pierluigi.fiorini@gmail.com>
 *
 * $BEGIN_LICENSE:MPL2$
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * $END_LICENSE$
 */

import QtQuick 2.10
import QtQuick.Controls 2.3
import QtQuick.Controls.Material 2.3
import QtQuick.Controls.Universal 2.3
import Fluid.Controls 1.0
import "Pages/Controls"

Tab {
    title: qsTr("Controls")

    Pane {
        id: listPane
        anchors {
            left: parent.left
            top: parent.top
            bottom: parent.bottom
        }
        width: 200
        padding: 0
        z: 2

        Material.background: "white"
        Material.elevation: 1

        Universal.background: Universal.accent

        ListView {
            id: listView
            anchors.fill: parent
            currentIndex: 0
            model: ListModel {
                ListElement { title: qsTr("BottomSheet"); source: "qrc:/qml/Pages/Controls/BottomSheetPage.qml" }
                ListElement { title: qsTr("Card"); source: "qrc:/qml/Pages/Controls/CardPage.qml" }
                ListElement { title: qsTr("Dialogs"); source: "qrc:/qml/Pages/Controls/DialogsPage.qml" }
                ListElement { title: qsTr("DatePicker"); source: "qrc:/qml/Pages/Controls/DatePicker.qml" }
                ListElement { title: qsTr("DateTimePicker"); source: "qrc:/qml/Pages/Controls/DateTimePicker.qml" }
                ListElement { title: qsTr("FAB"); source: "qrc:/qml/Pages/Controls/ActionButtonPage.qml" }
                ListElement { title: qsTr("ListItem"); source: "qrc:/qml/Pages/Controls/ListItemPage.qml" }
                ListElement { title: qsTr("NavigationDrawer"); source: "qrc:/qml/Pages/Controls/NavDrawerPage.qml" }
                ListElement { title: qsTr("Overlay"); source: "qrc:/qml/Pages/Controls/OverlayPage.qml" }
                ListElement { title: qsTr("Placeholder"); source: "qrc:/qml/Pages/Controls/Placeholder.qml" }
                ListElement { title: qsTr("Search"); source: "qrc:/qml/Pages/Controls/Search.qml" }
                ListElement { title: qsTr("SnackBar"); source: "qrc:/qml/Pages/Controls/SnackBarPage.qml" }
                ListElement { title: qsTr("TimePicker"); source: "qrc:/qml/Pages/Controls/TimePicker.qml" }
                ListElement { title: qsTr("Wave"); source: "qrc:/qml/Pages/Controls/WavePage.qml" }
            }
            header: Subheader {
                text: qsTr("Demos")
            }
            delegate: ListItem {
                text: model.title
                highlighted: ListView.isCurrentItem
                onClicked: {
                    listView.currentIndex = index
                    stackView.push(model.source)
                }
            }

            ScrollBar.vertical: ScrollBar {}
        }
    }

    StackView {
        id: stackView
        anchors {
            left: listPane.right
            top: parent.top
            right: parent.right
            bottom: parent.bottom
        }
        initialItem: BottomSheetPage {}
    }
}
