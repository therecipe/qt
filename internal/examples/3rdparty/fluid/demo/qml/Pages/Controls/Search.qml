/*
 * This file is part of Fluid.
 *
 * Copyright (C) 2018 Magnus Groß <magnus.gross21@gmail.com>
 *
 * $BEGIN_LICENSE:MPL2$
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * $END_LICENSE$
 */

import QtQuick 2.10
import QtQuick.Controls 2.3
import QtQuick.Controls.Material 2.3
import Fluid.Controls 1.0 as FluidControls

Item {
    FluidControls.SearchBar {
        id: searchbar
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
        visible: searchbar.searchResults.count !== 0
        anchors.top: searchbar.bottom
        x: searchbar.x
        width: parent.width
        height: parent.height
        model: searchbar.searchResults
        clip: true
        delegate: Item {
            width: parent.width
            height: 128
            FluidControls.Card {
                anchors.fill: parent
                anchors.margins: 8
                FluidControls.Ripple {
                    anchors.fill: parent
                    onClicked: console.log(model.uri)
                }
                Column {
                    anchors.fill: parent
                    anchors.margins: 16
                    spacing: 16
                    FluidControls.TitleLabel {
                        text: model.title
                    }
                    FluidControls.BodyLabel {
                        width: parent.width
                        text: model.body
                        wrapMode: Text.WordWrap
                    }
                }
            }
        }
    }
}
