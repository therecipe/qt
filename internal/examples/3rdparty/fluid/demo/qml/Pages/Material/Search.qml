/*
 * This file is part of Fluid.
 *
 * Copyright (C) 2017 Magnus Groß <magnus.gross21@gmail.com>
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
import Fluid.Controls 1.0
import Fluid.Material 1.0

Item {
    SearchBar {
        id: searchbar
        ListModel {
            id: searchResults
        }
        onSearchTextChanged: {
            if ("fluid".startsWith(searchText)) {
                searchSuggestions.append({text: "fluid"});
                searchSuggestions.append({text: "fluid demo"});
                searchSuggestions.append({text: "fluid interface"});
                searchSuggestions.append({text: "fluid layout"});
            } else if ("liri".startsWith(searchText)) {
                searchSuggestions.append({text: "liri"});
                searchSuggestions.append({text: "liri os"});
            }
        }
        onSearch: {
            searchResults.append({title: "example search result", body: "This text can go over multiple lines and automatically wraps accordingly. This section is used to provide a short description of the search result.", uri: query+"0"})
            searchResults.append({title: "second example search result", body: "Description", uri: query+"1"})
            searchResults.append({title: "third example search result", body: "Description", uri: query+"2"})
        }
    }
    ListView {
        id: resultsListView
        visible: searchResults.count != 0
        anchors.top: searchbar.bottom
        x: searchbar.x
        width: parent.width
        height: parent.height
        model: searchResults
        clip: true
        delegate: Item {
            width: parent.width
            height: 128
            Card {
                anchors.fill: parent
                anchors.margins: 8
                Ripple {
                    anchors.fill: parent
                    onClicked: console.log(model.uri)
                }
                Column {
                    anchors.fill: parent
                    anchors.margins: 16
                    spacing: 16
                    TitleLabel {
                        text: model.title
                    }
                    BodyLabel {
                        width: parent.width
                        text: model.body
                        wrapMode: Text.WordWrap
                    }
                }
            }
        }
    }
}
