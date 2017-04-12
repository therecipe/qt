/****************************************************************************
**
** Copyright (C) 2016 The Qt Company Ltd.
** Contact: https://www.qt.io/licensing/
**
** This file is part of the QtCanvas3D module of the Qt Toolkit.
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

Qt.include("gl-matrix.js")

//! [0]
Qt.include("ThreeJSLoader.js")
//! [0]

var gl;

var texturedShaderProgram = 0;
var vertexShader = 0;
var fragmentShader = 0;

var vertexPositionAttribute;
var textureCoordAttribute;
var vertexNormalAttribute;

var pMatrixUniform;
var mvMatrixUniform;
var nMatrixUniform;
var textureSamplerUniform;
var eyeUniform;

var modelOneTexture = 0;
var modelTwoTexture = 0;
var modelThreeTexture = 0;
var modelFourTexture = 0;
var modelFiveTexture = 0;

var vMatrix  = mat4.create();
var mMatrix  = mat4.create();
var mvMatrix = mat4.create();
var pMatrix  = mat4.create();
var nMatrix  = mat4.create();

var fov = degToRad(45);
var eye = [0, 1, 1];
var light = [0, 1, 1];

var posOne = [0, 0, 0];
var posTwo = [0.3, 0, 0];
var posThree = [-0.1, 0, 0.25];
var posFour = [0.1, 0, -0.45];
var posFive = [0, -0.14, 0];
var posSix = [-1.2, -0.28, 0.0];
var posSeven = [0.5, -0.28, 0.9];
var posEight = [0.5, -0.28, -0.9];
var posNine = [0.55, 0.09, -1.0];
var posTen = [1.0, 0.09, -0.7];

var rotOne = degToRad(90);
var rotTwo = degToRad(-80);
var rotThree = degToRad(15);
var rotFour = degToRad(40);
var rotFive = degToRad(60);

var drawMode = 0;
var canvas3d;
var isLogEnabled = false;

function log(message) {
    if (isLogEnabled)
        console.log(message)
}

function Model() {
    this.verticesVBO = 0;
    this.normalsVBO  = 0;
    this.texCoordVBO = 0;
    this.indexVBO    = 0;
    this.count       = 0;
}

var modelOne = new Model();
var modelTwo = new Model();
var modelThree = new Model();
var modelFour = new Model();
var modelFive = new Model();
var stateDumpExt;

function initializeGL(canvas) {
    canvas3d = canvas
    log("initializeGL...")
    try {
        gl = canvas.getContext("canvas3d", {depth:true, antialias:true, alpha:false});
        log("   Received context: "+gl);

        stateDumpExt = gl.getExtension("QTCANVAS3D_gl_state_dump");
        if (stateDumpExt)
            log("QTCANVAS3D_gl_state_dump extension found");
        else
            log("QTCANVAS3D_gl_state_dump extension NOT found");

        var contextConfig = gl.getContextAttributes();
        log("   Depth: "+contextConfig.alpha);
        log("   Stencil: "+contextConfig.stencil);
        log("   Antialiasing: "+contextConfig.antialias);
        log("   Premultiplied alpha: "+contextConfig.premultipliedAlpha);
        log("   Preserve drawingbuffer: "+contextConfig.preserveDrawingBuffer);
        log("   Prefer Low Power To High Performance: "+contextConfig.preferLowPowerToHighPerformance);
        log("   Fail If Major Performance Caveat: "+contextConfig.failIfMajorPerformanceCaveat);

        // Setup the OpenGL state
        gl.enable(gl.DEPTH_TEST);
        gl.disable(gl.CULL_FACE);
        gl.enable(gl.BLEND);
        gl.enable(gl.DEPTH_TEST);
        gl.depthMask(true);
        gl.blendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);

        gl.clearColor(0.9, 0.9, 0.9, 1.0);
        gl.clearDepth(1.0);

        // Set viewport
        gl.viewport(0, 0,
                    canvas.width * canvas.devicePixelRatio,
                    canvas.height * canvas.devicePixelRatio);

        // Initialize the shader program
        initShaders();

        // Initialize buffers
        initBuffers();

        // Load textures
        gl.pixelStorei(gl.UNPACK_FLIP_Y_WEBGL, true);
        loadTextures();

        // Load JSON models
        loadJSONModels();

        log("...initializeGL");
    } catch(e) {
        console.log("...initializeGL FAILURE!");
        console.log(""+e);
        console.log(""+e.message);
    }
}

function resizeGL(canvas)
{
    var pixelRatio = canvas.devicePixelRatio;
    canvas.pixelSize = Qt.size(canvas.width * pixelRatio,
                               canvas.height * pixelRatio);
    if (gl)
        gl.viewport(0, 0,
                    canvas.width * canvas.devicePixelRatio,
                    canvas.height * canvas.devicePixelRatio);
}

function paintGL(canvas) {
    // draw
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);

    gl.useProgram(texturedShaderProgram);

    // Calculate the perspective projection
    mat4.perspective(pMatrix, fov, canvas.width / canvas.height, 0.1, 100.0);
    gl.uniformMatrix4fv(pMatrixUniform, false, pMatrix);

    //! [7]
    // Get the view matrix
    mat4.identity(vMatrix);
    eye = moveEye(canvas.xRot, canvas.yRot, canvas.distance);
    mat4.lookAt(vMatrix, eye, [0, 0, 0], [0, 1, 0]);
    //! [7]

    // Apply light position
    if (canvas3d.animatingLight === true)
        light = moveEye(canvas.lightX, canvas.lightY, canvas.lightDistance);
    else
        light = eye;
    gl.uniform3fv(eyeUniform, light);

    if (canvas3d.drawWireframe)
        drawMode = gl.LINES;
    else
        drawMode = gl.TRIANGLES;

    if (modelOne.count > 0 && modelOneTexture !== 0 ) {
        // Draw model one
        log("   model one count:"+modelOne.count+" texture:"+modelOneTexture.name);

        // Bind the correct buffers
        gl.bindBuffer(gl.ARRAY_BUFFER, modelOne.verticesVBO);
        gl.enableVertexAttribArray(vertexPositionAttribute);
        gl.vertexAttribPointer(vertexPositionAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelOne.normalsVBO);
        gl.enableVertexAttribArray(vertexNormalAttribute);
        gl.vertexAttribPointer(vertexNormalAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelOne.texCoordVBO);
        gl.enableVertexAttribArray(textureCoordAttribute);
        gl.vertexAttribPointer(textureCoordAttribute, 2, gl.FLOAT, false, 0, 0);

        gl.activeTexture(gl.TEXTURE0);
        gl.bindTexture(gl.TEXTURE_2D, modelOneTexture);
        gl.uniform1i(textureSamplerUniform, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posOne);

        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);

        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, modelOne.indexVBO);

        // Getting state dump is a synchronous operation, so only do it when logging is enabled
        if (isLogEnabled && stateDumpExt)
            log("GL STATE DUMP:\n"+stateDumpExt.getGLStateDump(stateDumpExt.DUMP_FULL));

        gl.drawElements(drawMode, modelOne.count, gl.UNSIGNED_SHORT, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posTwo);
        mat4.rotateY(mMatrix, mMatrix, rotTwo);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.drawElements(drawMode, modelOne.count, gl.UNSIGNED_SHORT, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posThree);
        mat4.rotateY(mMatrix, mMatrix, rotThree);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.drawElements(drawMode, modelOne.count, gl.UNSIGNED_SHORT, 0);
    }


    if (modelTwo.count > 0 && modelTwoTexture !== 0 ) {
        // Draw model two
        log("   model two count:"+modelTwo.count+" texture:"+modelTwoTexture.name);

        // Bind the correct buffers
        gl.bindBuffer(gl.ARRAY_BUFFER, modelTwo.verticesVBO);
        gl.enableVertexAttribArray(vertexPositionAttribute);
        gl.vertexAttribPointer(vertexPositionAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelTwo.normalsVBO);
        gl.enableVertexAttribArray(vertexNormalAttribute);
        gl.vertexAttribPointer(vertexNormalAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelTwo.texCoordVBO);
        gl.enableVertexAttribArray(textureCoordAttribute);
        gl.vertexAttribPointer(textureCoordAttribute, 2, gl.FLOAT, false, 0, 0);

        gl.activeTexture(gl.TEXTURE0);
        gl.bindTexture(gl.TEXTURE_2D, modelTwoTexture);
        gl.uniform1i(textureSamplerUniform, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posOne);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, modelTwo.indexVBO);
        gl.drawElements(drawMode, modelTwo.count, gl.UNSIGNED_SHORT, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posTwo);
        mat4.rotateY(mMatrix, mMatrix, rotTwo);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.drawElements(drawMode, modelTwo.count, gl.UNSIGNED_SHORT, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posThree);
        mat4.rotateY(mMatrix, mMatrix, rotThree);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.drawElements(drawMode, modelTwo.count, gl.UNSIGNED_SHORT, 0);
    }

    if (modelFour.count > 0 && modelFourTexture !== 0 ) {
        // Draw model four
        log("   model four count:"+modelFour.count+" texture:"+modelFourTexture.name);

        // Bind the correct buffers
        gl.bindBuffer(gl.ARRAY_BUFFER, modelFour.verticesVBO);
        gl.enableVertexAttribArray(vertexPositionAttribute);
        gl.vertexAttribPointer(vertexPositionAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelFour.normalsVBO);
        gl.enableVertexAttribArray(vertexNormalAttribute);
        gl.vertexAttribPointer(vertexNormalAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelFour.texCoordVBO);
        gl.enableVertexAttribArray(textureCoordAttribute);
        gl.vertexAttribPointer(textureCoordAttribute, 2, gl.FLOAT, false, 0, 0);

        gl.activeTexture(gl.TEXTURE0);
        gl.bindTexture(gl.TEXTURE_2D, modelFourTexture);
        gl.uniform1i(textureSamplerUniform, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posFive);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, modelFour.indexVBO);
        gl.drawElements(drawMode, modelFour.count, gl.UNSIGNED_SHORT, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posSix);
        mat4.rotateY(mMatrix, mMatrix, rotFour);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.drawElements(drawMode, modelFour.count, gl.UNSIGNED_SHORT, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posSeven);
        mat4.rotateY(mMatrix, mMatrix, rotOne);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.drawElements(drawMode, modelFour.count, gl.UNSIGNED_SHORT, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posEight);
        mat4.rotateY(mMatrix, mMatrix, rotFive);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.drawElements(drawMode, modelFour.count, gl.UNSIGNED_SHORT, 0);
    }

    if (modelFive.count > 0 && modelFiveTexture !== 0 ) {
        // Draw model five
        log("   model five count:"+modelFive.count+" texture:"+modelFiveTexture.name);

        // Bind the correct buffers
        gl.bindBuffer(gl.ARRAY_BUFFER, modelFive.verticesVBO);
        gl.enableVertexAttribArray(vertexPositionAttribute);
        gl.vertexAttribPointer(vertexPositionAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelFive.normalsVBO);
        gl.enableVertexAttribArray(vertexNormalAttribute);
        gl.vertexAttribPointer(vertexNormalAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelFive.texCoordVBO);
        gl.enableVertexAttribArray(textureCoordAttribute);
        gl.vertexAttribPointer(textureCoordAttribute, 2, gl.FLOAT, false, 0, 0);

        gl.activeTexture(gl.TEXTURE0);
        gl.bindTexture(gl.TEXTURE_2D, modelFiveTexture);
        gl.uniform1i(textureSamplerUniform, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posNine);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, modelFive.indexVBO);
        gl.drawElements(drawMode, modelFive.count, gl.UNSIGNED_SHORT, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posTen);
        mat4.rotateX(mMatrix, mMatrix, rotFour);
        mat4.rotateY(mMatrix, mMatrix, rotFive);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.drawElements(drawMode, modelFive.count, gl.UNSIGNED_SHORT, 0);
    }


    if (modelThree.count > 0 && modelThreeTexture !== 0 ) {
        // Draw model three (Includes transparency, must be drawn last)
        log("   model three count:"+modelThree.count+" texture:"+modelThreeTexture.name);

        // Bind the correct buffers
        gl.bindBuffer(gl.ARRAY_BUFFER, modelThree.verticesVBO);
        gl.enableVertexAttribArray(vertexPositionAttribute);
        gl.vertexAttribPointer(vertexPositionAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelThree.normalsVBO);
        gl.enableVertexAttribArray(vertexNormalAttribute);
        gl.vertexAttribPointer(vertexNormalAttribute, 3, gl.FLOAT, false, 0, 0);

        gl.bindBuffer(gl.ARRAY_BUFFER, modelThree.texCoordVBO);
        gl.enableVertexAttribArray(textureCoordAttribute);
        gl.vertexAttribPointer(textureCoordAttribute, 2, gl.FLOAT, false, 0, 0);

        gl.activeTexture(gl.TEXTURE0);
        gl.bindTexture(gl.TEXTURE_2D, modelThreeTexture);
        gl.uniform1i(textureSamplerUniform, 0);

        // Calculate the modelview matrix
        mat4.identity(mMatrix);
        mat4.translate(mMatrix, mMatrix, posFour);
        // Calculate normal matrix before scaling, to keep lighting in order
        // Scale normal matrix with distance instead
        mat4.copy(nMatrix, mMatrix);
        mat4.scale(nMatrix, nMatrix, [canvas.distance, canvas.distance, canvas.distance]);
        mat4.invert(nMatrix, nMatrix);
        mat4.transpose(nMatrix, nMatrix);
        gl.uniformMatrix4fv(nMatrixUniform, false, nMatrix);
        // Scale the modelview matrix, and apply the matrix
        mat4.scale(mMatrix, mMatrix, [canvas.itemSize, canvas.itemSize, canvas.itemSize]);
        mat4.multiply(mvMatrix, vMatrix, mMatrix);
        gl.uniformMatrix4fv(mvMatrixUniform, false, mvMatrix);

        // Draw the model
        gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, modelThree.indexVBO);
        gl.drawElements(drawMode, modelThree.count, gl.UNSIGNED_SHORT, 0);
    }
}

//! [8]
function moveEye(xRot, yRot, distance) {
    var xAngle = degToRad(xRot);
    var yAngle = degToRad(yRot);

    var zPos = distance * Math.cos(xAngle) * Math.cos(yAngle);
    var xPos = distance * Math.sin(xAngle) * Math.cos(yAngle);
    var yPos = distance * Math.sin(yAngle);

    return [-xPos, yPos, zPos];
}
//! [8]

//! [3]
function handleLoadedModel(jsonObj) {
    log("handleLoadedModel...");
    var modelData = parseJSON3DModel(jsonObj, "");

    if (modelOne.count === 0)
        fillModel(modelData, modelOne);
    else if (modelTwo.count === 0)
        fillModel(modelData, modelTwo);
    //! [3]
    else if (modelThree.count === 0)
        fillModel(modelData, modelThree);
    else if (modelFour.count === 0)
        fillModel(modelData, modelFour);
    else if (modelFive.count === 0)
        fillModel(modelData, modelFive);

    log("...handleLoadedModel");
}

//! [4]
function fillModel(modelData, model) {
    log("   fillModel...");
    log("   "+model.verticesVBO.name);
    gl.bindBuffer(gl.ARRAY_BUFFER, model.verticesVBO);
    gl.bufferData(gl.ARRAY_BUFFER,
                  new Float32Array(modelData.vertices),
                  gl.STATIC_DRAW);
    log("   "+model.normalsVBO.name);
    if (isLogEnabled && stateDumpExt)
        log("GL STATE DUMP:\n"+stateDumpExt.getGLStateDump(stateDumpExt.DUMP_VERTEX_ATTRIB_ARRAYS_BIT || stateDumpExt.DUMP_VERTEX_ATTRIB_ARRAYS_CONTENTS_BIT));

    gl.bindBuffer(gl.ARRAY_BUFFER, model.normalsVBO);
    gl.bufferData(gl.ARRAY_BUFFER,
                  new Float32Array(modelData.normals),
                  gl.STATIC_DRAW);

    log("   "+model.texCoordVBO.name);
    gl.bindBuffer(gl.ARRAY_BUFFER, model.texCoordVBO);
    gl.bufferData(gl.ARRAY_BUFFER,
                  new Float32Array(modelData.texCoords[0]),
                  gl.STATIC_DRAW);

    log("   "+model.indexVBO.name);
    gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, model.indexVBO);
    gl.bufferData(gl.ELEMENT_ARRAY_BUFFER,
                  new Uint16Array(modelData.indices),
                  gl.STATIC_DRAW);

    model.count = modelData.indices.length;
    log("   ...fillModel");
}
//! [4]

function degToRad(degrees) {
    return degrees * Math.PI / 180;
}

function initShaders()
{
    log("   initShaders...")

    vertexShader = getShader(gl,
                             "attribute highp vec3 aVertexNormal;   \
                              attribute highp vec3 aVertexPosition; \
                              attribute highp vec2 aTextureCoord;   \

                              uniform highp mat4 uNormalMatrix;     \
                              uniform mat4 uMVMatrix;               \
                              uniform mat4 uPMatrix;                \
                              uniform vec3 eyePos;                  \

                              varying highp vec2 vTextureCoord;     \
                              varying highp vec4 vLighting;         \

                              void main(void) {                     \
                                 gl_Position = uPMatrix * uMVMatrix * vec4(aVertexPosition, 1.0);                   \
                                 vTextureCoord = aTextureCoord;                                                     \
                                 highp vec4 ambientLight = vec4(0.5, 0.5, 0.5, 1.0);                                \
                                 highp vec4 directionalLightColor = vec4(1.0, 1.0, 1.0, 1.0);                       \
                                 highp vec3 directionalVector = eyePos;                                             \
                                 highp vec4 transformedNormal = uNormalMatrix * vec4(aVertexNormal, 1.0);           \

                                 highp float directional = max(dot(transformedNormal.xyz, directionalVector), 0.0); \
                                 vLighting = ambientLight + (directionalLightColor * directional);                  \
                             }", gl.VERTEX_SHADER);

    fragmentShader = getShader(gl,
                               "varying highp vec2 vTextureCoord;   \
                                varying highp vec4 vLighting;       \

                                uniform sampler2D uSampler;         \

                                void main(void) {                   \
                                    mediump vec4 texelColor = texture2D(uSampler, vTextureCoord);   \
                                    gl_FragColor = vec4(texelColor * vLighting);                    \
                                }", gl.FRAGMENT_SHADER);

    texturedShaderProgram = gl.createProgram();
    texturedShaderProgram.name = "texturedShaderProgram";
    gl.attachShader(texturedShaderProgram, vertexShader);
    gl.attachShader(texturedShaderProgram, fragmentShader);
    gl.linkProgram(texturedShaderProgram);

    if (!gl.getProgramParameter(texturedShaderProgram, gl.LINK_STATUS)) {
        console.log("Could not initialize shaders");
        console.log(gl.getProgramInfoLog(texturedShaderProgram));
    }

    gl.useProgram(texturedShaderProgram);

    // look up where the vertex data needs to go.
    vertexPositionAttribute = gl.getAttribLocation(texturedShaderProgram, "aVertexPosition");
    vertexPositionAttribute.name = "aVertexPosition";
    gl.enableVertexAttribArray(vertexPositionAttribute);
    vertexNormalAttribute = gl.getAttribLocation(texturedShaderProgram, "aVertexNormal");
    vertexNormalAttribute.name = "aVertexNormal";
    gl.enableVertexAttribArray(vertexNormalAttribute);
    textureCoordAttribute = gl.getAttribLocation(texturedShaderProgram, "aTextureCoord");
    textureCoordAttribute.name = "aTextureCoord";
    gl.enableVertexAttribArray(textureCoordAttribute);

    pMatrixUniform = gl.getUniformLocation(texturedShaderProgram, "uPMatrix");
    pMatrixUniform.name = "uPMatrix";
    mvMatrixUniform = gl.getUniformLocation(texturedShaderProgram, "uMVMatrix");
    mvMatrixUniform.name = "uMVMatrix";
    textureSamplerUniform = gl.getUniformLocation(texturedShaderProgram, "uSampler")
    textureSamplerUniform.name = "uSampler";
    nMatrixUniform = gl.getUniformLocation(texturedShaderProgram, "uNormalMatrix");
    nMatrixUniform.name = "uNormalMatrix";
    eyeUniform = gl.getUniformLocation(texturedShaderProgram, "eyePos");
    eyeUniform.name = "eyePos";
    log("   ...initShaders");
}

//! [1]
function initBuffers() {
    modelOne.verticesVBO = gl.createBuffer();
    modelOne.verticesVBO.name = "modelOne.verticesVBO";
    modelOne.normalsVBO  = gl.createBuffer();
    modelOne.normalsVBO.name = "modelOne.normalsVBO";
    modelOne.texCoordVBO = gl.createBuffer();
    modelOne.texCoordVBO.name = "modelOne.texCoordVBO";
    modelOne.indexVBO    = gl.createBuffer();
    modelOne.indexVBO.name = "modelOne.indexVBO";

    modelTwo.verticesVBO = gl.createBuffer();
    modelTwo.verticesVBO.name = "modelTwo.verticesVBO";
    modelTwo.normalsVBO  = gl.createBuffer();
    modelTwo.normalsVBO.name = "modelTwo.normalsVBO";
    modelTwo.texCoordVBO = gl.createBuffer();
    modelTwo.texCoordVBO.name = "modelTwo.texCoordVBO";
    modelTwo.indexVBO    = gl.createBuffer();
    modelTwo.indexVBO.name = "modelTwo.indexVBO";

    //! [1]
    modelThree.verticesVBO = gl.createBuffer();
    modelThree.verticesVBO.name = "modelThree.verticesVBO";
    modelThree.normalsVBO  = gl.createBuffer();
    modelThree.normalsVBO.name = "modelThree.normalsVBO";
    modelThree.texCoordVBO = gl.createBuffer();
    modelThree.texCoordVBO.name = "modelThree.texCoordVBO";
    modelThree.indexVBO    = gl.createBuffer();
    modelThree.indexVBO.name = "modelThree.indexVBO";

    modelFour.verticesVBO = gl.createBuffer();
    modelFour.verticesVBO.name = "modelFour.verticesVBO";
    modelFour.normalsVBO  = gl.createBuffer();
    modelFour.normalsVBO.name = "modelFour.normalsVBO";
    modelFour.texCoordVBO = gl.createBuffer();
    modelFour.texCoordVBO.name = "modelFour.texCoordVBO";
    modelFour.indexVBO    = gl.createBuffer();
    modelFour.indexVBO.name = "modelFour.indexVBO";

    modelFive.verticesVBO = gl.createBuffer();
    modelFive.verticesVBO.name = "modelFive.verticesVBO";
    modelFive.normalsVBO  = gl.createBuffer();
    modelFive.normalsVBO.name = "modelFive.normalsVBO";
    modelFive.texCoordVBO = gl.createBuffer();
    modelFive.texCoordVBO.name = "modelFive.texCoordVBO";
    modelFive.indexVBO    = gl.createBuffer();
    modelFive.indexVBO.name = "modelFive.indexVBO";
}

//! [5]
function loadTextures() {
    // Load the first texture
    var goldImage = TextureImageFactory.newTexImage();
    goldImage.name = "goldImage";
    goldImage.imageLoaded.connect(function() {
        log("    creating model one texture");
        modelOneTexture = gl.createTexture();
        modelOneTexture.name = "modelOneTexture";
        gl.bindTexture(gl.TEXTURE_2D, modelOneTexture);
        gl.texImage2D(gl.TEXTURE_2D,    // target
                      0,                // level
                      gl.RGBA,          // internalformat
                      gl.RGBA,          // format
                      gl.UNSIGNED_BYTE, // type
                      goldImage);       // pixels
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST);
        gl.generateMipmap(gl.TEXTURE_2D);
    });
    goldImage.imageLoadingFailed.connect(function() {
        console.log("Texture load FAILED, "+goldImage.errorString);
    });
    goldImage.src = "qrc:///gold.jpg";
    log("   texture one source set")

    // Load the second texture
    var woodBoxImage = TextureImageFactory.newTexImage();
    woodBoxImage.name = "woodBoxImage";
    woodBoxImage.imageLoaded.connect(function() {
        log("    creating model two texture");
        modelTwoTexture = gl.createTexture();
        modelTwoTexture.name = "modelTwoTexture";
        gl.bindTexture(gl.TEXTURE_2D, modelTwoTexture);
        gl.texImage2D(gl.TEXTURE_2D,    // target
                      0,                // level
                      gl.RGBA,          // internalformat
                      gl.RGBA,          // format
                      gl.UNSIGNED_BYTE, // type
                      woodBoxImage);    // pixels
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST);
        gl.generateMipmap(gl.TEXTURE_2D);
    });
    woodBoxImage.imageLoadingFailed.connect(function() {
        console.log("Texture load FAILED, "+woodBoxImage.errorString);
    });
    woodBoxImage.src = "qrc:///woodbox.jpg";
    log("   texture two source set")
    //! [5]

    // Load the third texture
    var bushImage = TextureImageFactory.newTexImage();
    bushImage.name = "bushImage";
    bushImage.imageLoaded.connect(function() {
        log("    creating model three texture");
        modelThreeTexture = gl.createTexture();
        modelThreeTexture.name = "modelThreeTexture";
        gl.bindTexture(gl.TEXTURE_2D, modelThreeTexture);
        gl.texImage2D(gl.TEXTURE_2D,    // target
                      0,                // level
                      gl.RGBA,          // internalformat
                      gl.RGBA,          // format
                      gl.UNSIGNED_BYTE, // type
                      bushImage);    // pixels
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST);
        gl.generateMipmap(gl.TEXTURE_2D);
    });
    bushImage.imageLoadingFailed.connect(function() {
        console.log("Texture load FAILED, "+bushImage.errorString);
    });
    bushImage.src = "qrc:///bush.png";
    log("   texture three source set")

    // Load the fourth texture
    var palletImage = TextureImageFactory.newTexImage();
    palletImage.name = "palletImage";
    palletImage.imageLoaded.connect(function() {
        log("    creating model four texture");
        modelFourTexture = gl.createTexture();
        modelFourTexture.name = "modelFourTexture";
        gl.bindTexture(gl.TEXTURE_2D, modelFourTexture);
        gl.texImage2D(gl.TEXTURE_2D,    // target
                      0,                // level
                      gl.RGBA,          // internalformat
                      gl.RGBA,          // format
                      gl.UNSIGNED_BYTE, // type
                      palletImage);     // pixels
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST);
        gl.generateMipmap(gl.TEXTURE_2D);
    });
    palletImage.imageLoadingFailed.connect(function() {
        console.log("Texture load FAILED, "+palletImage.errorString);
    });
    palletImage.src = "qrc:///pallet.jpg";
    log("   texture four source set")

    // Load the fifth texture
    var rockImage = TextureImageFactory.newTexImage();
    rockImage.name = "rockImage";
    rockImage.imageLoaded.connect(function() {
        log("    creating model five texture");
        modelFiveTexture = gl.createTexture();
        modelFiveTexture.name = "modelFiveTexture";
        gl.bindTexture(gl.TEXTURE_2D, modelFiveTexture);
        gl.texImage2D(gl.TEXTURE_2D,    // target
                      0,                // level
                      gl.RGBA,          // internalformat
                      gl.RGBA,          // format
                      gl.UNSIGNED_BYTE, // type
                      rockImage);       // pixels
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
        gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST);
        gl.generateMipmap(gl.TEXTURE_2D);
    });
    rockImage.imageLoadingFailed.connect(function() {
        console.log("Texture load FAILED, "+rockImage.errorString);
    });
    rockImage.src = "qrc:///rock.jpg";
    log("   texture five source set")
}

//! [2]
function loadJSONModels() {
    // Load the first model
    var request = new XMLHttpRequest();
    request.open("GET", "gold.json");
    request.onreadystatechange = function () {
        if (request.readyState === XMLHttpRequest.DONE) {
            handleLoadedModel(JSON.parse(request.responseText));
        }
    }
    request.send();
    log("   XMLHttpRequest sent for model one")

    // Load the second model
    var request2 = new XMLHttpRequest();
    request2.open("GET", "woodbox.json");
    request2.onreadystatechange = function () {
        if (request2.readyState === XMLHttpRequest.DONE) {
            handleLoadedModel(JSON.parse(request2.responseText));
        }
    }
    request2.send();
    log("   XMLHttpRequest sent for model two")
    //! [2]

    // Load the third model
    var request3 = new XMLHttpRequest();
    request3.open("GET", "bush.json");
    request3.onreadystatechange = function () {
        if (request3.readyState === XMLHttpRequest.DONE) {
            handleLoadedModel(JSON.parse(request3.responseText));
        }
    }
    request3.send();
    log("   XMLHttpRequest sent for model three")

    // Load the fourth model
    var request4 = new XMLHttpRequest();
    request4.open("GET", "pallet.json");
    request4.onreadystatechange = function () {
        if (request4.readyState === XMLHttpRequest.DONE) {
            handleLoadedModel(JSON.parse(request4.responseText));
        }
    }
    request4.send();
    log("   XMLHttpRequest sent for model four")

    // Load the fifth model
    var request5 = new XMLHttpRequest();
    request5.open("GET", "rock.json");
    request5.onreadystatechange = function () {
        if (request5.readyState === XMLHttpRequest.DONE) {
            handleLoadedModel(JSON.parse(request5.responseText));
        }
    }
    request5.send();
    log("   XMLHttpRequest sent for model five")
}

function getShader(gl, str, type) {
    var shader = gl.createShader(type);
    gl.shaderSource(shader, str);
    gl.compileShader(shader);

    if (!gl.getShaderParameter(shader, gl.COMPILE_STATUS)) {
        console.log("JS:Shader compile failed");
        console.log(gl.getShaderInfoLog(shader));
        return null;
    }

    return shader;
}
