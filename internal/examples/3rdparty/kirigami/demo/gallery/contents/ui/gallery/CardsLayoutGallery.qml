/*
 *   Copyright 2018 Marco Martin <mart@kde.org>
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
import QtQuick.Controls 2.0 as Controls
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.4 as Kirigami

Kirigami.ScrollablePage {
    id: page

    title: "Cards Layout"

    actions.main: Kirigami.Action {
        iconName: "documentinfo"
        text: "Info"
        checkable: true
        onCheckedChanged: sheet.sheetOpen = checked;
        shortcut: "Alt+I"
    }

    Kirigami.OverlaySheet {
        id: sheet
        onSheetOpenChanged: page.actions.main.checked = sheetOpen
        header: RowLayout {
            Kirigami.Heading {
                Layout.fillWidth: true
                text: "Cards Layout"
            }
            Controls.ToolButton {
                text: "HIG..."
                enabled: false
                onClicked: Qt.openUrlExternally("")
            }
            Controls.ToolButton {
                text: "Source code..."
                onClicked: Qt.openUrlExternally("https://cgit.kde.org/kirigami.git/tree/examples/gallery/contents/ui/gallery/CardsLayoutGallery.qml")
            }
        }

        Controls.Label {
            property int implicitWidth: Kirigami.Units.gridUnit * 25
            wrapMode: Text.WordWrap
            text: "The Kirigami types AbstractCard and Card are used to implement the popular Card pattern used on many mobile and web platforms that is used to display a collection of information or actions.\n Besides the Card components, Kirigami offers also 3 kinds of views and positioners to help to present cards with beautiful and responsive layouts.\n\nIn this page, CardsLayout is presented, which should be used when the cards are not instantiated by a model or by a model which has always very few items (In the case of a big model CardsListView or CardsGridview should be used instead). They are presented as a grid of two columns which will remain centered if the application is really wide, or become a single column if there is not enough space for two columns, such as a mobile phone screen.\nA CardsLayout should always be contained within a ColumnLayout."
        }
    }

    ColumnLayout {
        Kirigami.CardsLayout {
            id: layout
            Kirigami.AbstractCard {
                Layout.fillHeight: true
                header: Kirigami.Heading {
                    text: "AbstractCard"
                    level: 2
                }
                contentItem: Controls.Label {
                    wrapMode: Text.WordWrap
                    text: "An AbstractCard is the simplest form of card. It's just a rectangle with a shadow, which can contain any Item in it. It can also have items assigned to the Header or Footer properties. In this case a Kirigami.Heading is its header and a Label with WordWrap is the contentItem."
                }
            }

            Kirigami.AbstractCard {
                showClickFeedback: true
                contentItem: Controls.Label {
                    wrapMode: Text.WordWrap
                    text: "This is an AbstractCard with a Label with WordWrap in it and nothing else, it's the simplest form Cards can be found in.\nAn AbstractCard can be clicked itself, with the usual onClicked signal handler and the showClickFeedback property can be used if the click should show any kind of visual feedback. It is recommended to set it to true if you plan to make the card reactive on any kind of click."
                }
                onClicked: showPassiveNotification("Card clicked")
            }

            Kirigami.Card {
                actions: [
                    Kirigami.Action {
                        text: "Action1"
                        icon.name: "add-placemark"
                    },
                    Kirigami.Action {
                        text: "Action2"
                        icon.name: "address-book-new-symbolic"
                    }
                ]
                banner {
                    imageSource: "../banner.jpg"
                    title: "Card"
                }
                contentItem: Controls.Label {
                    wrapMode: Text.WordWrap
                    text: "This is an instance of the Card type: it can optionally have an image, a title and an icon assigned to its banner group property, one or all of the properties together. A Card can also have Actions that will appear in the footer."
                }
            }


            Kirigami.Card {
                actions: [
                    Kirigami.Action {
                        text: "Action1"
                        icon.name: "add-placemark"
                    },
                    Kirigami.Action {
                        text: "Action2"
                        icon.name: "address-book-new-symbolic"
                    },
                    Kirigami.Action {
                        text: "Action3"
                        icon.name: "add-placemark"
                    },
                    Kirigami.Action {
                        text: "Action4"
                        icon.name: "address-book-new-symbolic"
                    },
                    Kirigami.Action {
                        text: "Action5"
                        icon.name: "add-placemark"
                    },
                    Kirigami.Action {
                        text: "Action6"
                        icon.name: "address-book-new-symbolic"
                    }
                ]
                banner {
                    imageSource: "../banner.jpg"
                    title: "Title Alignment"
                    titleAlignment: Qt.AlignLeft | Qt.AlignBottom
                }
                contentItem: Controls.Label {
                    wrapMode: Text.WordWrap
                    text: "The title can be aligned to all corners or centered with a combination of Qt.Alignment flags.\n When there are too many actions, they go in an overflow menu."
                }
            }

            Kirigami.Card {
                actions: [
                    Kirigami.Action {
                        text: "Action1"
                        icon.name: "add-placemark"
                    },
                    Kirigami.Action {
                        text: "Action2"
                        icon.name: "address-book-new-symbolic"
                    },
                    Kirigami.Action {
                        text: "Action3"
                        icon.name: "add-placemark"
                    }
                ]
                banner {
                    iconSource: "applications-graphics"
                    title: "Title only"
                }
                contentItem: Controls.Label {
                    wrapMode: Text.WordWrap
                    text: "The Banner can contain only the title and/or an icon, even if there is no banner image."
                }
            }

            Kirigami.Card {
                banner.imageSource: "../banner.jpg"

                header: Rectangle {
                    color: Qt.rgba(0,0,0,0.3)
                    implicitWidth: headerLayout.implicitWidth
                    implicitHeight: headerLayout.implicitHeight - avatarIcon.height/2
                    ColumnLayout {
                        id: headerLayout
                        anchors {
                            left: parent.left
                            right: parent.right
                        }
                        Controls.Label {
                            Layout.fillWidth: true
                            padding: Kirigami.Units.largeSpacing

                            color: "white"
                            wrapMode: Text.WordWrap
                            text: "It's possible to have custom contents overlapping the image, for cases where a more personalised design is needed."
                        }
                        Rectangle {
                            id: avatarIcon
                            color: "steelblue"
                            radius: width
                            Layout.alignment: Qt.AlignHCenter
                            Layout.preferredWidth: Kirigami.Units.iconSizes.huge
                            Layout.preferredHeight: Kirigami.Units.iconSizes.huge
                        }
                    }
                }
                contentItem: Controls.Label {
                    wrapMode: Text.WordWrap
                    topPadding: avatarIcon.height/2
                    text: "It's possible to customize the look and feel for Cards too, if the no padding behavior for headers is needed. This is usually discouraged in order to have greater consistency, but in some cases the design requires a more fancy layout, as shown in this example of a Card. If a custom header is used, the title and icon in the banner property shouldn't be used. If a custom footer is used (which is discouraged), actions shouldn't be used."
                }
                footer: RowLayout {
                    Controls.Label {
                        Layout.fillWidth: true
                        text: "Custom footer"
                    }
                    Controls.Button {
                        text: "Ok"
                    }
                }
            }

            Kirigami.Card {
                headerOrientation: Qt.Horizontal
                actions: [
                    Kirigami.Action {
                        text: "Action1"
                        icon.name: "add-placemark"
                    },
                    Kirigami.Action {
                        text: "Action2"
                        icon.name: "address-book-new-symbolic"
                    }
                ]
                banner {
                    imageSource: "../banner.jpg"
                    title: "Title"
                }
                contentItem: Controls.Label {
                    wrapMode: Text.WordWrap
                    text: "A card can optionally have horizontal orientation.\n In this case will be wider than tall, so is fit to be used also in a ColumnLayout.\nIf you need to put it in a CardsLayout, it will have by default a columnSpan of 2 (which can be overridden)."
                }
            }
        }
    }
}
