/*
 * This file is part of Fluid.
 *
 * Copyright (C) 2017 Pier Luigi Fiorini <pierluigi.fiorini@gmail.com>
 *
 * $BEGIN_LICENSE:MPL2$
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * $END_LICENSE$
 */

import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtQuick.Controls.Universal 2.0
import Fluid.Controls 1.0
import "Pages/Basic"

Tab {
    title: qsTr("Basic")

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
                ListElement { title: qsTr("Button"); source: "qrc:/qml/Pages/Basic/ButtonPage.qml" }
                ListElement { title: qsTr("CheckBox"); source: "qrc:/qml/Pages/Basic/CheckBoxPage.qml" }
                ListElement { title: qsTr("RadioButton"); source: "qrc:/qml/Pages/Basic/RadioButtonPage.qml" }
                ListElement { title: qsTr("Switch"); source: "qrc:/qml/Pages/Basic/SwitchPage.qml" }
                ListElement { title: qsTr("ProgressBar"); source: "qrc:/qml/Pages/Basic/ProgressBarPage.qml" }
                ListElement { title: qsTr("BusyIndicator"); source: "qrc:/qml/Pages/Basic/BusyIndicatorPage.qml" }
                ListElement { title: qsTr("Slider"); source: "qrc:/qml/Pages/Basic/SliderPage.qml" }
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
        initialItem: ButtonPage {}
    }
}
