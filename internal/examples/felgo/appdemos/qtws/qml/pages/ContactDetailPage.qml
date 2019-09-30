import QtQuick 2.0
import QtQuick.Layouts 1.1
import Felgo 3.0
import "../common"

FlickablePage {
  id: page
  title: "Contact Info"

  property string contactId
  property var contact
  property var dataTextSize: sp(12)

  flickable.contentWidth: width
  flickable.contentHeight: contentCol.height

  GridLayout {
    id: contentCol
    x: dp(Theme.navigationBar.defaultBarItemPadding)
    Layout.preferredWidth: implicitWidth
    Layout.maximumWidth: parent.width - dp(Theme.navigationBar.defaultBarItemPadding)
    columns: 2
    rowSpacing: columnSpacing * 0.5
    columnSpacing: dp(Theme.navigationBar.defaultBarItemPadding)
    anchors.horizontalCenter: parent.horizontalCenter

    Item {
      Layout.columnSpan: 2
      Layout.fillWidth: true
      height: parent.rowSpacing
    }

    AppText {
      text: "Name:"
      visible: appTextName.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextName
      Layout.fillWidth: true
      text: contact && contact.name ? contact.name : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== ""
      font.pixelSize: page.dataTextSize
    }

    AppText {
      text: "Email:"
      visible: appTextEmail.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextEmail
      Layout.fillWidth: true
      text: contact && contact.email ? contact.email : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== ""
      font.pixelSize: page.dataTextSize
    }

    AppText {
      text: "Phone:"
      visible: appTextPhone.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextPhone
      Layout.fillWidth: true
      text: contact && contact.work_phone ? contact.work_phone : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== "" && text !== "0"
      font.pixelSize: page.dataTextSize
    }

    AppText {
      text: "Company:"
      visible: appTextCompany.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextCompany
      Layout.fillWidth: true
      text: contact && contact.company ? contact.company : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== ""
      font.pixelSize: page.dataTextSize
    }

    AppText {
      text: "Position:"
      visible: appTextJobTitle.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextJobTitle
      Layout.fillWidth: true
      text: contact && contact.job_title ? contact.job_title : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== ""
      font.pixelSize: page.dataTextSize
    }

    AppText {
      text: "Address:"
      visible: appTextAddress.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextAddress
      Layout.fillWidth: true
      text: contact && contact.address ? contact.address : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== ""
      font.pixelSize: page.dataTextSize
    }

    AppText {
      text: "City:"
      visible: appTextCity.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextCity
      Layout.fillWidth: true
      text: contact && contact.city ? contact.city : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== ""
      font.pixelSize: page.dataTextSize
    }

    AppText {
      text: "Province:"
      visible: appTextProvince.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextProvince
      Layout.fillWidth: true
      text: contact && contact.province ? contact.province : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== ""
      font.pixelSize: page.dataTextSize
    }

    AppText {
      text: "Country:"
      visible: appTextCountry.visible
      font.pixelSize: page.dataTextSize
      font.bold: true
      font.weight: Font.Bold
      font.family: Theme.boldFont.name
    }

    AppText {
      id: appTextCountry
      Layout.fillWidth: true
      text: contact && contact.country ? contact.country : ""
      wrapMode: AppText.WrapAtWordBoundaryOrAnywhere
      visible: text !== ""
      font.pixelSize: page.dataTextSize
    }

    Item {
      Layout.fillWidth: true
      Layout.columnSpan: 2
      width: parent.width
      height: parent.rowSpacing
    }

    AppButton {
      text: "Store Contact"
      Layout.columnSpan: 2
      Layout.fillWidth: true
      flat: false
      anchors.horizontalCenter: parent.horizontalCenter
      onClicked: storeContact()
    }

    AppButton {
      text: "Remove from Contact List"
      Layout.columnSpan: 2
      Layout.fillWidth: true
      flat: true
      anchors.horizontalCenter: parent.horizontalCenter
      onClicked: NativeDialog.confirm("Confirm Removal", "Are you sure want to remove this contact?", function(accept) {
        if(accept) {
          logic.removeContact(page.contactId)
          navigationStack.pop()
        }
      })
      verticalMargin: 0
    }

  }

  function storeContact() {
    // get company data
    var company = contact && contact.company ? contact.company : ""
    var adress = contact && contact.address ? contact.address : ""
    var city = contact && contact.city ? contact.city : ""
    var province = contact && contact.province ? contact.province : ""
    var country = contact && contact.country ? contact.country : ""

    // get contact data
    var name = contact && contact.name ? contact.name : ""
    var phone = contact && contact.work_phone ? contact.work_phone : ""
    var email = contact && contact.email ? contact.email : ""
    var job = contact && contact.job_title ? contact.job_title : ""

    // build vCard
    var vCard = "BEGIN:VCARD
VERSION:2.1\n";

    // add plain or encoded name to vCard (encode if non-latin chars are present)
    if(vCardEncodingRequired(name)) {
      name = encodeQuotedPrintable(name)
      vCard += "N;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:"+name+";;;;
FN;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:"+name+"\n";
    } else {
      vCard += "N:"+name+";;;;
FN:"+name+"\n";
    }

    // add company
    if(vCardEncodingRequired(company))
      vCard += "ORG;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:"+encodeQuotedPrintable(company)+"\n";
    else
      vCard += "ORG:"+company+"\n";

    // add job title
    if(vCardEncodingRequired(company))
      vCard += "TITLE;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:"+encodeQuotedPrintable(job)+"\n";
    else
      vCard += "TITLE:"+job+"\n";

    // add phone and email to vCard
    if(phone != "")
      vCard += "TEL;WORK:"+phone+"\n"
    if(email != "")
      vCard += "EMAIL;WORK:"+email+"\n"

    // add adress to vCard
    if(adress != "" || city != "" || province != "" || country != "") {
      // add plain or encoded adress (encode if non-latin chars are present)
      if(vCardEncodingRequired(adress) || vCardEncodingRequired(city) || vCardEncodingRequired(province) || vCardEncodingRequired(country)) {
        adress = encodeQuotedPrintable(adress)
        city = encodeQuotedPrintable(city)
        province = encodeQuotedPrintable(province)
        country = encodeQuotedPrintable(country)
        vCard += "ADR;WORK;CHARSET=UTF-8;ENCODING=QUOTED-PRINTABLE:;;"+adress+";"+city+";"+province+";;"+country+"\n";
      }
      else
        vCard += "ADR;WORK:;;"+adress+";"+city+";"+province+";;"+country+"\n";
    }

    vCard += "END:VCARD";

    // store vCard
    var success = nativeUtils.storeContacts(vCard)
    if(system.isPlatform(System.IOS)) {
      // handle success/error on iOS give feedback to user
      if(success)
        NativeDialog.confirm("Contact stored.", "", function() {}, false)
      else
        NativeDialog.confirm("Could not store contact",
                             "The app does not have permission to access the AddressBook, please allow access in the device settings.",
                             function() {}, false)
    }
    else if (system.isPlatform(System.Android)) {
      // only react to error on Android, if successful the device will open the vcard data file
      if(!success) {
        NativeDialog.confirm("Could not store contact",
                             "Error when trying to store the vCard to the user documents folder.",
                             function() {}, false)
      }
    }
    else {
      // platform not supported
      NativeDialog.confirm("Could not store contact", "Storing contacts is only possible on iOS and Android.", function() {}, false)
    }
  }

  // function for encoding string in quoted-printable vor vCard
  function encodeQuotedPrintable(text) {
    var charsPerLine = 69 // output lines in quoted printable hold 70 chars max
    var result = ""
    var curLineLength = 0

    // convert string to utf-8
    var utf8 = unescape(encodeURIComponent(text));

    // convert all chars of line to quoted printable hex representation
    for(var i = 0; i < utf8.length; i++) {
      var charCode = utf8.charCodeAt(i)
      var hexChar =  charCode.toString(16).toUpperCase() // hex value of character
      if(hexChar.length == 1)
        hexChar = "0"+hexChar

      if((curLineLength + hexChar.length + 1) > charsPerLine) {
        curLineLength = 0
        result += "=\n" // invisible line break in vCard
      }

      curLineLength += (hexChar.length + 1)
      result += ("=" + hexChar)
    }

    return result
  }

  // function to check whether a string must be encoded for vCard format
  function vCardEncodingRequired(str) {
    var regExNonLatin = /[^\u0000-\u007f]/;
    // encoding is required if non lating string or multi-line text
    return regExNonLatin.test(str) || str.indexOf("\n") > 0
  }
}
