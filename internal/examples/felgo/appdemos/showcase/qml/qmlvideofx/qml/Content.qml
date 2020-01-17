/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the Qt Mobility Components.
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

import QtQuick 2.1

Rectangle {
    id: root
    property alias effect: effectLoader.item
    property alias gripSize: divider.gripSize
    property string effectSource
    property real volume: 0.5

    Divider {
        id: divider
        visible: false
        z: 1.0
        onValueChanged: updateDivider()
    }

    ShaderEffectSource {
        id: theSource
        smooth: true
        hideSource: true
    }

    Loader {
        id: contentLoader
    }

    Loader {
        id: effectLoader
        source: effectSource
    }

    onWidthChanged: {
        if (effectLoader.item)
            effectLoader.item.targetWidth = root.width
    }

    onHeightChanged: {
        if (effectLoader.item)
            effectLoader.item.targetHeight = root.height
    }

    onEffectSourceChanged: {
        effectLoader.source = effectSource
        effectLoader.item.parent = root
        effectLoader.item.targetWidth = root.width
        effectLoader.item.targetHeight = root.height
        updateSource()
        effectLoader.item.source = theSource
        divider.visible = effectLoader.item.divider
        updateDivider()
    }

    function init() {
        openImage("../../../assets/felgo-logo.png")
        root.effectSource = "EffectPassThrough.qml"
    }

    function updateDivider() {
        if (effectLoader.item && effectLoader.item.divider)
            effectLoader.item.dividerValue = divider.value
    }

    function updateSource() {
        if (contentLoader.item) {
            contentLoader.item.parent = root
            contentLoader.item.anchors.fill = root
            theSource.sourceItem = contentLoader.item
            if (effectLoader.item)
                effectLoader.item.anchors.fill = contentLoader.item
        }
    }

    function openImage(path) {
        stop()
        contentLoader.source = "ContentImage.qml"
        contentLoader.item.source = path
        updateSource()
    }

    function openVideo(path) {
        stop()
        contentLoader.source = "ContentVideo.qml"
        contentLoader.item.mediaSource = path
        contentLoader.item.volume = volume
        contentLoader.item.play()
        updateSource()
    }

    function openCamera() {
        stop()
        contentLoader.source = "ContentCamera.qml"
        updateSource()
    }

    function stop() {
        if (contentLoader.source == "ContentVideo.qml")
            contentLoader.item.stop()
        theSource.sourceItem = null
    }
}
