import QtQuick 2.0
import Felgo 3.0

SocialUserDelegate {
  id: cell
  height: userImageWithDoubleColumnText.height + dp(16)//dp(54) // 48dp is enought so it is clickable, make a bit bigger though

  property bool seperator: true
  property var defaultAction

  property alias actionButton: actionButton
  property alias text: userImageWithDoubleColumnText.text
  property alias subtext: userImageWithDoubleColumnText.subtext
  property alias userImageSource: userImageWithDoubleColumnText.userImageSource
  property alias locale: userImageWithDoubleColumnText.locale

  // for accessing customData objects
  // returns an empty json object if no data string is known
  property var userCustomData: modelData.customData ? JSON.parse(modelData.customData) : ({})
  function getCustomDataString(fieldName) {
    return userCustomData[fieldName] ? userCustomData[fieldName] : ""
  }

  Rectangle {
    anchors.fill: parent
    color: "#ffffffff"
  }

  UserImageWithDoubleColumnText {
    id: userImageWithDoubleColumnText
    text: modelData.text
    details: {
      var detail = []
      if(getCustomDataString("companyName") !== "")
        detail.push({"text" : getCustomDataString("companyName"), "icon" : IconType.building})
      if(getCustomDataString("jobFunction") !== "")
        detail.push({"text" : getCustomDataString("jobFunction"), "icon" : IconType.suitcase})
      if(getCustomDataString("qtInterest") !== "")
        detail.push({"text" : getCustomDataString("qtInterest"), "icon" : IconType.tag})
      if(getCustomDataString("qtExperience") !== "")
        detail.push({"text" : getCustomDataString("qtExperience"), "icon" : IconType.graduationcap})
      return detail
    }

    subtext: modelData['subText'] !== undefined ? modelData.subText: ""
    userImageSource: modelData && modelData.profile_picture || ""
    locale: modelData && modelData.locale || ""
    //height: cell.height
    anchors.left: parent.left
    anchors.right: actionButton.action !== undefined ? actionButton.left : parent.right
    anchors.leftMargin: x // padding left is defined in x pos of UserImageWithDoubleColumnText
    anchors.rightMargin: x // padding right is same as left margin
    anchors.verticalCenter: parent.verticalCenter
  }

  MouseArea {
    id: mouseArea
    anchors.fill: parent
    hoverEnabled: true

    onClicked: {
       parentPage.userSelected(gameNetworkUser, modelData) // trigger userSelected signal of userSearchPage
    }

    onPressed: {
      cell.opacity = 0.75
    }
    onReleased: {
      // we don't do an if check here because the selectable property might have changed after a mousedown but before a mousup
      // and in that case we still want the button to return it's highlight state to normal
      cell.opacity = 1
    }
    onCanceled: {
      // onCanceled is called if the element is dragged inside a ListView, which is the case e.g. in the FriendsListView
      cell.opacity = 1
    }
  }

  SocialActionButton{
    id: actionButton
    anchors.right: cell.right
    textSize: sp(15)
    anchors.rightMargin: dp(16)
    anchors.verticalCenter: cell.verticalCenter

    property var action: modelData.action === undefined ? defaultAction : modelData.action

    onClicked: {
      if(actionButton.action.text===qsTr("Accept")) {
        actionButton.enabled = false
        gameNetworkItem.sendFriendResponse(gameNetworkUser.userId, function(success) {
          actionButton.visible = false
          if(!!parentPage)
            parentPage.searchUsers() // trigger searchUsers function of page that uses this delegate
        })
      }
    }

    onActionChanged: {
      if(action !== undefined){
        visible = action
        text = action.text
        backgroundColor = action['backgroundColor'] ? action.backgroundColor : socialViewItem.tintLightColor
      }
    }
  }

  Rectangle{
    visible: seperator
    height: 1 // make this 1dp instead?
    width: parent.width
    anchors.bottom: parent.bottom
    color: socialViewItem.separatorColor
  }
}
