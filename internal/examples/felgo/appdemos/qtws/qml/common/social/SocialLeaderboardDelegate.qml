import Felgo 3.0
import QtQuick 2.2

SocialUserDelegate {
  id: scoreItem

  // indent
  x: dp(Theme.navigationBar.defaultBarItemPadding)
  width: parent ? parent.width - x : 0
  height: dp(48)

  // if this is a friendsOnly leaderboard, do not show the highlighted friends
  // the parent is the Leaderboard.qml file Repeater
  property bool higlightFriendPlayersWithColor: contextItem && contextItem.friendsOnly && socialViewItem.higlightFriendPlayersWithColorEnabled || false

  property bool isLoggedInPlayer: (!!gameNetworkUser && gameNetworkUser.userId === gameNetworkItem.user.userId) ? true : false // without the true/false in the end, the error log would be "Unable to assign [undefined] to bool" if the value is undefined
  property bool isFriendPlayer: (higlightFriendPlayersWithColor && !!gameNetworkUser && gameNetworkUser.isFriend) ? true : false

  property real defaultSpacing: (scoreItem.x / 3)

  // for accessing customData objects
  // returns an empty json object if no data string is known
  property var userCustomData: gameNetworkUser.customData ? JSON.parse(gameNetworkUser.customData) : ({})
  function getCustomDataString(fieldName) {
    return userCustomData[fieldName] ? userCustomData[fieldName] : ""
  }

  // this doesnt look good! there is an action at all players when they are clicked, so dont show i
  Rectangle {
    width: parent ? parent.width + defaultSpacing : 0
    anchors{
      right: parent ? parent.right : undefined
      top: parent ? parent.top : undefined
      bottom: parent ? parent.bottom : undefined
    }

    // set the background of the own player to the tintColor, and the background of friends to tintFriendPlayerHighlightColor
    color: getHighlightColor()
    // only needed to make visible if it is the logged in player - default color is white anyway
    visible: isLoggedInPlayer || isFriendPlayer
    // it looks better if it is opaque
    opacity: visible ? 0.3 : 1
  }

  function getHighlightColor() {
    if(isLoggedInPlayer)
      return socialViewItem.tintLightColor
    else if(isFriendPlayer) {
      //console.debug("isFriend, tintFriendPlayerHighlightColor:", tintFriendPlayerHighlightColor)
      return socialViewItem.tintFriendPlayerHighlightColor
    } else
      return "white"
  }

  // the text (the player name) will be above the country image, which is intended
  SocialFlagImage {
    id: flagImage
    locale: gameNetworkUser && gameNetworkUser.locale || ""
    anchors.verticalCenter: parent ? parent.verticalCenter : undefined
    anchors.right: parent ? parent.right : undefined
    anchors.rightMargin: dp(Theme.navigationBar.defaultBarItemPadding)
    height: parent.height * 0.4
    visible: socialViewItem.countryCodeEnabled // can be globally disabled in the GameNetworkView
  }

  Row {
    id: row
    spacing: dp(Theme.navigationBar.defaultBarItemPadding)
    height: dp(38)
    anchors.verticalCenter: parent ? parent.verticalCenter : undefined

    Column {
      id: rank
      anchors.verticalCenter: parent.verticalCenter
      spacing: -dp(5)

      Text {
        id: pos
        anchors.horizontalCenter: parent.horizontalCenter
        // pos is the index in the array, thus add 1 to it to be better understandable for non-techies as well ;)
        text: modelData ? modelData.pos + 1 : "-1"// ? modelData.pos : index + 1 // when pagination is supported, also add the page number to the index

        // default size is 15
        font.pixelSize: sp(24)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
      }
      Row {
        spacing: dp(2)
        anchors.horizontalCenter: parent.horizontalCenter

        Icon {
          icon: IconType.trophy
          color: socialViewItem.tintColor
          size: dp(6)
          anchors.verticalCenter: parent.verticalCenter
          anchors.verticalCenterOffset: 1
        }

        Text {
          id: score
          // pos is the index in the array, thus add 1 to it to be better understandable for non-techies as well ;)
          text: modelData ? Math.floor(modelData.value) : 0

          // default size is 15
          font.pixelSize: sp(9)
          font.family: socialViewItem.bodyFontName
          font.bold: true
          color: socialViewItem.tintColor
          anchors.verticalCenter: parent.verticalCenter
        }
      }
    }

    SocialUserImage {
      id: userImage
      placeholderImage: "\uf007" // user
      width: row.height
      height: row.height
      anchors.verticalCenter: parent.verticalCenter
      source: modelData && modelData.user && modelData.user.profile_picture || ""
      // the flag shall not be displayed here, as it is shown on the right side anyway
      //locale: modelData && modelData.user && modelData.user.locale || ""
    }

    // TODO: if for example an Arabic character is used that takes more height, the Column will be higher and the profile image and this whole delegate height is higher!
    // prevent this by setting a fixed height (and make the username with "..." if too long or just wrap anywhere)?
    Column {
      id: col
      y: defaultSpacing
      spacing: dp(2)
      anchors.verticalCenter: parent.verticalCenter

      Text {
        id: userName
        // use the remaining space for the name of the user
        // maybe write "Me" if this is the logged in player!? - but it is better if one sees the own user name, because normally it will be "Player 123" or just the number!
        // so make a combination of the two!?
        text: !!gameNetworkUser ? gameNetworkItem.getDisplayNameFromUserName(modelData.user.name, modelData.user) : ""// -> this was before the modification of appending a "Player" before the id: modelData ? ( isLoggedInPlayer ? "Me (" + modelData.user.name + ")" : modelData.user.name) : ""
        //text: modelData ? ( isLoggedInPlayer ? "Me" : modelData.user.name) : ""
        // old, shows your own name:
        //text: modelData ? modelData.user.name : ""
        font.pixelSize: sp(14)
        font.family: socialViewItem.bodyFontName
        color: socialViewItem.bodyColor
        width: scoreItem.width - row.spacing * 3 - rank.width - userImage.width - (flagImage.visible && flagImage.locale !== "" ?  flagImage.width + row.spacing : 0)
        elide: Text.ElideRight
        wrapMode: Text.WrapAnywhere
        maximumLineCount: 1
      }

      Flow {
        width: userName.width
        spacing: dp(6)

        Repeater {
          model: [
            {"text" : getCustomDataString("companyName"), "icon" : IconType.building},
            {"text" : getCustomDataString("jobFunction"), "icon" : IconType.suitcase},
            {"text" : getCustomDataString("qtInterest"), "icon" : IconType.tag},
            {"text" : getCustomDataString("qtExperience"), "icon" : IconType.graduationcap}
          ]

          Row {
            spacing: dp(2)
            height: detailIcon.height
            visible: modelData.text

            Icon {
              id: detailIcon
              icon: modelData.icon
              color: socialViewItem.bodyLightColor
              size: dp(6)
              textItem.width: width
              anchors.verticalCenter: parent.verticalCenter
            }

            Text {
              font.pixelSize: sp(8)
              text: modelData.text
              color: socialViewItem.bodyLightColor
              anchors.verticalCenter: parent.verticalCenter
            }
          }
        }
      }

      /*Row {
        spacing: 4
        Repeater {
          model: [
            {"text" : getCustomDataString("companyName"), "icon" : IconType.building}
          ]

          Row {
            spacing: 2
            // 0 score would also translate to false, but we want to show the score anyway, thus also show if index is 0
            visible: modelData.text

            Icon {
              icon: modelData.icon
              color: bodyLightColor
              size: 6
              anchors.verticalCenter: parent.verticalCenter
            }

            Text {
              font.pixelSize: 9
              elide: Text.ElideRight
              text: modelData.text
              color: bodyLightColor
            }
          }
        }
      }

      Row {
        spacing: 4
        Repeater {
          model: [
            {"text" : getCustomDataString("jobFunction"), "icon" : IconType.suitcase},
            {"text" : getCustomDataString("qtInterest"), "icon" : IconType.tag},
            {"text" : getCustomDataString("qtExperience"), "icon" : IconType.graduationcap}
          ]

          Row {
            spacing: 2
            // 0 score would also translate to false, but we want to show the score anyway, thus also show if index is 0
            visible: modelData.text

            Icon {
              icon: modelData.icon
              color: bodyLightColor
              size: 6
              anchors.verticalCenter: parent.verticalCenter
            }

            Text {
              font.pixelSize: 9
              elide: Text.ElideRight
              text: modelData.text
              color: bodyLightColor
            }
          }
        }
      }*/

      /*Text {
        id: score
        font.pixelSize: 10
        font.family: socialViewItem.bodyFontName
        color: bodyLightColor
        text: function() {
          var score = modelData ? Math.round(modelData.value) : 0
          // + ", Company: " + getCustomDataString("companyName") + ", Job: " + getCustomDataString("jobFunction") : ""

        }
      }*/
    }// column

  }// Row

  MouseArea {
    anchors.fill: parent
    onClicked: {
      parentPage.scoreSelected(gameNetworkUser, modelData)
    }
  }
}// end of Item

