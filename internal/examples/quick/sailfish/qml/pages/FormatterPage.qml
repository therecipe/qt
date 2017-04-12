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
    SilicaListView {
        anchors.fill: parent
        spacing: Theme.paddingMedium

        header: PageHeader {
            title: "Formatter"
        }

        model: ListModel {
        }

        section {
            property: 'section'

            delegate: SectionHeader {
                text: section
                height: Theme.itemSizeExtraSmall
            }
        }

        VerticalScrollDecorator {}

        delegate: Item {
            x: Theme.horizontalPageMargin
            width: parent.width - 2*Theme.horizontalPageMargin
            height: childrenRect.height

            Label {
                id: origin
                text: 'Samantha & The Spam Factory'
                font.pixelSize: Theme.fontSizeMedium
                truncationMode: TruncationMode.Fade
                anchors {
                    left: parent.left
                    right: date.left
                    rightMargin: Theme.paddingSmall
                }
            }
            Label {
                id: date
                text: Format.formatDate(dateTime, Formatter.TimepointRelative)
                font.pixelSize: Theme.fontSizeExtraSmall
                horizontalAlignment: Text.AlignRight
                anchors {
                    right: parent.right
                    baseline: origin.baseline
                }
            }
            Label {
                id: body
                text: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.'
                font.pixelSize: Theme.fontSizeExtraSmall
                wrapMode: Text.WordWrap
                maximumLineCount: 2
                truncationMode: TruncationMode.Fade
                anchors {
                    top: origin.bottom
                    left: parent.left
                    right: parent.right
                }
            }
            Label {
                function timestamp() {
                    var txt = Format.formatDate(dateTime, Formatter.Timepoint)
                    var elapsed = Format.formatDate(dateTime, Formatter.DurationElapsed)
                    return txt + (elapsed ? ' (' + elapsed + ')' : '')
                }

                id: timestampLabel
                text: timestamp()
                font.pixelSize: Theme.fontSizeExtraSmall
                font.italic: true
                anchors {
                    top: body.bottom
                    topMargin: Theme.paddingSmall
                    left: parent.left
                }
            }
            Label {
                text: "- " + fileSize
                font.pixelSize: Theme.fontSizeExtraSmall * 3 / 4
                font.italic: true
                anchors {
                    top: body.bottom
                    topMargin: Theme.paddingSmall
                    left: timestampLabel.right
                    right: parent.right
                }
            }
        }

        Component.onCompleted: {
            var offsets = [ 0,
                            3 * 60 * 1000, // 3 minutes ago
                            63 * 60 * 1000, // 63 minutes ago
                            6.5 * 60 * 60 * 1000, // 6 hours, 30 minutes ago
                            23.5 * 60 * 60 * 1000, // 23 hours, 30 minutes ago
                            24.25 * 60 * 60 * 1000, // 1 day, 15 minutes ago
                            50 * 60 * 60 * 1000, // 2 days, 2 hours ago
                            100 * 60 * 60 * 1000, // 4 days, 4 hours ago
                            6.25 * 24 * 60 * 60 * 1000, // 6 days, 6 hours ago
                            7.75 * 24 * 60 * 60 * 1000, // 7 days, 18 hours ago
                            20 * 24 * 60 * 60 * 1000, // 20 days ago
                            120 * 24 * 60 * 60 * 1000, // 120 days ago
                            364 * 24 * 60 * 60 * 1000, // 364 days ago
                            366 * 24 * 60 * 60 * 1000, // 366 days ago
                            1000 * 24 * 60 * 60 * 1000 ] // 1000 days ago
            var fileSizes = [ 0, // 0 B
                            99, // 99 B
                            100, // 0.1 kB
                            102347, // 99.9 kB
                            102400, // 0.1 MB
                            104805170, // 99.9 MB
                            1048523570, // 999.95MB-1
                            1048523571, // 999.95MB
                            107320495307, // 99.9 GB
                            1073688136907, // 999.95GB-1
                            1073688136908 ] // 999.95GB

            var now = new Date().getTime()
            for (var i = 0, fileSize = 0; i < offsets.length; ++i) {
                var dateTime = new Date(now - offsets[i])
                model.append({ 'dateTime': dateTime,
                               'section': Format.formatDate(dateTime, Formatter.TimepointSectionRelative),
                               'fileSize': Format.formatFileSize(fileSizes[fileSize++]) })
                fileSize %= fileSizes.length;
            }
        }
    }
}
