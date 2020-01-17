
var newRoostCenterPos = Qt.point(0,0)
var randomValue
var coinCenterPos = Qt.point(0,0)
var coinPositionModifier = 0.7*scene.gridSize
var gridSizeHalf = scene.gridSize/2
var newWindowTopleftPos = Qt.point(0,0)
var roostUrl = Qt.resolvedUrl("../entities/Roost.qml")
var coinUrl = Qt.resolvedUrl("../entities/Coin.qml")
var windowUrl = Qt.resolvedUrl("../entities/HenhouseWindow.qml")
var newElementProperties = {}

function createRandomRowForRowNumber(rowNumber) {
    for(var i=0; i<roostColumns; i++) {
        // Performance optimization, do not create own values for mathrandom during runtime (garbarge)
        //if(Math.random() < platformCreationProbability ) {
        randomValue = Math.random()
        if(randomValue < platformCreationProbability ) {

            // Performance optimization, do not recreate objects during runtime (garbarge)
            //var newRoostCenterPos = Qt.point(i*gridSize + gridSize/2, );
            newRoostCenterPos.x = i*gridSize + gridSizeHalf
            newRoostCenterPos.y = rowNumber*gridSize + gridSizeHalf

            if(physicsWorld.bodyAt(newRoostCenterPos)) {
                console.debug("no Roost can be created because there is a window already");
                continue;
            }

            console.debug("creating a new Roost at position", i*gridSize + gridSize/2, ",", rowNumber*gridSize + gridSize/2);

            // Performance optimization, do not provide properties as own object during runtime (garbarge)
//            entityManager.createEntityFromUrlWithProperties(Qt.resolvedUrl("entities/Roost.qml"),
//                                 {"x": newRoostCenterPos.x,
//                                     "y": newRoostCenterPos.y
//                                 });
            newElementProperties.x = newRoostCenterPos.x
            newElementProperties.y = newRoostCenterPos.y
            entityManager.createEntityFromUrlWithProperties(roostUrl,newElementProperties)

            // create a coin in 30% of all created blocks
            randomValue = Math.random()
            if(randomValue < coinCreationPropability) {

                // look at 1 grid position above
                // Performance optimization, do not recreate objects during runtime (garbarge)
                // var coinCenterPos = Qt.point(newRoostCenterPos.x, newRoostCenterPos.y-scene.gridSize);
                coinCenterPos.x = newRoostCenterPos.x
                coinCenterPos.y = newRoostCenterPos.y-scene.gridSize

                // test if one grid above is an empty field (so if no block is built there) - if so, a coin can be created
                if(physicsWorld.bodyAt(coinCenterPos)) {
                    console.debug("there is a block above the to create block, don't create a coin here!")
                    continue;
                }


                /*entityManager.createEntityFromUrlWithProperties*/
                // Performance optimization, do not provide properties as own object during runtime (garbarge)
//                entityManager.createEntityFromUrlWithProperties(Qt.resolvedUrl("entities/Coin.qml"),
//                                     {"x": coinCenterPos.x,
//                                         "y": coinCenterPos.y+0.7*scene.gridSize // move slightly up, so it looks better
//                                     });
                newElementProperties.x = coinCenterPos.x
                newElementProperties.y = coinCenterPos.y+coinPositionModifier
                entityManager.createEntityFromUrlWithProperties(coinUrl,newElementProperties)
            }
        } else if(i < roostColumns-1 && randomValue < windowCreationProbability ) {


            // the i<columns-1 check is required, because the window has a size of 2 grids, and its anchor is top left

            //var newWindowTopleftPos = Qt.point(i*gridSize, rowNumber*gridSize);
            // Performance optimization, do not provide properties as own object during runtime (garbarge)
            newWindowTopleftPos.x = i*gridSize
            newWindowTopleftPos.y = rowNumber*gridSize

            console.debug("newWindowTopleftPos.y-lastWindowY:", newWindowTopleftPos.y-lastWindowY)
            // this avoids creating too many windows, so not possible to have more than 2 on a scene with this code!
            if(newWindowTopleftPos.y-lastWindowY < minimumWindowHeightDifference) {
                console.debug("difference between last Window and current to create one too small!")
                continue;
            }

            // this might happen if a window was already created at this position
            if(physicsWorld.bodyAt(newWindowTopleftPos)) {
                console.debug("body at position x:", newWindowTopleftPos.x, ", y:", newWindowTopleftPos.y, physicsWorld.bodyAt(newWindowTopleftPos))
                console.debug("there is a window at the position where to create a window, so no creation is done");
                continue;
            }

            //console.debug("creating a new Window... at position x:", newWindowTopleftPos.x, ", y:", newWindowTopleftPos.y, "lastWindowY:", lastWindowY);
            // Performance optimization, do not provide properties as own object during runtime (garbarge)
//            entityManager.createEntityFromUrlWithProperties(Qt.resolvedUrl("entities/HenhouseWindow.qml"),
//                                 {"x": newWindowTopleftPos.x,
//                                     "y": newWindowTopleftPos.y,
//                                     "z": 0 // put behind all others, except the background
//                                 });
            newElementProperties.x = newWindowTopleftPos.x
            newElementProperties.y = newWindowTopleftPos.y
            newElementProperties.z = 0 // put behind all others, except the background
            entityManager.createEntityFromUrlWithProperties(windowUrl,newElementProperties)

            lastWindowY = newWindowTopleftPos.y;
        }

    }
}
