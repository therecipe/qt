#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QActionGroup_AddAction(QtObjectPtr ptr, QtObjectPtr action);
int QActionGroup_IsEnabled(QtObjectPtr ptr);
int QActionGroup_IsExclusive(QtObjectPtr ptr);
int QActionGroup_IsVisible(QtObjectPtr ptr);
void QActionGroup_SetEnabled(QtObjectPtr ptr, int v);
void QActionGroup_SetExclusive(QtObjectPtr ptr, int v);
void QActionGroup_SetVisible(QtObjectPtr ptr, int v);
QtObjectPtr QActionGroup_NewQActionGroup(QtObjectPtr parent);
QtObjectPtr QActionGroup_AddAction3(QtObjectPtr ptr, QtObjectPtr icon, char* text);
QtObjectPtr QActionGroup_AddAction2(QtObjectPtr ptr, char* text);
QtObjectPtr QActionGroup_CheckedAction(QtObjectPtr ptr);
void QActionGroup_ConnectHovered(QtObjectPtr ptr);
void QActionGroup_DisconnectHovered(QtObjectPtr ptr);
void QActionGroup_RemoveAction(QtObjectPtr ptr, QtObjectPtr action);
void QActionGroup_SetDisabled(QtObjectPtr ptr, int b);
void QActionGroup_ConnectTriggered(QtObjectPtr ptr);
void QActionGroup_DisconnectTriggered(QtObjectPtr ptr);
void QActionGroup_DestroyQActionGroup(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif