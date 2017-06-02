import QtQuick 2.4
import QtQuick.Layouts 1.1
import Material 0.2

ColumnLayout {
    spacing: 0

    Repeater {
        model: 2

        Rectangle {
            Layout.fillHeight: true
            Layout.fillWidth: true
            Layout.minimumHeight: grid.height + dp(80)
            Layout.minimumWidth: grid.width + dp(80)
            color: index == 0 ? "#EEE" : "#333"

            GridLayout {
                id: grid
                anchors.centerIn: parent
                rowSpacing: dp(20)
                columnSpacing: dp(10)
                columns: 2

                Label {
                    text: "Normal"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Slider {
                    Layout.alignment: Qt.AlignCenter
                    value: 0.2
                    darkBackground: index == 1
                }

                Label {
                    text: "Tickmarks"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Slider {
                    Layout.alignment: Qt.AlignCenter
                    value: 40
                    tickmarksEnabled: true
                    stepSize: 20
                    minimumValue: 0
                    maximumValue: 100
                    darkBackground: index == 1
                }

                Label {
                    text: "Numeric Value Label"
                    Layout.alignment:  Qt.AlignBottom
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Slider {
                    Layout.alignment: Qt.AlignCenter
                    value: 60
                    tickmarksEnabled: true
                    numericValueLabel: true
                    stepSize: 20
                    minimumValue: 0
                    maximumValue: 100
                    darkBackground: index == 1
                }

                Label {
                    text: "Numeric Value Label + Active Focus on Press"
                    wrapMode: Text.WordWrap
                    Layout.alignment:  Qt.AlignBottom
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Slider {
                    Layout.alignment: Qt.AlignCenter
                    value: 80
                    focus: true
                    tickmarksEnabled: true
                    numericValueLabel: true
                    stepSize: 20
                    minimumValue: 0
                    maximumValue: 100
                    activeFocusOnPress: true
                    darkBackground: index == 1
                }

                Label {
                    text: "Customized Numeric Value Label"
                    wrapMode: Text.WordWrap
                    Layout.alignment:  Qt.AlignBottom
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Slider {
                    Layout.alignment: Qt.AlignCenter
                    numericValueLabel: true
                    stepSize: 1
                    minimumValue: 0
                    maximumValue: 100
                    alwaysShowValueLabel: true
                    knobLabel: value + "%"
                    knobDiameter: dp(42)
                    darkBackground: index == 1
                }

                Label {
                    text: "Disabled"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Slider {
                    enabled: false
                    Layout.alignment: Qt.AlignCenter
                    darkBackground: index == 1
                }
            }
        }
    }
}
