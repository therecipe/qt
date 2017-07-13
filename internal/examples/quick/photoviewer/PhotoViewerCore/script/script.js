.pragma library

function getWidth(string) {
    return (string.match(/width=\"([0-9]+)\"/))[1]
}

function getHeight(string) {
    return (string.match(/height=\"([0-9]+)\"/))[1]
}

function getImagePath(string) {
    var pattern = /src=\"http:\/\/(\S+)\"/
    return (string.match(pattern))[1]
}

function calculateScale(width, height, cellSize) {
    var widthScale = (cellSize * 1.0) / width
    var heightScale = (cellSize * 1.0) / height
    var scale = 0

    if (widthScale <= heightScale) {
        scale = widthScale;
    } else if (heightScale < widthScale) {
        scale = heightScale;
    }
    return scale;
}
