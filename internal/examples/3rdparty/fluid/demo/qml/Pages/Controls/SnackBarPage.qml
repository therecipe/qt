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
import Fluid.Controls 1.0 as FluidControls

Item {
    Column {
        anchors.centerIn: parent

        TextField {
            id: textField
            placeholderText: qsTr("Text")
            text: qsTr("Marked as read")
            width: 300
        }

        TextField {
            id: buttonTextField
            placeholderText: qsTr("Button Text")
            text: qsTr("Undo")
            width: 300
        }

        Switch {
            text: qsTr("Full Width")
            checked: snackBar.fullWidth
            onCheckedChanged: snackBar.fullWidth = checked
        }

        Row {
            spacing: 8

            Button {
                text: qsTr("Open")
                onClicked: snackBar.open(textField.text, buttonTextField.text)
            }

            Button {
                text: qsTr("Close")
                onClicked: snackBar.close()
            }
        }
    }

    FluidControls.SnackBar {
        id: snackBar
        onClicked: console.log("Snack bar button clicked")
    }
}
