#ifdef __cplusplus
extern "C" {
#endif

int QMenuBar_IsDefaultUp(void* ptr);
int QMenuBar_IsNativeMenuBar(void* ptr);
void QMenuBar_SetCornerWidget(void* ptr, void* widget, int corner);
void QMenuBar_SetDefaultUp(void* ptr, int v);
void QMenuBar_SetNativeMenuBar(void* ptr, int nativeMenuBar);
void* QMenuBar_NewQMenuBar(void* parent);
void* QMenuBar_ActionAt(void* ptr, void* pt);
void* QMenuBar_ActiveAction(void* ptr);
void* QMenuBar_AddAction(void* ptr, char* text);
void* QMenuBar_AddAction2(void* ptr, char* text, void* receiver, char* member);
void* QMenuBar_AddMenu(void* ptr, void* menu);
void* QMenuBar_AddMenu3(void* ptr, void* icon, char* title);
void* QMenuBar_AddMenu2(void* ptr, char* title);
void* QMenuBar_AddSeparator(void* ptr);
void QMenuBar_Clear(void* ptr);
void* QMenuBar_CornerWidget(void* ptr, int corner);
int QMenuBar_HeightForWidth(void* ptr, int v);
void QMenuBar_ConnectHovered(void* ptr);
void QMenuBar_DisconnectHovered(void* ptr);
void* QMenuBar_InsertMenu(void* ptr, void* before, void* menu);
void* QMenuBar_InsertSeparator(void* ptr, void* before);
void QMenuBar_SetActiveAction(void* ptr, void* act);
void QMenuBar_SetVisible(void* ptr, int visible);
void QMenuBar_ConnectTriggered(void* ptr);
void QMenuBar_DisconnectTriggered(void* ptr);
void QMenuBar_DestroyQMenuBar(void* ptr);

#ifdef __cplusplus
}
#endif