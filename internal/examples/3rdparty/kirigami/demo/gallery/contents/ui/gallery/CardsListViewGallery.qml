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

    title: "List view of simple cards"

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
                text: "Cards List View"
            }
            Controls.ToolButton {
                text: "HIG..."
                enabled: false
                onClicked: Qt.openUrlExternally("")
            }
            Controls.ToolButton {
                text: "Source code..."
                onClicked: Qt.openUrlExternally("https://cgit.kde.org/kirigami.git/tree/examples/gallery/contents/ui/gallery/CardsListViewGallery.qml")
            }
        }

        Controls.Label {
            property int implicitWidth: Kirigami.Units.gridUnit * 25
            wrapMode: Text.WordWrap
            text: "The Kirigami types AbstractCard and Card are used to implement the popular Card pattern used on many mobile and web platforms that is used to display a collection of information or actions.\n Besides the Card components, Kirigami offers also 3 kinds of views and positioners to help to present cards with beautiful and responsive layouts.\n\nIn this page, CardsListView is used to do a list view of AbstractCard subclasses with a custom layout inside.\n CardsListView should be used only with cards which can look good at any horizontal size, so it is recommended to use directly AbstractCard with an appropriate layout inside, because they are stretching for the whole list width.\nTherefore is discouraged to use it with the Card type, unless it has Horizontal as headerOrientation.\n The choice between using this view with AbstractCard or a normal ListView with AbstractListItem/BasicListItem is purely a choice based on aestetics alone."
        }
    }

    Kirigami.CardsListView {
        id: view
        model: 100

        delegate: Kirigami.AbstractCard {
            //NOTE: never put a Layout as contentItem as it will cause binding loops
            //SEE: https://bugreports.qt.io/browse/QTBUG-66826
            contentItem: Item {
                implicitWidth: delegateLayout.implicitWidth
                implicitHeight: delegateLayout.implicitHeight
                GridLayout {
                    id: delegateLayout
                    anchors {
                        left: parent.left
                        top: parent.top
                        right: parent.right
                        //IMPORTANT: never put the bottom margin
                    }
                    rowSpacing: Kirigami.Units.largeSpacing
                    columnSpacing: Kirigami.Units.largeSpacing
                    columns: width > Kirigami.Units.gridUnit * 20 ? 4 : 2
                    Kirigami.Icon {
                        source: "applications-graphics"
                        Layout.fillHeight: true
                        Layout.maximumHeight: Kirigami.Units.iconSizes.huge
                        Layout.preferredWidth: height
                    }
                    ColumnLayout {
                        Kirigami.Heading {
                            level: 2
                            text: "Product "+ modelData
                        }
                        Kirigami.Separator {
                            Layout.fillWidth: true
                        }
                        Controls.Label {
                            Layout.fillWidth: true
                            wrapMode: Text.WordWrap
                            text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id risus id augue euismod accumsan."
                        }
                    }
                    Controls.Button {
                        Layout.alignment: Qt.AlignRight|Qt.AlignVCenter
                        Layout.columnSpan: 2 
                        text: "Install"
                        onClicked: showPassiveNotification("Install for Product " + modelData + " clicked");
                    }
                }
            }
        }
    }
}
