/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd and/or its subsidiary(-ies).
** Contact: https://www.qt.io/licensing/
**
** This file is part of the Qt Mobility Components.
**
** $QT_BEGIN_LICENSE:BSD$
** Commercial License Usage
** Licensees holding valid commercial Qt licenses may use this file in
** accordance with the commercial license agreement provided with the
** Software or, alternatively, in accordance with the terms contained in
** a written agreement between you and The Qt Company. For licensing terms
** and conditions see https://www.qt.io/terms-conditions. For further
** information use the contact form at https://www.qt.io/contact-us.
**
** BSD License Usage
** Alternatively, you may use this file under the terms of the BSD license
** as follows:
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

// Based on http://blog.qt.io/blog/2011/03/22/the-convenient-power-of-qml-scene-graph/

uniform float dividerValue;
uniform float targetWidth;
uniform float targetHeight;
uniform float time;

uniform sampler2D source;
uniform lowp float qt_Opacity;
varying vec2 qt_TexCoord0;

const float PI = 3.1415926535;
const int ITER = 7;
const float RATE = 0.1;
uniform float amplitude;
uniform float n;
uniform float pixDens;

void main()
{
    vec2 uv = qt_TexCoord0.xy;
    vec2 tc = uv;
    vec2 p = vec2(-1.0 + 2.0 * (gl_FragCoord.x - (pixDens * 14.0)) / targetWidth, -(-1.0 + 2.0 * (gl_FragCoord.y - (pixDens * 29.0)) / targetHeight));
    float diffx = 0.0;
    float diffy = 0.0;
    vec4 col;
    if (uv.x < dividerValue) {
        for (int i=0; i<ITER; ++i) {
            float theta = float(i) * PI / float(ITER);
            vec2 r = vec2(cos(theta) * p.x + sin(theta) * p.y, -1.0 * sin(theta) * p.x + cos(theta) * p.y);
            float diff = (sin(2.0 * PI * n * (r.y + time * RATE)) + 1.0) / 2.0;
            diffx += diff * sin(theta);
            diffy += diff * cos(theta);
        }
        tc = 0.5*(vec2(1.0,1.0) + p) + amplitude * vec2(diffx, diffy);
    }
    gl_FragColor = qt_Opacity * texture2D(source, tc);
}
