/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: http://www.qt.io/licensing/
**
** This file is part of the examples of the Qt Toolkit.
**
** $QT_BEGIN_LICENSE:BSD$
** You may use this file under the terms of the BSD license as follows:
**
** "Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions are
** met:
**   * Redistributions of source code must retain the above copyright
**     notice, this list of conditions and the following disclaimer.
**   * Redistributions in binary form must reproduce the above copyright
**     notice, this list of conditions and the following disclaimer in
**     the documentation and/or other materials provided with the
**     distribution.
**   * Neither the name of The Qt Company Ltd nor the names of its
**     contributors may be used to endorse or promote products derived
**     from this software without specific prior written permission.
**
**
** THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
** "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
** LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
** A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
** OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
** SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
** LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
** OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE."
**
** $QT_END_LICENSE$
**
****************************************************************************/

import QtQuick 2.6
import QtQuick.Layouts 1.1
import QtQuick.Controls 2.0

Pane {
    padding: 0

    property var delegateComponentMap: {
        "ItemDelegate": itemDelegateComponent,
        "SwipeDelegate": swipeDelegateComponent,
        "CheckDelegate": checkDelegateComponent,
        "RadioDelegate": radioDelegateComponent,
        "SwitchDelegate": switchDelegateComponent
    }

    Component {
        id: itemDelegateComponent

        ItemDelegate {
            text: labelText
            width: parent.width
        }
    }

    Component {
        id: swipeDelegateComponent

        SwipeDelegate {
            id: swipeDelegate
            text: labelText
            width: parent.width

            onClicked: if (swipe.complete) view.model.remove(ourIndex)

            Component {
                id: removeComponent

                Rectangle {
                    color: swipeDelegate.swipe.complete && swipeDelegate.pressed ? "#333" : "#444"
                    width: parent.width
                    height: parent.height
                    clip: true

                    Label {
                        font.pixelSize: swipeDelegate.font.pixelSize
                        text: "Remove"
                        color: "white"
                        anchors.centerIn: parent
                    }
                }
            }

            swipe.left: removeComponent
            swipe.right: removeComponent
        }
    }

    Component {
        id: checkDelegateComponent

        CheckDelegate {
            text: labelText
            width: parent.width
        }
    }

    ButtonGroup {
        id: radioButtonGroup
    }

    Component {
        id: radioDelegateComponent

        RadioDelegate {
            text: labelText
            ButtonGroup.group: radioButtonGroup
        }
    }

    Component {
        id: switchDelegateComponent

        SwitchDelegate {
            text: labelText
        }
    }

    ColumnLayout {
        id: column
        spacing: 40
        anchors.fill: parent
        anchors.topMargin: 20

        Label {
            Layout.fillWidth: true
            wrapMode: Label.Wrap
            horizontalAlignment: Qt.AlignHCenter
            text: "Delegate controls are used as delegates in views such as ListView."
        }

        ListView {
            id: listView
            Layout.fillWidth: true
            Layout.fillHeight: true
            clip: true
            model: ListModel {
                ListElement { type: "ItemDelegate"; text: "ItemDelegate" }
                ListElement { type: "ItemDelegate"; text: "ItemDelegate" }
                ListElement { type: "ItemDelegate"; text: "ItemDelegate" }
                ListElement { type: "SwipeDelegate"; text: "SwipeDelegate" }
                ListElement { type: "SwipeDelegate"; text: "SwipeDelegate" }
                ListElement { type: "SwipeDelegate"; text: "SwipeDelegate" }
                ListElement { type: "CheckDelegate"; text: "CheckDelegate" }
                ListElement { type: "CheckDelegate"; text: "CheckDelegate" }
                ListElement { type: "CheckDelegate"; text: "CheckDelegate" }
                ListElement { type: "RadioDelegate"; text: "RadioDelegate" }
                ListElement { type: "RadioDelegate"; text: "RadioDelegate" }
                ListElement { type: "RadioDelegate"; text: "RadioDelegate" }
                ListElement { type: "SwitchDelegate"; text: "SwitchDelegate" }
                ListElement { type: "SwitchDelegate"; text: "SwitchDelegate" }
                ListElement { type: "SwitchDelegate"; text: "SwitchDelegate" }
            }

            section.property: "type"
            section.delegate: Pane {
                width: listView.width
                height: sectionLabel.implicitHeight + 20

                Label {
                    id: sectionLabel
                    text: section
                    anchors.centerIn: parent
                }
            }

            delegate: Loader {
                id: delegateLoader
                width: listView.width
                sourceComponent: delegateComponentMap[text]

                property string labelText: text
                property ListView view: listView
                property int ourIndex: index

                // Can't find a way to do this in the SwipeDelegate component itself,
                // so do it here instead.
                ListView.onRemove: SequentialAnimation {
                    PropertyAction {
                        target: delegateLoader
                        property: "ListView.delayRemove"
                        value: true
                    }
                    NumberAnimation {
                        target: item
                        property: "height"
                        to: 0
                        easing.type: Easing.InOutQuad
                    }
                    PropertyAction {
                        target: delegateLoader
                        property: "ListView.delayRemove"
                        value: false
                    }
                }
            }
        }
    }
}
