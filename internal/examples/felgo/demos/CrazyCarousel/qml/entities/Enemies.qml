import QtQuick 2.0
import Felgo 3.0

EntityBase {
    id: enemies
    entityType: "enemies"

    Enemy {
        id: enemy1
        width: 95.5; height: 112
        image.source: "../../assets/enemy1.png"
    }

    Enemy {
        id: enemy2
        x: 95.5
        width: 69.5; height: 112
        image.source: "../../assets/enemy2.png"
    }

    Enemy {
        id: enemy3
        x: 165
        width: 49.5; height: 112
        image.source: "../../assets/enemy3.png"
    }

    Enemy {
        id: enemy4
        x: 214.5
        width: 105.5; height: 112
        image.source: "../../assets/enemy4.png"
    }
}
