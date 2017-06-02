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
import Fluid.Material 1.0
import "../.."

Flickable {
    clip: true
    contentHeight: Math.max(layout.implicitHeight, height)

    ScrollBar.vertical: ScrollBar {}

    Column {
        id: layout
        anchors.fill: parent

        Repeater {
            model: 2

            StyledRectangle {
                //y: height * index
                width: parent.width
                height: parent.height / 2

                Column {
                    anchors.centerIn: parent

                    ActionButton {
                        iconName: "device/airplanemode_active"
                    }

                    ActionButton {
                        iconName: "navigation/check"

                        Material.elevation: 1
                    }

                    ActionButton {
                        iconName: "device/airplanemode_active"

                        Material.background: Material.primaryColor
                    }

                    ActionButton {
                        iconName: "navigation/check"

                        Material.elevation: 1
                        Material.background: Material.primaryColor
                    }
                }
            }
        }
    }
}
