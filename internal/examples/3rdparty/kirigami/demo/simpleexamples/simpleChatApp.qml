/*
 *   Copyright 2017 Marco Martin <mart@kde.org>
 *
 *   This program is free software; you can redistribute it and/or modify
 *   it under the terms of the GNU Library General Public License as
 *   published by the Free Software Foundation; either version 2 or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU Library General Public License for more details
 *
 *   You should have received a copy of the GNU Library General Public
 *   License along with this program; if not, write to the
 *   Free Software Foundation, Inc.,
 *   51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 */

import QtQuick 2.6
import QtQuick.Layouts 1.2
import QtQuick.Controls 2.2 as QQC2
import org.kde.kirigami 2.4 as Kirigami

Kirigami.ApplicationWindow {
    id: root

    header: Kirigami.ToolBarApplicationHeader {}
    //FIXME: perhaps the default logic for going widescreen should be refined upstream
    wideScreen: width > columnWidth * 3
    property int columnWidth: Kirigami.Units.gridUnit * 13
    property int footerHeight: Math.round(Kirigami.Units.gridUnit * 2.5)

    globalDrawer: Kirigami.GlobalDrawer {
        contentItem.implicitWidth: columnWidth
        title: "Chat App"
        titleIcon: "konversation"
        modal: true
        drawerOpen: false

        actions: [
            Kirigami.Action {
                text: "Rooms"
                iconName: "view-list-icons"
            },
            Kirigami.Action {
                text: "Contacts"
                iconName: "tag-people"
            },
            Kirigami.Action {
                text: "Search"
                iconName: "search"
            }
        ]
    }
    contextDrawer: Kirigami.OverlayDrawer {
        id: contextDrawer
        //they can depend on the page like that or be defined directly here
        edge: Qt.RightEdge
        modal: !root.wideScreen
        onModalChanged: drawerOpen = !modal
        handleVisible: applicationWindow == undefined ? false : applicationWindow().controlsVisible

        //here padding 0 as listitems look better without as opposed to any other control
        topPadding: 0
        bottomPadding: 0
        leftPadding: 0
        rightPadding: 0
        
        contentItem: ColumnLayout {
            
            readonly property int implicitWidth: root.columnWidth
            spacing: 0
            QQC2.Control {
                Layout.fillWidth: true
                background: Rectangle {
                    anchors.fill: parent
                    color: Kirigami.Theme.highlightColor
                    opacity: 0.8
                }

                padding: Kirigami.Units.gridUnit

                contentItem: ColumnLayout {
                    id: titleLayout
                    spacing: Kirigami.Units.gridUnit
                
                    RowLayout {
                        spacing: Kirigami.Units.gridUnit
                        Rectangle {
                            color: Kirigami.Theme.highlightedTextColor
                            radius: width
                            implicitWidth: Kirigami.Units.iconSizes.medium
                            implicitHeight: implicitWidth
                        }
                        ColumnLayout {
                            QQC2.Label {
                                Layout.fillWidth: true
                                color: Kirigami.Theme.highlightedTextColor
                                text: "KDE"
                            }
                            QQC2.Label {
                                Layout.fillWidth: true
                                color: Kirigami.Theme.highlightedTextColor
                                font.pointSize: Kirigami.Units.fontMetrics.font.pointSize * 0.8
                                text: "#kde: kde.org"
                            }
                        }
                    }
                    QQC2.Label {
                        Layout.fillWidth: true
                        color: Kirigami.Theme.highlightedTextColor
                        text: "Main room for KDE community, other rooms are listed at kde.org/rooms"
                        wrapMode: Text.WordWrap
                    }
                }
            }

            Kirigami.Separator {
                Layout.fillWidth: true
            }

            QQC2.ScrollView {
                Layout.fillWidth: true
                Layout.fillHeight: true
                ListView {
                    model: 50
                    delegate: Kirigami.BasicListItem {
                        label: "Person " + modelData
                        separatorVisible: false
                        reserveSpaceForIcon: false
                    }
                }
            }

            Kirigami.Separator {
                Layout.fillWidth: true
                Layout.maximumHeight: 1//implicitHeight
            }
            Kirigami.BasicListItem {
                label: "Group call"
                icon: "call-start"
                separatorVisible: false
            }
            Kirigami.BasicListItem {
                label: "Send Attachment"
                icon: "mail-attachment"
                separatorVisible: false
            }
        }
    }

    pageStack.defaultColumnWidth: columnWidth
    pageStack.initialPage: [channelsComponent, chatComponent]

    Component {
        id: channelsComponent
        Kirigami.ScrollablePage {
            title: "Channels"
            actions.main: Kirigami.Action {
                icon.name: "search"
                text: "Search"
            }
            background: Rectangle {
                anchors.fill: parent
                color: Kirigami.Theme.backgroundColor
            }
            footer: QQC2.ToolBar {
                height: root.footerHeight
                padding: Kirigami.Units.smallSpacing
                RowLayout {
                    anchors.fill: parent
                    spacing: Kirigami.Units.smallSpacing
                    //NOTE: icon support in tool button in Qt 5.11
                    QQC2.ToolButton {
                        Layout.fillHeight: true
                        //make it square
                        implicitWidth: height
                        Kirigami.Icon {
                            anchors.centerIn: parent
                            width: Kirigami.Units.iconSizes.smallMedium
                            height: width
                            source: "configure"
                        }
                        onClicked: root.pageStack.layers.push(secondLayerComponent);
                    }
                    QQC2.ComboBox {
                        Layout.fillWidth: true
                        Layout.fillHeight: true
                        model: ["First", "Second", "Third"]
                    }
                }
            }
            ListView {
                id: channelsList
                currentIndex: 2
                model: 30
                delegate: Kirigami.BasicListItem {
                    label: "#Channel " + modelData
                    checkable: true
                    checked: channelsList.currentIndex == index
                    separatorVisible: false
                    reserveSpaceForIcon: false
                }
            }
        }
    }

    Component {
        id: chatComponent
        Kirigami.ScrollablePage {
            title: "#KDE"
            actions {
                left: Kirigami.Action {
                    icon.name: "documentinfo"
                    text: "Channel info"
                }
                main: Kirigami.Action {
                    icon.name: "search"
                    text: "Search"
                }
            }
            actions.contextualActions: [
                Kirigami.Action {
                    text: "Room Settings"
                    iconName: "configure"
                    Kirigami.Action {
                        text: "Setting 1"
                    }
                    Kirigami.Action {
                        text: "Setting 2"
                    }
                },
                Kirigami.Action {
                    text: "Shared Media"
                    iconName: "document-share"
                    Kirigami.Action {
                        text: "Media 1"
                    }
                    Kirigami.Action {
                        text: "Media 2"
                    }
                    Kirigami.Action {
                        text: "Media 3"
                    }
                }
            ]
            background: Rectangle {
                anchors.fill: parent
                color: Kirigami.Theme.backgroundColor
            }
            footer: QQC2.Control {
                height: footerHeight
                padding: Kirigami.Units.smallSpacing
                background: Rectangle {
                    color: Kirigami.Theme.backgroundColor
                    Kirigami.Separator {
                        Rectangle {
                            anchors.fill: parent
                            color: Kirigami.Theme.focusColor
                            visible: chatTextInput.activeFocus
                        }
                        anchors {
                            left: parent.left
                            right: parent.right
                            top: parent.top
                        }
                    }
                }
                contentItem: RowLayout {
                    QQC2.TextField {
                        Layout.fillWidth: true
                        id: chatTextInput
                        background: Item {}
                    }
                    //NOTE: icon support in tool button in Qt 5.11
                    QQC2.ToolButton {
                        Layout.fillHeight: true
                        //make it square
                        implicitWidth: height
                        Kirigami.Icon {
                            anchors.centerIn: parent
                            width: Kirigami.Units.iconSizes.smallMedium
                            height: width
                            source: "go-next"
                        }
                    }
                }
            }

            ListView {
                id: channelsList
                verticalLayoutDirection: ListView.BottomToTop
                currentIndex: 2
                model: 30
                delegate: Item {
                    height: Kirigami.Units.gridUnit * 4
                    ColumnLayout {
                        x: Kirigami.Units.gridUnit
                        anchors.verticalCenter: parent.verticalCenter
                        QQC2.Label {
                            text: modelData % 2 ? "John Doe" : "John Applebaum"
                        }
                        QQC2.Label {
                            text: "Message " + modelData
                        }
                    }
                }
            }
        }
    }

    Component {
        id: secondLayerComponent
        Kirigami.Page {
            title: "Settings"
            background: Rectangle {
                color: Kirigami.Theme.backgroundColor
            }
            footer: QQC2.ToolBar {
                height: root.footerHeight
                QQC2.ToolButton {
                    Layout.fillHeight: true
                    //make it square
                    implicitWidth: height
                    Kirigami.Icon {
                        anchors.centerIn: parent
                        width: Kirigami.Units.iconSizes.smallMedium
                        height: width
                        source: "configure"
                    }
                    onClicked: root.pageStack.layers.pop();
                }
            }
        }
    }
}
