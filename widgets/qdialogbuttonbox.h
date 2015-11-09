#ifdef __cplusplus
extern "C" {
#endif

int QDialogButtonBox_CenterButtons(void* ptr);
int QDialogButtonBox_Orientation(void* ptr);
void QDialogButtonBox_SetCenterButtons(void* ptr, int center);
void QDialogButtonBox_SetOrientation(void* ptr, int orientation);
void QDialogButtonBox_SetStandardButtons(void* ptr, int buttons);
int QDialogButtonBox_StandardButtons(void* ptr);
void* QDialogButtonBox_NewQDialogButtonBox(void* parent);
void* QDialogButtonBox_NewQDialogButtonBox2(int orientation, void* parent);
void* QDialogButtonBox_NewQDialogButtonBox3(int buttons, void* parent);
void* QDialogButtonBox_NewQDialogButtonBox4(int buttons, int orientation, void* parent);
void QDialogButtonBox_ConnectAccepted(void* ptr);
void QDialogButtonBox_DisconnectAccepted(void* ptr);
void* QDialogButtonBox_AddButton3(void* ptr, int button);
void* QDialogButtonBox_AddButton2(void* ptr, char* text, int role);
void QDialogButtonBox_AddButton(void* ptr, void* button, int role);
void* QDialogButtonBox_Button(void* ptr, int which);
int QDialogButtonBox_ButtonRole(void* ptr, void* button);
void QDialogButtonBox_Clear(void* ptr);
void QDialogButtonBox_ConnectClicked(void* ptr);
void QDialogButtonBox_DisconnectClicked(void* ptr);
void QDialogButtonBox_ConnectHelpRequested(void* ptr);
void QDialogButtonBox_DisconnectHelpRequested(void* ptr);
void QDialogButtonBox_ConnectRejected(void* ptr);
void QDialogButtonBox_DisconnectRejected(void* ptr);
void QDialogButtonBox_RemoveButton(void* ptr, void* button);
int QDialogButtonBox_StandardButton(void* ptr, void* button);
void QDialogButtonBox_DestroyQDialogButtonBox(void* ptr);

#ifdef __cplusplus
}
#endif