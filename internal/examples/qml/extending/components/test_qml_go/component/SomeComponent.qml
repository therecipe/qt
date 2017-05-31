import Charts 1.0
import QtQuick 2.0

Rectangle {
    width: 300; height: 200

    color: "blue"

    PieChart {
        id: aPieChart
        anchors.centerIn: parent
        width: 100; height: 100
        name: "A simple pie chart"
        color: "red"
    }
}
