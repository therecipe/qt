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
import QtQuick.Layouts 1.3
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
                        text: qsTr("Horizontal / Single")
                    }

                    Slider {
                        from: 0.0
                        to: 1.0
                        value: 0.5
                    }

                    Slider {
                        from: 0.0
                        to: 1.0
                        value: 0.5
                        enabled: false
                    }

                    // Row 3

                    Label {
                        text: qsTr("Horizontal / Range")
                    }

                    RangeSlider {
                        from: 0.0
                        to: 1.0
                        first.value: 0.4
                        second.value: 0.6
                    }

                    RangeSlider {
                        from: 0.0
                        to: 1.0
                        first.value: 0.4
                        second.value: 0.6
                        enabled: false
                    }

                    // Row 4

                    Label {
                        text: qsTr("Vertical / Single")
                    }

                    Slider {
                        from: 0.0
                        to: 1.0
                        value: 0.5
                        orientation: Qt.Vertical
                    }

                    Slider {
                        from: 0.0
                        to: 1.0
                        value: 0.5
                        enabled: false
                        orientation: Qt.Vertical
                    }

                    // Row 5

                    Label {
                        text: qsTr("Vertical / Range")
                    }

                    RangeSlider {
                        from: 0.0
                        to: 1.0
                        first.value: 0.4
                        second.value: 0.6
                        orientation: Qt.Vertical
                    }

                    RangeSlider {
                        from: 0.0
                        to: 1.0
                        first.value: 0.4
                        second.value: 0.6
                        enabled: false
                        orientation: Qt.Vertical
                    }
                }
            }
        }
    }
}
