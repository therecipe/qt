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

var curVal = 0
var memory = 0
var lastOp = ""
var previousOperator = ""
var digits = ""

function disabled(op) {
    if (digits == "" && !((op >= "0" && op <= "9") || op == "."))
        return true
    else if (op == '=' && previousOperator.length != 1)
        return true
    else if (op == "." && digits.toString().search(/\./) != -1) {
        return true
    } else if (op == "√" &&  digits.toString().search(/-/) != -1) {
        return true
    } else {
        return false
    }
}

function digitPressed(op)
{
    if (disabled(op))
        return
    if (digits.toString().length >= display.maxDigits)
        return
    if (lastOp.toString().length == 1 && ((lastOp >= "0" && lastOp <= "9") || lastOp == ".") ) {
        digits = digits + op.toString()
        display.appendDigit(op.toString())
    } else {
        digits = op
        display.appendDigit(op.toString())
    }
    lastOp = op
}

function operatorPressed(op)
{
    if (disabled(op))
        return
    lastOp = op

    if (op == "±") {
            digits = Number(digits.valueOf() * -1)
            display.setDigit(display.displayNumber(digits))
            return
        }

    if (previousOperator == "+") {
        digits = Number(digits.valueOf()) + Number(curVal.valueOf())
    } else if (previousOperator == "−") {
        digits = Number(curVal.valueOf()) - Number(digits.valueOf())
    } else if (previousOperator == "×") {
        digits = Number(curVal) * Number(digits.valueOf())
    } else if (previousOperator == "÷") {
        digits = Number(curVal) / Number(digits.valueOf())
    }

    if (op == "+" || op == "−" || op == "×" || op == "÷") {
        previousOperator = op
        curVal = digits.valueOf()
        digits = ""
        display.displayOperator(previousOperator)
        return
    }

    if (op == "=") {
        display.newLine("=", digits.valueOf())
    }

    curVal = 0
    previousOperator = ""

    if (op == "1/x") {
        digits = (1 / digits.valueOf()).toString()
    } else if (op == "x^2") {
        digits = (digits.valueOf() * digits.valueOf()).toString()
    } else if (op == "Abs") {
        digits = (Math.abs(digits.valueOf())).toString()
    } else if (op == "Int") {
        digits = (Math.floor(digits.valueOf())).toString()
    } else if (op == "√") {
        digits = Number(Math.sqrt(digits.valueOf()))
        display.newLine("√", digits.valueOf())
    } else if (op == "mc") {
        memory = 0;
    } else if (op == "m+") {
        memory += digits.valueOf()
    } else if (op == "mr") {
        digits = memory.toString()
    } else if (op == "m-") {
        memory = digits.valueOf()
    } else if (op == "backspace") {
        digits = digits.toString().slice(0, -1)
        display.clear()
        display.appendDigit(digits)
    } else if (op == "Off") {
        Qt.quit();
    }

    // Reset the state on 'C' operator or after
    // an error occurred
    if (op == "C" || display.isError) {
        display.clear()
        curVal = 0
        memory = 0
        lastOp = ""
        digits = ""
    }
}

