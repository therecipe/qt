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
import QtQuick.Layouts 1.0
import Fluid.Controls 1.0
import "../.."

Flickable {
    clip: true
    contentHeight: Math.max(layout.implicitHeight, height)

    ScrollBar.vertical: ScrollBar {}

    ColumnLayout {
        id: layout
        anchors.fill: parent

        Repeater {
            model: 2

            StyledRectangle {
                Layout.fillWidth: true
                Layout.fillHeight: true
                Layout.minimumWidth: grid.width + 80
                Layout.minimumHeight: grid.height + 80

                GridLayout {
                    id: grid
                    anchors.centerIn: parent
                    columns: 2
                    rows: 4

                    // Row 1

                    TitleLabel {
                        text: qsTr("Enabled")

                        Layout.alignment: Qt.AlignHCenter
                    }

                    TitleLabel {
                        text: qsTr("Disabled")

                        Layout.alignment: Qt.AlignHCenter
                    }

                    // Row 2

                    Button {
                        text: qsTr("Button")
                    }

                    Button {
                        text: qsTr("Button")
                        enabled: false
                    }

                    // Row 3

                    Button {
                        text: qsTr("Checked")
                        checkable: false
                        checked: true
                    }

                    Button {
                        text: qsTr("Checked")
                        checkable: false
                        checked: true
                        enabled: false
                    }

                    // Row 4

                    Button {
                        text: qsTr("Flat")
                        flat: true
                    }

                    Button {
                        text: qsTr("Flat")
                        flat: true
                        enabled: false
                    }

                    // Row 5

                    Button {
                        text: qsTr("Highlighted")
                        highlighted: true
                    }

                    Button {
                        text: qsTr("Highlighted")
                        highlighted: true
                        enabled: false
                    }
                }
            }
        }
    }
}
