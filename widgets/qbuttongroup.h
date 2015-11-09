#ifdef __cplusplus
extern "C" {
#endif

void* QButtonGroup_NewQButtonGroup(void* parent);
void QButtonGroup_AddButton(void* ptr, void* button, int id);
void* QButtonGroup_Button(void* ptr, int id);
void* QButtonGroup_CheckedButton(void* ptr);
int QButtonGroup_CheckedId(void* ptr);
int QButtonGroup_Exclusive(void* ptr);
int QButtonGroup_Id(void* ptr, void* button);
void QButtonGroup_RemoveButton(void* ptr, void* button);
void QButtonGroup_SetExclusive(void* ptr, int v);
void QButtonGroup_SetId(void* ptr, void* button, int id);
void QButtonGroup_DestroyQButtonGroup(void* ptr);
void QButtonGroup_ConnectButtonClicked(void* ptr);
void QButtonGroup_DisconnectButtonClicked(void* ptr);
void QButtonGroup_ConnectButtonPressed(void* ptr);
void QButtonGroup_DisconnectButtonPressed(void* ptr);
void QButtonGroup_ConnectButtonReleased(void* ptr);
void QButtonGroup_DisconnectButtonReleased(void* ptr);
void QButtonGroup_ConnectButtonToggled(void* ptr);
void QButtonGroup_DisconnectButtonToggled(void* ptr);

#ifdef __cplusplus
}
#endif