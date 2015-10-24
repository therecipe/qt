#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDialogButtonBox_CenterButtons(QtObjectPtr ptr);
int QDialogButtonBox_Orientation(QtObjectPtr ptr);
void QDialogButtonBox_SetCenterButtons(QtObjectPtr ptr, int center);
void QDialogButtonBox_SetOrientation(QtObjectPtr ptr, int orientation);
void QDialogButtonBox_SetStandardButtons(QtObjectPtr ptr, int buttons);
int QDialogButtonBox_StandardButtons(QtObjectPtr ptr);
QtObjectPtr QDialogButtonBox_NewQDialogButtonBox(QtObjectPtr parent);
QtObjectPtr QDialogButtonBox_NewQDialogButtonBox2(int orientation, QtObjectPtr parent);
QtObjectPtr QDialogButtonBox_NewQDialogButtonBox3(int buttons, QtObjectPtr parent);
QtObjectPtr QDialogButtonBox_NewQDialogButtonBox4(int buttons, int orientation, QtObjectPtr parent);
void QDialogButtonBox_ConnectAccepted(QtObjectPtr ptr);
void QDialogButtonBox_DisconnectAccepted(QtObjectPtr ptr);
QtObjectPtr QDialogButtonBox_AddButton3(QtObjectPtr ptr, int button);
QtObjectPtr QDialogButtonBox_AddButton2(QtObjectPtr ptr, char* text, int role);
void QDialogButtonBox_AddButton(QtObjectPtr ptr, QtObjectPtr button, int role);
QtObjectPtr QDialogButtonBox_Button(QtObjectPtr ptr, int which);
int QDialogButtonBox_ButtonRole(QtObjectPtr ptr, QtObjectPtr button);
void QDialogButtonBox_Clear(QtObjectPtr ptr);
void QDialogButtonBox_ConnectClicked(QtObjectPtr ptr);
void QDialogButtonBox_DisconnectClicked(QtObjectPtr ptr);
void QDialogButtonBox_ConnectHelpRequested(QtObjectPtr ptr);
void QDialogButtonBox_DisconnectHelpRequested(QtObjectPtr ptr);
void QDialogButtonBox_ConnectRejected(QtObjectPtr ptr);
void QDialogButtonBox_DisconnectRejected(QtObjectPtr ptr);
void QDialogButtonBox_RemoveButton(QtObjectPtr ptr, QtObjectPtr button);
int QDialogButtonBox_StandardButton(QtObjectPtr ptr, QtObjectPtr button);
void QDialogButtonBox_DestroyQDialogButtonBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif