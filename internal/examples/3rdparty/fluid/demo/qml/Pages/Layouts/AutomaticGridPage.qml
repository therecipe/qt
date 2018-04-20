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
import Fluid.Layouts 1.0 as FluidLayouts
import "../.."

Flickable {
    clip: true
    contentHeight: Math.max(layout.implicitHeight, height)

    ScrollBar.vertical: ScrollBar {}

    FluidLayouts.AutomaticGrid {
        id: layout

        anchors.fill: parent

        cellWidth: 100
        cellHeight: cellWidth

        model: 250
        delegate: Rectangle {
            color: Qt.rgba(Math.random(), Math.random(), Math.random(), 1.0)
            width: 100
            height: width

            Text {
                anchors.centerIn: parent
                text: index + 1
            }
        }
    }
}
