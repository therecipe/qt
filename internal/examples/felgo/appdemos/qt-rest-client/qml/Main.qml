import Felgo 3.0
import QtQuick 2.7
import QtQuick.Layouts 1.3
import QtCharts 2.2
import "."

App {
    id: app

    // model
    DataModel {
      id: dataModel
      dispatcher: logic

      onLoadJsonDataFailed: {
        app.errorMsg = error
        logic.clearData()
      }
      onWeatherDataChanged: if(weatherData && weatherData.weatherForCity !== "") app.errorMsg = ""
    }

    // business logic
    Logic {
      id: logic
    }

    // view
    property string errorMsg: ""
    readonly property string weatherServiceAppId: "d8ed259735b17a417d92789cd24abae6";

    NavigationStack {

        // we display the json data in a list
        Page {
            id: page
            title: qsTr("REST with Qt QML and Felgo")
            ColumnLayout {
                anchors.fill: parent    // Important: otherwise search bar not visible when there is no text in UI
                anchors.margins: app.dp(16)

                SearchBar {
                    id: weatherSearchBar
                    focus: true
                    Layout.fillWidth: true
                    placeHolderText: qsTr("Enter city name")
                    onAccepted: logic.loadJsonData()
                }

                AppText {
                    text: app.errorMsg
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    visible: app.errorMsg !== undefined
                }

                AppText {
                    text: qsTr("Weather for %1").arg(dataModel.weatherData.weatherForCity)
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    color: Theme.tintColor
                    font.family: Theme.boldFont.name
                    font.bold: true
                    font.weight: Font.Bold
                    visible: dataModel.weatherAvailable
                }

                AppText {
                    text: qsTr("From: %1 %2").arg(dataModel.weatherData.weatherDate).arg(dataModel.weatherFromCache ? " (Cached)" : "")
                    Layout.fillWidth: true
                    visible: dataModel.weatherAvailable
                }

                AppText {
                    text: qsTr("Temperature: %1°").arg(dataModel.weatherData.weatherTemp)
                    Layout.fillWidth: true
                    visible: dataModel.weatherAvailable
                }

                AppText {
                    text: qsTr("Condition: %1").arg(dataModel.weatherData.weatherCondition)
                    Layout.fillWidth: true
                    visible: dataModel.weatherAvailable
                }

                Image {
                    source: dataModel.weatherData.weatherIconUrl
                    visible: dataModel.weatherAvailable
                }
                Item {
                    Layout.fillHeight: true
                }
            }

        }
    }
}
