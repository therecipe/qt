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
Qt.include("threex.planets.js")

var SUN = 0;
var MERCURY = 1;
var VENUS = 2;
var EARTH = 3;
var MARS = 4;
var JUPITER = 5;
var SATURN = 6;
var URANUS = 7;
var NEPTUNE = 8;
var NUM_SELECTABLE_PLANETS = 9;
var MOON = 9;
var SOLAR_SYSTEM = 100;

var camera, scene, renderer;
var planetCanvas, mouse, raycaster;

var daysPerFrame;
var daysPerFrameScale;
var planetScale;
var cameraDistance;

var objects = []; // Planet objects
var hitObjects = []; // Planet hit detection objects
var planets = []; // Planet data info

var commonGeometry;
var hitGeometry;
var solarDistance = 2600000;
var saturnOuterRadius = 120.700;
var uranusOuterRadius = 40;

var qmlView;

var oldFocusedPlanetPosition;
var oldCameraPosition;
var defaultCameraPosition;

var y = 2000;
var m = 1;
var D = 1;
// Time scale formula based on http://www.stjarnhimlen.se/comp/ppcomp.html
var startD = 367 * y - 7 * (y + (m + 9) / 12) / 4 + 275 * m / 9 + D - 730530;
var oldTimeD = startD;
var currTimeD = startD;

var auScale = 149597.870700; // AU in thousands of kilometers

var focusedScaling = false;
var focusedMinimumScale = 20;
var actualScale;

function initializeGL(canvas, eventSource, mainView) {

    planetCanvas = canvas;
    qmlView = mainView;

    camera = new THREE.PerspectiveCamera(45, canvas.width / canvas.height, 2500000, 20000000);
    defaultCameraPosition = new THREE.Vector3(solarDistance, solarDistance, solarDistance);
    camera.position.set(defaultCameraPosition.x, defaultCameraPosition.y, defaultCameraPosition.z);

    scene = new THREE.Scene();

    var starSphere = THREEx.Planets.createStarfield(8500000);
    scene.add(starSphere);

    var light = new THREE.PointLight(0x777777, 2);
    light.position.set(0, 0, 0);
    scene.add(light);

    scene.add(new THREE.AmbientLight(0x111111));

    loadPlanetData();
    createPlanets();
    setScale(1200);

    camera.lookAt(objects[0].position); // look at the Sun

    raycaster = new THREE.Raycaster();
    mouse = new THREE.Vector2();

    renderer = new THREE.Canvas3DRenderer(
                { canvas: canvas, antialias: true, devicePixelRatio: canvas.devicePixelRatio });

    renderer.setPixelRatio(canvas.devicePixelRatio);
    renderer.setSize(canvas.width, canvas.height);
    //! [5]
    eventSource.mouseDown.connect(onDocumentMouseDown);
    //! [5]
}

function loadPlanetData() {

    // Planet Data
    // radius - planet radius in millions of meters
    // tilt - planet axis angle
    // N1 N2 - longitude of the ascending node
    // i1 i2 - inclination to the ecliptic (plane of the Earth's orbit)
    // w1 w2 - argument of perihelion
    // a1 a2 - semi-major axis, or mean distance from Sun
    // e1 e2 - eccentricity (0=circle, 0-1=ellipse, 1=parabola)
    // M1 M2 - mean anomaly (0 at perihelion; increases uniformly with time)
    // period - sidereal rotation period
    // centerOfOrbit - the planet in the center of the orbit
    // (orbital elements based on http://www.stjarnhimlen.se/comp/ppcomp.html)

    var sun = { radius: 694.439, tilt: 63.87, period: 25.05 };
    planets.push(sun);
    var mercury = {
        radius: 2.433722, tilt: 0.04, N1: 48.3313, N2: 0.0000324587,
        i1: 7.0047, i2: 0.0000000500, w1: 29.1241, w2: 0.0000101444,
        a1: 0.387098, a2: 0, e1: 0.205635, e2: 0.000000000559,
        M1: 168.6562, M2: 4.0923344368, period: 58.646,
        centerOfOrbit: SUN
    };
    planets.push(mercury);
    var venus = {
        radius: 6.046079, tilt: 177.36, N1: 76.6799, N2: 0.0000246590,
        i1: 3.3946, i2: 0.0000000275, w1: 54.8910, w2: 0.0000138374,
        a1: 0.723330, a2: 0, e1: 0.006773, e2: -0.000000001302,
        M1: 48.0052, M2: 1.6021302244, period: 243.0185,
        centerOfOrbit: SUN
    };
    planets.push(venus);
    var earth = {
        radius: 6.371, tilt: 25.44, N1: 174.873, N2: 0,
        i1: 0.00005, i2: 0, w1: 102.94719, w2: 0,
        a1: 1, a2: 0, e1: 0.01671022, e2: 0,
        M1: 357.529, M2: 0.985608, period: 0.997,
        centerOfOrbit: SUN
    };
    planets.push(earth);
    var mars = {
        radius: 3.389372, tilt: 25.19, N1: 49.5574, N2: 0.0000211081,
        i1: 1.8497, i2: -0.0000000178, w1: 286.5016, w2: 0.0000292961,
        a1: 1.523688, a2: 0, e1: 0.093405, e2: 0.000000002516,
        M1: 18.6021, M2: 0.5240207766, period: 1.025957,
        centerOfOrbit: SUN
    };
    planets.push(mars);
    var jupiter = {
        radius: 71.41254, tilt: 3.13, N1: 100.4542, N2: 0.0000276854,
        i1: 1.3030, i2: -0.0000001557, w1: 273.8777, w2: 0.0000164505,
        a1: 5.20256, a2: 0, e1: 0.048498, e2: 0.000000004469,
        M1: 19.8950, M2: 0.0830853001, period: 0.4135,
        centerOfOrbit: SUN
    };
    planets.push(jupiter);
    var saturn = {
        radius: 60.19958, tilt: 26.73, N1: 113.6634, N2: 0.0000238980,
        i1: 2.4886, i2: -0.0000001081, w1: 339.3939, w2: 0.0000297661,
        a1: 9.55475, a2: 0, e1: 0.055546, e2: -0.000000009499,
        M1: 316.9670, M2: 0.0334442282, period: 0.4395,
        centerOfOrbit: SUN
    };
    planets.push(saturn);
    var uranus = {
        radius: 25.5286, tilt: 97.77, N1: 74.0005, N2: 0.000013978,
        i1: 0.7733, i2: 0.000000019, w1: 96.6612, w2: 0.000030565,
        a1: 19.18171, a2: -0.0000000155, e1: 0.047318, e2: 0.00000000745,
        M1: 142.5905, M2: 0.011725806, period: 0.71833,
        centerOfOrbit: SUN
    };
    planets.push(uranus);
    var neptune = {
        radius: 24.73859, tilt: 28.32, N1: 131.7806, N2: 0.000030173,
        i1: 1.7700, i2: -0.000000255, w1: 272.8461, w2: 0.000006027,
        a1: 30.05826, a2: 0.00000003313, e1: 0.008606, e2: 0.00000000215,
        M1: 260.2471, M2: 0.005995147, period: 0.6713,
        centerOfOrbit: SUN
    };
    planets.push(neptune);
    var moon = {
        radius: 1.5424, tilt: 28.32, N1: 125.1228, N2: -0.0529538083,
        i1: 5.1454, i2: 0, w1: 318.0634, w2: 0.1643573223,
        a1: 0.273, a2: 0, e1: 0.054900, e2: 0,
        M1: 115.3654, M2: 13.0649929509, period: 27.321582,
        centerOfOrbit: EARTH
    };
    planets.push(moon);

}

function createPlanets() {

    objects = [];

    commonGeometry = new THREE.BufferGeometry().fromGeometry(new THREE.SphereGeometry(1, 64, 64));
    hitGeometry = new THREE.BufferGeometry().fromGeometry(new THREE.SphereGeometry(1, 8, 8));

    var ringSegments = 70;
    var mesh, innerRadius, outerRadius, ring;

    for (var i = 0; i < planets.length; i ++) {
        switch (i) {
        case SUN:
            mesh = createSun(planets[i]["radius"]);
            mesh.position.set(0, 0, 0);
            break;
        case MERCURY:
            mesh = createPlanet(planets[i]["radius"], 0.005, 'images/mercurymap.jpg',
                                'images/mercurybump.jpg');
            break;
        case VENUS:
            mesh = createPlanet(planets[i]["radius"], 0.005, 'images/venusmap.jpg',
                                'images/venusbump.jpg');
            break;
        case EARTH:
            mesh = createPlanet(planets[i]["radius"], 0.05, 'images/earthmap1k.jpg',
                                'images/earthbump1k.jpg', 'images/earthspec1k.jpg');
            createEarthCloud(mesh);
            break;
        case MARS:
            mesh = createPlanet(planets[i]["radius"], 0.05, 'images/marsmap1k.jpg',
                                'images/marsbump1k.jpg');
            break;
        case JUPITER:
            mesh = createPlanet(planets[i]["radius"], 0.02, 'images/jupitermap.jpg',
                                'images/jupitermap.jpg');
            break;
        case SATURN:
            mesh = createPlanet(planets[i]["radius"], 0.05, 'images/saturnmap.jpg',
                                'images/saturnmap.jpg');
            innerRadius = (planets[i]["radius"] + 6.630) / planets[i]["radius"];
            outerRadius = (planets[i]["radius"] + saturnOuterRadius) / planets[i]["radius"];
            ring = createRing(innerRadius, outerRadius, ringSegments,
                                  'qrc:images/saturnringcolortrans.png');
            ring.receiveShadow = true;
            ring.castShadow = true;
            mesh.add(ring);
            break;
        case URANUS:
            mesh = createPlanet(planets[i]["radius"], 0.05, 'images/uranusmap.jpg',
                                'images/uranusmap.jpg');
            innerRadius = (planets[i]["radius"] + 2) / planets[i]["radius"];
            outerRadius = (planets[i]["radius"] + uranusOuterRadius) / planets[i]["radius"];
            ring = createRing(innerRadius, outerRadius, ringSegments,
                                  'qrc:images/uranusringcolortrans.png');
            ring.receiveShadow = true;
            ring.castShadow = true;
            mesh.add(ring);
            break;
        case NEPTUNE:
            mesh = createPlanet(planets[i]["radius"], 0.05, 'images/neptunemap.jpg',
                                'images/neptunemap.jpg');
            break;
        case MOON:
            mesh = createPlanet(planets[i]["radius"], 0.05, 'images/moonmap1k.jpg',
                                'images/moonbump1k.jpg');
            break;
        }

        objects.push(mesh);
        scene.add(mesh);

        // Create separate meshes for click detection
        var hitMesh = new THREE.Mesh(hitGeometry);
        hitMesh.visible = false;
        hitObjects.push(hitMesh);
        scene.add(hitMesh);
    }

}

function createSun(radius) {

    var textureLoader = new THREE.TextureLoader();
    var texture = textureLoader.load('images/sunmap.jpg');
    var material = new THREE.MeshBasicMaterial({ map: texture });
    var mesh = new THREE.Mesh(commonGeometry, material);
    mesh.scale.set(radius, radius, radius);

    mesh.receiveShadow = false;
    mesh.castShadow = false;

    return mesh;
}

function createPlanet(radius, bumpMapScale, mapTexture, bumpTexture, specularTexture) {

    var textureLoader = new THREE.TextureLoader();
    var material = new THREE.MeshPhongMaterial({
                                                   map: textureLoader.load(mapTexture),
                                                   bumpMap: textureLoader.load(bumpTexture),
                                                   bumpScale: bumpMapScale
                                               });

    if (specularTexture) {
        material.specularMap = textureLoader.load(specularTexture);
        material.specular = new THREE.Color('grey');
        material.shininess = 50.0;
    } else {
        material.shininess = 1.0;
    }

    var mesh = new THREE.Mesh(commonGeometry, material);
    mesh.scale.set(radius, radius, radius);

    return mesh;

}

function createEarthCloud(earthMesh) {

    var textureLoader = new THREE.TextureLoader();
    var material = new THREE.MeshPhongMaterial({
                                                   map: textureLoader.load('qrc:images/earthcloudmapcolortrans.png'),
                                                   side: THREE.BackSide,
                                                   transparent: true,
                                                   opacity: 0.8
                                               });
    var mesh = new THREE.Mesh(commonGeometry, material);

    var material2 = new THREE.MeshPhongMaterial({
                                                   map: textureLoader.load('qrc:images/earthcloudmapcolortrans.png'),
                                                   side: THREE.FrontSide,
                                                   transparent: true,
                                                   opacity: 0.8
                                               });
    var mesh2 = new THREE.Mesh(commonGeometry, material2);

    mesh.scale.set(1.02, 1.02, 1.02);
    earthMesh.add(mesh);
    mesh2.scale.set(1.02, 1.02, 1.02);
    earthMesh.add(mesh2);
}

function createRing(radius, width, height, texture) {

    var textureLoader = new THREE.TextureLoader();
    var geometry = new THREE.BufferGeometry().fromGeometry(
                new THREEx.Planets._RingGeometry(radius, width, height));

    var material = new THREE.MeshPhongMaterial({
                                                   map: textureLoader.load(texture),
                                                   side: THREE.DoubleSide,
                                                   transparent: true,
                                                   opacity: 0.8
                                               });
    material.map.minFilter = THREE.NearestFilter;
    var mesh = new THREE.Mesh(geometry, material);
    mesh.lookAt(new THREE.Vector3(0, 90, 0));

    return mesh;

}

function createStarfield(radius) {

    var textureLoader = new THREE.TextureLoader();
    var texture = textureLoader.load('images/galaxy_starfield.png')
    var material = new THREE.MeshBasicMaterial({
                                                   map: texture,
                                                   side: THREE.BackSide
                                               })
    var geometry = new THREE.BufferGeometry().fromGeometry(new THREE.SphereGeometry(radius, 32, 32));
    var mesh = new THREE.Mesh(geometry, material)

    return mesh

}

function onResizeGL(canvas) {

    if (camera === undefined) return;

    camera.aspect = canvas.width / canvas.height;
    camera.updateProjectionMatrix();
    renderer.setPixelRatio(canvas.devicePixelRatio);
    renderer.setSize(canvas.width, canvas.height);

}

function onSpeedChanged(value) {

    daysPerFrameScale = value;

}

function setScale(value, focused) {

    // Save actual scale in focus mode
    if (!focused)
        actualScale = value;

    // Limit minimum scaling in focus mode to avoid jitter caused by rounding errors
    if (value <= focusedMinimumScale && (focusedScaling || focused)) {
        planetScale = focusedMinimumScale;
    } else {
        planetScale = actualScale;
    }

    for (var i = 0; i < objects.length; i++) {
        var object = objects[i];
        // first reset scale
        var radius = planets[i]["radius"];
        object.scale.set(radius, radius, radius);
        if (i === SUN) {
            object.scale.multiplyScalar(planetScale / 100);
        } else {
            object.scale.multiplyScalar(planetScale);
        }
        hitObjects[i].scale.set(object.scale.x, object.scale.y, object.scale.z);
    }

}

function prepareFocusedPlanetAnimation() {

    oldCameraPosition = camera.position.clone();

    var planet = SUN;
    if (qmlView.oldPlanet !== SOLAR_SYSTEM)
        planet = qmlView.oldPlanet;
    oldFocusedPlanetPosition = objects[planet].position.clone();
    qmlView.oldPlanet = qmlView.focusedPlanet;

    if (qmlView.focusedPlanet !== SOLAR_SYSTEM && actualScale <= focusedMinimumScale) {
        // Limit minimum scaling in focus mode to avoid jitter caused by rounding errors
        planetScale = focusedMinimumScale;
        setScale(focusedMinimumScale, true);
        focusedScaling = true;
    } else if (focusedScaling === true) {
        // Restore normal scaling
        focusedScaling = false;
        setScale(actualScale);
    }

    calculateLookAtOffset();
    calculateCameraOffset();

}

function setCameraDistance(distance) {

    cameraDistance = distance;

}

function calculateLookAtOffset() {

    var offset = oldFocusedPlanetPosition.clone();

    var planet = 0;
    if (qmlView.focusedPlanet !== SOLAR_SYSTEM)
        planet = qmlView.oldPlanet;

    var focusedPlanetPosition = objects[planet].position.clone();
    offset.sub(focusedPlanetPosition);

    qmlView.xLookAtOffset = offset.x;
    qmlView.yLookAtOffset = offset.y;
    qmlView.zLookAtOffset = offset.z;

}

function calculateCameraOffset() {

    var offset = oldCameraPosition.clone();

    var planet = 0;
    if (qmlView.focusedPlanet !== SOLAR_SYSTEM)
        planet = qmlView.focusedPlanet;

    var newCameraPosition = getNewCameraPosition(getOuterRadius(planet));

    if (qmlView.focusedPlanet !== SUN)
        offset.sub(newCameraPosition);

    if (qmlView.focusedPlanet === SUN && qmlView.oldPlanet === SOLAR_SYSTEM) {
        qmlView.xCameraOffset = Math.abs(offset.x);
        qmlView.yCameraOffset = Math.abs(offset.y);
        qmlView.zCameraOffset = Math.abs(offset.z);
    } else { // from a planet to another
        qmlView.xCameraOffset = offset.x;
        qmlView.yCameraOffset = offset.y;
        qmlView.zCameraOffset = offset.z;
    }

}

function getNewCameraPosition( radius ) {

    var position;
    if (qmlView.focusedPlanet === SOLAR_SYSTEM) {
        position = defaultCameraPosition.clone();
        position.multiplyScalar(cameraDistance);
    } else if (qmlView.focusedPlanet === SUN) {
        position = new THREE.Vector3(radius * planetScale * 2,
                                     radius * planetScale * 2,
                                     radius * planetScale * 2);
        position.multiplyScalar(cameraDistance);

    } else {
        var vec1 = objects[qmlView.focusedPlanet].position.clone();
        var vec2 = new THREE.Vector3(0, 1, 0);
        vec1.normalize();
        vec2.cross(vec1);
        vec2.multiplyScalar(radius * planetScale * cameraDistance * 4);
        vec2.add(objects[qmlView.focusedPlanet].position);
        vec1.set(0, radius * planetScale, 0);
        vec2.add(vec1);
        position = vec2;
    }
    return position;
}

function onDocumentMouseDown(x, y) {

    // Mouse selection for planets and Solar system, not for the Moon.
    // Intersection tests are done against a set of cruder hit objects instead of
    // actual planet meshes, as checking a lot of faces can be slow.
    mouse.set((x / planetCanvas.width) * 2 - 1, - (y / planetCanvas.height ) * 2 + 1);

    raycaster.setFromCamera(mouse, camera);

    var intersects = [];
    var i = 0;
    var objectCount = hitObjects.length - 1; // -1 excludes the moon, which is the last object
    while (i < objectCount) {
        // Update hitObject position
        var objectPos = objects[i].position;
        var hitObject = hitObjects[i];
        hitObject.position.set(objectPos.x, objectPos.y, objectPos.z);
        hitObject.updateMatrixWorld();

        hitObject.raycast( raycaster, intersects );

        i++;
    }
    intersects.sort( raycaster.ascSort );

    var selectedPlanet;

    if (intersects.length > 0) {
        var intersect = intersects[0];

        i = 0;
        while (i < objectCount) {
            if (intersect.object === hitObjects[i]) {
                selectedPlanet = i;
                break;
            }
            i++;
        }
        if (selectedPlanet < NUM_SELECTABLE_PLANETS) {
            qmlView.focusedPlanet = selectedPlanet;
            // Limit minimum scaling in focus mode to avoid jitter caused by rounding errors
            if (actualScale <= focusedMinimumScale) {
                planetScale = focusedMinimumScale;
                setScale(focusedMinimumScale, true);
            }
            focusedScaling = true;
        }
    } else {
        qmlView.focusedPlanet = SOLAR_SYSTEM;
        // Restore normal scaling
        if (focusedScaling === true) {
            focusedScaling = false;
            setScale(actualScale);
        }
    }

}

function paintGL(canvas) {

    if (qmlView.focusedPlanet === SOLAR_SYSTEM)
        daysPerFrame = daysPerFrameScale * 10;
    else
        daysPerFrame = daysPerFrameScale * planets[qmlView.focusedPlanet]["period"] / 100;

    // Advance the time in days
    oldTimeD = currTimeD;
    currTimeD = currTimeD + daysPerFrame;
    var deltaTimeD = currTimeD - oldTimeD;

    // Position the planets orbiting the sun
    for (var i = 1; i < objects.length; i ++) {
        var object = objects[i];
        var planet = planets[i];

        // Bumpmaps of mercury, venus, jupiter and moon need special handling
        if (i == MERCURY || i == VENUS || i == JUPITER || i == MOON)
            object.material.bumpScale = 0.03 * planetScale;
        else
            object.material.bumpScale = 0.3 * planetScale;

        // Calculate the planet orbital elements from the current time in days
        var N =  (planet["N1"] + planet["N2"] * currTimeD) * Math.PI / 180;
        var iPlanet = (planet["i1"] + planet["i2"] * currTimeD) * Math.PI / 180;
        var w =  (planet["w1"] + planet["w2"] * currTimeD) * Math.PI / 180;
        var a = planet["a1"] + planet["a2"] * currTimeD;
        var e = planet["e1"] + planet["e2"] * currTimeD;
        var M = (planet["M1"] + planet["M2"] * currTimeD) * Math.PI / 180;
        var E = M + e * Math.sin(M) * (1.0 + e * Math.cos(M));

        var xv = a * (Math.cos(E) - e);
        var yv = a * (Math.sqrt(1.0 - e * e) * Math.sin(E));
        var v = Math.atan2(yv, xv);

        // Calculate the distance (radius)
        var r = Math.sqrt(xv * xv + yv * yv);

        // From http://www.davidcolarusso.com/astro/
        // Modified to compensate for the right handed coordinate system of OpenGL
        var xh = r * (Math.cos(N) * Math.cos(v + w)
                      - Math.sin(N) * Math.sin(v + w) * Math.cos(iPlanet));
        var zh = -r * (Math.sin(N) * Math.cos(v + w)
                       + Math.cos(N) * Math.sin(v + w) * Math.cos(iPlanet));
        var yh = r * (Math.sin(w + v) * Math.sin(iPlanet));

        // Apply the position offset from the center of orbit to the bodies
        var centerOfOrbit = objects[planet["centerOfOrbit"]];
        object.position.set(centerOfOrbit.position.x + xh * auScale,
                            centerOfOrbit.position.y + yh * auScale,
                            centerOfOrbit.position.z + zh * auScale);

        // Calculate and apply the appropriate axis tilt to the bodies
        // and rotate them around the axis
        var radians = planet["tilt"] * Math.PI / 180; // tilt in radians
        object.rotation.order = 'ZXY';
        object.rotation.x = 0;
        object.rotation.y += (deltaTimeD / planet["period"]) * 2 * Math.PI;
        object.rotation.z = radians;
    }

    // rotate the Sun
    var sun = objects[SUN];
    sun.rotation.order = 'ZXY';
    sun.rotation.x = 0;
    sun.rotation.y += (deltaTimeD / planets[SUN]["period"]) * 2 * Math.PI;
    sun.rotation.z = planets[SUN]["tilt"] * Math.PI / 180; // tilt in radians

    // calculate the outer radius of the focused item
    var outerRadius = getOuterRadius(qmlView.focusedPlanet);

    // get the appropriate near plane position for the camera and animate it with QML animations
    qmlView.cameraNear = outerRadius;

    camera.near = qmlView.cameraNear;
    camera.updateProjectionMatrix();

    // Calculate and set camera position
    var cameraPosition = getNewCameraPosition(outerRadius);
    var cameraOffset = new THREE.Vector3(qmlView.xCameraOffset,
                                         qmlView.yCameraOffset,
                                         qmlView.zCameraOffset);
    cameraPosition.add(cameraOffset);
    camera.position.set(cameraPosition.x, cameraPosition.y, cameraPosition.z);

    // Calculate and set camera look-at point
    var lookAtPlanet = SUN;
    if (qmlView.focusedPlanet !== SOLAR_SYSTEM)
        lookAtPlanet = qmlView.focusedPlanet;
    var cameraLookAt = objects[lookAtPlanet].position.clone();
    var lookAtOffset = new THREE.Vector3(qmlView.xLookAtOffset,
                                         qmlView.yLookAtOffset,
                                         qmlView.zLookAtOffset);
    cameraLookAt.add(lookAtOffset);
    camera.lookAt(cameraLookAt);

    // Render the scene
    renderer.render(scene, camera);

}

function getOuterRadius( planet ) {

    var outerRadius = solarDistance;
    if (planet !== SOLAR_SYSTEM) {
        outerRadius = planets[planet]["radius"];
        if (planet === SATURN) {
            outerRadius =+ saturnOuterRadius;
        } else if (planet === URANUS) {
            outerRadius =+ uranusOuterRadius;
        } else if (planet === SUN) {
            outerRadius = planets[planet]["radius"] / 100;
        }
    }

    return outerRadius;
}
