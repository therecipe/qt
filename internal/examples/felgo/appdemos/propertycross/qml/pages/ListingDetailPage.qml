import QtQuick 2.0
import Felgo 3.0

FlickablePage {
  property var model: ({})

  title: qsTr("Property Details")

  readonly property bool isFavorite: dataModel.isFavorite(model)

  rightBarItem: IconButtonBarItem {
    icon: isFavorite ? IconType.heart : IconType.hearto
    onClicked: {
      logic.toggleFavorite(model)
    }
  }

  clip: true //clip content behind navigation bar when scrolling

  flickable.contentWidth: parent.width
  flickable.contentHeight: contentCol.height + contentPadding
  flickable.bottomMargin: contentPadding

  Column {
    id: contentCol
    y: contentPadding
    anchors.left: parent.left
    anchors.right: parent.right
    anchors.margins: contentPadding
    spacing: contentPadding

    AppText {
      text: model.price_formatted
      width: parent.width
      wrapMode: Text.WrapAtWordBoundaryOrAnywhere
      font.pixelSize: sp(24)
    }

    AppText {
      text: model.title
      width: parent.width
      wrapMode: Text.WrapAtWordBoundaryOrAnywhere
      font.pixelSize: sp(20)
    }

    AppImage {
      source: model.img_url
      width: parent.width
      height: model && width * model.img_height / model.img_width || 0
      anchors.horizontalCenter: parent.horizontalCenter
    }

    AppText {
      width: parent.width
      wrapMode: Text.WrapAtWordBoundaryOrAnywhere
      text: "%1 bed, %2 bathrooms".arg(model.bedroom_number).arg(model.bathroom_number)
    }

    AppText {
      width: parent.width
      wrapMode: Text.WrapAtWordBoundaryOrAnywhere
      text: model.summary
    }
  }
}

