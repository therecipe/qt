/*
 * QML Material - An application framework implementing Material Design.
 * Copyright (C) 2015 Michael Spencer <sonrisesoftware@gmail.com>
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
import Material 0.2
import Material.ListItems 0.1 as ListItem

Column {
    id: column

    anchors.fill: parent

    ListItem.Subheader {
        text: "Custom icons with different colors"
    }

    Row {
        Item {
            width: dp(56)
            height: dp(56)

            Icon {
                anchors.centerIn: parent
                source: Qt.resolvedUrl("images/weather-sunset.svg")
                color: Theme.light.iconColor
            }
        }

        Item {
            width: dp(56)
            height: dp(56)

            Icon {
                anchors.centerIn: parent
                source: Qt.resolvedUrl("images/weather-pouring.svg")
                color: Theme.light.iconColor
            }
        }

        Rectangle {
            width: dp(56)
            height: dp(56)
            color: "#333"
            Icon {
                anchors.centerIn: parent
                source: Qt.resolvedUrl("images/weather-sunset.svg")
                color: Theme.dark.iconColor
            }
        }

        Rectangle {
            width: dp(56)
            height: dp(56)
            color: "#333"
            Icon {
                anchors.centerIn: parent
                source: Qt.resolvedUrl("images/weather-pouring.svg")
                color: Theme.dark.iconColor
            }
        }
    }

    ListItem.Subheader {
        text: "Non-colorized custom icons"
    }

    Row {
        Item {
            width: dp(56)
            height: dp(40)

            Icon {
                anchors.centerIn: parent
                source: Qt.resolvedUrl("images/go-last.color.svg")
                color: Theme.light.iconColor
            }
        }

        Item {
            width: dp(56)
            height: dp(40)

            Icon {
                anchors.centerIn: parent
                source: Qt.resolvedUrl("images/list-add.color.svg")
                color: Theme.light.iconColor
            }
        }
    }

    ListItem.Subheader {
        text: "Custom icons in IconButtons"
    }

    Row {
        Item {
            width: dp(56)
            height: dp(40)

            IconButton {
                anchors.centerIn: parent
                iconSource: Qt.resolvedUrl("images/weather-sunset.svg")
                color: Theme.light.iconColor
            }
        }

        Item {
            width: dp(56)
            height: dp(40)

            IconButton {
                anchors.centerIn: parent
                iconSource: Qt.resolvedUrl("images/list-add.color.svg")
                color: Theme.light.iconColor
            }
        }
    }
}
