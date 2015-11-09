#ifdef __cplusplus
extern "C" {
#endif

void* QMacToolBarItem_NewQMacToolBarItem(void* parent);
void QMacToolBarItem_ConnectActivated(void* ptr);
void QMacToolBarItem_DisconnectActivated(void* ptr);
void QMacToolBarItem_DestroyQMacToolBarItem(void* ptr);
int QMacToolBarItem_Selectable(void* ptr);
void QMacToolBarItem_SetIcon(void* ptr, void* icon);
void QMacToolBarItem_SetSelectable(void* ptr, int selectable);
void QMacToolBarItem_SetStandardItem(void* ptr, int standardItem);
void QMacToolBarItem_SetText(void* ptr, char* text);
int QMacToolBarItem_StandardItem(void* ptr);
char* QMacToolBarItem_Text(void* ptr);

#ifdef __cplusplus
}
#endif