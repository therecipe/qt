#ifdef __cplusplus
extern "C" {
#endif

int QGroupBox_Alignment(void* ptr);
int QGroupBox_IsCheckable(void* ptr);
int QGroupBox_IsChecked(void* ptr);
int QGroupBox_IsFlat(void* ptr);
void QGroupBox_SetAlignment(void* ptr, int alignment);
void QGroupBox_SetCheckable(void* ptr, int checkable);
void QGroupBox_SetChecked(void* ptr, int checked);
void QGroupBox_SetFlat(void* ptr, int flat);
void QGroupBox_SetTitle(void* ptr, char* title);
char* QGroupBox_Title(void* ptr);
void* QGroupBox_NewQGroupBox(void* parent);
void* QGroupBox_NewQGroupBox2(char* title, void* parent);
void QGroupBox_ConnectClicked(void* ptr);
void QGroupBox_DisconnectClicked(void* ptr);
void QGroupBox_ConnectToggled(void* ptr);
void QGroupBox_DisconnectToggled(void* ptr);
void QGroupBox_DestroyQGroupBox(void* ptr);

#ifdef __cplusplus
}
#endif