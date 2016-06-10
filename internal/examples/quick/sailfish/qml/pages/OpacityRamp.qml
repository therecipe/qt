/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Matt Vogt <matthew.vogt@jollamobile.com>
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
        contentHeight: column.height

        VerticalScrollDecorator {}
        Column {
            id: column
            width: parent.width

            PageHeader { title: "Opacity Ramp" }

            Row {
                spacing: Theme.paddingMedium
                anchors.horizontalCenter: parent.horizontalCenter

                IconButton {
                    icon.source: "image://theme/icon-m-left"
                    onClicked: {
                        effect.direction = OpacityRamp.RightToLeft
                        directionVal.text = "Direction: Right to Left"
                    }
                }
                IconButton {
                    icon.source: "image://theme/icon-m-down"
                    onClicked: {
                        effect.direction = OpacityRamp.TopToBottom
                        directionVal.text = "Direction: Top to Bottom"
                    }
                }
                IconButton {
                    icon.source: "image://theme/icon-m-up"
                    onClicked: {
                        effect.direction = OpacityRamp.BottomToTop
                        directionVal.text = "Direction: Bottom to Top"
                    }
                }
                IconButton {
                    icon.source: "image://theme/icon-m-right"
                    onClicked: {
                        effect.direction = OpacityRamp.LeftToRight
                        directionVal.text = "Direction: Left to Right"
                    }
                }
            }
            Label {
                id: directionVal
                anchors.horizontalCenter: parent.horizontalCenter
                text: "Direction: Left to Right"
            }
            Slider {
                id: slope
                width: parent.width
                anchors.horizontalCenter: parent.horizontalCenter
                value: 2.0
                stepSize: 0.02
                minimumValue: 0.5
                maximumValue: 4.0
            }
            Label {
                id: slopeVal
                anchors.horizontalCenter: parent.horizontalCenter
                text: "Slope: " + slope.value.toFixed(2)
            }
            Slider {
                id: offset
                width: parent.width
                value: 0.5
                stepSize: 0.01
            }
            Label {
                id: offsetVal
                anchors.horizontalCenter: parent.horizontalCenter
                text: "Offset: " + offset.value.toFixed(2)
            }
            Item {
                width: parent.width
                height: subject.height + Theme.itemSizeSmall

                Rectangle {
                    id: subject
                    anchors.centerIn: parent
                    width: childrenRect.width
                    height: childrenRect.height
                    color: Theme.secondaryHighlightColor

                    Image {
                        source: "image://theme/icon-launcher-component-gallery"
                        fillMode: Image.PreserveAspectFit
                        width: implicitWidth * 2
                        height: implicitHeight * 2
                    }
                }
                OpacityRampEffect {
                    id: effect
                    slope: slope.value
                    offset: offset.value
                    sourceItem: subject
                }
            }
            SectionHeader { text: "Label fading" }

            Item {
                height: Theme.paddingMedium; width: parent.width
            }

            Label {
                width: slider.value
                x: Theme.horizontalPageMargin
                truncationMode: TruncationMode.Fade
                text: "Abcdefghijlkmnopqrstuvxyz0123456789!W#$%&'()*+,-./:;<=>?\|"
            }
            Slider {
                id: slider

                minimumValue: 1
                maximumValue: parent.width - 2 * Theme.paddingLarge
                value: Math.round(maximumValue / 2)
                valueText: Math.round(value)
                width: parent.width
                label: "Label width"
            }
        }
    }
}
