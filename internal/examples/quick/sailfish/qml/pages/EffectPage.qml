/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Mikko Harju <mikko.harju@jollamobile.com>
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
    id: page

    SilicaFlickable {
        contentHeight: column.height + Theme.paddingLarge*2
        anchors.fill: parent
        VerticalScrollDecorator {}
        Column {
            id: column
            spacing: Theme.paddingMedium
            y: Theme.paddingLarge
            width: page.width

            PageHeader {
                title: "Effects"
            }
            GlassItem {
                id: effect
                objectName: "menuitem"
                height: Theme.paddingLarge
                width: page.width
                falloffRadius: Math.exp(fpixw.value)
                radius: Math.exp(pixw.value)
                color: Theme.highlightColor
                cache: false
            }

            GlassItem {
                id: button
                anchors.horizontalCenter: parent.horizontalCenter
                width: 180
                height: 27
                falloffRadius: Math.exp(fpixw.value)
                radius: Math.exp(pixw.value)
                ratio: 0.0
                cache: false
            }
            GlassItem {
                id: slidertrack
                x: Math.round(Screen.width/8)
                width: page.width-2*x
                height: 27
                falloffRadius: Math.exp(fpixw.value)
                radius: Math.exp(pixw.value)
                dimmed: true
                ratio: 0.0
                cache: false
                GlassItem {
                    id: slider
                    property real _width : slidertrack.width
                    on_WidthChanged: {
                        widthAnimation.stop()
                        widthAnimation.to = _width
                        widthAnimation.start()
                    }
                    NumberAnimation on width {
                        id: widthAnimation
                        from: slider.height
                        duration: 5000
                        running: Qt.application.active
                        loops: Animation.Infinite
                    }
                    height: 27
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                    ratio: 0.0
                    cache: false
                }
            }
            Row {
                height: Theme.itemSizeMedium
                anchors.horizontalCenter: parent.horizontalCenter
                GlassItem {
                    dimmed: !true
                    falloffRadius: true ? undefined : 0.075
                    cache: false
                }
                GlassItem {
                    dimmed: !false
                    falloffRadius: false ? undefined : 0.075
                    cache: false
                }
            }
            Row {
                height: Theme.itemSizeMedium
                anchors.horizontalCenter: parent.horizontalCenter
                GlassItem { id: defaultItem }
                GlassItem {
                    scale: 0.2
                    rotation: 30
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                    cache: false
                }
                GlassItem {
                    width: 200
                    height: Theme.itemSizeMedium
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                    cache: false
                }
            }
            Row {
                height: Theme.itemSizeMedium
                anchors.horizontalCenter: parent.horizontalCenter
                GlassItem {
                    color: "red"
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                    cache: false
                }
                GlassItem {
                    color: "green"
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                    cache: false
                }
                GlassItem {
                    color: "blue"
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                    cache: false
                }
                GlassItem {
                    color: "white"
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                    cache: false
                }
                GlassItem {
                    color: Theme.highlightColor
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                    cache: false
                }
            }
            Row {
                anchors.horizontalCenter: parent.horizontalCenter
                GlassItem {
                    id: blink
                    property int c
                    Timer {
                        repeat: true
                        running: Qt.application.active
                        interval: 1000
                        onTriggered: blink.c++
                    }
                    onCChanged: {
                        var colors = ["red", "yellow", "green", "blue"]
                        blink.color = colors[c % colors.length]
                    }
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                }
                GlassItem {
                    id: dim
                    Timer {
                        repeat: true
                        running: Qt.application.active
                        interval: 1000
                        onTriggered: dim.dimmed = !dim.dimmed
                    }
                    falloffRadius: Math.exp(fpixw.value)
                    radius: Math.exp(pixw.value)
                }
            }
            Slider {
                id: pixw
                width: parent.width
                minimumValue: -10.0
                maximumValue: 0.0
                value: Math.log(defaultItem.radius)
                valueText: Math.exp(value).toFixed(4)
                label: "radius"
            }
            Slider {
                id: fpixw
                anchors.horizontalCenter: parent.horizontalCenter
                width: parent.width
                minimumValue: -10.0
                maximumValue: 0.0
                value: Math.log(defaultItem.falloffRadius)
                valueText: Math.exp(value).toFixed(4)
                label: "falloff"
            }
        }
    }
}
