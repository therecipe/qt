/*
 * QML Material - An application framework implementing Material Design.
 * Copyright (C) 2015 Michael Spencer <sonrisesoftware@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as
 * published by the Free Software Foundation, either version 2.1 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */
import QtQuick 2.4
import Material 0.2
import Material.ListItems 0.1 as ListItem

TabbedPage {
    id: page
    title: "Page Title that is very long to demonstrate eliding titles in ActionBar"

    actionBar.backgroundColor: Palette.colors.grey['200']
    actionBar.decorationColor: Palette.colors.grey['300']

    onGoBack: {
        confirmationDialog.show()
        event.accepted = true
    }

    Tab {
        title: "Overview"
        iconName: "action/home"

        Rectangle {
            color: Palette.colors.red["200"]

            Button {
                anchors.centerIn: parent
                darkBackground: true
                text: "Go to tab 3"
                onClicked: page.selectedTab = 2
            }
        }
    }

    Tab {
        title: "Projects"

        Rectangle { color: Palette.colors.purple["200"] }
    }

    Tab {
        title: "Inbox"

        Rectangle { color: Palette.colors.orange["200"] }
    }

    Tab {
        title: "Disabled Tab"
        enabled: false
        Rectangle { color: Palette.colors.purple["200"] }
    }

    Dialog {
        id: confirmationDialog

        title: "Do you want to go back?"
        positiveButtonText: "Go Back"
        negativeButtonText: "Cancel"

        onAccepted: page.forcePop()
    }
}
