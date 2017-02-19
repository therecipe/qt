/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: http://www.qt.io/licensing/
**
** This file is part of the Qt3D module of the Qt Toolkit.
**
** $QT_BEGIN_LICENSE:LGPL3$
** Commercial License Usage
** Licensees holding valid commercial Qt licenses may use this file in
** accordance with the commercial license agreement provided with the
** Software or, alternatively, in accordance with the terms contained in
** a written agreement between you and The Qt Company. For licensing terms
** and conditions see http://www.qt.io/terms-conditions. For further
** information use the contact form at http://www.qt.io/contact-us.
**
** GNU Lesser General Public License Usage
** Alternatively, this file may be used under the terms of the GNU Lesser
** General Public License version 3 as published by the Free Software
** Foundation and appearing in the file LICENSE.LGPLv3 included in the
** packaging of this file. Please review the following information to
** ensure the GNU Lesser General Public License version 3 requirements
** will be met: https://www.gnu.org/licenses/lgpl.html.
**
** GNU General Public License Usage
** Alternatively, this file may be used under the terms of the GNU
** General Public License version 2.0 or later as published by the Free
** Software Foundation and appearing in the file LICENSE.GPL included in
** the packaging of this file. Please review the following information to
** ensure the GNU General Public License version 2.0 requirements will be
** met: http://www.gnu.org/licenses/gpl-2.0.html.
**
** $QT_END_LICENSE$
**
****************************************************************************/

import QtQuick 2.0
import QtQuick.Scene3D 2.0
import QtQuick.Layouts 1.2
import QtMultimedia 5.0

Item {
    id: mainview
    width: 1215
    height: 720
    visible: true
    property bool isHoverEnabled: false

    property variant magnitudeArray: null
    property int millisecondsPerBar: 68
    property string magnitudeDataSourceFile: "qrc:/music/visualization.raw"
    property int mediaLatencyOffset: 68

    state: "stopped"
    states: [
        State {
            name: "playing"
            PropertyChanges {
                target: playButtonImage
                source: {
                    if (playButtonMouseArea.containsMouse)
                        "qrc:/images/pausehoverpressed.png"
                    else
                        "qrc:/images/pausenormal.png"
                }
            }
            PropertyChanges {
                target: stopButtonImage
                source: "qrc:/images/stopnormal.png"
            }
        },
        State {
            name: "paused"
            PropertyChanges {
                target: playButtonImage
                source: {
                    if (playButtonMouseArea.containsMouse)
                        "qrc:/images/playhoverpressed.png"
                    else
                        "qrc:/images/playnormal.png"
                }
            }
            PropertyChanges {
                target: stopButtonImage
                source: "qrc:/images/stopnormal.png"
            }
        },
        State {
            name: "stopped"
            PropertyChanges {
                target: playButtonImage
                source: "qrc:/images/playnormal.png"
            }
            PropertyChanges {
                target: stopButtonImage
                source: "qrc:/images/stopdisabled.png"
            }
        }
    ]

    Component.onCompleted: isHoverEnabled = touchSettings.isHoverEnabled()

    //![0]
    MediaPlayer {
        id: mediaPlayer
        autoPlay: true
        volume: 0.5
        source: "qrc:/music/tiltshifted_lost_neon_sun.mp3"
        //![0]

        onStatusChanged: {
            if (status == MediaPlayer.EndOfMedia) //{
                state = "stopped"
        }

        onError: console.error("error with audio " + mediaPlayer.error)

        onDurationChanged: {
            // Load the pre-calculated magnitude data for the visualizer
            var request = new XMLHttpRequest()
            request.responseType = 'arraybuffer'
            request.onreadystatechange = function() {
                    if (request.readyState === XMLHttpRequest.DONE) {
                        if (request.status == 200 || request.status == 0) {
                            var arrayBuffer = request.response
                            if (arrayBuffer) {
                                magnitudeArray = new Uint16Array(arrayBuffer)
                                visualizer.startVisualization()
                              }
                        } else {
                            console.warn("Couldn't load magnitude data for bars.")
                        }
                        request = null
                    }
                };

            request.open('GET', magnitudeDataSourceFile, true)
            request.send(null)
        }

        function getNextAudioLevel(offsetMs) {
            if (magnitudeArray === null)
                return 0.0;

            // Calculate the integer index position in to the magnitude array
            var index = ((mediaPlayer.position + offsetMs) /
                         mainview.millisecondsPerBar) | 0;

            if (index < 0 || index >= magnitudeArray.length)
                return 0.0;

            return (magnitudeArray[index] / 63274.0);
        }
    }

    Image {
        id: coverImage
        anchors.fill: parent
        source: "qrc:/images/albumcover.png"
    }

    //![1]
    Scene3D {
        anchors.fill: parent

        Visualizer {
            id: visualizer
            animationState: mainview.state
            numberOfBars: 120
            barRotationTimeMs: 8160 // 68 ms per bar
        }
    }
    //![1]

    Rectangle {
        id: blackBottomRect
        color: "black"
        width: parent.width
        height: 0.14 * mainview.height
        anchors.bottom: parent.bottom
    }

    // Duration of played content
    Text {
        text: formatDuration(mediaPlayer.position)
        color: "#80C342"
        x: parent.width / 6
        y: mainview.height - mainview.height / 8
        font.pixelSize: 12
    }

    // Duration of the content left
    Text {
        text: "-" + formatDuration(mediaPlayer.duration - mediaPlayer.position)
        color: "#80C342"
        x: parent.width - parent.width / 6
        y: mainview.height - mainview.height / 8
        font.pixelSize: 12
    }

    function formatDuration(milliseconds) {
        var minutes = Math.floor(milliseconds / 60000)
        milliseconds -= minutes * 60000
        var seconds = milliseconds / 1000
        seconds = Math.round(seconds)
        if (seconds < 10)
            return minutes + ":0" + seconds
        else
            return minutes + ":" + seconds
    }

    property int buttonHorizontalMargin: 10
    Rectangle {
        id: playButton
        height: 54
        width: 54
        anchors.bottom: parent.bottom
        anchors.bottomMargin: width
        x: parent.width / 2 - width - buttonHorizontalMargin
        color: "transparent"

        Image {
            id: playButtonImage
            source: "qrc:/images/pausenormal.png"
        }

        MouseArea {
            id: playButtonMouseArea
            anchors.fill: parent
            hoverEnabled: isHoverEnabled
            onClicked: {
                if (mainview.state == 'paused' || mainview.state == 'stopped')
                    mainview.state = 'playing'
                else
                    mainview.state = 'paused'
            }
            onEntered: {
                if (mainview.state == 'playing')
                    playButtonImage.source = "qrc:/images/pausehoverpressed.png"
                else
                    playButtonImage.source = "qrc:/images/playhoverpressed.png"
            }
            onExited: {
                if (mainview.state == 'playing')
                    playButtonImage.source = "qrc:/images/pausenormal.png"
                else
                    playButtonImage.source = "qrc:/images/playnormal.png"
            }
        }
    }

    Rectangle {
        id: stopButton
        height: 54
        width: 54
        anchors.bottom: parent.bottom
        anchors.bottomMargin: width
        x: parent.width / 2 + buttonHorizontalMargin
        color: "transparent"

        Image {
            id: stopButtonImage
            source: "qrc:/images/stopnormal.png"
        }

        MouseArea {
            anchors.fill: parent
            hoverEnabled: isHoverEnabled
            onClicked: mainview.state = 'stopped'
            onEntered: {
                if (mainview.state != 'stopped')
                    stopButtonImage.source = "qrc:/images/stophoverpressed.png"
            }
            onExited: {
                if (mainview.state != 'stopped')
                    stopButtonImage.source = "qrc:/images/stopnormal.png"
            }
        }
    }
}
