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

// Based on "Graphics Shaders: Theory and Practice" (http://cgeducation.org/ShadersBook/)

uniform float dividerValue;
uniform float mixLevel;
uniform float resS;
uniform float resT;

uniform sampler2D source;
uniform lowp float qt_Opacity;
varying vec2 qt_TexCoord0;

void main()
{
    vec2 uv = qt_TexCoord0.xy;
    vec4 c = vec4(0.0);
    if (uv.x < dividerValue) {
        vec2 st = qt_TexCoord0.st;
        vec3 irgb = texture2D(source, st).rgb;
        vec2 stp0 = vec2(1.0 / resS, 0.0);
        vec2 st0p = vec2(0.0       , 1.0 / resT);
        vec2 stpp = vec2(1.0 / resS, 1.0 / resT);
        vec2 stpm = vec2(1.0 / resS, -1.0 / resT);
        const vec3 W = vec3(0.2125, 0.7154, 0.0721);
        float i00   = dot(texture2D(source, st).rgb, W);
        float im1m1 = dot(texture2D(source, st-stpp).rgb, W);
        float ip1p1 = dot(texture2D(source, st+stpp).rgb, W);
        float im1p1 = dot(texture2D(source, st-stpm).rgb, W);
        float ip1m1 = dot(texture2D(source, st+stpm).rgb, W);
        float im10  = dot(texture2D(source, st-stp0).rgb, W);
        float ip10  = dot(texture2D(source, st+stp0).rgb, W);
        float i0m1  = dot(texture2D(source, st-st0p).rgb, W);
        float i0p1  = dot(texture2D(source, st+st0p).rgb, W);
        float h = -1.0*im1p1 - 2.0*i0p1 - 1.0*ip1p1 + 1.0*im1m1 + 2.0*i0m1 + 1.0*ip1m1;
        float v = -1.0*im1m1 - 2.0*im10 - 1.0*im1p1 + 1.0*ip1m1 + 2.0*ip10 + 1.0*ip1p1;
        float mag = 1.0 - length(vec2(h, v));
        vec3 target = vec3(mag, mag, mag);
        c = vec4(target, 1.0);
    } else {
        c = texture2D(source, qt_TexCoord0);
    }
    gl_FragColor = qt_Opacity * c;
}
