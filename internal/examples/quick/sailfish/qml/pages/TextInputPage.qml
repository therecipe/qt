/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Joona Petrell <joona.petrell@jollamobile.com>
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
    id: textInputPage

    property var textAlignment: undefined

    SilicaFlickable {
        anchors.fill: parent
        contentHeight: column.height + Theme.paddingLarge

        VerticalScrollDecorator {}

        Column {
            id: column
            width: parent.width

            PageHeader { title: "Text input" }

            SectionHeader {
                text: "Text fields"
            }

            TextField {
                focus: true
                label: "Normal"
                placeholderText: label
                width: parent.width
                horizontalAlignment: textAlignment
                EnterKey.iconSource: "image://theme/icon-m-enter-next"
                EnterKey.onClicked: passwordField.focus = true
            }

            PasswordField {
                id: passwordField
                width: parent.width
                horizontalAlignment: textAlignment
                EnterKey.iconSource: "image://theme/icon-m-enter-next"
                EnterKey.onClicked: numberField.focus = true
            }

            TextField {
                id: numberField
                width: parent.width
                inputMethodHints: Qt.ImhFormattedNumbersOnly
                label: "Number"
                placeholderText: label
                horizontalAlignment: textAlignment
                EnterKey.iconSource: "image://theme/icon-m-enter-next"
                EnterKey.onClicked: phoneField.focus = true
            }

            TextField {
                id: phoneField
                width: parent.width
                inputMethodHints: Qt.ImhDialableCharactersOnly
                label: "Phone number"
                placeholderText: label
                placeholderColor: Theme.highlightColor
                color: Theme.highlightColor
                horizontalAlignment: textAlignment
                EnterKey.iconSource: "image://theme/icon-m-enter-close"
                EnterKey.onClicked: focus = false
            }

            SectionHeader {
                text: "Search field"
            }

            SearchField {
                width: parent.width
            }

            SectionHeader {
                text: "Text area"
            }

            TextArea {
                id: textArea
                width: parent.width
                label: "Multi-line text"
                placeholderText: label
            }

            SectionHeader {
                text: "Enter key"
            }

            ComboBox {
                id: enterKeyCombobox
                readonly property bool iconMode: currentIndex === 1
                readonly property bool textMode: currentIndex === 2

                width: parent.width
                label: "Mode"
                currentIndex: 1

                menu: ContextMenu {
                    MenuItem { text: "None" }
                    MenuItem { text: "Icon" }
                    MenuItem { text: "Text" }
                }
            }

            Row {
                width: parent.width
                TextSwitch {
                    id: enabled; checked: true; text: "Enabled"; width: parent.width/2
                }
                TextSwitch {
                    id: highlighted; text: "Highlighted"; width: parent.width/2
                }
            }

            TextField {
                width: parent.width
                label: "Change options above"
                placeholderText: label
                horizontalAlignment: textAlignment
                focusOutBehavior: FocusBehavior.KeepFocus
                EnterKey.iconSource: enterKeyCombobox.iconMode ? "image://theme/icon-m-enter-next" : ""
                EnterKey.text: enterKeyCombobox.textMode ? "Text" : ""
                EnterKey.enabled: enabled.checked
                EnterKey.highlighted: highlighted.checked
                EnterKey.onClicked: parent.focus = true
            }

            SectionHeader {
                text: "Custom"
            }

            TextField {
                width: parent.width
                readOnly: true
                label: "Read only"
                text: label
                horizontalAlignment: textAlignment
            }

            TextField {
                width: parent.width
                readOnly: true
                focusOnClick: true
                label: "Read only, but focusable"
                text: label
                horizontalAlignment: textAlignment
                EnterKey.onClicked: parent.focus = true
            }

            TextField {
                width: parent.width
                inputMethodHints: Qt.ImhNoPredictiveText
                validator: RegExpValidator { regExp: /^[a-zA-Z]{3,}$/ }
                label: "Validated input"
                placeholderText: "Enter three or more characters"
                horizontalAlignment: textAlignment
                EnterKey.enabled: acceptableInput
                EnterKey.iconSource: "image://theme/icon-m-enter-next"
                EnterKey.onClicked: hiddenField.focus = true
            }

            TextField {
                id: hiddenField
                width: parent.width
                placeholderText: "Label hidden"
                horizontalAlignment: textAlignment
                labelVisible: false
                EnterKey.iconSource: "image://theme/icon-m-enter-close"
                EnterKey.onClicked: focus = false
            }

            SectionHeader {
                text: "Debug"
            }

            ComboBox {
                width: parent.width
                label: "Horizontal alignment"
                onCurrentIndexChanged: {
                    switch (currentIndex) {
                    case 0:
                        textAlignment = undefined
                        break
                    case 1:
                        textAlignment = TextInput.AlignLeft
                        break
                    case 2:
                        textAlignment = TextInput.AlignRight
                        break
                    case 3:
                        textAlignment = TextInput.AlignHCenter
                        break
                    }
                }

                menu: ContextMenu {
                    MenuItem { text: "Follows text" }
                    MenuItem { text: "Left" }
                    MenuItem { text: "Right" }
                    MenuItem { text: "HorizontalCenter" }
                }
            }
        }
    }
}
