// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6
import QtGraphicalEffects 1.0

Image {
    id: image
    property string imageName: ""
    onImageNameChanged: {
        calculatePath()
    }
    property int imageSize: 0
    onImageSizeChanged: {
        if(imageName.length > 0) {
            calculatePath()
        }
    }
    // default: primary color
    property alias color: colorOverlay.color
    //trick: to be triggered if folder changed
    property string currentIconFolder: iconFolder
    onCurrentIconFolderChanged: {
        if(imageName.length > 0) {
            calculatePath()
        }
    }
    ColorOverlay {
        id: colorOverlay
        anchors.fill: image
        source: image
        color: primaryColor
    }
    function calculatePath() {
        var path = "qrc:/images/"+currentIconFolder
        switch(imageSize) {
            case 18:
                path += "/x18/"
                break;
            case 36:
                path += "/x36/"
                break;
            case 48:
                path += "/x48/"
                break;
            default:
                path += "/"
        } // switch
        path += imageName
        source = path
    }
}
