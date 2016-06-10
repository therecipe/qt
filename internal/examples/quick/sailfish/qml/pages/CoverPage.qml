/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Martin Jones <martin.jones@jollamobile.com>
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
    id: coverPage

    property string message: "Hello, Covers!"

    function displayAction(text) {
        message = text
        clearActionTimer.start()
    }
    onStatusChanged: {
        if (status == PageStatus.Deactivating) {
            mainWindow.cover = "private/DefaultCover.qml"
            oneCoverAction.enabled = false
            twoCoverActions.enabled = false
            changeCoverActionA.enabled = false
            changeCoverActionB.enabled = false
        }
    }

    Timer {
        id: clearActionTimer
        onTriggered: coverPage.message = "Hello, Covers!"
    }

    CoverActionList {
        id: twoCoverActions

        enabled: false

        CoverAction {
            iconSource: "image://theme/icon-cover-pause"
            onTriggered: displayAction("Pause")
        }

        CoverAction {
            iconSource: "image://theme/icon-cover-next"
            onTriggered: displayAction("Next")
        }
    }

    CoverActionList {
        id: oneCoverAction

        enabled: false

        CoverAction {
            property bool next: true
            iconSource: next ? "image://theme/icon-cover-next" : "image://theme/icon-cover-pause"
            onTriggered: {
                if (next) {
                    displayAction("Next")
                } else {
                    displayAction("Pause")
                }
                next = !next
            }
        }
    }

    CoverActionList {
        id: changeCoverActionA

        enabled: false

        CoverAction {
            iconSource: "image://theme/icon-cover-next"
            onTriggered: {
                changeCoverActionA.enabled = false
                changeCoverActionB.enabled = true
                displayAction("A")
            }
        }
    }

    CoverActionList {
        id: changeCoverActionB

        enabled: false

        CoverAction {
            iconSource: "image://theme/icon-cover-pause"
            onTriggered: {
                changeCoverActionB.enabled = false
                changeCoverActionA.enabled = true
                displayAction("B")
            }
        }
    }

    SilicaFlickable {
        anchors.fill: parent
        contentHeight: column.height
        Column {
            id: column
            width: parent.width
            PageHeader {
                title: "Covers"
            }

            ComboBox {
                id: coverCombo
                label: "Select cover:"
                width: parent.width
                menu: ContextMenu {
                    MenuItem {
                        // Do not set the default cover in an application,
                        // the path may change without prior notice.
                        // The default cover is included only to provide an idea
                        // what a coverless application will show in the switcher.
                        text: "Default cover (no actions)"
                        onClicked: {
                            mainWindow.cover = "private/DefaultCover.qml"
                        }
                    }
                    MenuItem {
                        text: "Component-based blue"
                        onClicked: {
                            mainWindow.cover = blueCover
                        }
                    }
                    MenuItem {
                        text: "Inline green"
                        onClicked: {
                            mainWindow.cover = greenCover
                        }
                    }
                    MenuItem {
                        text: "External file red (no actions)"
                        onClicked: {
                            mainWindow.cover = Qt.resolvedUrl("CoverExample.qml")
                        }
                    }
                    MenuItem {
                        id: changingCover
                        text: "Altering blue and green"
                    }
                    MenuItem {
                        text: "No cover"
                        onClicked: {
                            mainWindow.cover = null
                        }
                    }
                }
            }

            ComboBox {
                label: "Actions:"
                width: parent.width
                menu: ContextMenu {
                    MenuItem {
                        text: "No actions"
                        onClicked: {
                            oneCoverAction.enabled = false
                            twoCoverActions.enabled = false
                            changeCoverActionA.enabled = false
                            changeCoverActionB.enabled = false
                        }
                    }
                    MenuItem {
                        text: "One action (next)"
                        onClicked: {
                            oneCoverAction.enabled = true
                            twoCoverActions.enabled = false
                            changeCoverActionA.enabled = false
                            changeCoverActionB.enabled = false
                        }
                    }
                    MenuItem {
                        text: "Two actions (pause, next)"
                        onClicked: {
                            twoCoverActions.enabled = true
                            oneCoverAction.enabled = false
                            changeCoverActionA.enabled = false
                            changeCoverActionB.enabled = false
                        }
                    }
                    MenuItem {
                        text: "Dynamically changing actions"
                        onClicked: {
                            twoCoverActions.enabled = false
                            oneCoverAction.enabled = false
                            changeCoverActionB.enabled = false
                            changeCoverActionA.enabled = true
                        }
                    }
                }
            }
        }
    }

    Timer {
        property bool active: Qt.application.active
        repeat: true
        triggeredOnStart: true
        onActiveChanged: {
            if (coverCombo.currentItem == changingCover)
                running = !active
        }

        onTriggered: {
            if (mainWindow.cover !== greenCover)
                mainWindow.cover = greenCover
            else
                mainWindow.cover = blueCover
        }
    }

    Rectangle {
        id: greenCover

        objectName: "greenCover"
        visible: false // ApplicationWindow will override this
        anchors.fill: parent
        color: "green"
        Label {
            anchors.centerIn: parent
            color: "black"
            text: coverPage.message
        }
    }

    Component {
        id: blueCover

        Rectangle {
            anchors.fill: parent
            color: Qt.rgba(0,0,1,0.4)
            Label {
                anchors.centerIn: parent
                text: coverPage.message
            }
        }
    }
}

