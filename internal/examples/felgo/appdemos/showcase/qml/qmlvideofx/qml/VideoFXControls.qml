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

Item {
  id: root
  anchors.fill: parent

  property int itemWidth: dp(60)
  property int itemHeight: dp(36)
  property int scaledMargin: dp(4)
  property int fontSize: sp(12)
  readonly property real gripSize: dp(20)


  property VideoFXEffect targetEffect
  readonly property real parameterPanelHeight: parameterPanel.height + 2 * scaledMargin

  signal effectLoaded(string name, string source)

  Rectangle {
    color: "black"
    width: parent.width
    anchors.bottom: parent.bottom
    height: parameterPanelHeight
  }

  ParameterPanel {
    id: parameterPanel
    anchors {
      left: parent.left
      right: effectName.left
      bottom: parent.bottom
      margins: scaledMargin
    }
    gripSize: root.gripSize
  }

  Button {
    id: effectName
    anchors {
      right: parent.right
      bottom: parent.bottom
      margins: scaledMargin
    }

    text: "No effect"
    width: itemWidth * 2
    height: itemHeight
    onClicked: {
      effectName.visible = false
      listview.visible = true
      lvbg.visible = true
    }
    color: "#303030"
  }

  Rectangle {
    id: lvbg
    width: itemWidth * 2
    height: Math.min(parent.height, listview.contentHeight)
    color: "black"
    opacity: 0.8
    visible: false

    anchors {
      right: parent.right
      bottom: parent.bottom
      margins: scaledMargin
    }

    ListView {
      id: listview
      width: itemWidth * 2
      height: parent.height
      anchors.bottom: parent.bottom
      visible: false

      model: EffectSelectionList {}
      delegate: effectDelegate

      clip: true
      focus: true

      Component {
        id: effectDelegate
        Button {
          text: name
          width: itemWidth * 2
          onClicked: loadEffect(name, source)
        }
      }
    }
  }

  // load effect
  function loadEffect(name, source) {
    targetEffect.effectSource = source
    listview.visible = false
    lvbg.visible = false
    effectName.text = name
    effectName.visible = true
    parameterPanel.model = targetEffect.effect.parameters
    effectLoaded(name, source)
  }
}
