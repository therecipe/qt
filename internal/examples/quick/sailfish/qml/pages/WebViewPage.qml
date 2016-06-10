/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Raine Makelainen <raine.makelainen@jollamobile.com>
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
import QtWebKit 3.0
import Sailfish.Silica 1.0

Page {
    id: webViewPage

    onStatusChanged: {
        if (status == PageStatus.Active) {
            pageStack.pushAttached(attachedPage)
        }
    }

    SilicaWebView {
        id: webView
        anchors {
            top: parent.top
            left: parent.left
            right: parent.right
            bottom: navigationColumn.top
        }

        opacity: 0
        onLoadingChanged: {
            switch (loadRequest.status)
            {
            case WebView.LoadSucceededStatus:
                opacity = 1
                break
            case WebView.LoadFailedStatus:
                opacity = 0
                viewPlaceHolder.errorString = loadRequest.errorString
                break
            default:
                opacity = 0
                break
            }
        }

        FadeAnimation on opacity {}
        PullDownMenu {
            MenuItem {
                text: "Reload"
                onClicked: webView.reload()
            }
        }
    }

    ViewPlaceholder {
        id: viewPlaceHolder
        property string errorString

        enabled: webView.opacity === 0 && !webView.loading
        text: "Web content load error: " + errorString
        hintText: "Check network connectivity and pull down to reload"
    }

    Column {
        id: navigationColumn
        x: Theme.horizontalPageMargin
        width: parent.width - 2 * Theme.horizontalPageMargin
        anchors.bottom: urlField.top

        Label {
            text: "Back navigation " + (webViewPage.backNavigation ? "enabled." : "disabled.")
        }

        Label {
            text: "Forward navigation " + (webViewPage.forwardNavigation ? "enabled." : "disabled.")
        }
    }

    TextField {
        id: urlField
        anchors {
            left: parent.left
            right: parent.right
            bottom: parent.bottom
        }
        inputMethodHints: Qt.ImhUrlCharactersOnly
        text: "http://sailfishos.org"
        label: webView.title
        EnterKey.onClicked: {
            if (text.match(/http:\/\//) != null) {
                webView.url = text
            } else {
                webView.url = "http://" + text
            }
            parent.focus = true
        }
    }

    Component {
        id: attachedPage
        Page {
            Label {
                id: mainLabel
                x: Theme.paddingLarge
                width: parent.width - Theme.paddingLarge * 2
                horizontalAlignment: Text.AlignHCenter
                anchors.verticalCenter: parent.verticalCenter
                wrapMode: Text.Wrap
                font {
                    pixelSize: Theme.fontSizeLarge
                    family: Theme.fontFamilyHeading
                }
                color: Theme.primaryColor
                text: "This is an attached page and shows how forward navigation goes to disabled state based on web content size."
            }
        }
    }

    Component.onCompleted: {
        webView.url = urlField.text
    }
}
