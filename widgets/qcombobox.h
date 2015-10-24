#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QComboBox_Count(QtObjectPtr ptr);
char* QComboBox_CurrentData(QtObjectPtr ptr, int role);
int QComboBox_CurrentIndex(QtObjectPtr ptr);
char* QComboBox_CurrentText(QtObjectPtr ptr);
int QComboBox_DuplicatesEnabled(QtObjectPtr ptr);
int QComboBox_HasFrame(QtObjectPtr ptr);
int QComboBox_InsertPolicy(QtObjectPtr ptr);
int QComboBox_IsEditable(QtObjectPtr ptr);
int QComboBox_MaxCount(QtObjectPtr ptr);
int QComboBox_MaxVisibleItems(QtObjectPtr ptr);
int QComboBox_MinimumContentsLength(QtObjectPtr ptr);
int QComboBox_ModelColumn(QtObjectPtr ptr);
void QComboBox_SetCompleter(QtObjectPtr ptr, QtObjectPtr completer);
void QComboBox_SetCurrentIndex(QtObjectPtr ptr, int index);
void QComboBox_SetCurrentText(QtObjectPtr ptr, char* text);
void QComboBox_SetDuplicatesEnabled(QtObjectPtr ptr, int enable);
void QComboBox_SetEditable(QtObjectPtr ptr, int editable);
void QComboBox_SetFrame(QtObjectPtr ptr, int v);
void QComboBox_SetIconSize(QtObjectPtr ptr, QtObjectPtr size);
void QComboBox_SetInsertPolicy(QtObjectPtr ptr, int policy);
void QComboBox_SetMaxCount(QtObjectPtr ptr, int max);
void QComboBox_SetMaxVisibleItems(QtObjectPtr ptr, int maxItems);
void QComboBox_SetMinimumContentsLength(QtObjectPtr ptr, int characters);
void QComboBox_SetModelColumn(QtObjectPtr ptr, int visibleColumn);
void QComboBox_SetSizeAdjustPolicy(QtObjectPtr ptr, int policy);
void QComboBox_SetValidator(QtObjectPtr ptr, QtObjectPtr validator);
int QComboBox_SizeAdjustPolicy(QtObjectPtr ptr);
QtObjectPtr QComboBox_NewQComboBox(QtObjectPtr parent);
void QComboBox_ConnectActivated(QtObjectPtr ptr);
void QComboBox_DisconnectActivated(QtObjectPtr ptr);
void QComboBox_AddItem2(QtObjectPtr ptr, QtObjectPtr icon, char* text, char* userData);
void QComboBox_AddItem(QtObjectPtr ptr, char* text, char* userData);
void QComboBox_AddItems(QtObjectPtr ptr, char* texts);
void QComboBox_Clear(QtObjectPtr ptr);
void QComboBox_ClearEditText(QtObjectPtr ptr);
QtObjectPtr QComboBox_Completer(QtObjectPtr ptr);
void QComboBox_ConnectCurrentIndexChanged(QtObjectPtr ptr);
void QComboBox_DisconnectCurrentIndexChanged(QtObjectPtr ptr);
void QComboBox_ConnectCurrentTextChanged(QtObjectPtr ptr);
void QComboBox_DisconnectCurrentTextChanged(QtObjectPtr ptr);
void QComboBox_ConnectEditTextChanged(QtObjectPtr ptr);
void QComboBox_DisconnectEditTextChanged(QtObjectPtr ptr);
int QComboBox_Event(QtObjectPtr ptr, QtObjectPtr event);
int QComboBox_FindData(QtObjectPtr ptr, char* data, int role, int flags);
int QComboBox_FindText(QtObjectPtr ptr, char* text, int flags);
void QComboBox_HidePopup(QtObjectPtr ptr);
void QComboBox_ConnectHighlighted(QtObjectPtr ptr);
void QComboBox_DisconnectHighlighted(QtObjectPtr ptr);
char* QComboBox_InputMethodQuery(QtObjectPtr ptr, int query);
void QComboBox_InsertItem2(QtObjectPtr ptr, int index, QtObjectPtr icon, char* text, char* userData);
void QComboBox_InsertItem(QtObjectPtr ptr, int index, char* text, char* userData);
void QComboBox_InsertItems(QtObjectPtr ptr, int index, char* list);
void QComboBox_InsertSeparator(QtObjectPtr ptr, int index);
char* QComboBox_ItemData(QtObjectPtr ptr, int index, int role);
QtObjectPtr QComboBox_ItemDelegate(QtObjectPtr ptr);
char* QComboBox_ItemText(QtObjectPtr ptr, int index);
QtObjectPtr QComboBox_LineEdit(QtObjectPtr ptr);
QtObjectPtr QComboBox_Model(QtObjectPtr ptr);
void QComboBox_RemoveItem(QtObjectPtr ptr, int index);
QtObjectPtr QComboBox_RootModelIndex(QtObjectPtr ptr);
void QComboBox_SetEditText(QtObjectPtr ptr, char* text);
void QComboBox_SetItemData(QtObjectPtr ptr, int index, char* value, int role);
void QComboBox_SetItemDelegate(QtObjectPtr ptr, QtObjectPtr delegate);
void QComboBox_SetItemIcon(QtObjectPtr ptr, int index, QtObjectPtr icon);
void QComboBox_SetItemText(QtObjectPtr ptr, int index, char* text);
void QComboBox_SetLineEdit(QtObjectPtr ptr, QtObjectPtr edit);
void QComboBox_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QComboBox_SetRootModelIndex(QtObjectPtr ptr, QtObjectPtr index);
void QComboBox_SetView(QtObjectPtr ptr, QtObjectPtr itemView);
void QComboBox_ShowPopup(QtObjectPtr ptr);
QtObjectPtr QComboBox_Validator(QtObjectPtr ptr);
QtObjectPtr QComboBox_View(QtObjectPtr ptr);
void QComboBox_DestroyQComboBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif