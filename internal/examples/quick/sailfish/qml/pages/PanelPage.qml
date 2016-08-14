/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Andrew den Exter <andrew.den.exter@jollamobile.com>
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

    DockedPanel {
        id: controlPanel

        width: page.isPortrait ? parent.width : Theme.itemSizeExtraLarge + Theme.paddingLarge
        height: page.isPortrait ? Theme.itemSizeExtraLarge + Theme.paddingLarge : parent.height

        dock: page.isPortrait ? Dock.Top : Dock.Left

        Flow {
            width: isPortrait ? undefined : Theme.itemSizeExtraLarge

            anchors.centerIn: parent

            Switch {
                icon.source: "image://theme/icon-m-shuffle"
            }
            Switch {
                icon.source: "image://theme/icon-m-repeat"
            }
            Switch {
                icon.source: "image://theme/icon-m-share"
            }
        }
    }

    Drawer {
        id: drawer

        anchors.fill: parent
        dock: page.isPortrait ? Dock.Top : Dock.Left

        background: SilicaListView {
            anchors.fill: parent
            model: 5

            header: PageHeader { title: "Drawer" }

            PullDownMenu {
                MenuItem {
                    text: "Option 1"
                }
                MenuItem {
                    text: "Option 2"
                }
            }
            VerticalScrollDecorator {}

            delegate: ListItem {
                id: listItem

                Label {
                    x: Theme.horizontalPageMargin
                    text: "List Item " + modelData
                    anchors.verticalCenter: parent.verticalCenter
                    color: listItem.highlighted ? Theme.highlightColor : Theme.primaryColor
                }
            }
        }

        SilicaFlickable {
            anchors {
                fill: parent
                leftMargin: page.isPortrait ? 0 : controlPanel.visibleSize
                topMargin: page.isPortrait ? controlPanel.visibleSize : 0
                rightMargin: page.isPortrait ? 0 : progressPanel.visibleSize
                bottomMargin: page.isPortrait ? progressPanel.visibleSize : 0
            }

            contentHeight: column.height + Theme.paddingLarge

            VerticalScrollDecorator {}

            MouseArea {
                enabled: drawer.open
                anchors.fill: column
                onClicked: drawer.open = false
            }

            Column {
                id: column
                spacing: Theme.paddingLarge
                width: parent.width
                enabled: !drawer.opened

                PageHeader { title: "Panels and sections" }

                SectionHeader {
                    text: "Panels"
                }
                Button {
                    text: controlPanel.open ? "Hide controls" : "Show controls"
                    onClicked: controlPanel.open = !controlPanel.open
                    anchors.horizontalCenter: parent.horizontalCenter
                }

                Button {
                    text: progressPanel.open ? "Hide progress" : "Show progress"
                    onClicked: progressPanel.open = !progressPanel.open
                    anchors.horizontalCenter: parent.horizontalCenter
                }

                Button {
                    text: "Open drawer"
                    onClicked: drawer.open = true
                    anchors.horizontalCenter: parent.horizontalCenter
                }

                SectionHeader {
                    text: "Expanding sections"
                }
                ExpandingSectionGroup {
                    currentIndex: 0

                    Repeater {
                        model: 5

                        ExpandingSection {
                            id: section

                            property int sectionIndex: model.index
                            title: "Section " + (model.index + 1)

                            content.sourceComponent: Column {
                                width: section.width

                                Repeater {
                                    model: (section.sectionIndex + 1) * 2

                                    TextSwitch {
                                        text: "Option " + (model.index + 1)
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    DockedPanel {
        id: progressPanel

        width: page.isPortrait ? parent.width : Theme.itemSizeExtraLarge + Theme.paddingLarge
        height: page.isPortrait ? Theme.itemSizeExtraLarge + Theme.paddingLarge : parent.height

        dock: page.isPortrait ? Dock.Bottom : Dock.Right

        ProgressCircle {
            id: progressCircle

            anchors.centerIn: parent

            NumberAnimation on value {
                from: 0
                to: 1
                duration: 1000
                running: progressPanel.expanded
                loops: Animation.Infinite
            }
        }
    }
}
