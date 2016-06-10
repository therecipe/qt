/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Chris Adams <chris.adams@jollamobile.com>
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
import QtFeedback 5.0
import Sailfish.Silica 1.0

Page {
    id: page

    ThemeEffect {
        id: buttonBuzz
        effect: ThemeEffect.Press
    }

    ThemeEffect {
        id: keypadBuzz
        effect: ThemeEffect.PressWeak
    }

    ThemeEffect {
        id: longBuzz
        effect: ThemeEffect.PressStrong
    }

    HapticsEffect {
        id: customBuzz
        intensity: intensitySlider.value
        duration: Math.round(durationSlider.value)
    }

    HapticsEffect {
        id: periodicBuzz
        intensity: 0.9
        duration: 2000
        attackTime: 500
        fadeTime: 1000
        attackIntensity: 0.5
        fadeIntensity: 0.1
        period: 20
    }

    SilicaFlickable {
        anchors.fill: parent
        contentHeight: column.height + Theme.paddingLarge
        contentWidth: parent.width

        VerticalScrollDecorator {}
        Column {
            id: column

            width: parent.width
            spacing: Theme.paddingLarge

            PageHeader {
                title: "Haptics"
            }

            SectionHeader {
                text: "Theme effects, onPressed"
            }
            Button {
                text: "Long"
                anchors.horizontalCenter: parent.horizontalCenter
                onPressed: longBuzz.play()
            }
            Button {
                text: "Button"
                anchors.horizontalCenter: parent.horizontalCenter
                onPressed: buttonBuzz.play()
            }
            Button {
                text: "Keypad"
                anchors.horizontalCenter: parent.horizontalCenter
                onPressed: keypadBuzz.play()
            }
            Button {
                text: "Periodic"
                anchors.horizontalCenter: parent.horizontalCenter
                onPressed: periodicBuzz.start()
            }
            SectionHeader {
                text: "Custom effects"
            }
            Label {
                x: Theme.horizontalPageMargin
                text: "Effect duration:"
            }
            Slider {
                id: durationSlider
                width: page.width
                value: 20
                valueText: Math.round(durationSlider.value)
                minimumValue: 1
                maximumValue: 500
            }
            Label {
                x: Theme.horizontalPageMargin
                text: "Effect intensity:"
            }
            Slider {
                id: intensitySlider
                width: page.width
                value: 0.5
                valueText: Math.round(intensitySlider.value)
                minimumValue: 0.1
                maximumValue: 1.0
            }
            Button {
                text: "onPressed"
                anchors.horizontalCenter: parent.horizontalCenter
                onPressed: customBuzz.start()
            }
            Button {
                anchors.horizontalCenter: parent.horizontalCenter
                text: "onReleased"
                onReleased: customBuzz.start()
            }
            Button {
                text: "onClicked"
                anchors.horizontalCenter: parent.horizontalCenter
                onClicked: customBuzz.start()
            }
            Button {
                text: "onPressed, onReleased"
                anchors.horizontalCenter: parent.horizontalCenter
                onPressed: customBuzz.start()
                onReleased: customBuzz.start()
            }
        }
    }
}
