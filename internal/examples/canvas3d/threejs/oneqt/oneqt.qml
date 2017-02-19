/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the QtCanvas3D module of the Qt Toolkit.
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
import QtCanvas3D 1.0
import QtQuick.Layouts 1.1

Item {
    id: mainview
    width: 1280
    height: 768
    visible: true
    focus: true

    Keys.onPressed: {
        if (event.key === Qt.Key_1) imageCube.state = 'image1';
        else if (event.key === Qt.Key_2) imageCube.state = 'image2';
        else if (event.key === Qt.Key_3) imageCube.state = 'image3';
        else if (event.key === Qt.Key_4) imageCube.state = 'image4';
        else if (event.key === Qt.Key_5) imageCube.state = 'image5';
        else if (event.key === Qt.Key_6) imageCube.state = 'image6';
    }

    //! [0]
    ImageCube {
        id: imageCube
        width: 512 * (parent.width / 1280)
        height: 512 * (parent.height / 768)
        anchors.bottom: parent.bottom
        anchors.right: parent.right
        //! [0]
        angleOffset: -180 / 8.0
        backgroundColor: "#FCFCFC"
        state: "image6"
        image1: "qrc:/textures/devices.png"
        image2: "qrc:/textures/iot.png"
        image3: "qrc:/textures/embedded.png"
        image4: "qrc:/textures/dataviz.jpg"
        image5: "qrc:/textures/multiscreen.png"
        image6: "qrc:/textures/puzzle-pieces.png"

        onStateChanged: {
            if (imageCube.state == "image1") {
                page1Button.selected = true;
                page2Button.selected = false;
                page3Button.selected = false;
                page4Button.selected = false;
                page5Button.selected = false;
                page6Button.selected = false;
                info1.visible = true;
                info2.visible = false;
                info4.visible = false;
                info3.visible = false;
                info5.visible = false;
                info6.visible = false;
            } else if (imageCube.state == "image2") {
                page1Button.selected = false;
                page2Button.selected = true;
                page3Button.selected = false;
                page4Button.selected = false;
                page5Button.selected = false;
                page6Button.selected = false;
                info1.visible = false;
                info2.visible = true;
                info4.visible = false;
                info3.visible = false;
                info5.visible = false;
                info6.visible = false;
            } else if (imageCube.state == "image3") {
                page1Button.selected = false;
                page2Button.selected = false;
                page3Button.selected = true;
                page4Button.selected = false;
                page5Button.selected = false;
                page6Button.selected = false;
                info1.visible = false;
                info2.visible = false;
                info3.visible = true;
                info4.visible = false;
                info5.visible = false;
                info6.visible = false;
            } else if (imageCube.state == "image4") {
                page1Button.selected = false;
                page2Button.selected = false;
                page3Button.selected = false;
                page4Button.selected = true;
                page5Button.selected = false;
                page6Button.selected = false;
                info1.visible = false;
                info2.visible = false;
                info3.visible = false;
                info4.visible = true;
                info5.visible = false;
                info6.visible = false;
            } else if (imageCube.state == "image5") {
                page1Button.selected = false;
                page2Button.selected = false;
                page3Button.selected = false;
                page4Button.selected = false;
                page5Button.selected = true;
                page6Button.selected = false;
                info1.visible = false;
                info2.visible = false;
                info3.visible = false;
                info4.visible = false;
                info5.visible = true;
                info6.visible = false;
            } else if (imageCube.state == "image6") {
                page1Button.selected = false;
                page2Button.selected = false;
                page3Button.selected = false;
                page4Button.selected = false;
                page5Button.selected = false;
                page6Button.selected = true;
                info1.visible = false;
                info2.visible = false;
                info3.visible = false;
                info4.visible = false;
                info5.visible = false;
                info6.visible = true;
            }
        }
    }

    Rectangle {
        id: menuBar
        anchors.top: parent.top
        width: parent.width
        height: 52
        color: "#ffffff"
        RowLayout {
            spacing: 28

            Image {
                id: qtLogo
                source: "qrc:/textures/qtlogosmall.png"
                Layout.minimumWidth: 133
                Layout.minimumHeight:52
                Layout.preferredWidth: 133
                Layout.preferredHeight: 52
                Layout.maximumWidth: 133
                Layout.maximumHeight: 52
            }

            Navibutton {
                id: page1Button
                text: "Device Creation"
                stateTarget: imageCube
                stateSelect: "image1"
                Layout.minimumHeight:52
                Layout.preferredHeight: 52
                Layout.maximumHeight: 52
            }

            Navibutton {
                id: page2Button
                text: "IoT"
                stateTarget: imageCube
                stateSelect: "image2"
                Layout.minimumHeight:52
                Layout.preferredHeight: 52
                Layout.maximumHeight: 52
            }

            Navibutton {
                id: page3Button
                text: "Rapid Development"
                stateTarget: imageCube
                stateSelect: "image3"
                Layout.minimumHeight:52
                Layout.preferredHeight: 52
                Layout.maximumHeight: 52
            }

            Navibutton {
                id: page4Button
                text: "Modern UX"
                stateTarget: imageCube
                stateSelect: "image4"
                Layout.minimumHeight:52
                Layout.preferredHeight: 52
                Layout.maximumHeight: 52
            }

            Navibutton {
                id: page5Button
                text: "Cross Platform"
                stateTarget: imageCube
                stateSelect: "image5"
                Layout.minimumHeight:52
                Layout.preferredHeight: 52
                Layout.maximumHeight: 52
            }

            Navibutton {
                id: page6Button
                text: "In the Box"
                stateTarget: imageCube
                stateSelect: "image6"
                Layout.minimumHeight:52
                Layout.preferredHeight: 52
                Layout.maximumHeight: 52
            }
        }
    }

    Rectangle {
        id: separator
        anchors.top: menuBar.bottom
        width: parent.width
        height: 4
        color: "#e6e6e6"
    }

    InfoSheet {
        id: info1
        width: parent.width
        anchors.top: separator.bottom
        anchors.left: parent.left
        visible: false
        headingText1: "Easily Create "
        headingText2: "Powerful & Connected Devices"
        text: "We believe modern embedded development must include a cross-platform user<br>"+
              "experience and that your tech strategy should be based on easy creation of<br>"+
              "connected devices and UIs that run anywhere on any embedded platform including<br>"+
              "RTOS – making your and your end users’ life easier. With Qt, you can do this and<br>"+
              "more."
    }

    InfoSheet {
        id: info2
        visible: false
        width: parent.width
        anchors.top: separator.bottom
        anchors.left: parent.left
        headingText1: "Write & Recycle "
        headingText2: "Internet of Things"
        text: "A key focus for us is to help you get your embedded device to market quickly. You<br>"+
              "can write and recycle Qt application and device UI code to run on all your target<br>"+
              "devices. You can take your applications everywhere: embedded, desktop and mobile<br>"+
              "platforms. Qt lets you future-proof your “things” by making them platform<br>"+
              "independent. Should you want diversity between platforms, like a responsive UI<br>"+
              "design for different screen sizes, this is simple to implement with Qt, as well."

    }

    InfoSheet {
        id: info3
        visible: false
        width: parent.width
        anchors.top: separator.bottom
        anchors.left: parent.left

        headingText1: "Rapid Embedded "
        headingText2: "Prototyping & Deployment"

        text: "We don’t want you wasting hundreds of man hours just setting up your embedded<br>"+
              "toolchains. Prototyping on a real device can start immediately upon installation with<br>"+
              "our fully pre-configured software stack, Boot to Qt. We provide full embedded<br>"+
              "tooling for direct device deployment, on-device debugging and profiling, and the<br>"+
              "needed tools to customize your stack."
    }

    InfoSheet {
        id: info4
        visible: false
        width: parent.width
        anchors.top: separator.bottom
        anchors.left: parent.left

        headingText1: "Modern UX with "
        headingText2: "Top Performance"
        text: "Your productivity is at the core of what drives us. We made creating embedded<br>"+
              "devices agile and painless without sacrificing maximum native performance. You get<br>"+
              "to write your application using high level C++ libraries with no need to worry about<br>"+
              "nasty platform details. Using Qt Creator IDE and with a variety of UI approaches to<br>"+
              "choose from you can create the optimal UX for your end users."
    }

    InfoSheet {
        id: info5
        visible: false
        width: parent.width
        anchors.top: separator.bottom
        anchors.left: parent.left

        headingText1: "Cross-platform "
        headingText2: "is Our Specialty"
        text: "We make cross-platform application development easy. Target all the screens in your<br>"+
              "end users’ lives. You only need to write and maintain one code base regardless of<br>"+
              "what kind of and how many target platforms you might have and we’re talking about<br>"+
              "all major operating systems here. No need for separate implementations for<br>"+
              "different user devices. Qt makes your time-to-market faster, technology strategy<br>"+
              "simpler and future-proof, consequently reducing costs."
    }

    InfoSheet {
        id: info6
        visible: false
        width: parent.width
        anchors.top: separator.bottom
        anchors.left: parent.left

        headingText1: "What’s in the Box "
        headingText2: "Everything You Need"
        text: "Qt combines functionality with productivity. You can amaze your users with stunning<br>"+
              "UIs and native performance by developing your desktop and multiscreen<br>"+
              "applications with Qt.<br>"+
              "<ul>"+
              "<li>C++ library classes – comprehensive, highly intuitive, and modularized</li>"+
              "<li>Declarative programming technology – exquisite UI design with Qt Quick</li>"+
              "<li>Tooling – productive and professional development with Qt Creator IDE</li>"+
              "</ul><br>"+
              "Qt saves you development time, adds efficiency and ultimately shortens your<br>"+
              "time-to-market."
    }
}
