/*
 *   Copyright 2016 Marco Martin <mart@kde.org>
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
import QtQuick.Layouts 1.2
import org.kde.kirigami 2.0 as Kirigami

Kirigami.AbstractApplicationWindow {
    id: root
    width: 500
    height: 800
    visible: true

    globalDrawer: Kirigami.GlobalDrawer {
        title: "Widget gallery"
        titleIcon: "applications-graphics"

        actions: [
            Kirigami.Action {
                text: "View"
                iconName: "view-list-icons"
                Kirigami.Action {
                        text: "action 1"
                }
                Kirigami.Action {
                        text: "action 2"
                }
                Kirigami.Action {
                        text: "action 3"
                }
            },
            Kirigami.Action {
                text: "Sync"
                iconName: "folder-sync"
                Kirigami.Action {
                        text: "action 4"
                }
                Kirigami.Action {
                        text: "action 5"
                }
            },
            Kirigami.Action {
                text: "Checkable"
                iconName: "view-list-details"
                checkable: true
                checked: false
                onTriggered: {
                    print("Action checked:" + checked)
                }
            },
            Kirigami.Action {
                text: "Settings"
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

    header: Kirigami.ApplicationHeader {}

    pageStack: Controls.StackView {
        anchors.fill: parent
        property int currentIndex: 0
        focus: true
        onCurrentIndexChanged: {
            if (depth > currentIndex+1) {
                pop(get(currentIndex));
            }
        }
        onDepthChanged: {
            currentIndex = depth-1;
        }
        initialItem: mainPageComponent

        Keys.onReleased: {
            if (event.key == Qt.Key_Back ||
            (event.key === Qt.Key_Left && (event.modifiers & Qt.AltModifier))) {
                event.accepted = true;
                if (root.contextDrawer && root.contextDrawer.drawerOpen) {
                    root.contextDrawer.close();
                } else if (root.globalDrawer && root.globalDrawer.drawerOpen) {
                    root.globalDrawer.close();
                } else {
                    var backEvent = {accepted: false}
                    if (root.pageStack.currentIndex >= 1) {
                        root.pageStack.currentItem.backRequested(backEvent);
                        if (!backEvent.accepted) {
                            if (root.pageStack.depth > 1) {
                                root.pageStack.currentIndex = Math.max(0, root.pageStack.currentIndex - 1);
                                backEvent.accepted = true;
                            } else {
                                Qt.quit();
                            }
                        }
                    }

                    if (!backEvent.accepted) {
                        Qt.quit();
                    }
                }
            }
        }
    }


    Component {
        id: settingsComponent
        Kirigami.Page {
            title: "Settings"
            objectName: "settingsPage"
            Rectangle {
                anchors.fill: parent
            }
        }
    }

    //Main app content
    Component {
        id: mainPageComponent
        MultipleColumnsGallery {}
    }

}
