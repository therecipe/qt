#ifdef __cplusplus
extern "C" {
#endif

int QCompleter_CaseSensitivity(void* ptr);
int QCompleter_CompletionColumn(void* ptr);
int QCompleter_CompletionMode(void* ptr);
char* QCompleter_CompletionPrefix(void* ptr);
int QCompleter_CompletionRole(void* ptr);
int QCompleter_FilterMode(void* ptr);
int QCompleter_MaxVisibleItems(void* ptr);
int QCompleter_ModelSorting(void* ptr);
void QCompleter_SetCaseSensitivity(void* ptr, int caseSensitivity);
void QCompleter_SetCompletionColumn(void* ptr, int column);
void QCompleter_SetCompletionMode(void* ptr, int mode);
void QCompleter_SetCompletionPrefix(void* ptr, char* prefix);
void QCompleter_SetCompletionRole(void* ptr, int role);
void QCompleter_SetFilterMode(void* ptr, int filterMode);
void QCompleter_SetMaxVisibleItems(void* ptr, int maxItems);
void QCompleter_SetModelSorting(void* ptr, int sorting);
void QCompleter_SetWrapAround(void* ptr, int wrap);
int QCompleter_WrapAround(void* ptr);
void* QCompleter_NewQCompleter2(void* model, void* parent);
void* QCompleter_NewQCompleter(void* parent);
void* QCompleter_NewQCompleter3(char* list, void* parent);
void QCompleter_ConnectActivated(void* ptr);
void QCompleter_DisconnectActivated(void* ptr);
void QCompleter_Complete(void* ptr, void* rect);
int QCompleter_CompletionCount(void* ptr);
void* QCompleter_CompletionModel(void* ptr);
char* QCompleter_CurrentCompletion(void* ptr);
void* QCompleter_CurrentIndex(void* ptr);
int QCompleter_CurrentRow(void* ptr);
void QCompleter_ConnectHighlighted(void* ptr);
void QCompleter_DisconnectHighlighted(void* ptr);
void* QCompleter_Model(void* ptr);
char* QCompleter_PathFromIndex(void* ptr, void* index);
void* QCompleter_Popup(void* ptr);
int QCompleter_SetCurrentRow(void* ptr, int row);
void QCompleter_SetModel(void* ptr, void* model);
void QCompleter_SetPopup(void* ptr, void* popup);
void QCompleter_SetWidget(void* ptr, void* widget);
char* QCompleter_SplitPath(void* ptr, char* path);
void* QCompleter_Widget(void* ptr);
void QCompleter_DestroyQCompleter(void* ptr);

#ifdef __cplusplus
}
#endif