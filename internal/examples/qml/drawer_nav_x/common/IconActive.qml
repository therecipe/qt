// ekke (Ekkehard Gentz) @ekkescorner
import QtQuick 2.6

Image {
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
    //trick: to be triggered if folder changed
    property string currentIconFolder: iconFolder
    onCurrentIconFolderChanged: {
        if(imageName.length > 0) {
            calculatePath()
        }
    }
    opacity: iconActiveOpacity
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
