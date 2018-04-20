/*
 *   Copyright 2015 Marco Martin <mart@kde.org>
 *
 *   This program is free software; you can redistribute it and/or modify
 *   it under the terms of the GNU Library General Public License as
 *   published by the Free Software Foundation; either version 2 or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU Library General Public License for more details
 *
 *   You should have received a copy of the GNU Library General Public
 *   License along with this program; if not, write to the
 *   Free Software Foundation, Inc.,
 *   51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */

import QtQuick 2.5
import QtQuick.Layouts 1.1
import QtQuick.Controls 2.0 as QQC2
import QtQuick.Window 2.2
import org.kde.kirigami 2.4 as Kirigami

Kirigami.ScrollablePage {
    id: page

    title: "Size Units"

    GridLayout {
        id: themetest
        columns: 2
        anchors.margins: Kirigami.Units.gridUnit / 2

        Kirigami.Heading {
            text: "Screen"
            Layout.columnSpan: 2
            level: 2
        }

        QQC2.Label {
            text: "Geometry (pixels):"
        }
        QQC2.Label {
            text: page.width + "x" + page.height
        }

        QQC2.Label {
            text: "Geometry (gridUnits):"
        }
        QQC2.Label {
            text: Math.round(page.width / Kirigami.Units.gridUnit) + "x" + Math.round(page.height / Kirigami.Units.gridUnit)
        }

        QQC2.Label {
            text: "Units.gridUnit:"
        }
        QQC2.Label {
            text: Kirigami.Units.gridUnit
        }

        QQC2.Label {
            text: "Units.devicePixelRatio:"
        }
        QQC2.Label {
            text: Kirigami.Units.devicePixelRatio.toFixed(2)
        }
        QQC2.Label {
            text: "Screen.devicePixelRatio:"
        }
        QQC2.Label {
            text: Screen.devicePixelRatio
        }

        Kirigami.Heading {
            text: "Font Metrics"
            level: 2
            Layout.columnSpan: 2
        }

        QQC2.Label {
            text: "FontMetrics pointSize:"
        }
        QQC2.Label {
            text: Kirigami.Units.fontMetrics.font.pointSize
        }

        QQC2.Label {
            text: "FontMetrics pixelSize:"
        }
        QQC2.Label {
            text: Number(Kirigami.Units.fontMetrics.font.pixelSize).toFixed(2)
        }

        QQC2.Label {
            text: "Height of default font:"
        }
        QQC2.Label {
            text: Number(font.pixelSize / Screen.pixelDensity).toFixed(2) + "mm"
        }

        Kirigami.Heading {
            text: "Sizes"
            Layout.columnSpan: 2
            level: 2
        }

        QQC2.Label {
            text: "Pixel density:"
        }
        QQC2.Label {
            text: Number(Screen.pixelDensity).toFixed(2)
        }

        QQC2.Label {
            text: "2cm x 2cm square:"
        }
        Rectangle {
            width: Math.round(Screen.pixelDensity * 20)
            height: Math.round(Screen.pixelDensity * 20)
            color: "black"
        }

        QQC2.Label {
            Layout.columnSpan: 2
            Layout.fillHeight: true
        }
    }
}
