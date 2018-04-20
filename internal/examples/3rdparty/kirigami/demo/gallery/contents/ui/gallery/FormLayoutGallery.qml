

import QtQuick 2.6
import QtQuick.Layouts 1.2
import QtQuick.Controls 2.2
import org.kde.kirigami 2.4 as Kirigami

Kirigami.ScrollablePage {
    title: "Form Layout"

    Kirigami.FormLayout {
        id: layout

        TextField {
            Kirigami.FormData.label: "Label:"
        }
        TextField {
        }
        TextField {
            Kirigami.FormData.label:"Lo&nger label:"
        }
        Kirigami.Separator {
            Kirigami.FormData.isSection: true
        }
        TextField {
            Kirigami.FormData.label: "After separator:"
        }
        ComboBox {
            Kirigami.FormData.label: "Combo:"
            model: ["First", "Second", "Third"]
        }
        CheckBox {
            checked: true
            text: "Option"
        }
        Kirigami.Separator {
            Kirigami.FormData.isSection: true
            Kirigami.FormData.label: "Section title"
        }
        TextField {
            Kirigami.FormData.label: "Label:"
        }
        Item {
            width:1
            height:1
            Kirigami.FormData.isSection: true
        }
        TextField {
            Kirigami.FormData.label: "Section without line:"
        }
        TextField {
        }
        Item {
            width:1
            height:1
            Kirigami.FormData.isSection: true
            Kirigami.FormData.label: "Section with title without line"
        }
        TextField {
            Kirigami.FormData.label: "Title:"
        }
        TextField {
            Kirigami.FormData.label: "Checkable label"
            Kirigami.FormData.checkable: true
            enabled: Kirigami.FormData.checked
        }
        ColumnLayout {
            Layout.rowSpan: 3
            Kirigami.FormData.label: "Label for radios:"
            RadioButton {
                checked: true
                text: "One"
            }
            RadioButton {
                text: "Two"
            }
            RadioButton {
                text: "Three"
            }
        }
        Button {
            text: item ? "Remove Field" : "Add Field"
            property TextField item
            onClicked: {
                if (item) {
                    item.destroy();
                } else {
                    item = dyncomponent.createObject(layout);
                }
            }
            Component {
                id: dyncomponent
                TextField {
                    Kirigami.FormData.label: "Generated Title:"
                }
            }
        }
    }
}
