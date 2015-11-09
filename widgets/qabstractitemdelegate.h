#ifdef __cplusplus
extern "C" {
#endif

void QAbstractItemDelegate_ConnectCloseEditor(void* ptr);
void QAbstractItemDelegate_DisconnectCloseEditor(void* ptr);
void QAbstractItemDelegate_ConnectCommitData(void* ptr);
void QAbstractItemDelegate_DisconnectCommitData(void* ptr);
void* QAbstractItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index);
void QAbstractItemDelegate_DestroyEditor(void* ptr, void* editor, void* index);
int QAbstractItemDelegate_EditorEvent(void* ptr, void* event, void* model, void* option, void* index);
int QAbstractItemDelegate_HelpEvent(void* ptr, void* event, void* view, void* option, void* index);
void QAbstractItemDelegate_Paint(void* ptr, void* painter, void* option, void* index);
void QAbstractItemDelegate_SetEditorData(void* ptr, void* editor, void* index);
void QAbstractItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index);
void QAbstractItemDelegate_ConnectSizeHintChanged(void* ptr);
void QAbstractItemDelegate_DisconnectSizeHintChanged(void* ptr);
void QAbstractItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index);
void QAbstractItemDelegate_DestroyQAbstractItemDelegate(void* ptr);

#ifdef __cplusplus
}
#endif