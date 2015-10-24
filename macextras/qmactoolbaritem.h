#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMacToolBarItem_NewQMacToolBarItem(QtObjectPtr parent);
void QMacToolBarItem_ConnectActivated(QtObjectPtr ptr);
void QMacToolBarItem_DisconnectActivated(QtObjectPtr ptr);
void QMacToolBarItem_DestroyQMacToolBarItem(QtObjectPtr ptr);
int QMacToolBarItem_Selectable(QtObjectPtr ptr);
void QMacToolBarItem_SetIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QMacToolBarItem_SetSelectable(QtObjectPtr ptr, int selectable);
void QMacToolBarItem_SetStandardItem(QtObjectPtr ptr, int standardItem);
void QMacToolBarItem_SetText(QtObjectPtr ptr, char* text);
int QMacToolBarItem_StandardItem(QtObjectPtr ptr);
char* QMacToolBarItem_Text(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif