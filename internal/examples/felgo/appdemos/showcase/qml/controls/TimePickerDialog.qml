import Felgo 3.0
import QtQuick 2.5

Dialog {
  id: timePickerDialog
  contentHeight: dp(48) * 2.2
  negativeActionLabel: qsTr("Cancel")
  title: qsTr("Set Time")
  positiveActionLabel: qsTr("OK")
  onCanceled: { timePickerDialog.close(); time = initialTime }
  onAccepted: { timePickerDialog.close() }

  property var time: pickerFrom.time
  property var initialTime
  onIsOpenChanged: {
    if(isOpen) {
      time = Qt.binding(function(){ return pickerFrom.time })
      initialTime = pickerFrom.time
    }
  }

  TimePicker {
    id: pickerFrom
    anchors.horizontalCenter: parent.horizontalCenter
    width: parent.width - 2 * dp(Theme.navigationBar.defaultBarItemPadding)
    height: parent.height
    fontFamily: Theme.normalFont.name
  }

  function openWith(time) {
    pickerFrom.setTime(time)
    open()
  }
}
