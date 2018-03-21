import QtQuick 2.7  //Item

import Theme 1.0    //Theme
import View 1.0     //StackItemHeader

StackItemHeader {
  image: "qrc:/qml/assets/ic_insert_drive_file_black_24px.svg"
  title: "Files"

   contentItem: Item {
    anchors.fill: parent

    HeaderButton {
      anchors {
        top: parent.top
        right: send.left
        bottom: parent.bottom
      }
      width: Theme.minimumWidth * 0.1

      text: "Create Allowance"
      code: "create"
      image: "qrc:/qml/assets/ic_create_black_24px.svg"
    }

    HeaderButton {
      id: send

      anchors {
        top: parent.top
        right: recover.left
        bottom: parent.bottom
      }
      width: Theme.minimumWidth * 0.1

      text: "Upload Files"
      code: "files"
      image: "qrc:/qml/assets/ic_cloud_upload_black_24px.svg"
    }

    HeaderButton {
      id: recover

      anchors {
        top: parent.top
        right: parent.right
        bottom: parent.bottom
      }
      width: Theme.minimumWidth * 0.1

      text: "Upload Folder"
      code: "folder"
      image: "qrc:/qml/assets/ic_create_new_folder_black_24px.svg"
    }
  }
}
