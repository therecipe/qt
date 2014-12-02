#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QDialogButtonBox_WinLayout();
int QDialogButtonBox_MacLayout();
int QDialogButtonBox_KdeLayout();
int QDialogButtonBox_GnomeLayout();
int QDialogButtonBox_InvalidRole();
int QDialogButtonBox_AcceptRole();
int QDialogButtonBox_RejectRole();
int QDialogButtonBox_DestructiveRole();
int QDialogButtonBox_ActionRole();
int QDialogButtonBox_HelpRole();
int QDialogButtonBox_YesRole();
int QDialogButtonBox_NoRole();
int QDialogButtonBox_ApplyRole();
int QDialogButtonBox_ResetRole();
int QDialogButtonBox_Ok();
int QDialogButtonBox_Open();
int QDialogButtonBox_Save();
int QDialogButtonBox_Cancel();
int QDialogButtonBox_Close();
int QDialogButtonBox_Discard();
int QDialogButtonBox_Apply();
int QDialogButtonBox_Reset();
int QDialogButtonBox_RestoreDefaults();
int QDialogButtonBox_Help();
int QDialogButtonBox_SaveAll();
int QDialogButtonBox_Yes();
int QDialogButtonBox_YesToAll();
int QDialogButtonBox_No();
int QDialogButtonBox_NoToAll();
int QDialogButtonBox_Abort();
int QDialogButtonBox_Retry();
int QDialogButtonBox_Ignore();
//Public Functions
QtObjectPtr QDialogButtonBox_New_QWidget(QtObjectPtr parent);
QtObjectPtr QDialogButtonBox_New_Orientation_QWidget(int orientation, QtObjectPtr parent);
QtObjectPtr QDialogButtonBox_New_StandardButton_QWidget(int buttons, QtObjectPtr parent);
QtObjectPtr QDialogButtonBox_New_StandardButton_Orientation_QWidget(int buttons, int orientation, QtObjectPtr parent);
void QDialogButtonBox_Destroy(QtObjectPtr ptr);
void QDialogButtonBox_AddButton_QAbstractButton_ButtonRole(QtObjectPtr ptr, QtObjectPtr button, int role);
QtObjectPtr QDialogButtonBox_AddButton_String_ButtonRole(QtObjectPtr ptr, char* text, int role);
QtObjectPtr QDialogButtonBox_AddButton_StandardButton(QtObjectPtr ptr, int button);
QtObjectPtr QDialogButtonBox_Button_StandardButton(QtObjectPtr ptr, int which);
int QDialogButtonBox_ButtonRole_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button);
int QDialogButtonBox_CenterButtons(QtObjectPtr ptr);
void QDialogButtonBox_Clear(QtObjectPtr ptr);
int QDialogButtonBox_Orientation(QtObjectPtr ptr);
void QDialogButtonBox_RemoveButton_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button);
void QDialogButtonBox_SetCenterButtons_Bool(QtObjectPtr ptr, int center);
void QDialogButtonBox_SetOrientation_Orientation(QtObjectPtr ptr, int orientation);
void QDialogButtonBox_SetStandardButtons_StandardButton(QtObjectPtr ptr, int buttons);
int QDialogButtonBox_StandardButton_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button);
int QDialogButtonBox_StandardButtons(QtObjectPtr ptr);
//Signals
void QDialogButtonBox_ConnectSignalAccepted(QtObjectPtr ptr);
void QDialogButtonBox_DisconnectSignalAccepted(QtObjectPtr ptr);
void QDialogButtonBox_ConnectSignalHelpRequested(QtObjectPtr ptr);
void QDialogButtonBox_DisconnectSignalHelpRequested(QtObjectPtr ptr);
void QDialogButtonBox_ConnectSignalRejected(QtObjectPtr ptr);
void QDialogButtonBox_DisconnectSignalRejected(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
