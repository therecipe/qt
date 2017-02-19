/*
 * The MIT License
 * Copyright &copy; 2010-2014 three.js authors
 * Copyright &copy; 2014 Digia Plc and/or its subsidiary(-ies).
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

function GLModel() {
    this.vertices = [];
    this.normals = [];
    this.texCoords = [];
    this.indices = [];
}

function Geometry() {
    this.vertices = [];
    this.faceVertexUvs = [];
    this.faces = [];
}

function Face3( a, b, c, normal, color, materialIndex ) {

    this.a = a;
    this.b = b;
    this.c = c;

    this.normal = normal instanceof Vector3 ? normal : new Vector3();
    this.vertexNormals = normal instanceof Array ? normal : [ ];

    this.color = color instanceof Array ? color :  [ ];
    this.vertexColors = color instanceof Array ? color : [];

    this.vertexTangents = [];

    this.materialIndex = materialIndex !== undefined ? materialIndex : 0;

    this.centroid = new Vector3()
};

function Vector2( x, y ) {
    this.x = x || 0;
    this.y = y || 0;
};

function Vector3( x, y, z ) {
    this.x = x || 0;
    this.y = y || 0;
    this.z = z || 0;
};

Vector3.prototype = {

    constructor: Vector3,

    set: function ( x, y, z ) {

        this.x = x;
        this.y = y;
        this.z = z;

        return this;
    },

    copy: function ( v ) {

        this.x = v.x;
        this.y = v.y;
        this.z = v.z;

        return this;

    }
};

Face3.prototype = {

    constructor: Face3,

    clone: function () {

        var face = new Face3( this.a, this.b, this.c );

        face.normal.copy( this.normal );
        face.color.copy( this.color );
        face.centroid.copy( this.centroid );

        face.materialIndex = this.materialIndex;

        var i, il;
        for ( i = 0, il = this.vertexNormals.length; i < il; i ++ ) face.vertexNormals[ i ] = this.vertexNormals[ i ].clone();
        for ( i = 0, il = this.vertexColors.length; i < il; i ++ ) face.vertexColors[ i ] = this.vertexColors[ i ].clone();
        for ( i = 0, il = this.vertexTangents.length; i < il; i ++ ) face.vertexTangents[ i ] = this.vertexTangents[ i ].clone();

        return face;
    }
};

function parseJSON3DModel( json, texturePath )
{
    var formatVersion = json.metadata.formatVersion;

    if (formatVersion < 4.0 ) {
        var i, j;
        var geometry = new Geometry();
        var scale = ( json.scale !== undefined ) ? 1.0 / json.scale : 1.0;
        parseModel( json, geometry, scale );

        // Translate to model we can use from GL
        var glModel = new GLModel();


        glModel.texCoords[0] = [];
        for (var faceIdx = 0; faceIdx < geometry.faces.length; faceIdx++) {
            var face = geometry.faces[faceIdx];
            // Array indices
            var fidx = faceIdx * 3;
            glModel.indices.push(fidx);
            glModel.indices.push(fidx + 1);
            glModel.indices.push(fidx + 2);

            var v1 = geometry.vertices[face.a];
            var v2 = geometry.vertices[face.b];
            var v3 = geometry.vertices[face.c];

            glModel.vertices.push(v1.x);
            glModel.vertices.push(v1.y);
            glModel.vertices.push(v1.z);
            glModel.vertices.push(v2.x);
            glModel.vertices.push(v2.y);
            glModel.vertices.push(v2.z);
            glModel.vertices.push(v3.x);
            glModel.vertices.push(v3.y);
            glModel.vertices.push(v3.z);

            // Materials
            // face.materialIndex

            // Only parse first layer of UVs for this face
            for(var uvLayer = 0; uvLayer < geometry.faceVertexUvs.length; uvLayer++) {
                var faceUVs = geometry.faceVertexUvs[uvLayer][faceIdx];

                glModel.texCoords[uvLayer].push(faceUVs[0].x);
                glModel.texCoords[uvLayer].push(faceUVs[0].y);
                glModel.texCoords[uvLayer].push(faceUVs[1].x);
                glModel.texCoords[uvLayer].push(faceUVs[1].y);
                glModel.texCoords[uvLayer].push(faceUVs[2].x);
                glModel.texCoords[uvLayer].push(faceUVs[2].y);
            }

            // Normal
            if (face.vertexNormals !== undefined) {
                // Per vertex normals
                var vrtNormals = face.vertexNormals;

                var v1 = vrtNormals[0];
                var v2 = vrtNormals[1];
                var v3 = vrtNormals[2];

                glModel.normals.push(v1.x);
                glModel.normals.push(v1.y);
                glModel.normals.push(v1.z);
                glModel.normals.push(v2.x);
                glModel.normals.push(v2.y);
                glModel.normals.push(v2.z);
                glModel.normals.push(v3.x);
                glModel.normals.push(v3.y);
                glModel.normals.push(v3.z);
            } else if (face.normal !== undefined) {
                // Per face normal
                glModel.normals.push(face.normal.x);
                glModel.normals.push(face.normal.y);
                glModel.normals.push(face.normal.z);
                glModel.normals.push(face.normal.x);
                glModel.normals.push(face.normal.y);
                glModel.normals.push(face.normal.z);
                glModel.normals.push(face.normal.x);
                glModel.normals.push(face.normal.y);
                glModel.normals.push(face.normal.z);
            }
        }
    }

    return glModel;
};

function parseModel( json, geometry, scale ) {

    function isBitSet( value, position )
    {
        return value & ( 1 << position );
    }

    var i, j, fi,

            offset, zLength,

            colorIndex, normalIndex, uvIndex, materialIndex,

            type,
            isQuad,
            hasMaterial,
            hasFaceVertexUv,
            hasFaceNormal, hasFaceVertexNormal,
            hasFaceColor, hasFaceVertexColor,

            vertex, face, faceA, faceB, color, hex, normal,

            uvLayer, uv, u, v,

            faces = json.faces,
            vertices = json.vertices,
            normals = json.normals,
            colors = json.colors,

            nUvLayers = 0;

    if ( json.uvs !== undefined ) {
        // disregard empty arrays
        for ( i = 0; i < json.uvs.length; i++ ) {
            if ( json.uvs[ i ].length > 0 ) nUvLayers ++;
        }

        for ( i = 0; i < nUvLayers; i++ ) {
            geometry.faceVertexUvs[ i ] = [];
        }
    }

    offset = 0;
    zLength = vertices.length;

    while ( offset < zLength ) {
        vertex = new Vector3();

        vertex.x = vertices[ offset ++ ] * scale;
        vertex.y = vertices[ offset ++ ] * scale;
        vertex.z = vertices[ offset ++ ] * scale;

        geometry.vertices.push( vertex );
    }

    offset = 0;
    zLength = faces.length;

    while ( offset < zLength ) {
        type = faces[ offset ++ ];

        isQuad              = isBitSet( type, 0 );
        hasMaterial         = isBitSet( type, 1 );
        hasFaceVertexUv     = isBitSet( type, 3 );
        hasFaceNormal       = isBitSet( type, 4 );
        hasFaceVertexNormal = isBitSet( type, 5 );
        hasFaceColor	    = isBitSet( type, 6 );
        hasFaceVertexColor  = isBitSet( type, 7 );

        // console.log("type", type, "bits", isQuad, hasMaterial, hasFaceVertexUv, hasFaceNormal, hasFaceVertexNormal, hasFaceColor, hasFaceVertexColor);

        if ( isQuad ) {
            faceA = new Face3();
            faceA.a = faces[ offset ];
            faceA.b = faces[ offset + 1 ];
            faceA.c = faces[ offset + 3 ];

            faceB = new Face3();
            faceB.a = faces[ offset + 1 ];
            faceB.b = faces[ offset + 2 ];
            faceB.c = faces[ offset + 3 ];

            offset += 4;

            if ( hasMaterial ) {
                materialIndex = faces[ offset ++ ];
                faceA.materialIndex = materialIndex;
                faceB.materialIndex = materialIndex;
            }

            // to get face <=> uv index correspondence
            fi = geometry.faces.length;

            if ( hasFaceVertexUv ) {
                for ( i = 0; i < nUvLayers; i++ ) {
                    uvLayer = json.uvs[ i ];
                    geometry.faceVertexUvs[ i ][ fi ] = [];
                    geometry.faceVertexUvs[ i ][ fi + 1 ] = []

                    for ( j = 0; j < 4; j ++ ) {
                        uvIndex = faces[ offset ++ ];

                        u = uvLayer[ uvIndex * 2 ];
                        v = uvLayer[ uvIndex * 2 + 1 ];

                        uv = new Vector2( u, v );

                        if ( j !== 2 ) geometry.faceVertexUvs[ i ][ fi ].push( uv );
                        if ( j !== 0 ) geometry.faceVertexUvs[ i ][ fi + 1 ].push( uv );
                    }
                }
            }

            if ( hasFaceNormal ) {
                normalIndex = faces[ offset ++ ] * 3;
                faceA.normal.set(
                            normals[ normalIndex ++ ],
                            normals[ normalIndex ++ ],
                            normals[ normalIndex ]
                            );
                faceB.normal.copy( faceA.normal );
            }

            if ( hasFaceVertexNormal ) {
                for ( i = 0; i < 4; i++ ) {
                    normalIndex = faces[ offset ++ ] * 3;
                    normal = new Vector3(
                                normals[ normalIndex ++ ],
                                normals[ normalIndex ++ ],
                                normals[ normalIndex ]
                                );

                    if ( i !== 2 ) faceA.vertexNormals.push( normal );
                    if ( i !== 0 ) faceB.vertexNormals.push( normal );
                }
            }

            if ( hasFaceColor ) {
                colorIndex = faces[ offset ++ ];
                hex = colors[ colorIndex ];
                //                    faceA.color.setHex( hex );
                //                    faceB.color.setHex( hex );
            }

            if ( hasFaceVertexColor ) {
                for ( i = 0; i < 4; i++ ) {
                    colorIndex = faces[ offset ++ ];
                    hex = colors[ colorIndex ];

                    //                        if ( i !== 2 ) faceA.vertexColors.push( new THREE.Color( hex ) );
                    //                        if ( i !== 0 ) faceB.vertexColors.push( new THREE.Color( hex ) );
                }
            }

            geometry.faces.push( faceA );
            geometry.faces.push( faceB );

        } else {
            face = new Face3();
            face.a = faces[ offset ++ ];
            face.b = faces[ offset ++ ];
            face.c = faces[ offset ++ ];

            if ( hasMaterial ) {
                materialIndex = faces[ offset ++ ];
                face.materialIndex = materialIndex;
            }

            // to get face <=> uv index correspondence

            fi = geometry.faces.length;

            if ( hasFaceVertexUv ) {

                for ( i = 0; i < nUvLayers; i++ ) {

                    uvLayer = json.uvs[ i ];

                    geometry.faceVertexUvs[ i ][ fi ] = [];

                    for ( j = 0; j < 3; j ++ ) {

                        uvIndex = faces[ offset ++ ];

                        u = uvLayer[ uvIndex * 2 ];
                        v = uvLayer[ uvIndex * 2 + 1 ];

                        uv = new Vector2( u, v );

                        geometry.faceVertexUvs[ i ][ fi ].push( uv );

                    }

                }

            }

            if ( hasFaceNormal ) {

                normalIndex = faces[ offset ++ ] * 3;

                face.normal.set(
                            normals[ normalIndex ++ ],
                            normals[ normalIndex ++ ],
                            normals[ normalIndex ]
                            );

            }

            if ( hasFaceVertexNormal ) {

                for ( i = 0; i < 3; i++ ) {

                    normalIndex = faces[ offset ++ ] * 3;

                    normal = new Vector3(
                                normals[ normalIndex ++ ],
                                normals[ normalIndex ++ ],
                                normals[ normalIndex ]
                                );

                    face.vertexNormals.push( normal );

                }

            }


            if ( hasFaceColor ) {
                //                    colorIndex = faces[ offset ];
                offset++;
                //                    face.color.setHex( colors[ colorIndex ] );
            }


            if ( hasFaceVertexColor ) {

                for ( i = 0; i < 3; i++ ) {

                    //                        colorIndex = faces[offset];
                    offset++;
                    //                        face.vertexColors.push( new THREE.Color( colors[ colorIndex ] ) );

                }

            }

            geometry.faces.push( face );

        }

    }
};
