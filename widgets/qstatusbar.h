#ifdef __cplusplus
extern "C" {
#endif

int QStatusBar_IsSizeGripEnabled(void* ptr);
void QStatusBar_SetSizeGripEnabled(void* ptr, int v);
void* QStatusBar_NewQStatusBar(void* parent);
void QStatusBar_AddPermanentWidget(void* ptr, void* widget, int stretch);
void QStatusBar_AddWidget(void* ptr, void* widget, int stretch);
void QStatusBar_ClearMessage(void* ptr);
char* QStatusBar_CurrentMessage(void* ptr);
int QStatusBar_InsertPermanentWidget(void* ptr, int index, void* widget, int stretch);
int QStatusBar_InsertWidget(void* ptr, int index, void* widget, int stretch);
void QStatusBar_ConnectMessageChanged(void* ptr);
void QStatusBar_DisconnectMessageChanged(void* ptr);
void QStatusBar_RemoveWidget(void* ptr, void* widget);
void QStatusBar_ShowMessage(void* ptr, char* message, int timeout);
void QStatusBar_DestroyQStatusBar(void* ptr);

#ifdef __cplusplus
}
#endif