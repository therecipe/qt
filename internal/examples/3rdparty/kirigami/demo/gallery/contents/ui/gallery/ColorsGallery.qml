/*
 *   Copyright 2016 Aleix Pol Gonzalez <aleixpol@kde.org>
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

import QtQuick 2.0
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.0

ScrollablePage {
    title: "Colors"

   // leftPadding: 20//Units.gridUnit
    function endsWith(subject, predicate) {
        var idx = subject.indexOf(predicate);
        return idx + predicate.length == subject.length;
    }
    function getkeys() {
        var keys = [];
        for(var v in Theme) {
            if (endsWith(v, "Color"))
                keys.push(v);
        }
        keys.sort();
        return keys;
    }

    Column {
        Flow {
            id: view
            anchors.horizontalCenter: parent.horizontalCenter
            width: Math.floor(parent.width / (Units.gridUnit * 9)) * (Units.gridUnit * 9)
            Repeater {
                model: getkeys()

                delegate: ColumnLayout {
                    width: Units.gridUnit * 9
                    Rectangle {
                        Layout.alignment: Qt.AlignHCenter
                        width: Units.gridUnit * 7
                        height: Units.gridUnit * 3
                        color: Theme[modelData]
                        border {
                            width: 1
                            color: "black"
                        }
                    }
                    Label {
                        Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
                        text: modelData
                    }
                    Item {
                        width: 1
                        height: Units.smallSpacing
                    }
                }
            }
        }
    }
}
