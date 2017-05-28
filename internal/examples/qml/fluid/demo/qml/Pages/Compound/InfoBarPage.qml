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
import Fluid.Controls 1.0

Item {
    Button {
        anchors.centerIn: parent
        text: qsTr("Open")
        onClicked: infoBar.open(qsTr("Message sent"))
    }

    InfoBar {
        id: infoBar
        buttonText: qsTr("OK")
    }
}
