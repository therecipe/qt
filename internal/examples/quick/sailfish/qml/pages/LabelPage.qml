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
    id: page
    SilicaFlickable {
        anchors.fill: parent
        contentWidth: parent.width
        contentHeight: col.height + Theme.paddingLarge

        VerticalScrollDecorator {}

        Column {
            id: col
            spacing: Theme.paddingLarge
            width: parent.width
            PageHeader {
                title: "Labels"
            }
            SectionHeader {
                text: "Text styling"
            }
            Label {
                text: "Header text"
                anchors.horizontalCenter: parent.horizontalCenter
                font.family: Theme.fontFamilyHeading
            }
            Label {
                text: "Normal text"
                anchors.horizontalCenter: parent.horizontalCenter
            }
            Label {
                text: "Secondary text"
                color: Theme.secondaryColor
                anchors.horizontalCenter: parent.horizontalCenter
            }
            Label {
                text: "Highlighted text"
                color: Theme.highlightColor
                anchors.horizontalCenter: parent.horizontalCenter
            }
            Label {
                text: "Secondary highlighted text"
                color: Theme.secondaryHighlightColor
                anchors.horizontalCenter: parent.horizontalCenter
            }
            Rectangle {
                color: Theme.highlightBackgroundColor
                anchors.horizontalCenter: parent.horizontalCenter
                height: Theme.itemSizeSmall
                width: page.width
                Label {
                    text: "Normal text with background"
                    anchors.centerIn: parent
                }
            }
            SectionHeader {
                text: "Text eliding"
            }
            Label {
                text: "Text that should not be elided or faded out"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignLeft
                color: Theme.secondaryHighlightColor
            }
            Label {
                text: "Text that should be elided off the right end"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth*0.9)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignLeft
                truncationMode: TruncationMode.Elide
                color: Theme.secondaryHighlightColor
            }
            Label {
                text: "Text that should be elided off the left end"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth*0.9)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignRight
                truncationMode: TruncationMode.Elide
                color: Theme.secondaryHighlightColor
            }
            Label {
                text: "Text that should be faded out rather than elided"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth*0.9)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignLeft
                truncationMode: TruncationMode.Fade
                color: Theme.secondaryHighlightColor
            }
            Label {
                text: "Text that should be faded out from the left"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth*0.9)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignRight
                truncationMode: TruncationMode.Fade
                color: Theme.secondaryHighlightColor
            }
        }
    }
}
