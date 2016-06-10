/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Bea Lam <bea.lam@jollamobile.com>
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

    SilicaListView {
        anchors.fill: parent

        header: PageHeader {
            title: "Combo Box"
        }

        model: VisualItemModel {

            // set currentIndex to change the selected value
            ComboBox {
                width: page.width
                label: "Screen brightness"
                currentIndex: 1

                menu: ContextMenu {
                    MenuItem { text: "automatic" }
                    MenuItem { text: "manual" }
                    MenuItem { text: "Long Long Long Long selection" }
                }
            }

            ComboBox {
                width: page.width
                label: "A Long Long Long Long Long Label"


                menu: ContextMenu {
                    MenuItem { text: "Short" }
                    MenuItem { text: "Quite the Medium" }
                    MenuItem { text: "Long Long Long Long Long selection" }
                }
            }

            // if there are many options, they will open
            // in a separate dialog
            ComboBox {
                width: page.width
                label: "Setting (many options)"
                menu: ContextMenu {
                    MenuItem { text: "option a" }
                    MenuItem { text: "option b" }
                    MenuItem { text: "option c" }
                    MenuItem { text: "option d" }
                    MenuItem { text: "option e" }
                    MenuItem { text: "option f" }
                    MenuItem { text: "option g" }
                    MenuItem { text: "option h" }
                    MenuItem { text: "option i" }
                    MenuItem { text: "option j" }
                    MenuItem { text: "option k" }
                    MenuItem { text: "option l" }
                    MenuItem { text: "option m" }
                    MenuItem { text: "option n" }
                    MenuItem { text: "option o" }
                    MenuItem { text: "option p" }
                    MenuItem { text: "option q" }
                    MenuItem { text: "option r" }
                    MenuItem { text: "option s" }
                    MenuItem { text: "option t" }
                    MenuItem { text: "option u" }
                    MenuItem { text: "option v" }
                    MenuItem { text: "option w" }
                    MenuItem { text: "option x" }
                    MenuItem { text: "option y" }
                    MenuItem { text: "option z" }
                }
            }

            Row {
                ComboBox {
                    width: page.width / 2
                    label: "Mini combo 1"

                    menu: ContextMenu {
                        MenuItem { text: "an option" }
                        MenuItem { text: "yet another option" }
                        MenuItem { text: "yet another option with a long label" }
                    }
                }

                ComboBox {
                    width: page.width / 2
                    label: "Mini combo 2"

                    menu: ContextMenu {
                        MenuItem { text: "an option" }
                        MenuItem { text: "yet another option" }
                        MenuItem { text: "yet another option with a long label" }
                    }
                }
            }
            ComboBox {
                width: page.width
                label: "Preferred direction"
                currentIndex: 1

                menu: ContextMenu {
                    MenuItem { text: "Up" }
                    MenuItem { text: "Down" }
                    MenuItem { text: "Left" }
                    MenuItem { text: "Right" }
                }
                description: "This combobox comes with an extra description."
            }
        }
    }
}
