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

Kirigami.ApplicationWindow {
    id: root
    width: 500
    height: 800
    visible: true


    header: Kirigami.ApplicationHeader {
        headerStyle: Kirigami.ApplicationHeaderStyle.TabBar
    }



    pageStack.initialPage: mainPageComponent


    Component.onCompleted: {
        pageStack.push(mainPageComponent);
        pageStack.push(mainPageComponent);
        pageStack.currentIndex = 0;
    }

    //Main app content
    Component {
        id: mainPageComponent
        MultipleColumnsGallery {}
    }

}
