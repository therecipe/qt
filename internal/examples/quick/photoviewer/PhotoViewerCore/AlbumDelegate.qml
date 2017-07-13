/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the QtQml module of the Qt Toolkit.
**
** $QT_BEGIN_LICENSE:BSD$
** Commercial License Usage
** Licensees holding valid commercial Qt licenses may use this file in
** accordance with the commercial license agreement provided with the
** Software or, alternatively, in accordance with the terms contained in
** a written agreement between you and The Qt Company. For licensing terms
** and conditions see https://www.qt.io/terms-conditions. For further
** information use the contact form at https://www.qt.io/contact-us.
**
** BSD License Usage
** Alternatively, you may use this file under the terms of the BSD license
** as follows:
**
** "Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions are
** met:
**   * Redistributions of source code must retain the above copyright
**     notice, this list of conditions and the following disclaimer.
**   * Redistributions in binary form must reproduce the above copyright
**     notice, this list of conditions and the following disclaimer in
**     the documentation and/or other materials provided with the
**     distribution.
**   * Neither the name of The Qt Company Ltd nor the names of its
**     contributors may be used to endorse or promote products derived
**     from this software without specific prior written permission.
**
**
** THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
** "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
** LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
** A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
** OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
** SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
** LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
** OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE."
**
** $QT_END_LICENSE$
**
****************************************************************************/

import QtQuick 2.0
import QtQuick.XmlListModel 2.0
import QtQml.Models 2.1

Component {
    id: albumDelegate
    Package {

        Item {
            Package.name: 'browser'
            GridView {
                id: photosGridView; model: visualModel.parts.grid; width: mainWindow.width; height: mainWindow.height - 21
                x: 0; y: 21; cellWidth: 160; cellHeight: 153; interactive: false
                onCurrentIndexChanged: photosListView.positionViewAtIndex(currentIndex, ListView.Contain)
            }
        }

        Item {
            Package.name: 'fullscreen'
            ListView {
                id: photosListView; model: visualModel.parts.list; orientation: Qt.Horizontal
                width: mainWindow.width; height: mainWindow.height; interactive: false
                onCurrentIndexChanged: photosGridView.positionViewAtIndex(currentIndex, GridView.Contain)
                highlightRangeMode: ListView.StrictlyEnforceRange; snapMode: ListView.SnapOneItem
            }
        }

        Item {
            Package.name: 'album'
            id: albumWrapper; width: 210; height: 220

            DelegateModel {
                id: visualModel; delegate: PhotoDelegate { }
                model: RssModel { id: rssModel; tags: tag }
            }

            BusyIndicator {
                id: busyIndicator
                anchors { centerIn: parent; verticalCenterOffset: -20 }
                on: rssModel.status != XmlListModel.Ready
            }

            PathView {
                id: photosPathView; model: visualModel.parts.stack; pathItemCount: 5
                visible: !busyIndicator.visible
                anchors.centerIn: parent; anchors.verticalCenterOffset: -30
                path: Path {
                    PathAttribute { name: 'z'; value: 9999.0 }
                    PathLine { x: 1; y: 1 }
                    PathAttribute { name: 'z'; value: 0.0 }
                }
            }

            MouseArea {
                anchors.fill: parent
                onClicked: mainWindow.editMode ? photosModel.remove(index) : albumWrapper.state = 'inGrid'
            }

            Tag {
                anchors { horizontalCenter: parent.horizontalCenter; bottom: parent.bottom; bottomMargin: 10 }
                frontLabel: tag; backLabel: qsTr("Remove"); flipped: mainWindow.editMode
                onTagChanged: rssModel.tags = tag
                onBackClicked: if (mainWindow.editMode) photosModel.remove(index);
            }

            states: [
            State {
                name: 'inGrid'
                PropertyChanges { target: photosGridView; interactive: true }
                PropertyChanges { target: albumsShade; opacity: 1 }
                PropertyChanges { target: backButton; onClicked: albumWrapper.state = ''; y: 6 }
            },
            State {
                name: 'fullscreen'; extend: 'inGrid'
                PropertyChanges { target: photosGridView; interactive: false }
                PropertyChanges { target: photosListView; interactive: true }
                PropertyChanges { target: photosShade; opacity: 1 }
                PropertyChanges { target: backButton; y: -backButton.height - 8 }
            }
            ]

            GridView.onAdd: NumberAnimation {
                target: albumWrapper; properties: "scale"; from: 0.0; to: 1.0; easing.type: Easing.OutQuad
            }
            GridView.onRemove: SequentialAnimation {
                PropertyAction { target: albumWrapper; property: "GridView.delayRemove"; value: true }
                NumberAnimation { target: albumWrapper; property: "scale"; from: 1.0; to: 0.0; easing.type: Easing.OutQuad }
                PropertyAction { target: albumWrapper; property: "GridView.delayRemove"; value: false }
            }

            transitions: [
            Transition {
                from: '*'; to: 'inGrid'
                SequentialAnimation {
                    NumberAnimation { properties: 'opacity'; duration: 250 }
                    PauseAnimation { duration: 350 }
                    NumberAnimation { target: backButton; properties: "y"; duration: 200; easing.type: Easing.OutQuad }
                }
            },
            Transition {
                from: 'inGrid'; to: '*'
                NumberAnimation { properties: "y,opacity"; easing.type: Easing.OutQuad; duration: 300 }
            }
            ]
        }
    }
}
