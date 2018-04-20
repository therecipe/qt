/*
 * This file is part of Fluid.
 *
 * Copyright (C) 2018 Pier Luigi Fiorini <pierluigi.fiorini@gmail.com>
 * Copyright (C) 2018 Michael Spencer <sonrisesoftware@gmail.com>
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
import QtQuick.Layouts 1.3
import Fluid.Core 1.0 as FluidCore
import Fluid.Controls 1.0 as FluidControls

FluidControls.ApplicationWindow {
    id: window

    visible: true

    width: 1024
    height: 800

    title: qsTr("Fluid Demo")

    appBar.maxActionCount: 3

    Material.primary: Material.LightBlue
    Material.accent: Material.Blue

    Universal.accent: Universal.Cobalt

    FluidControls.NavigationListView {
        id: navDrawer

        topContent: Image {
            source: FluidCore.Device.isMobile ? "qrc:/images/materialbg.png" : ""

            Layout.fillWidth: true
            Layout.preferredHeight: FluidCore.Device.isMobile ? 200 : window.header.height
        }

        actions: [
            FluidControls.Action {
                text: qsTr("Action 1")
                icon.source: FluidControls.Utils.iconUrl("action/info")
                onTriggered: console.log("action1 triggered")
            },
            FluidControls.Action {
                text: qsTr("Action 2")
                icon.source: FluidControls.Utils.iconUrl("action/info")
                hasDividerAfter: true
                onTriggered: console.log("action2 triggered")
            },
            FluidControls.Action {
                text: qsTr("Action 3")
                icon.source: FluidControls.Utils.iconUrl("action/info")
                onTriggered: console.log("action3 triggered")
            },
            FluidControls.Action {
                text: qsTr("Action 4")
                icon.source: FluidControls.Utils.iconUrl("action/info")
                onTriggered: console.log("action4 triggered")
            },
            FluidControls.Action {
                text: qsTr("Action 5")
                icon.source: FluidControls.Utils.iconUrl("action/info")
                visible: false
                onTriggered: console.log("action5 triggered")
            }
        ]
    }

    initialPage: FluidControls.TabbedPage {
        title: window.title

        leftAction: FluidControls.Action {
            icon.source: FluidControls.Utils.iconUrl("navigation/menu")
            onTriggered: navDrawer.open()
        }

        actions: [
            FluidControls.Action {
                text: qsTr("Dummy error")
                icon.source: FluidControls.Utils.iconUrl("alert/warning")
                toolTip: qsTr("Show a dummy error")
                onTriggered: console.log("Dummy error")
            },
            FluidControls.Action {
                text: qsTr("Colors")
                icon.source: FluidControls.Utils.iconUrl("image/color_lens")
                toolTip: qsTr("Pick a color")
                onTriggered: console.log("Colors")
            },
            FluidControls.Action {
                text: qsTr("Settings")
                icon.source: FluidControls.Utils.iconUrl("action/settings")
                toolTip: qsTr("Settings")
                hoverAnimation: true
                onTriggered: console.log("Settings clicked")
            },
            FluidControls.Action {
                text: qsTr("This should not be visible")
                icon.source: FluidControls.Utils.iconUrl("alert/warning")
                visible: false
            },
            FluidControls.Action {
                text: qsTr("Language")
                icon.source: FluidControls.Utils.iconUrl("action/language")
                enabled: false
            },
            FluidControls.Action {
                text: qsTr("Accounts")
                icon.source: FluidControls.Utils.iconUrl("action/account_circle")
            }
        ]

        BasicComponents {}
        LayoutComponents {}
        Controls {}
        Style {}
    }
}
