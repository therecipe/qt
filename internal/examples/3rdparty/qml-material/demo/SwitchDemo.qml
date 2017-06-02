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
                rowSpacing: dp(40)
                columnSpacing: dp(40)
                columns: 3

                // Empty filler
                Item { width: 1; height: 1 }

                Label {
                    text: "Normal"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Label {
                    text: "Disabled"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Label {
                    text: "On"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Switch {
                    checked: true
                    darkBackground: index == 1
                }

                Switch {
                    checked: true
                    enabled: false
                    darkBackground: index == 1
                }

                Label {
                    text: "Off"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Switch {
                    darkBackground: index == 1
                }

                Switch {
                    enabled: false
                    darkBackground: index == 1
                }
            }
        }
    }
}
