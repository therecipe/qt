#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCompleter_CaseSensitivity(QtObjectPtr ptr);
int QCompleter_CompletionColumn(QtObjectPtr ptr);
int QCompleter_CompletionMode(QtObjectPtr ptr);
char* QCompleter_CompletionPrefix(QtObjectPtr ptr);
int QCompleter_CompletionRole(QtObjectPtr ptr);
int QCompleter_FilterMode(QtObjectPtr ptr);
int QCompleter_MaxVisibleItems(QtObjectPtr ptr);
int QCompleter_ModelSorting(QtObjectPtr ptr);
void QCompleter_SetCaseSensitivity(QtObjectPtr ptr, int caseSensitivity);
void QCompleter_SetCompletionColumn(QtObjectPtr ptr, int column);
void QCompleter_SetCompletionMode(QtObjectPtr ptr, int mode);
void QCompleter_SetCompletionPrefix(QtObjectPtr ptr, char* prefix);
void QCompleter_SetCompletionRole(QtObjectPtr ptr, int role);
void QCompleter_SetFilterMode(QtObjectPtr ptr, int filterMode);
void QCompleter_SetMaxVisibleItems(QtObjectPtr ptr, int maxItems);
void QCompleter_SetModelSorting(QtObjectPtr ptr, int sorting);
void QCompleter_SetWrapAround(QtObjectPtr ptr, int wrap);
int QCompleter_WrapAround(QtObjectPtr ptr);
QtObjectPtr QCompleter_NewQCompleter2(QtObjectPtr model, QtObjectPtr parent);
QtObjectPtr QCompleter_NewQCompleter(QtObjectPtr parent);
QtObjectPtr QCompleter_NewQCompleter3(char* list, QtObjectPtr parent);
void QCompleter_ConnectActivated(QtObjectPtr ptr);
void QCompleter_DisconnectActivated(QtObjectPtr ptr);
void QCompleter_Complete(QtObjectPtr ptr, QtObjectPtr rect);
int QCompleter_CompletionCount(QtObjectPtr ptr);
QtObjectPtr QCompleter_CompletionModel(QtObjectPtr ptr);
char* QCompleter_CurrentCompletion(QtObjectPtr ptr);
QtObjectPtr QCompleter_CurrentIndex(QtObjectPtr ptr);
int QCompleter_CurrentRow(QtObjectPtr ptr);
void QCompleter_ConnectHighlighted(QtObjectPtr ptr);
void QCompleter_DisconnectHighlighted(QtObjectPtr ptr);
QtObjectPtr QCompleter_Model(QtObjectPtr ptr);
char* QCompleter_PathFromIndex(QtObjectPtr ptr, QtObjectPtr index);
QtObjectPtr QCompleter_Popup(QtObjectPtr ptr);
int QCompleter_SetCurrentRow(QtObjectPtr ptr, int row);
void QCompleter_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QCompleter_SetPopup(QtObjectPtr ptr, QtObjectPtr popup);
void QCompleter_SetWidget(QtObjectPtr ptr, QtObjectPtr widget);
char* QCompleter_SplitPath(QtObjectPtr ptr, char* path);
QtObjectPtr QCompleter_Widget(QtObjectPtr ptr);
void QCompleter_DestroyQCompleter(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif