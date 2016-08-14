/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Martin Jones <martin.jones@jollamobile.com>
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
    id: mainPage

    ListModel {
        id: pagesModel

        ListElement {
            page: "ButtonPage.qml"
            title: "Buttons"
            subtitle: "Assorted Button variants"
            section: "Controls"
        }
        ListElement {
            page: "ComboBoxPage.qml"
            title: "Combo Box"
            subtitle: "ComboBox component"
            section: "Controls"
        }
        ListElement {
            page: "MenuPage.qml"
            title: "Menus"
            subtitle: "Set of miscellaneous components"
            section: "Controls"
        }
        ListElement {
            page: "SliderPage.qml"
            title: "Slider"
            subtitle: "Slider can be used either as progress indicator or for selecting a value in a given range"
            section: "Controls"
        }
        ListElement {
            page: "ProgressPage.qml"
            title: "Progress Indicators"
            subtitle: "Progress indicator components"
            section: "Controls"
        }
        ListElement {
            page: "FormatterPage.qml"
            title: "Formatter"
            subtitle: "Text Formatting"
            section: "Text"
        }
        ListElement {
            page: "LabelPage.qml"
            title: "Labels"
            subtitle: "Assorted labels"
            section: "Text"
        }
        ListElement {
            page: "TextInputPage.qml"
            title: "Text Input"
            subtitle: "TextField and TextArea components"
            section: "Text"
        }
        ListElement {
            page: "WizardPage.qml"
            title: "Account creation wizard"
            subtitle: "Account creation wizard"
            section: "Text"
        }
        ListElement {
            page: "CoverPage.qml"
            title: "Cover View"
            subtitle: "Cover is visible in switcher"
            section: "View"
        }
        ListElement {
            page: "DialogPage.qml"
            title: "Dialogs"
            subtitle: "Assorted dialogs"
            section: "View"
        }
        ListElement {
            page: "OrientationPage.qml"
            title: "Orientation"
            subtitle: "How to manage the window orientation"
            section: "View"
        }
        ListElement {
            page: "PageStackPage.qml"
            title: "Page Stack"
            subtitle: "How to manage the page stack"
            section: "View"
        }
        ListElement {
            page: "InteractionHintPage.qml"
            title: "Touch hints"
            subtitle: "Animations to guide user"
            section: "View"
        }
        ListElement {
            page: "PanelPage.qml"
            title: "Panels and Sections"
            subtitle: "Panel and section components"
            section: "View"
        }

        ListElement {
            page: "WebViewPage.qml"
            title: "Web View"
            subtitle: "How to use web view"
            section: "View"
        }

        ListElement {
            page: "EffectPage.qml"
            title: "Effects"
            subtitle: "Glass effects and shaders"
            section: "Styling and Effects"
        }
        ListElement {
            page: "HapticPage.qml"
            title: "Haptics"
            subtitle: "Haptic Effects"
            section: "Styling and Effects"
        }
        ListElement {
            page: "OpacityRamp.qml"
            title: "Opacity Ramp"
            subtitle: "Opacity Ramp properties"
            section: "Styling and Effects"
        }
        ListElement {
            page: "SearchPage.qml"
            title: "Search"
            subtitle: "Search page shows how to implement in-application search"
            section: "Example"
        }
    }
    SilicaListView {
        id: listView
        anchors.fill: parent
        model: pagesModel
        header: PageHeader { title: "Components" }
        section {
            property: 'section'
            delegate: SectionHeader {
                text: section
            }
        }
        delegate: BackgroundItem {
            width: listView.width
            Label {
                id: firstName
                text: model.title
                color: highlighted ? Theme.highlightColor : Theme.primaryColor
                anchors.verticalCenter: parent.verticalCenter
                x: Theme.horizontalPageMargin
            }
            onClicked: pageStack.push(Qt.resolvedUrl(page))
        }
        VerticalScrollDecorator {}
    }
}





