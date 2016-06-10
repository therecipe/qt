/****************************************************************************************
**
** Copyright (C) 2014 Jolla Ltd.
** Contact: Joona Petrell <joona.petrell@jollamobile.com>
** All rights reserved.
** 
** This file is part of Sailfish Silica UI component package.
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

    onStatusChanged: {
        if (status == PageStatus.Active) {
            pageStack.pushAttached(attachedPage)
        }
    }
    SilicaFlickable {
        anchors.fill: parent
        contentHeight: column.height
        PullDownMenu {
            MenuItem {
                text: "Clear"
                onClicked: {
                    swipeCombobox.currentIndex = 0
                    flickCombobox.currentIndex = 0
                    pullCombobox.currentIndex = 0
                    tapCombobox.currentIndex = 0
                }
            }
        }

        Column {
            id: column
            width: root.width
            PageHeader { title: "Touch hints" }
            ComboBox {
                id: swipeCombobox

                width: parent.width
                label: "Swipe hint"
                menu: ContextMenu {
                    onActivated: {
                        flickCombobox.currentIndex = 0
                        pullCombobox.currentIndex = 0
                        tapCombobox.currentIndex = 0
                    }

                    MenuItem { text: "None" }
                    MenuItem { text: "Top" }
                    MenuItem { text: "Bottom" }
                    MenuItem { text: "Left" }
                    MenuItem { text: "Right" }
                }

                onCurrentIndexChanged: update()

                function update() {
                    hint.stop()
                    tapHint.stop()

                    hint.interactionMode = TouchInteraction.EdgeSwipe

                    switch (currentIndex) {
                    case 0: // "None"
                        break
                    case 1: // "Top"
                        hint.direction = TouchInteraction.Down
                        hint.start()
                        break
                    case 2: // "Bottom"
                        hint.direction = TouchInteraction.Up
                        hint.start()
                        break
                    case 3: // "Left"
                        hint.direction = TouchInteraction.Right
                        hint.start()
                        break
                    case 4: // "Right"
                        hint.direction = TouchInteraction.Left
                        hint.start()
                        break
                    default:
                        console.log("Unknown option")
                        break
                    }
                }
            }
            ComboBox {
                id: flickCombobox

                currentIndex: 4
                width: parent.width
                label: "Flick hint"
                menu: ContextMenu {
                    onActivated: {
                        swipeCombobox.currentIndex = 0
                        pullCombobox.currentIndex = 0
                        tapCombobox.currentIndex = 0
                    }

                    MenuItem { text: "None" }
                    MenuItem { text: "Up" }
                    MenuItem { text: "Down" }
                    MenuItem { text: "Left" }
                    MenuItem { text: "Right" }
                }

                Component.onCompleted: update()
                onCurrentIndexChanged: update()

                function update() {
                    hint.stop()
                    tapHint.stop()

                    hint.interactionMode = TouchInteraction.Swipe

                    switch (currentIndex) {
                    case 0: // "None"
                        break
                    case 1: // "Up"
                        hint.direction = TouchInteraction.Up
                        hint.start()
                        break
                    case 2: // "Down"
                        hint.direction = TouchInteraction.Down
                        hint.start()
                        break
                    case 3: // "Left"
                        hint.direction = TouchInteraction.Left
                        hint.start()
                        break
                    case 4: // "Right"
                        hint.direction = TouchInteraction.Right
                        hint.start()
                        break
                    default:
                        console.log("Unknown option")
                        break
                    }
                }
            }
            ComboBox {
                id: pullCombobox

                width: parent.width
                label: "Pull hint"
                menu: ContextMenu {
                    onActivated: {
                        swipeCombobox.currentIndex = 0
                        flickCombobox.currentIndex = 0
                        tapCombobox.currentIndex = 0
                    }

                    MenuItem { text: "None" }
                    MenuItem { text: "Up" }
                    MenuItem { text: "Down" }
                }

                onCurrentIndexChanged: update()

                function update() {
                    hint.stop()
                    tapHint.stop()

                    hint.interactionMode = TouchInteraction.Pull

                    switch (currentIndex) {
                    case 0: // "None"
                        break
                    case 1: // "Up"
                        hint.direction = TouchInteraction.Up
                        hint.start()
                        break
                    case 2: // "Down"
                        hint.direction = TouchInteraction.Down
                        hint.start()
                        break
                    default:
                        console.log("Unknown option")
                        break
                    }
                }
            }
            ComboBox {
                id: tapCombobox

                width: parent.width
                label: "Tap hint"
                menu: ContextMenu {
                    onActivated: {
                        swipeCombobox.currentIndex = 0
                        flickCombobox.currentIndex = 0
                        pullCombobox.currentIndex = 0
                    }

                    MenuItem { text: "None" }
                    MenuItem { text: "Tap" }
                }

                onCurrentIndexChanged: update()

                function update() {
                    hint.stop()
                    tapHint.stop()

                    switch (currentIndex) {
                    case 0: // "None"
                        break
                    case 1: // "Tap"
                        tapHint.start()
                        break
                    default:
                        console.log("Unknown option")
                        break
                    }
                }
            }

            Item {
                width: parent.width
                height: Screen.height
            }
            InfoLabel { text: "More content :)" }
            Item {
                width: parent.width
                height: Screen.height/3
            }
        }

        VerticalScrollDecorator {}
    }
    InteractionHintLabel {
        anchors.bottom: parent.bottom
        opacity: (swipeCombobox._menuOpen || flickCombobox._menuOpen) ? 0.0 : 1.0
        Behavior on opacity { FadeAnimation {} }
        text: {
            switch (swipeCombobox.currentIndex) {
            case 1:
                return "Swipe from top to close the application"
            case 2:
                return "Swipe from bottom to access events view"
            case 3:
            case 4:
                return "Swipe from either edge to go back to home"
            }
            switch (flickCombobox.currentIndex) {
            case 1:
            case 2:
                return "Flick up and down to scroll for more content"
            case 3:
                return "Flick left to access attached page"
            case 4:
                return "Flick right to return to previous page"
            }
            switch (pullCombobox.currentIndex) {
            case 1:
                return "Pull up to access pulley menu"
            case 2:
                return "Pull down to access pulley menu"
            }
            switch (tapCombobox.currentIndex) {
            case 1:
                return "Tap"
            }
            return "Select example gesture hint"
        }
    }
    TouchInteractionHint {
        id: hint
        loops: Animation.Infinite
    }

    TapInteractionHint {
        id: tapHint
        loops: Animation.Infinite
        anchors.centerIn: parent
    }

    Page {
        id: attachedPage
        InfoLabel {
            text: "Attached page"
            anchors.centerIn: parent
        }
    }
}
