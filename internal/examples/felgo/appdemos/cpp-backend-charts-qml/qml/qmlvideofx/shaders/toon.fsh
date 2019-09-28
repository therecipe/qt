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

// Based on http://www.geeks3d.com/20101029/shader-library-pixelation-post-processing-effect-glsl/

uniform float dividerValue;
uniform float threshold;
uniform float resS;
uniform float resT;
uniform float magTol;
uniform float quantize;

uniform sampler2D source;
uniform lowp float qt_Opacity;
varying vec2 qt_TexCoord0;

void main()
{
    vec4 color = vec4(1.0, 0.0, 0.0, 1.1);
    vec2 uv = qt_TexCoord0.xy;
    if (uv.x < dividerValue) {
        vec2 st = qt_TexCoord0.st;
        vec3 rgb = texture2D(source, st).rgb;
        vec2 stp0 = vec2(1.0/resS,  0.0);
        vec2 st0p = vec2(0.0     ,  1.0/resT);
        vec2 stpp = vec2(1.0/resS,  1.0/resT);
        vec2 stpm = vec2(1.0/resS, -1.0/resT);
        float i00 =   dot( texture2D(source, st).rgb, vec3(0.2125,0.7154,0.0721));
        float im1m1 = dot( texture2D(source, st-stpp).rgb, vec3(0.2125,0.7154,0.0721));
        float ip1p1 = dot( texture2D(source, st+stpp).rgb, vec3(0.2125,0.7154,0.0721));
        float im1p1 = dot( texture2D(source, st-stpm).rgb, vec3(0.2125,0.7154,0.0721));
        float ip1m1 = dot( texture2D(source, st+stpm).rgb, vec3(0.2125,0.7154,0.0721));
        float im10 =  dot( texture2D(source, st-stp0).rgb, vec3(0.2125,0.7154,0.0721));
        float ip10 =  dot( texture2D(source, st+stp0).rgb, vec3(0.2125,0.7154,0.0721));
        float i0m1 =  dot( texture2D(source, st-st0p).rgb, vec3(0.2125,0.7154,0.0721));
        float i0p1 =  dot( texture2D(source, st+st0p).rgb, vec3(0.2125,0.7154,0.0721));
        float h = -1.*im1p1 - 2.*i0p1 - 1.*ip1p1  +  1.*im1m1 + 2.*i0m1 + 1.*ip1m1;
        float v = -1.*im1m1 - 2.*im10 - 1.*im1p1  +  1.*ip1m1 + 2.*ip10 + 1.*ip1p1;
        float mag = sqrt(h*h + v*v);
        if (mag > magTol) {
            color = vec4(0.0, 0.0, 0.0, 1.0);
        }
        else {
            rgb.rgb *= quantize;
            rgb.rgb += vec3(0.5, 0.5, 0.5);
            ivec3 irgb = ivec3(rgb.rgb);
            rgb.rgb = vec3(irgb) / quantize;
            color = vec4(rgb, 1.0);
        }
    } else {
        color = texture2D(source, uv);
    }
    gl_FragColor = qt_Opacity * color;
}
