/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Matt Vogt <matt.vogt@jollamobile.com>
** All rights reserved.
** 
** This file is part of Sailfish Silica UI component package
**
** You may use this file under the terms of BSD license as follows:
**
** Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions are met:
**     * Redistributions of source code must retain the above copyright
**       notice, this list of conditions and the following disclaimer.
**     * Redistributions in binary form must reproduce the above copyright
**       notice, this list of conditions and the following disclaimer in the
**       documentation and/or other materials provided with the distribution.
**     * Neither the name of the Jolla Ltd nor the
**       names of its contributors may be used to endorse or promote products
**       derived from this software without specific prior written permission.
** 
** THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
** ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
** WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
** DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR
** ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
** (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
** LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
** ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
** SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**
****************************************************************************************/

import QtQuick 2.0
import Sailfish.Silica 1.0

Page {
    id: root

    property bool backStepping: true
    property int operationType: PageStackAction.Animated

    function switchOperationType() {
        operationType = (operationType === PageStackAction.Immediate ? PageStackAction.Animated : PageStackAction.Immediate)
    }

    Column {
        width: parent.width
        spacing: Theme.paddingLarge

        PageHeader { title: 'Page Stack' }

        Button {
            text: 'Push'
            anchors.horizontalCenter: parent.horizontalCenter
            onClicked: pageStack.push(nextPage, {}, operationType)
        }
        Column {
            // No spacing
            width: parent.width
            TextSwitch {
                anchors.horizontalCenter: parent.horizontalCenter
                text: 'Back-stepping'
                automaticCheck: false
                checked: backStepping
                onClicked: backStepping = !root.backStepping
            }
            TextSwitch {
                anchors.horizontalCenter: parent.horizontalCenter
                text: 'Immediate'
                automaticCheck: false
                checked: operationType === PageStackAction.Immediate
                onClicked: switchOperationType()
            }
        }
    }

    Component {
        id: nextPage
        Page {
            backNavigation: backStepping

            SilicaFlickable {
                anchors.fill: parent
                contentHeight: column.height

                Column {
                    id: column
                    width: parent.width
                    spacing: Theme.paddingLarge

                    PageHeader { title: 'Depth: ' + pageStack.depth }

                    Button {
                        text: 'Push'
                        anchors.horizontalCenter: parent.horizontalCenter
                        onClicked: pageStack.push(nextPage, {}, operationType)
                    }
                    Button {
                        text: 'Pop'
                        anchors.horizontalCenter: parent.horizontalCenter
                        onClicked: pageStack.pop(undefined, operationType)
                    }
                    Button {
                        text: 'Replace'
                        anchors.horizontalCenter: parent.horizontalCenter
                        onClicked: pageStack.replace(nextPage, {}, operationType)
                    }
                    Button {
                        text: 'Push attached'
                        anchors.horizontalCenter: parent.horizontalCenter
                        onClicked: pageStack.pushAttached(nextPage, {})
                    }
                    Button {
                        text: 'Pop attached'
                        anchors.horizontalCenter: parent.horizontalCenter
                        onClicked: pageStack.popAttached(undefined, operationType)
                        enabled: pageStack._currentContainer.attachedContainer || pageStack._currentContainer.attached
                        opacity: enabled ? 1.0 : 0.4
                        Behavior on opacity { FadeAnimation {} }
                    }
                    Button {
                        text: 'Move Up'
                        anchors.horizontalCenter: parent.horizontalCenter
                        onClicked: pageStack.navigateForward(operationType)
                        enabled: pageStack._currentContainer.attachedContainer
                        opacity: enabled ? 1.0 : 0.4
                        Behavior on opacity { FadeAnimation {} }
                    }
                    Button {
                        text: 'Move Down'
                        anchors.horizontalCenter: parent.horizontalCenter
                        onClicked: pageStack.navigateBack(operationType)
                        enabled: pageStack._currentContainer.attached
                        opacity: enabled ? 1.0 : 0.4
                        Behavior on opacity { FadeAnimation {} }
                    }
                    Column {
                        // No spacing
                        width: parent.width
                        TextSwitch {
                            anchors.horizontalCenter: parent.horizontalCenter
                            text: 'Back-stepping'
                            automaticCheck: false
                            checked: backStepping
                            onClicked: backStepping = !root.backStepping
                        }
                        TextSwitch {
                            anchors.horizontalCenter: parent.horizontalCenter
                            text: 'Immediate'
                            automaticCheck: false
                            checked: operationType === PageStackAction.Immediate
                            onClicked: switchOperationType()
                        }
                    }
                }
            }
        }
    }
}
