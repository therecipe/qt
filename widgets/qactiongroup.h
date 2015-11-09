#ifdef __cplusplus
extern "C" {
#endif

void* QActionGroup_AddAction(void* ptr, void* action);
int QActionGroup_IsEnabled(void* ptr);
int QActionGroup_IsExclusive(void* ptr);
int QActionGroup_IsVisible(void* ptr);
void QActionGroup_SetEnabled(void* ptr, int v);
void QActionGroup_SetExclusive(void* ptr, int v);
void QActionGroup_SetVisible(void* ptr, int v);
void* QActionGroup_NewQActionGroup(void* parent);
void* QActionGroup_AddAction3(void* ptr, void* icon, char* text);
void* QActionGroup_AddAction2(void* ptr, char* text);
void* QActionGroup_CheckedAction(void* ptr);
void QActionGroup_ConnectHovered(void* ptr);
void QActionGroup_DisconnectHovered(void* ptr);
void QActionGroup_RemoveAction(void* ptr, void* action);
void QActionGroup_SetDisabled(void* ptr, int b);
void QActionGroup_ConnectTriggered(void* ptr);
void QActionGroup_DisconnectTriggered(void* ptr);
void QActionGroup_DestroyQActionGroup(void* ptr);

#ifdef __cplusplus
}
#endif