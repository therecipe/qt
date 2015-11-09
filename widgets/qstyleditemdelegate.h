#ifdef __cplusplus
extern "C" {
#endif

void* QStyledItemDelegate_NewQStyledItemDelegate(void* parent);
void* QStyledItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index);
char* QStyledItemDelegate_DisplayText(void* ptr, void* value, void* locale);
void* QStyledItemDelegate_ItemEditorFactory(void* ptr);
void QStyledItemDelegate_Paint(void* ptr, void* painter, void* option, void* index);
void QStyledItemDelegate_SetEditorData(void* ptr, void* editor, void* index);
void QStyledItemDelegate_SetItemEditorFactory(void* ptr, void* factory);
void QStyledItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index);
void QStyledItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index);
void QStyledItemDelegate_DestroyQStyledItemDelegate(void* ptr);

#ifdef __cplusplus
}
#endif