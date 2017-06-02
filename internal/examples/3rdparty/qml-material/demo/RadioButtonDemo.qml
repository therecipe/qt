import QtQuick 2.4
import QtQuick.Layouts 1.1
import QtQuick.Controls 1.3 as QuickControls
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
                columnSpacing: dp(20)
                columns: 2

                QuickControls.ExclusiveGroup { id: optionGroup }

                Label {
                    Layout.alignment : Qt.AlignHCenter
                    text: "Normal"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                Label {
                    Layout.alignment : Qt.AlignHCenter
                    text: "Disabled"
                    color: index == 0 ? Theme.light.textColor : Theme.dark.textColor
                }

                RadioButton {
                    checked: true
                    text: "Option 1"
                    darkBackground: index == 1
                    canToggle: true
                    exclusiveGroup: optionGroup
                }

                RadioButton {
                    checked: true
                    enabled: false
                    text: "Disabled"
                    darkBackground: index == 1
                }

                RadioButton {
                    text: "Option 2"
                    darkBackground: index == 1
                    canToggle: true
                    exclusiveGroup: optionGroup
                }

                RadioButton {
                    enabled: false
                    text: "Disabled"
                    darkBackground: index == 1
                }
            }
        }
    }
}
