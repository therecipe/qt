/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Vesa-Matti Hartikainen <vesa-matti.hartikainen@jollamobile.com>
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

        VerticalScrollDecorator {}
        Column {
            id: column
            spacing: Theme.paddingLarge
            width: parent.width
            PageHeader { title: "Slider" }
            Label {
                text: "Select value: " + userSlider.value.toFixed(3)
                anchors.horizontalCenter: parent.horizontalCenter
                color: Theme.highlightColor
                font.family: Theme.fontFamilyHeading
            }
            Slider {
                id: userSlider
                width: parent.width
                anchors.horizontalCenter: parent.horizontalCenter
                value: 0.5
            }
            Slider {
                id: changingSlider
                value: 30
                minimumValue: 0
                maximumValue: 300
                enabled: false
                width: parent.width
                handleVisible: highlightToggle.checked
                valueText : valueIndicatorToggle.checked ? value : ""
                label: labelToggle.checked ? labelToggle.label : ""
                Rectangle { anchors.fill: parent; color: "red"; opacity: 0.3; z:-1; visible: sizeToggle.checked }
            }
            Column {
                // No spacing in this column
                width: parent.width
                TextSwitch {
                    id: sliderUpdateToggle
                    text: "Update slider"

                    Timer {
                        interval: 100
                        repeat: true
                        onTriggered: changingSlider.value = (changingSlider.value + 10) % 300
                        running: sliderUpdateToggle.checked
                    }
                }
                TextSwitch {
                    id: highlightToggle
                    text: "Show handle"
                }
                TextSwitch {
                    id: valueIndicatorToggle
                    text: "Value indicator"
                }
                TextSwitch {
                    id: labelToggle
                    property string label: "ÄÖqgjÅAB"
                    text: "Label"
                }
                TextSwitch {
                    id: sizeToggle
                    text: "Show size"
                }
            }
            Slider {
                value: -50
                minimumValue:-40
                maximumValue:20
                stepSize: stepSwitch.checked ? 10 : 0.25
                width: parent.width
                valueText: value.toFixed(2)
                label: "Custom"

            }
            TextSwitch {
                id: stepSwitch
                text: "Bigger steps"
            }
        }
    }
}
