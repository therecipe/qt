/*
 * This file is part of Fluid.
 *
 * Copyright (C) 2017 Pier Luigi Fiorini <pierluigi.fiorini@gmail.com>
 * Copyright (C) 2017 Michael Spencer <sonrisesoftware@gmail.com>
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
                    columns: 3
                    rows: 3

                    // Row 1

                    Item {
                        width: 1
                        height: 1
                    }

                    TitleLabel {
                        text: qsTr("Enabled")

                        Layout.alignment: Qt.AlignHCenter
                    }

                    TitleLabel {
                        text: qsTr("Disabled")

                        Layout.alignment: Qt.AlignHCenter
                    }

                    // Row 2

                    Label {
                        text: qsTr("On")
                    }

                    CheckBox {
                        checked: true
                        text: qsTr("CheckBox")
                    }

                    CheckBox {
                        checked: true
                        enabled: false
                        text: qsTr("CheckBox")
                    }

                    // Row 3

                    Label {
                        text: qsTr("Off")
                    }

                    CheckBox {
                        text: qsTr("CheckBox")
                    }

                    CheckBox {
                        text: qsTr("CheckBox")
                        enabled: false
                    }
                }
            }
        }
    }
}
