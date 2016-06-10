/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
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
    SilicaFlickable {
        anchors.fill: parent
        contentHeight: column.height + Theme.paddingLarge

        // Why is this necessary?
        contentWidth: parent.width

        VerticalScrollDecorator {}
        Column {
            id: column
            spacing: Theme.paddingLarge
            width: parent.width
            PageHeader {
                title: "Buttons"
            }
            SectionHeader {
                text: "Icon buttons"
            }
            Row {
                id: iconButtons
                spacing: Theme.paddingLarge
                anchors.horizontalCenter: parent.horizontalCenter
                property bool playing
                IconButton {
                    icon.source: "image://theme/icon-l-clear"
                }
                IconButton {
                    id: pause
                    icon.source: "image://theme/icon-l-pause"
                    onClicked: iconButtons.playing = false
                    enabled: iconButtons.playing
                }
                IconButton {
                    id: play
                    icon.source: "image://theme/icon-l-play"
                    onClicked: iconButtons.playing = true
                    enabled: !iconButtons.playing
                }
            }
            SectionHeader {
                text: "Text buttons"
            }
            Row {
                spacing: Theme.paddingLarge
                anchors.horizontalCenter: parent.horizontalCenter
                Button {
                    text: "Call"
                }
                Button {
                    text: "SMS"
                }
            }
            Button {
                text: "Medium"
                anchors.horizontalCenter: parent.horizontalCenter
                preferredWidth: Theme.buttonWidthMedium
            }
            Button {
                text: "Large (Disabled)"
                enabled: false
                anchors.horizontalCenter: parent.horizontalCenter
                preferredWidth: Theme.buttonWidthLarge
            }

            SectionHeader {
                text: "Icon switches"
            }
            Row {
                spacing: Theme.itemSizeSmall
                anchors.horizontalCenter: parent.horizontalCenter
                Switch {
                    icon.source: "image://theme/icon-m-shuffle"
                    enabled: false
                }
                Switch {
                    icon.source: "image://theme/icon-m-repeat"
                }
                Switch {
                    icon.source: "image://theme/icon-m-share"
                    onCheckedChanged: { busy = true; busyTimer.start() }
                    Timer {
                        id: busyTimer
                        interval: 4200
                        onTriggered: parent.busy = false
                    }
                }
            }
            SectionHeader {
                text: "Text switches"
            }
            Column {
                // No spacing in this column
                width: parent.width
                TextSwitch {
                    text: "Simple switch"
                }
                TextSwitch {
                    text: "Disabled switch"
                    enabled: false
                }
                TextSwitch {
                    text: "Switch with description"
                    description: "This switch has a description to explain the effect of the check action"
                }
                TextSwitch {
                    text: "This switch has a very long text property wrapping over multiple lines. The first line will be centered on the indicator."
                }
                TextSwitch {
                    text: "Switch with busy state"
                    onCheckedChanged: { busy = true; textBusyTimer.start() }
                    Timer {
                        id: textBusyTimer
                        interval: 4700
                        onTriggered: parent.busy = false
                    }
                }
                IconTextSwitch {
                    text: "Switch with an icon"
                    description: "This switch has both a textual label and an icon."
                    icon.source: "image://theme/icon-m-gps"
                }
            }
        }
    }
}
