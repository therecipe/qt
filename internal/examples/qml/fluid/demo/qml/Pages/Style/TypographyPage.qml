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
import QtQuick.Layouts 1.0
import QtQuick.Controls 2.0
import Fluid.Controls 1.0
import "../.."

Flickable {
    clip: true
    contentHeight: Math.max(layout.implicitHeight, height)

    ScrollBar.vertical: ScrollBar {}

    ColumnLayout {
        id: layout
        anchors.fill: parent
        anchors.margins: Units.mediumSpacing
        spacing: Units.smallSpacing

        DisplayLabel {
            level: 4
            text: "Display 4"
        }

        DisplayLabel {
            level: 3
            text: "Display 3"
        }

        DisplayLabel {
            level: 2
            text: "Display 2"
        }

        DisplayLabel {
            level: 1
            text: "Display 1"
        }

        HeadlineLabel {
            text: "Headline"
        }

        TitleLabel {
            text: "Title"
        }

        SubheadingLabel {
            text: "Subheading"
        }

        BodyLabel {
            level: 2
            text: "Body 2"
        }

        BodyLabel {
            level: 1
            text: "Body 1"
        }

        CaptionLabel {
            text: "Caption"
        }

        Label {
            text: "Label"
        }

        Item {
            Layout.fillHeight: true
        }
    }
}
