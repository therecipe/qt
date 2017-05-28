/*
 * This file is part of Fluid.
 *
 * Copyright (C) 2017 Pier Luigi Fiorini <pierluigi.fiorini@gmail.com>
 *
 * $BEGIN_LICENSE:MPL2$
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * $END_LICENSE$
 */

import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import Fluid.Controls 1.0 as FluidControls
import Fluid.Material 1.0 as FluidMaterial

Item {
    Column {
        anchors.centerIn: parent

        Button {
            text: qsTr("List-style BottomSheet")
            onClicked: listBottomSheet.open()
        }

        Button {
            text: qsTr("Long List-style BottomSheet")
            onClicked: longListBottomSheet.open()
        }

        Button {
            text: qsTr("Grid-style BottomSheet")
            onClicked: gridBottomSheet.open()
        }

        Button {
            text: qsTr("Custom BottomSheet")
            onClicked: customBottomSheet.open()
        }
    }

    FluidMaterial.BottomSheetList {
        id: listBottomSheet
        title: qsTr("Save As")
        actions: [
            FluidControls.Action {
                text: qsTr("Folder")
                iconName: "file/folder"
            },
            FluidControls.Action {
                text: qsTr("New Folder")
                iconName: "file/create_new_folder"
            },
            FluidControls.Action {
                text: qsTr("Shared Folder")
                iconName: "file/folder_shared"
            },
            FluidControls.Action {
                text: qsTr("Cloud")
                iconName: "file/cloud"
            },
            FluidControls.Action {
                text: qsTr("Email Attachment")
                iconName: "file/attachment"
            },
            FluidControls.Action {
                text: qsTr("Upload")
                iconName: "file/file_upload"
            }
        ]
    }

    FluidMaterial.BottomSheetList {
        id: longListBottomSheet
        title: qsTr("Save As")
        actions: [
            FluidControls.Action {
                text: qsTr("Folder")
                iconName: "file/folder"
            },
            FluidControls.Action {
                text: qsTr("New Folder")
                iconName: "file/create_new_folder"
            },
            FluidControls.Action {
                text: qsTr("Shared Folder")
                iconName: "file/folder_shared"
            },
            FluidControls.Action {
                text: qsTr("Cloud")
                iconName: "file/cloud"
            },
            FluidControls.Action {
                text: qsTr("Email Attachment")
                iconName: "file/attachment"
            },
            FluidControls.Action {
                text: qsTr("Upload")
                iconName: "file/file_upload"
                hasDividerAfter: true
            },
            FluidControls.Action {
                text: qsTr("Placeholder 1")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 2")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 3")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 4")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 5")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 6")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 7")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 8")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 9")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 10")
                iconName: "file/cloud_done"
            }
        ]
    }

    FluidMaterial.BottomSheetGrid {
        id: gridBottomSheet
        actions: [
            FluidControls.Action {
                text: qsTr("Folder")
                iconName: "file/folder"
            },
            FluidControls.Action {
                text: qsTr("New Folder")
                iconName: "file/create_new_folder"
            },
            FluidControls.Action {
                text: qsTr("Shared Folder")
                iconName: "file/folder_shared"
            },
            FluidControls.Action {
                text: qsTr("Cloud")
                iconName: "file/cloud"
            },
            FluidControls.Action {
                text: qsTr("Email Attachment")
                iconName: "file/attachment"
            },
            FluidControls.Action {
                text: qsTr("Upload")
                iconName: "file/file_upload"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 1")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 2")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 3")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 4")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 5")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 6")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 7")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 8")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 9")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 10")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 11")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 12")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 13")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 14")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 15")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 16")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 17")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 18")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 19")
                iconName: "file/cloud_done"
            },
            FluidControls.Action {
                text: qsTr("Placeholder 20")
                iconName: "file/cloud_done"
            }
        ]
    }

    FluidMaterial.BottomSheet {
        id: customBottomSheet

        Column {
            width: parent.width

            Pane {
                width: parent.width
                padding: 16

                Column {
                    spacing: 8

                    FluidControls.TitleLabel {
                        text: "freedom"
                    }

                    FluidControls.BodyLabel {
                        text: "/ˈfriːdəm/"
                        color: Material.secondaryTextColor
                    }
                }

                Material.background: Material.color(Material.Yellow, Material.Shade800)
            }

            Pane {
                width: parent.width
                implicitHeight: 100
                padding: 16

                Column {
                    FluidControls.SubheadingLabel {
                        text: "noun"
                        color: Material.secondaryTextColor
                    }

                    FluidControls.BodyLabel {
                        text: "the right to live in the way you want without being controlled by anyone else"
                    }
                }
            }
        }
    }
}
