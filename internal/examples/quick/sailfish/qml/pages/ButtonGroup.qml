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

Item {
    id: root

    implicitWidth: childrenRect.width
    implicitHeight: childrenRect.height

    property Item checkedButton

    property int __silica_buttongroup

    property list<Item> _connectedChildren

    function reset() {
        var firstCheckable = null

        var existing = []
        if (_connectedChildren !== undefined) {
            for (var i = 0; i < _connectedChildren.length; ++i) {
                if (_connectedChildren[i] !== null) {
                    existing.push(_connectedChildren[i])
                }
            }
        }

        _traverseChildren(root, function(child) {
            if (child.hasOwnProperty('__silica_textswitch')) {
                var connected
                for (var i = 0; i < existing.length; ++i) {
                    if (existing[i] === child) {
                        connected = true
                        break
                    }
                }
                if (!connected) {
                    child.automaticCheck = false
                    child.clicked.connect(function(){ _childClicked(child) })
                    existing.push(child)
                }

                if (child.checked && checkedButton === null) {
                    checkedButton = child
                } else if (firstCheckable === null) {
                    firstCheckable = child
                }
            }
        })

        _connectedChildren = [ existing ]

        if (checkedButton === null && firstCheckable !== null) {
            checkedButton = firstCheckable
        }
        if (checkedButton) {
            _childClicked(checkedButton)
        }
    }

    function _childClicked(clickedChild) {
        checkedButton = clickedChild

        _traverseChildren(root, function(child) {
            if (child.hasOwnProperty('__silica_textswitch')) {
                child.checked = (child === checkedButton)
            }
        })
    }

    function _traverseChildren(obj, cb) {
        for (var i = 0; i < obj.children.length; ++i) {
            var child = obj.children[i]
            if (!child.hasOwnProperty('__silica_buttongroup')) {
                cb(child)
                if (child.hasOwnProperty('children') && (typeof child.children === 'object')) {
                    _traverseChildren(child, cb)
                }
            }
        }
    }

    Component.onCompleted: reset()
}
