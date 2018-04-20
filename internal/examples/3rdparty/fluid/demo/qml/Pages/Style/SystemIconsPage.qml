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
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.3
import QtQuick.Controls.Material 2.3
import Fluid.Controls 1.0 as FluidControls

Page {
    header: Row {
        RadioButton {
            id: lightRadio
            text: qsTr("Light")
            checked: true
        }

        RadioButton {
            id: darkRadio
            text: qsTr("Dark")
        }
    }

    Material.theme: lightRadio.checked ? Material.Light : Material.Dark

    ScrollView {
        id: scrollView
        anchors.fill: parent
        clip: true

        GridLayout {
            anchors.fill: parent
            columns: (scrollView.width * 0.8) / 48
            columnSpacing: 16
            rowSpacing: 16

            FluidControls.Icon {
                name: "text-editor-symbolic"
                size: 48
            }

            FluidControls.Icon {
                name: "weather-few-clouds-symbolic"
                size: 48
            }

            FluidControls.Icon {
                name: "system-software-install-symbolic"
                size: 48
            }

            FluidControls.Icon {
                name: "system-users-symbolic"
                size: 48
            }

            FluidControls.Icon {
                name: "accessories-calculator"
                size: 48
            }

            FluidControls.Icon {
                name: "accessories-character-map"
                size: 48
            }

            FluidControls.Icon {
                name: "accessories-dictionary"
                size: 48
            }

            FluidControls.Icon {
                name: "accessories-text-editor"
                size: 48
            }
        }
    }
}
