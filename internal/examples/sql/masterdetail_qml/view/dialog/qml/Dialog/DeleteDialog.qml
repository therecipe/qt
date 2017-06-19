import QtQuick.Dialogs 1.2  //MessageDialog

import Dialog 1.0

DeleteDialogController {

  MessageDialog {
    id: deleteDialogDismissed
    title: "Delete Album"
    text: "Select the album you want to delete."
    icon: StandardIcon.Information
  }

  MessageDialog {
    id: deleteDialog
    title: "Delete Album"
    text: ""
    icon: StandardIcon.Question
    standardButtons: StandardButton.Yes | StandardButton.No

    onYes: deleteAlbumCommand()
    onNo: deleteDialogDismissed.visible = true
  }

  onDeleteAlbumShowRequest: {
    deleteDialog.text = "Are you sure you want to delete '"+title+"' by '"+artist+"'?"
    deleteDialog.visible = true
  }
}
