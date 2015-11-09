#ifdef __cplusplus
extern "C" {
#endif

int QItemDelegate_HasClipping(void* ptr);
void QItemDelegate_SetClipping(void* ptr, int clip);
void* QItemDelegate_NewQItemDelegate(void* parent);
void* QItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index);
void* QItemDelegate_ItemEditorFactory(void* ptr);
void QItemDelegate_Paint(void* ptr, void* painter, void* option, void* index);
void QItemDelegate_SetEditorData(void* ptr, void* editor, void* index);
void QItemDelegate_SetItemEditorFactory(void* ptr, void* factory);
void QItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index);
void QItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index);
void QItemDelegate_DestroyQItemDelegate(void* ptr);

#ifdef __cplusplus
}
#endif