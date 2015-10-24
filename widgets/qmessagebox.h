#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QMessageBox_ButtonMask_Type();
char* QMessageBox_DetailedText(QtObjectPtr ptr);
int QMessageBox_Icon(QtObjectPtr ptr);
char* QMessageBox_InformativeText(QtObjectPtr ptr);
void QMessageBox_SetDetailedText(QtObjectPtr ptr, char* text);
void QMessageBox_SetIcon(QtObjectPtr ptr, int v);
void QMessageBox_SetIconPixmap(QtObjectPtr ptr, QtObjectPtr pixmap);
void QMessageBox_SetInformativeText(QtObjectPtr ptr, char* text);
void QMessageBox_SetStandardButtons(QtObjectPtr ptr, int buttons);
void QMessageBox_SetText(QtObjectPtr ptr, char* text);
void QMessageBox_SetTextFormat(QtObjectPtr ptr, int format);
void QMessageBox_SetTextInteractionFlags(QtObjectPtr ptr, int flags);
int QMessageBox_StandardButtons(QtObjectPtr ptr);
char* QMessageBox_Text(QtObjectPtr ptr);
int QMessageBox_TextFormat(QtObjectPtr ptr);
int QMessageBox_TextInteractionFlags(QtObjectPtr ptr);
QtObjectPtr QMessageBox_NewQMessageBox2(int icon, char* title, char* text, int buttons, QtObjectPtr parent, int f);
QtObjectPtr QMessageBox_NewQMessageBox(QtObjectPtr parent);
void QMessageBox_QMessageBox_About(QtObjectPtr parent, char* title, char* text);
void QMessageBox_QMessageBox_AboutQt(QtObjectPtr parent, char* title);
QtObjectPtr QMessageBox_AddButton3(QtObjectPtr ptr, int button);
QtObjectPtr QMessageBox_AddButton2(QtObjectPtr ptr, char* text, int role);
void QMessageBox_AddButton(QtObjectPtr ptr, QtObjectPtr button, int role);
QtObjectPtr QMessageBox_Button(QtObjectPtr ptr, int which);
void QMessageBox_ConnectButtonClicked(QtObjectPtr ptr);
void QMessageBox_DisconnectButtonClicked(QtObjectPtr ptr);
int QMessageBox_ButtonRole(QtObjectPtr ptr, QtObjectPtr button);
QtObjectPtr QMessageBox_CheckBox(QtObjectPtr ptr);
QtObjectPtr QMessageBox_ClickedButton(QtObjectPtr ptr);
int QMessageBox_QMessageBox_Critical(QtObjectPtr parent, char* title, char* text, int buttons, int defaultButton);
QtObjectPtr QMessageBox_DefaultButton(QtObjectPtr ptr);
QtObjectPtr QMessageBox_EscapeButton(QtObjectPtr ptr);
int QMessageBox_Exec(QtObjectPtr ptr);
int QMessageBox_QMessageBox_Information(QtObjectPtr parent, char* title, char* text, int buttons, int defaultButton);
void QMessageBox_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member);
int QMessageBox_QMessageBox_Question(QtObjectPtr parent, char* title, char* text, int buttons, int defaultButton);
void QMessageBox_RemoveButton(QtObjectPtr ptr, QtObjectPtr button);
void QMessageBox_SetCheckBox(QtObjectPtr ptr, QtObjectPtr cb);
void QMessageBox_SetDefaultButton(QtObjectPtr ptr, QtObjectPtr button);
void QMessageBox_SetDefaultButton2(QtObjectPtr ptr, int button);
void QMessageBox_SetEscapeButton(QtObjectPtr ptr, QtObjectPtr button);
void QMessageBox_SetEscapeButton2(QtObjectPtr ptr, int button);
void QMessageBox_SetVisible(QtObjectPtr ptr, int visible);
void QMessageBox_SetWindowModality(QtObjectPtr ptr, int windowModality);
void QMessageBox_SetWindowTitle(QtObjectPtr ptr, char* title);
int QMessageBox_StandardButton(QtObjectPtr ptr, QtObjectPtr button);
int QMessageBox_QMessageBox_Warning(QtObjectPtr parent, char* title, char* text, int buttons, int defaultButton);
void QMessageBox_DestroyQMessageBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif