/********************************************************************************
**
** The MIT License
**
** Copyright © 2014 threex.planets Jerome Etienne
**
** Permission is hereby granted, free of charge, to any person obtaining a copy
** of this software and associated documentation files (the "Software"), to deal
** in the Software without restriction, including without limitation the rights
** to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
** copies of the Software, and to permit persons to whom the Software is
** furnished to do so, subject to the following conditions:
**
** The above copyright notice and this permission notice shall be included in
** all copies or substantial portions of the Software.
**
** THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
** IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
** FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
** AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
** LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
** OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
** THE SOFTWARE.
**
********************************************************************************/

var THREEx = THREEx || {}

THREEx.Planets	= {}

THREEx.Planets.baseURL	= '../'

// from http://planetpixelemporium.com/

THREEx.Planets.createSun	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var texture	= new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/sunmap.jpg')
    var material	= new THREE.MeshPhongMaterial({
        map	: texture,
        bumpMap	: texture,
        bumpScale: 0.05,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}

THREEx.Planets.createMercury	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var material	= new THREE.MeshPhongMaterial({
        map	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/mercurymap.jpg'),
        bumpMap	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/mercurybump.jpg'),
        bumpScale: 0.005,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}

THREEx.Planets.createVenus	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var material	= new THREE.MeshPhongMaterial({
        map	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/venusmap.jpg'),
        bumpMap	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/venusbump.jpg'),
        bumpScale: 0.005,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}

THREEx.Planets.createEarth	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var material	= new THREE.MeshPhongMaterial({
        map		: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/earthmap1k.jpg'),
        bumpMap		: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/earthbump1k.jpg'),
        bumpScale	: 0.05,
        specularMap	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/earthspec1k.jpg'),
        specular	: new THREE.Color('grey'),
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}

THREEx.Planets.createEarthCloud	= function(){
    // create destination canvas
    var canvasResult	= document.createElement('canvas')
    canvasResult.width	= 1024
    canvasResult.height	= 512
    var contextResult	= canvasResult.getContext('2d')

    // load earthcloudmap
    var imageMap	= new Image();
    imageMap.addEventListener("load", function() {

        // create dataMap ImageData for earthcloudmap
        var canvasMap	= document.createElement('canvas')
        canvasMap.width	= imageMap.width
        canvasMap.height= imageMap.height
        var contextMap	= canvasMap.getContext('2d')
        contextMap.drawImage(imageMap, 0, 0)
        var dataMap	= contextMap.getImageData(0, 0, canvasMap.width, canvasMap.height)

        // load earthcloudmaptrans
        var imageTrans	= new Image();
        imageTrans.addEventListener("load", function(){
            // create dataTrans ImageData for earthcloudmaptrans
            var canvasTrans		= document.createElement('canvas')
            canvasTrans.width	= imageTrans.width
            canvasTrans.height	= imageTrans.height
            var contextTrans	= canvasTrans.getContext('2d')
            contextTrans.drawImage(imageTrans, 0, 0)
            var dataTrans		= contextTrans.getImageData(0, 0, canvasTrans.width, canvasTrans.height)
            // merge dataMap + dataTrans into dataResult
            var dataResult		= contextMap.createImageData(canvasMap.width, canvasMap.height)
            for(var y = 0, offset = 0; y < imageMap.height; y++){
                for(var x = 0; x < imageMap.width; x++, offset += 4){
                    dataResult.data[offset+0]	= dataMap.data[offset+0]
                    dataResult.data[offset+1]	= dataMap.data[offset+1]
                    dataResult.data[offset+2]	= dataMap.data[offset+2]
                    dataResult.data[offset+3]	= 255 - dataTrans.data[offset+0]
                }
            }
            // update texture with result
            contextResult.putImageData(dataResult,0,0)
            material.map.needsUpdate = true;
        })
        imageTrans.src	= THREEx.Planets.baseURL+'images/earthcloudmaptrans.jpg';
    }, false);
    imageMap.src	= THREEx.Planets.baseURL+'images/earthcloudmap.jpg';

    var geometry	= new THREE.SphereGeometry(0.51, 32, 32)
    var material	= new THREE.MeshPhongMaterial({
        map		: new THREE.Texture(canvasResult),
        side		: THREE.DoubleSide,
        transparent	: true,
        opacity		: 0.8,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}


THREEx.Planets.createMoon	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var material	= new THREE.MeshPhongMaterial({
        map	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/moonmap1k.jpg'),
        bumpMap	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/moonbump1k.jpg'),
        bumpScale: 0.002,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}

THREEx.Planets.createMars	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var material	= new THREE.MeshPhongMaterial({
        map	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/marsmap1k.jpg'),
        bumpMap	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/marsbump1k.jpg'),
        bumpScale: 0.05,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}

THREEx.Planets.createJupiter	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var texture	= new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/jupitermap.jpg')
    var material	= new THREE.MeshPhongMaterial({
        map	: texture,
        bumpMap	: texture,
        bumpScale: 0.02,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}


THREEx.Planets.createSaturn	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var texture	= new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/saturnmap.jpg')
    var material	= new THREE.MeshPhongMaterial({
        map	: texture,
        bumpMap	: texture,
        bumpScale: 0.05,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}


THREEx.Planets.createSaturnRing	= function(){
    // create destination canvas
    var canvasResult	= document.createElement('canvas')
    canvasResult.width	= 915
    canvasResult.height	= 64
    var contextResult	= canvasResult.getContext('2d')

    // load earthcloudmap
    var imageMap	= new Image();
    imageMap.addEventListener("load", function() {

        // create dataMap ImageData for earthcloudmap
        var canvasMap	= document.createElement('canvas')
        canvasMap.width	= imageMap.width
        canvasMap.height= imageMap.height
        var contextMap	= canvasMap.getContext('2d')
        contextMap.drawImage(imageMap, 0, 0)
        var dataMap	= contextMap.getImageData(0, 0, canvasMap.width, canvasMap.height)

        // load earthcloudmaptrans
        var imageTrans	= new Image();
        imageTrans.addEventListener("load", function(){
            // create dataTrans ImageData for earthcloudmaptrans
            var canvasTrans		= document.createElement('canvas')
            canvasTrans.width	= imageTrans.width
            canvasTrans.height	= imageTrans.height
            var contextTrans	= canvasTrans.getContext('2d')
            contextTrans.drawImage(imageTrans, 0, 0)
            var dataTrans		= contextTrans.getImageData(0, 0, canvasTrans.width, canvasTrans.height)
            // merge dataMap + dataTrans into dataResult
            var dataResult		= contextMap.createImageData(canvasResult.width, canvasResult.height)
            for(var y = 0, offset = 0; y < imageMap.height; y++){
                for(var x = 0; x < imageMap.width; x++, offset += 4){
                    dataResult.data[offset+0]	= dataMap.data[offset+0]
                    dataResult.data[offset+1]	= dataMap.data[offset+1]
                    dataResult.data[offset+2]	= dataMap.data[offset+2]
                    dataResult.data[offset+3]	= 255 - dataTrans.data[offset+0]/4
                }
            }
            // update texture with result
            contextResult.putImageData(dataResult,0,0)
            material.map.needsUpdate = true;
        })
        imageTrans.src	= THREEx.Planets.baseURL+'images/saturnringpattern.gif';
    }, false);
    imageMap.src	= THREEx.Planets.baseURL+'images/saturnringcolor.jpg';

    var geometry	= new THREEx.Planets._RingGeometry(0.55, 0.75, 64);
    var material	= new THREE.MeshPhongMaterial({
        map		: new THREE.Texture(canvasResult),
        // map		: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/ash_uvgrid01.jpg'),
        side		: THREE.DoubleSide,
        transparent	: true,
        opacity		: 0.8,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    mesh.lookAt(new THREE.Vector3(0.5,-4,1))
    return mesh
}


THREEx.Planets.createUranus	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var texture	= new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/uranusmap.jpg')
    var material	= new THREE.MeshPhongMaterial({
        map	: texture,
        bumpMap	: texture,
        bumpScale: 0.05,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}

THREEx.Planets.createUranusRing	= function(){
    // create destination canvas
    var canvasResult	= document.createElement('canvas')
    canvasResult.width	= 1024
    canvasResult.height	= 72
    var contextResult	= canvasResult.getContext('2d')

    // load earthcloudmap
    var imageMap	= new Image();
    imageMap.addEventListener("load", function() {

        // create dataMap ImageData for earthcloudmap
        var canvasMap	= document.createElement('canvas')
        canvasMap.width	= imageMap.width
        canvasMap.height= imageMap.height
        var contextMap	= canvasMap.getContext('2d')
        contextMap.drawImage(imageMap, 0, 0)
        var dataMap	= contextMap.getImageData(0, 0, canvasMap.width, canvasMap.height)

        // load earthcloudmaptrans
        var imageTrans	= new Image();
        imageTrans.addEventListener("load", function(){
            // create dataTrans ImageData for earthcloudmaptrans
            var canvasTrans		= document.createElement('canvas')
            canvasTrans.width	= imageTrans.width
            canvasTrans.height	= imageTrans.height
            var contextTrans	= canvasTrans.getContext('2d')
            contextTrans.drawImage(imageTrans, 0, 0)
            var dataTrans		= contextTrans.getImageData(0, 0, canvasTrans.width, canvasTrans.height)
            // merge dataMap + dataTrans into dataResult
            var dataResult		= contextMap.createImageData(canvasResult.width, canvasResult.height)
            for(var y = 0, offset = 0; y < imageMap.height; y++){
                for(var x = 0; x < imageMap.width; x++, offset += 4){
                    dataResult.data[offset+0]	= dataMap.data[offset+0]
                    dataResult.data[offset+1]	= dataMap.data[offset+1]
                    dataResult.data[offset+2]	= dataMap.data[offset+2]
                    dataResult.data[offset+3]	= 255 - dataTrans.data[offset+0]/2
                }
            }
            // update texture with result
            contextResult.putImageData(dataResult,0,0)
            material.map.needsUpdate = true;
        })
        imageTrans.src	= THREEx.Planets.baseURL+'images/uranusringtrans.gif';
    }, false);
    imageMap.src	= THREEx.Planets.baseURL+'images/uranusringcolour.jpg';

    var geometry	= new THREEx.Planets._RingGeometry(0.55, 0.75, 64);
    var material	= new THREE.MeshPhongMaterial({
        map		: new THREE.Texture(canvasResult),
        // map		: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/ash_uvgrid01.jpg'),
        side		: THREE.DoubleSide,
        transparent	: true,
        opacity		: 0.8,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    mesh.lookAt(new THREE.Vector3(0.5,-4,1))
    return mesh
}


THREEx.Planets.createNeptune	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var texture	= new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/neptunemap.jpg')
    var material	= new THREE.MeshPhongMaterial({
        map	: texture,
        bumpMap	: texture,
        bumpScale: 0.05,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}



THREEx.Planets.createPluto	= function(){
    var geometry	= new THREE.SphereGeometry(0.5, 32, 32)
    var material	= new THREE.MeshPhongMaterial({
        map	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/plutomap1k.jpg'),
        bumpMap	: new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/plutobump1k.jpg'),
        bumpScale: 0.005,
    })
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}

THREEx.Planets.createStarfield	= function(radius){
    var texture	= new THREE.TextureLoader().load(THREEx.Planets.baseURL+'images/galaxy_starfield.png')
    var material	= new THREE.MeshBasicMaterial({
        map	: texture,
        side	: THREE.BackSide
    })
    var geometry	= new THREE.SphereGeometry(radius, 32, 32)
    var mesh	= new THREE.Mesh(geometry, material)
    return mesh
}


//////////////////////////////////////////////////////////////////////////////////
//		comment								//
//////////////////////////////////////////////////////////////////////////////////

/**
 * change the original from three.js because i needed different UV
 *
 * @author Kaleb Murphy
 * @author jerome etienne
 *
 * Minor changes
 * @author Titta Heikkala
 */
THREEx.Planets._RingGeometry = function ( innerRadius, outerRadius, thetaSegments ) {

    THREE.Geometry.call( this )

    innerRadius	= innerRadius || 0
    outerRadius	= outerRadius || 50
    thetaSegments	= thetaSegments	|| 8

    var normal	= new THREE.Vector3( 0, 0, 1 )

    for(var i = 0; i < thetaSegments; i++ ){
        var angleLo	= (i / thetaSegments) *Math.PI*2
        var angleHi	= ((i+1) / thetaSegments) *Math.PI*2

        var vertex1	= new THREE.Vector3(innerRadius * Math.cos(angleLo), innerRadius * Math.sin(angleLo), 0);
        var vertex2	= new THREE.Vector3(outerRadius * Math.cos(angleLo), outerRadius * Math.sin(angleLo), 0);
        var vertex3	= new THREE.Vector3(innerRadius * Math.cos(angleHi), innerRadius * Math.sin(angleHi), 0);
        var vertex4	= new THREE.Vector3(outerRadius * Math.cos(angleHi), outerRadius * Math.sin(angleHi), 0);

        this.vertices.push( vertex1 );
        this.vertices.push( vertex2 );
        this.vertices.push( vertex3 );
        this.vertices.push( vertex4 );


        var vertexIdx	= i * 4;

        // Create the first triangle
        var face = new THREE.Face3(vertexIdx + 0, vertexIdx + 1, vertexIdx + 2, normal);
        var uvs = []

        var uv = new THREE.Vector2(0, 0)
        uvs.push(uv)
        var uv = new THREE.Vector2(1, 0)
        uvs.push(uv)
        var uv = new THREE.Vector2(0, 1)
        uvs.push(uv)

        this.faces.push(face);
        this.faceVertexUvs[0].push(uvs);

        // Create the second triangle
        var face = new THREE.Face3(vertexIdx + 2, vertexIdx + 1, vertexIdx + 3, normal);
        var uvs = []

        var uv = new THREE.Vector2(0, 1)
        uvs.push(uv)
        var uv = new THREE.Vector2(1, 0)
        uvs.push(uv)
        var uv = new THREE.Vector2(1, 1)
        uvs.push(uv)

        this.faces.push(face);
        this.faceVertexUvs[0].push(uvs);
    }

    this.computeFaceNormals();

    this.boundingSphere = new THREE.Sphere( new THREE.Vector3(), outerRadius );

};
THREEx.Planets._RingGeometry.prototype = Object.create( THREE.Geometry.prototype );

