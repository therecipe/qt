/****************************************************************************
**
** Copyright (C) 2015 The Qt Company Ltd.
** Contact: http://www.qt.io/licensing/
**
** This file is part of the examples of the Qt Toolkit.
**
** $QT_BEGIN_LICENSE:BSD$
** You may use this file under the terms of the BSD license as follows:
**
** "Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions are
** met:
**   * Redistributions of source code must retain the above copyright
**     notice, this list of conditions and the following disclaimer.
**   * Redistributions in binary form must reproduce the above copyright
**     notice, this list of conditions and the following disclaimer in
**     the documentation and/or other materials provided with the
**     distribution.
**   * Neither the name of The Qt Company Ltd nor the names of its
**     contributors may be used to endorse or promote products derived
**     from this software without specific prior written permission.
**
**
** THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
** "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
** LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
** A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
** OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
** SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
** LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
** OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE."
**
** $QT_END_LICENSE$
**
****************************************************************************/

import QtQuick 2.0
import QtQuick.Window 2.0

Item {
    id: display
    property real fontSize: Math.floor(Screen.pixelDensity * 5.0)
    property bool enteringDigits: false
    property int maxDigits: (width / fontSize) + 1
    property string displayedOperand
    property string errorString: qsTr("ERROR")
    property bool isError: displayedOperand === errorString

    function displayOperator(operator)
    {
        listView.model.append({ "operator": operator, "operand": "" })
        enteringDigits = true
        listView.positionViewAtEnd()
    }

    function newLine(operator, operand)
    {
        displayedOperand = displayNumber(operand)
        listView.model.append({ "operator": operator, "operand": displayedOperand })
        enteringDigits = false
        listView.positionViewAtEnd()
    }

    function appendDigit(digit)
    {
        if (!enteringDigits)
            listView.model.append({ "operator": "", "operand": "" })
        var i = listView.model.count - 1;
        listView.model.get(i).operand = listView.model.get(i).operand + digit;
        enteringDigits = true
        listView.positionViewAtEnd()
    }

    function setDigit(digit)
    {
        var i = listView.model.count - 1;
        listView.model.get(i).operand = digit;
        listView.positionViewAtEnd()
    }

    function clear()
    {
        displayedOperand = ""
        if (enteringDigits) {
            var i = listView.model.count - 1
            if (i >= 0)
                listView.model.remove(i)
            enteringDigits = false
        }
    }

    // Returns a string representation of a number that fits in
    // display.maxDigits characters, trying to keep as much precision
    // as possible. If the number cannot be displayed, returns an
    // error string.
    function displayNumber(num) {
        if (typeof(num) != "number")
            return errorString;

        var intNum = parseInt(num);
        var intLen = intNum.toString().length;

        // Do not count the minus sign as a digit
        var maxLen = num < 0 ? maxDigits + 1 : maxDigits;

        if (num.toString().length <= maxLen) {
            if (isFinite(num))
                return num.toString();
            return errorString;
        }

        // Integer part of the number is too long - try
        // an exponential notation
        if (intNum == num || intLen > maxLen - 3) {
            var expVal = num.toExponential(maxDigits - 6).toString();
            if (expVal.length <= maxLen)
                return expVal;
        }

        // Try a float presentation with fixed number of digits
        var floatStr = parseFloat(num).toFixed(maxDigits - intLen - 1).toString();
        if (floatStr.length <= maxLen)
            return floatStr;

        return errorString;
    }

    Item {
        id: theItem
        width: parent.width + 32
        height: parent.height

        Rectangle {
            id: rect
            x: 16
            color: "white"
            height: parent.height
            width: display.width - 16
        }
        Image {
            anchors.right: rect.left
            source: "images/paper-edge-left.png"
            height: parent.height
            fillMode: Image.TileVertically
        }
        Image {
            anchors.left: rect.right
            source: "images/paper-edge-right.png"
            height: parent.height
            fillMode: Image.TileVertically
        }

        Image {
            id: grip
            source: "images/paper-grip.png"
            anchors.horizontalCenter: parent.horizontalCenter
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 20
        }

        ListView {
            id: listView
            x: 16; y: 30
            width: display.width
            height: display.height - 50 - y
            delegate: Item {
                height: display.fontSize * 1.1
                width: parent.width
                Text {
                    id: operator
                    x: 6
                    font.pixelSize: display.fontSize
                    color: "#6da43d"
                    text: model.operator
                }
                Text {
                    id: operand
                    font.pixelSize: display.fontSize
                    anchors.right: parent.right
                    anchors.rightMargin: 22
                    text: model.operand
                }
            }
            model: ListModel { }
        }

    }

}
