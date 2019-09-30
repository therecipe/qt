import QtQuick 2.0
import Felgo 3.0

SocialUserDelegate {
  height: isLocalUser ? localUserCol.height : otherUserCol.height

  // if the user did not set any customData yet, this is an empty string. this var makes sure it always is a valid object (an empty one if there is no string)
  // TODO: add this to GameNetworkUser
  property var userCustomData: gameNetworkUser.customData ? JSON.parse(gameNetworkUser.customData) : ({})
  property bool isLocalUser: gameNetworkUser.userId === gameNetworkItem.user.userId

  function updateCustomDataField(fieldName, newText) {
    console.debug("updateCustomDataField", fieldName, "with new text:", newText)
    // make a copy first, otherwise we would remove the binding from userCustomData
    var newUserData = JSON.parse(JSON.stringify(userCustomData))
    newUserData[fieldName]= newText
    // this updates the userCustomData property below!
    // we need to stringify, otherwise the result cannot be parsed
    // TODO: add a check to VPGN::updateUserCustomData that assures the type is not an object! currently there is this error: qrc:///qml/Felgo/gamenetwork/FelgoGameNetwork.qml:2407: Error: Cannot assign QJSValue to QString
    gameNetworkItem.updateUserCustomData(JSON.stringify(newUserData))
  }

  function resetAllCustomDataFields() {
    var newUserData = {}
    gameNetworkItem.updateUserCustomData(JSON.stringify(newUserData))
  }

  // Column for Local User
  Column {
    id: localUserCol
    x: dp(Theme.navigationBar.defaultBarItemPadding)
    width: parent.width - x
    spacing: dp(Theme.navigationBar.defaultBarItemPadding)
    // only show editable fields if the local user is the same like the gamenetwork user
    visible: isLocalUser

    Text {
      text: qsTr("Enter these optional fields so other conference attendees can connect with you. The more infos you provide, the better others can find & connect with you.\n\nWith the Business Meet feature, you can search & filter the conference attendees. Simply enter one of the fields in the search bar of the business meet view, like the company name(e.g. 'Felgo'), the job function (e.g. 'developer'), the interest (e.g. 'mobile') or the experience level ('Qt Veteran').\n\nThe fields you enter will be displayed in the business meet view and also in the leaderboard.")
      width: parent.width // consider margin right + left
      wrapMode: Text.WrapAtWordBoundaryOrAnywhere
      font.pixelSize: sp(13)
      font.family: socialViewItem.bodyFontName
      color: socialViewItem.bodyColor
    }

    Row {
      spacing: dp(5)

      Icon {
        icon: IconType.building
        size: dp(12)
        anchors.verticalCenter: parent.verticalCenter
        width: dp(20)
      }

      Text {
        text: qsTr("Company") + ": " + userCustomData["companyName"]
        font.pixelSize: sp(13)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        anchors.verticalCenter: parent.verticalCenter
      }
    }

    SocialViewButton {
      text: qsTr("Set Company Name")
      font.pixelSize: sp(13)
      onClicked: {
        NativeDialog.inputText(qsTr("What is your company name?"), "", qsTr("Your company name"), userCustomData["companyName"], function(ok, text) {
          if(ok) {
            updateCustomDataField("companyName", text)
            amplitude.logEvent("Update Profile",{"field" : "companyName", "value" : text})
          }
        })
      }
    }

    Row {
      spacing: dp(5)

      Icon {
        icon: IconType.suitcase
        size: dp(12)
        anchors.verticalCenter: parent.verticalCenter
        width: dp(20)
      }

      Text {
        text: qsTr("Position: ") + userCustomData["jobFunction"] || ""
        font.pixelSize: sp(13)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        anchors.verticalCenter: parent.verticalCenter
      }
    }
    SocialViewButton {
      text: qsTr("Set Job Function")
      font.pixelSize: sp(13)
      onClicked: {
        dialogStatus = "jobFunction"
        // on desktop, simulate a change of jobFunction between the first 2 options
        if(system.desktopPlatform) {
          var testIndex = 1
          if(userCustomData["jobFunction"] == "Manager") { // a === would not work here
            testIndex = 0
          }
          nativeUtils.alertSheetFinished(testIndex)
        } else {
          nativeUtils.displayAlertSheet("", ["Developer", "Manager", "Product Manager", "Designer", "Other"], true)
        }
      }
    }

    Row {
      spacing: dp(5)

      Icon {
        icon: IconType.tag
        size: dp(12)
        anchors.verticalCenter: parent.verticalCenter
        width: dp(20)
      }

      Text {
        text: qsTr("Main Qt Interest: ") + userCustomData["qtInterest"] || ""
        font.pixelSize: sp(13)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        anchors.verticalCenter: parent.verticalCenter
      }
    }
    SocialViewButton {
      text: qsTr("Set Qt Interest")
      font.pixelSize: sp(13)
      onClicked: {
        dialogStatus = "qtInterest"
        // on desktop, simulate a change between the first 2 options
        if(system.desktopPlatform) {
          var testIndex = 1
          if(userCustomData["qtInterest"] == "Mobile") { // a === would not work here
            testIndex = 0
          }
          nativeUtils.alertSheetFinished(testIndex)
        } else {
          nativeUtils.displayAlertSheet("What is Your Main Qt Interest?", ["Desktop", "Mobile", "Automotive", "Automation", "Embedded", "Other"], true)
        }
      }
    }

    Row {
      spacing: dp(5)

      Icon {
        icon: IconType.graduationcap
        size: dp(12)
        anchors.verticalCenter: parent.verticalCenter
        width: dp(20)
      }

      Text {
        text: qsTr("Qt XP Level: ") + userCustomData["qtExperience"] || ""
        font.pixelSize: sp(13)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        anchors.verticalCenter: parent.verticalCenter
      }
    }
    SocialViewButton {
      text: qsTr("Set Qt XP Level")
      font.pixelSize: sp(13)
      onClicked: {
        dialogStatus = "qtExperience"
        // on desktop, simulate a change between the first 2 options
        if(system.desktopPlatform) {
          var testIndex = 1
          if(userCustomData["qtExperience"] == "Qt Intermediate") { // a === would not work here
            testIndex = 0
          }
          nativeUtils.alertSheetFinished(testIndex)
        } else {
          // 3 levels: <1y Qt Newcomer, 1-3yrs Qt Intermediate, 3+yrs Qt Veteran
          nativeUtils.displayAlertSheet("How long did you use Qt?", ["Less than 1 year", "1-3 years", "More than 3 years"], true)
        }
      }
    }

    SocialViewButton {
      text: qsTr("Reset All Business Meet Fields")
      font.pixelSize: sp(13)
      onClicked: {
        NativeDialog.confirm(qsTr("Reset Fields"), qsTr("Do you really want to reset all fields for businees meetings? You will probably get less connections at the conference if you do so."), function(ok) {
          if(ok) {
            resetAllCustomDataFields()
            amplitude.logEvent("Reset Custom Fields")
          }
        })
      }
    }
  }

  // Column for Other Users, display custom data fields
  Column {
    id: otherUserCol
    visible: !isLocalUser
    x: dp(Theme.navigationBar.defaultBarItemPadding)
    width: parent.width - x
    spacing: dp(Theme.navigationBar.defaultBarItemPadding)

    Row {
      spacing: dp(5)
      visible: !system.publishBuild || userCustomData["companyName"] !== undefined && userCustomData["companyName"] != ""

      Icon {
        icon: IconType.building
        size: dp(12)
        anchors.verticalCenter: parent.verticalCenter
        width: dp(20)
      }

      Text {
        text: qsTr("Company: ") + userCustomData["companyName"] || ""
        font.pixelSize: sp(13)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        anchors.verticalCenter: parent.verticalCenter
      }
    }
    Row {
      spacing: dp(5)
      visible: !system.publishBuild || (userCustomData["jobFunction"] !== undefined && userCustomData["jobFunction"] != "")

      Icon {
        icon: IconType.suitcase
        size: dp(12)
        anchors.verticalCenter: parent.verticalCenter
        width: dp(20)
      }

      Text {
        text: qsTr("Position: ") + userCustomData["jobFunction"] || ""
        font.pixelSize: sp(13)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        anchors.verticalCenter: parent.verticalCenter
      }
    }
    Row {
      spacing: dp(5)
      visible: !system.publishBuild || (userCustomData["qtInterest"] !== undefined && userCustomData["qtInterest"] != "")

      Icon {
        icon: IconType.tag
        size: dp(12)
        anchors.verticalCenter: parent.verticalCenter
        width: dp(20)
      }

      Text {
        text: qsTr("Main Qt Interest: ") + userCustomData["qtInterest"] || ""
        font.pixelSize: sp(13)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        anchors.verticalCenter: parent.verticalCenter
      }
    }
    Row {
      spacing: dp(5)
      visible: !system.publishBuild || (userCustomData["qtExperience"] !== undefined && userCustomData["qtExperience"] != "")

      Icon {
        icon: IconType.graduationcap
        size: dp(12)
        anchors.verticalCenter: parent.verticalCenter
        width: dp(20)
      }

      Text {
        text: qsTr("Qt XP Level: ") + userCustomData["qtExperience"] || ""
        font.pixelSize: sp(13)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        anchors.verticalCenter: parent.verticalCenter
      }
    }
  }
}
