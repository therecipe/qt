import Felgo 3.0
import QtQuick 2.0
import QtMultimedia 5.5
import QZXing 2.3 // barcode scanning
import "../common"

ListPage {
  id: page
  title: "QR Contacts"
  rightBarItem: IconButtonBarItem {
    icon: IconType.envelope
    visible: page.listView.count > 0
    onClicked: sendContacts()
  }

  model: JsonListModel {
    id: listModel
    source: prepareContactsModel(dataModel.contacts)
  }

  delegate: SimpleRow {
   property var modelData: listModel.get(index)
   text: modelData.name
   detailText: modelData.pretty_scan_timestamp
   onSelected: page.navigationStack.push(Qt.resolvedUrl("ContactDetailPage.qml"),{ contactId: modelData["scan_tag"], contact: modelData })
  }

  emptyText.wrapMode: Text.WordWrap
  emptyText.font.pixelSize: sp(16) // default is 24
  emptyText.text: qsTr("No contacts added yet.\n\n1.) Scan the conference badge of contacts you meet to auto-fill the email, name, phone number and company infos.\n\n2.) Export all contacts as .csv file and import in your CRM. Or store the contacts on your phone.")
  listView.topMargin: scanHeader.height

  Rectangle {
    id: scanHeader
    width: parent.width
    height: scanButton.height + dp(Theme.navigationBar.defaultBarItemPadding)
    color: app.secondaryTintColor

    Row {
      anchors.bottom: parent.bottom
      anchors.horizontalCenter: parent.horizontalCenter
      anchors.verticalCenter: parent.verticalCenter

      IconButton {
        icon: IconType.qrcode
        anchors.verticalCenter: parent.verticalCenter
      }

      AppButton {
        id: scanButton
        text: qsTr("Scan QR Code to Add Contact")
        anchors.verticalCenter: parent.verticalCenter
        flat: true
      }
    }

    MouseArea {
      anchors.fill: parent
      onClicked: {
        amplitude.logEvent("Click Scan Card")
        page.navigationStack.push(contactScanPage)
      }
    }
  }

  // Component for Contact Scan page
  Component {
    id: contactScanPage
    ContactScanPage {

    }
  }

  // prepareContactsModel - sort contacts by scan time
  function prepareContactsModel(contacts) {

    if(dataModel.contacts && dataModel.contacts.length > 0)
      logic.clearContacts()

    if(!contacts)
      return []

    var sortedContacts = []
    for(var tag in contacts) {
      var dateObj = new Date()
      dateObj.setTime(contacts[tag]["scan_timestamp"])
      contacts[tag]["pretty_scan_timestamp"] = dateObj.toDateString() + ", " + dateObj.toTimeString()

      sortedContacts.push(contacts[tag])
    }

    // sort contacts
    sortedContacts = sortedContacts.sort(function(a, b) {
        return (a.scan_timestamp < b.scan_timestamp) - (a.scan_timestamp > b.scan_timestamp)
    })

    return sortedContacts
  }

  // sencContacts - create string CSV of contacts and send via email
  function sendContacts() {
    if(!dataModel.contacts)
      return

    var csv = ""
    for(var tag in dataModel.contacts) {
      var contact = dataModel.contacts[tag]
      var colSeparator = ";"
      var rowSeparator = "\n"

      // create headline before first contact
      if(csv == "") {
        csv += "First Name"+colSeparator+"Last Name"+colSeparator+"Name"+colSeparator+"Email"+colSeparator+"Phone"+colSeparator+
            "Company"+colSeparator+"Position"+colSeparator+"Address"+colSeparator+"City"+colSeparator+"Province"+colSeparator+"Country"+rowSeparator
      }

      // add contact
      csv += contact.first_name+colSeparator
          +contact.last_name+colSeparator
          +contact.name+colSeparator
          +contact.email+colSeparator
          +contact.work_phone+colSeparator
          +contact.company+colSeparator
          +contact.job_title+colSeparator
          +contact.address+colSeparator
          +contact.city+colSeparator
          +contact.province+colSeparator
          +contact.country+rowSeparator
    }

    nativeUtils.sendEmail("", "Qt World Summit 2017 - Contacts", "Your QtWS 2017 Contacts: \n\n"+csv)

    amplitude.logEvent("Send Scanned Contacts")
  }
}
