/****************************************************************************
**
** Copyright (C) 2012 Digia Plc and/or its subsidiary(-ies).
** Contact: http://www.qt-project.org/legal
**
** This file is part of the QML Presentation System.
**
** $QT_BEGIN_LICENSE:LGPL$
** Commercial License Usage
** Licensees holding valid commercial Qt licenses may use this file in
** accordance with the commercial license agreement provided with the
** Software or, alternatively, in accordance with the terms contained in
** a written agreement between you and Digia.  For licensing terms and
** conditions see http://qt.digia.com/licensing.  For further information
** use the contact form at http://qt.digia.com/contact-us.
**
** GNU Lesser General Public License Usage
** Alternatively, this file may be used under the terms of the GNU Lesser
** General Public License version 2.1 as published by the Free Software
** Foundation and appearing in the file LICENSE.LGPL included in the
** packaging of this file.  Please review the following information to
** ensure the GNU Lesser General Public License version 2.1 requirements
** will be met: http://www.gnu.org/licenses/old-licenses/lgpl-2.1.html.
**
** In addition, as a special exception, Digia gives you certain additional
** rights.  These rights are described in the Digia Qt LGPL Exception
** version 1.1, included in the file LGPL_EXCEPTION.txt in this package.
**
** GNU General Public License Usage
** Alternatively, this file may be used under the terms of the GNU
** General Public License version 3.0 as published by the Free Software
** Foundation and appearing in the file LICENSE.GPL included in the
** packaging of this file.  Please review the following information to
** ensure the GNU General Public License version 3.0 requirements will be
** met: http://www.gnu.org/copyleft/gpl.html.
**
**
** $QT_END_LICENSE$
**
****************************************************************************/


import Qt.labs.presentation 1.0
import QtQuick 2.0

OpacityTransitionPresentation
{
    Slide {
        title: "Custom Transitions, 1"
        centeredText: "The Presentation {} Element has a 'switchSlides(from, to, forward)' function which will be called whenever a transition between slides should take place."
    }

    Slide {
        title: "Custom Transitions, 2"
        centeredText:
            "Re-implement this function and add your own transition for the slides. "
    }

    Slide {
        title: "Custom Transitions, 3"
        content: [
            "'from' slide",
            " opacity fades out",
            " scale increases slightly",
            "'to' slide",
            " opacity fades in",
            " scale increases from small to normal",
            "All with a fairly simple ParallelAnimation {}"
        ]
    }

    Slide {
        title: "Custom Transition, 4"
        centeredText: "The transition works whichever way you go...\nLets go back to the beginning..."
    }

}
