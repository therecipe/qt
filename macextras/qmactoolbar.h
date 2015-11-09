#ifdef __cplusplus
extern "C" {
#endif

void* QMacToolBar_NewQMacToolBar(void* parent);
void* QMacToolBar_NewQMacToolBar2(char* identifier, void* parent);
void* QMacToolBar_AddAllowedItem(void* ptr, void* icon, char* text);
void* QMacToolBar_AddItem(void* ptr, void* icon, char* text);
void QMacToolBar_AddSeparator(void* ptr);
void QMacToolBar_AttachToWindow(void* ptr, void* window);
void QMacToolBar_DetachFromWindow(void* ptr);
void QMacToolBar_DestroyQMacToolBar(void* ptr);

#ifdef __cplusplus
}
#endif