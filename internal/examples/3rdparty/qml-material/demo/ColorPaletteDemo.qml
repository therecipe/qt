/*
 * QML Material - An application framework implementing Material Design.
 * Copyright (C) 2015 Ricardo Vieira <ricardo.vieira@tecnico.ulisboa.pt>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as
 * published by the Free Software Foundation, either version 2.1 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

import QtQuick 2.4
import QtQuick.Layouts 1.1
import Material 0.2

Item {
    implicitHeight: palette.height

    GridLayout {
        id: palette
        anchors.centerIn: parent
        width: parent.width
        columns: parseInt(parent.width / dp(300)) || 1
        rowSpacing: dp(20)

        Repeater {
            model: [ "Red", "Pink", "Purple", "Deep Purple", "Indigo",
                     "Blue", "Light Blue", "Cyan", "Teal", "Green",
                     "Light Green", "Lime", "Yellow", "Amber", "Orange",
                     "Deep Orange", "Grey", "Blue Grey", "Brown" ]

            ColumnLayout {
                Layout.alignment: Qt.AlignCenter
                spacing: dp(5)
                property var currentColor: modelData.charAt(0).toLowerCase() +
                                           modelData.slice(1).replace(" ", "")

                Rectangle {
                    width: dp(300)
                    height: dp(126)
                    color: Palette.colors[currentColor]["500"]

                    Label {
                        anchors.top: parent.top
                        anchors.left: parent.left
                        anchors.margins: dp(16)
                        text: modelData
                        style: "body2"
                        color: Theme.lightDark(parent.color,
                                               Theme.light.textColor,
                                               Theme.dark.textColor)
                    }

                    Label {
                        anchors.left: parent.left
                        anchors.bottom: parent.bottom
                        anchors.margins: dp(16)
                        text: "500"
                        style: "body2"
                        color: Theme.lightDark(parent.color,
                                               Theme.light.textColor,
                                               Theme.dark.textColor)
                    }

                    Label {
                        anchors.right: parent.right
                        anchors.bottom: parent.bottom
                        anchors.margins: dp(16)
                        text: Palette.colors[currentColor]["500"]
                        style: "body2"
                        color: Theme.lightDark(parent.color,
                                               Theme.light.textColor,
                                               Theme.dark.textColor)
                    }
                }

                ColumnLayout {
                    spacing: 0

                    Repeater {
                        model: ["100", "200", "300", "400", "600", "700", "800", "900"]

                        Rectangle {
                            width: dp(300)
                            height: dp(46)
                            color: Palette.colors[currentColor][modelData]

                            Label {
                                anchors.left: parent.left
                                anchors.margins: dp(16)
                                anchors.verticalCenter: parent.verticalCenter
                                text: modelData
                                style: "body2"
                                color: Theme.lightDark(parent.color,
                                                       Theme.light.textColor,
                                                       Theme.dark.textColor)
                            }

                            Label {
                                anchors.right: parent.right
                                anchors.margins: dp(16)
                                anchors.verticalCenter: parent.verticalCenter
                                text: Palette.colors[currentColor][modelData]
                                style: "body2"
                                color: Theme.lightDark(parent.color,
                                                       Theme.light.textColor,
                                                       Theme.dark.textColor)
                            }
                        }
                    }
                }

                ColumnLayout {
                    spacing: 0
                    visible: typeof Palette.colors[currentColor]["A100"] != 'undefined'

                    Repeater {
                        model: visible ? ["A100", "A200", "A400", "A700"] : 0

                        Rectangle {
                            width: dp(300)
                            height: dp(46)
                            color: Palette.colors[currentColor][modelData]

                            Label {
                                anchors.left: parent.left
                                anchors.margins: dp(16)
                                anchors.verticalCenter: parent.verticalCenter
                                text: modelData
                                style: "body2"
                                color: Theme.lightDark(parent.color,
                                                       Theme.light.textColor,
                                                       Theme.dark.textColor)
                            }

                            Label {
                                anchors.right: parent.right
                                anchors.margins: dp(16)
                                anchors.verticalCenter: parent.verticalCenter
                                text: Palette.colors[currentColor][modelData]
                                style: "body2"
                                color: Theme.lightDark(parent.color,
                                                       Theme.light.textColor,
                                                       Theme.dark.textColor)
                            }
                        }
                    }
                }
            }
        }
    }
}
