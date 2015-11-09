#ifdef __cplusplus
extern "C" {
#endif

int QComboBox_Count(void* ptr);
void* QComboBox_CurrentData(void* ptr, int role);
int QComboBox_CurrentIndex(void* ptr);
char* QComboBox_CurrentText(void* ptr);
int QComboBox_DuplicatesEnabled(void* ptr);
int QComboBox_HasFrame(void* ptr);
int QComboBox_InsertPolicy(void* ptr);
int QComboBox_IsEditable(void* ptr);
int QComboBox_MaxCount(void* ptr);
int QComboBox_MaxVisibleItems(void* ptr);
int QComboBox_MinimumContentsLength(void* ptr);
int QComboBox_ModelColumn(void* ptr);
void QComboBox_SetCompleter(void* ptr, void* completer);
void QComboBox_SetCurrentIndex(void* ptr, int index);
void QComboBox_SetCurrentText(void* ptr, char* text);
void QComboBox_SetDuplicatesEnabled(void* ptr, int enable);
void QComboBox_SetEditable(void* ptr, int editable);
void QComboBox_SetFrame(void* ptr, int v);
void QComboBox_SetIconSize(void* ptr, void* size);
void QComboBox_SetInsertPolicy(void* ptr, int policy);
void QComboBox_SetMaxCount(void* ptr, int max);
void QComboBox_SetMaxVisibleItems(void* ptr, int maxItems);
void QComboBox_SetMinimumContentsLength(void* ptr, int characters);
void QComboBox_SetModelColumn(void* ptr, int visibleColumn);
void QComboBox_SetSizeAdjustPolicy(void* ptr, int policy);
void QComboBox_SetValidator(void* ptr, void* validator);
int QComboBox_SizeAdjustPolicy(void* ptr);
void* QComboBox_NewQComboBox(void* parent);
void QComboBox_ConnectActivated(void* ptr);
void QComboBox_DisconnectActivated(void* ptr);
void QComboBox_AddItem2(void* ptr, void* icon, char* text, void* userData);
void QComboBox_AddItem(void* ptr, char* text, void* userData);
void QComboBox_AddItems(void* ptr, char* texts);
void QComboBox_Clear(void* ptr);
void QComboBox_ClearEditText(void* ptr);
void* QComboBox_Completer(void* ptr);
void QComboBox_ConnectCurrentIndexChanged(void* ptr);
void QComboBox_DisconnectCurrentIndexChanged(void* ptr);
void QComboBox_ConnectCurrentTextChanged(void* ptr);
void QComboBox_DisconnectCurrentTextChanged(void* ptr);
void QComboBox_ConnectEditTextChanged(void* ptr);
void QComboBox_DisconnectEditTextChanged(void* ptr);
int QComboBox_Event(void* ptr, void* event);
int QComboBox_FindData(void* ptr, void* data, int role, int flags);
int QComboBox_FindText(void* ptr, char* text, int flags);
void QComboBox_HidePopup(void* ptr);
void QComboBox_ConnectHighlighted(void* ptr);
void QComboBox_DisconnectHighlighted(void* ptr);
void* QComboBox_InputMethodQuery(void* ptr, int query);
void QComboBox_InsertItem2(void* ptr, int index, void* icon, char* text, void* userData);
void QComboBox_InsertItem(void* ptr, int index, char* text, void* userData);
void QComboBox_InsertItems(void* ptr, int index, char* list);
void QComboBox_InsertSeparator(void* ptr, int index);
void* QComboBox_ItemData(void* ptr, int index, int role);
void* QComboBox_ItemDelegate(void* ptr);
char* QComboBox_ItemText(void* ptr, int index);
void* QComboBox_LineEdit(void* ptr);
void* QComboBox_Model(void* ptr);
void QComboBox_RemoveItem(void* ptr, int index);
void* QComboBox_RootModelIndex(void* ptr);
void QComboBox_SetEditText(void* ptr, char* text);
void QComboBox_SetItemData(void* ptr, int index, void* value, int role);
void QComboBox_SetItemDelegate(void* ptr, void* delegate);
void QComboBox_SetItemIcon(void* ptr, int index, void* icon);
void QComboBox_SetItemText(void* ptr, int index, char* text);
void QComboBox_SetLineEdit(void* ptr, void* edit);
void QComboBox_SetModel(void* ptr, void* model);
void QComboBox_SetRootModelIndex(void* ptr, void* index);
void QComboBox_SetView(void* ptr, void* itemView);
void QComboBox_ShowPopup(void* ptr);
void* QComboBox_Validator(void* ptr);
void* QComboBox_View(void* ptr);
void QComboBox_DestroyQComboBox(void* ptr);

#ifdef __cplusplus
}
#endif