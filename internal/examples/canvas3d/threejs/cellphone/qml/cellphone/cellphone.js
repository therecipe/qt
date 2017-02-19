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

Qt.include("three.js")

var camera, scene, renderer, cubeCamera;
var caseMesh, frontMesh, iconMesh;
var caseColor = new THREE.Color("#000000");
var meshesReady = false;
var canvasTextureProvider = null;
var zeroVector = new THREE.Vector3(0, 0, 0);
var cameraLight;
var sphereImage, iconImage;
var caseMeshFile, frontMeshFile, iconMeshFile;

function initializeGL(canvas, textureSource) {
    scene = new THREE.Scene();

    cubeCamera = new THREE.CubeCamera(5, 100, 512);
    cubeCamera.renderTarget.texture.minFilter = THREE.LinearMipMapLinearFilter;
    scene.add(cubeCamera);

    // Background sphere
    var sphere = new THREE.Mesh(
                new THREE.SphereGeometry(40, 16, 16),
                new THREE.MeshBasicMaterial());
    sphere.side = THREE.BackSide;
    sphere.scale.x = -1;
    scene.add(sphere);
    var textureLoader = new THREE.TextureLoader();
    if (sphereImage) {
        var sphereTexture = textureLoader.load(sphereImage, updateEnvMap);
        sphereTexture.mapping = THREE.UVMapping;
        sphereTexture.minFilter = THREE.LinearFilter;
        sphere.material.map = sphereTexture;
    } else {
        sphere.material.color = new THREE.Color("black");
    }

    camera = new THREE.PerspectiveCamera( 75, canvas.width / canvas.height, 0.001, 1000 );

    var light = new THREE.AmbientLight( 0x666666 );
    scene.add( light );

    cameraLight = new THREE.DirectionalLight( 0xffffff, 1 );
    cameraLight.position.y = 1.5;
    scene.add( cameraLight );

    var caseMaterial = new THREE.MeshPhongMaterial({ envMap: cubeCamera.renderTarget,
                                                   combine: THREE.MultiplyOperation,
                                                   reflectivity: 0.35,
                                                   specular: "#555555" });
    //! [0]
    var frontTexture = new THREE.QtQuickItemTexture( textureSource );
    //! [0]
    //! [1]
    var frontMaterial = new THREE.MeshPhongMaterial( { map: frontTexture } );
    //! [1]

    var iconMaterial = new THREE.MeshPhongMaterial( { transparent:true } );
    if (iconImage) {
        var iconTexture = textureLoader.load(iconImage);
        iconTexture.minFilter = THREE.LinearFilter;
        iconMaterial.map = iconTexture;
    } else {
        iconMaterial.opacity = 0;
    }

    renderer = new THREE.Canvas3DRenderer(
                { canvas: canvas, antialias: true, devicePixelRatio: canvas.devicePixelRatio });
    renderer.setSize( canvas.width, canvas.height );

    // The cellphone meshes were created using a third party tool (Blender).
    // They were exported from Blender as Wavefront OBJ files and then converted into JSON format
    // using the OBJ -> JSON conversion script provided by three.js (convert_obj_three.py).
    var loader = new THREE.JSONLoader();
    loader.load( caseMeshFile, function ( geometry, materials ) {
        geometry.computeVertexNormals();
        var bufferGeometry = new THREE.BufferGeometry();
        bufferGeometry.fromGeometry(geometry);
        caseMesh = new THREE.Mesh( bufferGeometry, caseMaterial );
        if (iconMesh && frontMesh)
            meshesReady = true;
        caseMesh.material.color = caseColor;
        updateEnvMap();
        scene.add( caseMesh );
    } );
    loader.load( frontMeshFile, function ( geometry, materials ) {
        var bufferGeometry = new THREE.BufferGeometry();
        bufferGeometry.fromGeometry(geometry);
        frontMesh = new THREE.Mesh( bufferGeometry, frontMaterial );
        if (iconMesh && caseMesh)
            meshesReady = true;
        scene.add( frontMesh );
    } );
    loader.load( iconMeshFile, function ( geometry, materials ) {
        var bufferGeometry = new THREE.BufferGeometry();
        bufferGeometry.fromGeometry(geometry);
        iconMesh = new THREE.Mesh( bufferGeometry, iconMaterial );
        if (caseMesh && frontMesh)
            meshesReady = true;
        scene.add( iconMesh );
    } );
}

function updateEnvMap() {
    if (caseMesh) {
        // Take a snapshot of the scene using cube camera to create environment map for case
        caseMesh.material.envMap = cubeCamera.renderTarget;
        cubeCamera.updateCubeMap(renderer, scene);
        caseMesh.material.needsUpdate = true;
    }
}

function setSphereTexture(image) {
    sphereImage = image;
}

function setIconTexture(image) {
    iconImage = image;
}

function setMeshFiles(caseFile, frontFile, iconFile) {
    caseMeshFile = caseFile;
    frontMeshFile = frontFile;
    iconMeshFile = iconFile;
}

function setCaseColor(color) {
    caseColor.set(color);
    if (caseMesh)
        caseMesh.material.color = caseColor;
}

function resizeGL(canvas) {
    camera.aspect = canvas.width / canvas.height;
    camera.updateProjectionMatrix();

    renderer.setPixelRatio(canvas.devicePixelRatio);
    renderer.setSize( canvas.width, canvas.height );
}

function degToRad(degrees) {
    return degrees * Math.PI / 180;
}

function paintGL(canvas) {
    if (meshesReady) {
        var cameraRad = degToRad(canvas.cameraAngle);
        var lightRad = cameraRad - 0.8;
        caseMesh.rotation.x = degToRad(canvas.xRotAnim);
        caseMesh.rotation.y = degToRad(canvas.yRotAnim);
        caseMesh.rotation.z = cameraRad + degToRad(canvas.zRotAnim);
        frontMesh.rotation.set(caseMesh.rotation.x, caseMesh.rotation.y, caseMesh.rotation.z);
        iconMesh.rotation.set(caseMesh.rotation.x, caseMesh.rotation.y, caseMesh.rotation.z);
        camera.position.x = canvas.distance * Math.sin(cameraRad);
        camera.position.z = canvas.distance * Math.cos(cameraRad);
        cameraLight.position.x = (canvas.distance + 2) * Math.sin(lightRad);
        cameraLight.position.z = (canvas.distance + 2) * Math.cos(lightRad);
        camera.lookAt(zeroVector);
    }

    renderer.render( scene, camera );

}
