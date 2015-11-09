#ifdef __cplusplus
extern "C" {
#endif

int QMessageBox_ButtonMask_Type();
char* QMessageBox_DetailedText(void* ptr);
int QMessageBox_Icon(void* ptr);
char* QMessageBox_InformativeText(void* ptr);
void QMessageBox_SetDetailedText(void* ptr, char* text);
void QMessageBox_SetIcon(void* ptr, int v);
void QMessageBox_SetIconPixmap(void* ptr, void* pixmap);
void QMessageBox_SetInformativeText(void* ptr, char* text);
void QMessageBox_SetStandardButtons(void* ptr, int buttons);
void QMessageBox_SetText(void* ptr, char* text);
void QMessageBox_SetTextFormat(void* ptr, int format);
void QMessageBox_SetTextInteractionFlags(void* ptr, int flags);
int QMessageBox_StandardButtons(void* ptr);
char* QMessageBox_Text(void* ptr);
int QMessageBox_TextFormat(void* ptr);
int QMessageBox_TextInteractionFlags(void* ptr);
void* QMessageBox_NewQMessageBox2(int icon, char* title, char* text, int buttons, void* parent, int f);
void* QMessageBox_NewQMessageBox(void* parent);
void QMessageBox_QMessageBox_About(void* parent, char* title, char* text);
void QMessageBox_QMessageBox_AboutQt(void* parent, char* title);
void* QMessageBox_AddButton3(void* ptr, int button);
void* QMessageBox_AddButton2(void* ptr, char* text, int role);
void QMessageBox_AddButton(void* ptr, void* button, int role);
void* QMessageBox_Button(void* ptr, int which);
void QMessageBox_ConnectButtonClicked(void* ptr);
void QMessageBox_DisconnectButtonClicked(void* ptr);
int QMessageBox_ButtonRole(void* ptr, void* button);
void* QMessageBox_CheckBox(void* ptr);
void* QMessageBox_ClickedButton(void* ptr);
int QMessageBox_QMessageBox_Critical(void* parent, char* title, char* text, int buttons, int defaultButton);
void* QMessageBox_DefaultButton(void* ptr);
void* QMessageBox_EscapeButton(void* ptr);
int QMessageBox_Exec(void* ptr);
int QMessageBox_QMessageBox_Information(void* parent, char* title, char* text, int buttons, int defaultButton);
void QMessageBox_Open(void* ptr, void* receiver, char* member);
int QMessageBox_QMessageBox_Question(void* parent, char* title, char* text, int buttons, int defaultButton);
void QMessageBox_RemoveButton(void* ptr, void* button);
void QMessageBox_SetCheckBox(void* ptr, void* cb);
void QMessageBox_SetDefaultButton(void* ptr, void* button);
void QMessageBox_SetDefaultButton2(void* ptr, int button);
void QMessageBox_SetEscapeButton(void* ptr, void* button);
void QMessageBox_SetEscapeButton2(void* ptr, int button);
void QMessageBox_SetVisible(void* ptr, int visible);
void QMessageBox_SetWindowModality(void* ptr, int windowModality);
void QMessageBox_SetWindowTitle(void* ptr, char* title);
int QMessageBox_StandardButton(void* ptr, void* button);
int QMessageBox_QMessageBox_Warning(void* parent, char* title, char* text, int buttons, int defaultButton);
void QMessageBox_DestroyQMessageBox(void* ptr);

#ifdef __cplusplus
}
#endif