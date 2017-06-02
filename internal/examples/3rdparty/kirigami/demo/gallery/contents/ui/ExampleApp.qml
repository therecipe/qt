/*
 *   Copyright 2015 Marco Martin <mart@kde.org>
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

import QtQuick 2.1
import QtQuick.Controls 2.0 as Controls
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.1 as Kirigami
import "gallery"

Kirigami.ApplicationWindow {
    id: root

    header: Kirigami.ApplicationHeader {}

    globalDrawer: Kirigami.GlobalDrawer {
        title: "Widget gallery"
        titleIcon: "applications-graphics"
        bannerImageSource: "banner.jpg"

        actions: [
            Kirigami.Action {
                text: "Submenu 1"
                iconName: "view-list-icons"
                Kirigami.Action {
                        text: "Action 1"
                        onTriggered: showPassiveNotification(text + " clicked")
                }
                Kirigami.Action {
                        text: "Action 2"
                        onTriggered: showPassiveNotification(text + " clicked")
                }
                Kirigami.Action {
                        text: "Action 3"
                        onTriggered: showPassiveNotification(text + " clicked")
                }
            },
            Kirigami.Action {
                text: "Submenu 2"
                iconName: "folder-sync"
                Kirigami.Action {
                        text: "Action 4"
                        onTriggered: showPassiveNotification(text + " clicked")
                }
                Kirigami.Action {
                        text: "Action 5"
                        onTriggered: showPassiveNotification(text + " clicked")
                }
            },
            Kirigami.Action {
                text: "Checkable"
                iconName: "view-list-details"
                checkable: true
                checked: false
                onTriggered: {
                    showPassiveNotification("Action checked: " + checked)
                }
            },
            Kirigami.Action {
                text: "Open A Page"
                iconName: "configure"
                checkable: true
                //Need to do this, otherwise it breaks the bindings
                property bool current: pageStack.currentItem ? pageStack.currentItem.objectName == "settingsPage" : false
                onCurrentChanged: {
                    checked = current;
                }
                onTriggered: {
                    pageStack.push(settingsComponent);
                }
            }
            ]

        Controls.CheckBox {
            checked: true
            text: "Option 1"
        }
        Controls.CheckBox {
            text: "Option 2"
        }
        Controls.CheckBox {
            text: "Option 3"
        }
        Controls.Slider {
            Layout.fillWidth: true
            value: 0.5
        }
    }
    contextDrawer: Kirigami.ContextDrawer {
        id: contextDrawer
    }

    pageStack.initialPage: mainPageComponent

    Component {
        id: settingsComponent
        Kirigami.Page {
            title: "Settings"
            objectName: "settingsPage"
            Rectangle {
                anchors.fill: parent
                Controls.Button {
                    anchors.centerIn: parent
                    text: "Remove Page"
                    onClicked: applicationWindow().pageStack.pop();
                }
            }
        }
    }

    //Main app content
    Component {
        id: mainPageComponent
        MainPage {}
    }

}
