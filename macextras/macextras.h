#ifdef __cplusplus
extern "C" {
#endif

char* QMacPasteboardMime_ConvertorName(void* ptr);
int QMacPasteboardMime_Count(void* ptr, void* mimeData);
char* QMacPasteboardMime_FlavorFor(void* ptr, char* mime);
char* QMacPasteboardMime_MimeFor(void* ptr, char* flav);
void QMacPasteboardMime_DestroyQMacPasteboardMime(void* ptr);
char* QMacPasteboardMime_ObjectNameAbs(void* ptr);
void QMacPasteboardMime_SetObjectNameAbs(void* ptr, char* name);
void* QMacToolBar_NewQMacToolBar(void* parent);
void* QMacToolBar_NewQMacToolBar2(char* identifier, void* parent);
void* QMacToolBar_AddAllowedItem(void* ptr, void* icon, char* text);
void* QMacToolBar_AddItem(void* ptr, void* icon, char* text);
void QMacToolBar_AddSeparator(void* ptr);
void QMacToolBar_AttachToWindow(void* ptr, void* window);
void QMacToolBar_DetachFromWindow(void* ptr);
void QMacToolBar_DestroyQMacToolBar(void* ptr);
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