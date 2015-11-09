#ifdef __cplusplus
extern "C" {
#endif

int QCheckBox_IsTristate(void* ptr);
void QCheckBox_SetTristate(void* ptr, int y);
void* QCheckBox_NewQCheckBox(void* parent);
void* QCheckBox_NewQCheckBox2(char* text, void* parent);
int QCheckBox_CheckState(void* ptr);
void QCheckBox_SetCheckState(void* ptr, int state);
void QCheckBox_ConnectStateChanged(void* ptr);
void QCheckBox_DisconnectStateChanged(void* ptr);
void QCheckBox_DestroyQCheckBox(void* ptr);

#ifdef __cplusplus
}
#endif