#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QAbstractItemDelegate_ConnectCloseEditor(QtObjectPtr ptr);
void QAbstractItemDelegate_DisconnectCloseEditor(QtObjectPtr ptr);
void QAbstractItemDelegate_ConnectCommitData(QtObjectPtr ptr);
void QAbstractItemDelegate_DisconnectCommitData(QtObjectPtr ptr);
QtObjectPtr QAbstractItemDelegate_CreateEditor(QtObjectPtr ptr, QtObjectPtr parent, QtObjectPtr option, QtObjectPtr index);
void QAbstractItemDelegate_DestroyEditor(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr index);
int QAbstractItemDelegate_EditorEvent(QtObjectPtr ptr, QtObjectPtr event, QtObjectPtr model, QtObjectPtr option, QtObjectPtr index);
int QAbstractItemDelegate_HelpEvent(QtObjectPtr ptr, QtObjectPtr event, QtObjectPtr view, QtObjectPtr option, QtObjectPtr index);
void QAbstractItemDelegate_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr index);
void QAbstractItemDelegate_SetEditorData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr index);
void QAbstractItemDelegate_SetModelData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr model, QtObjectPtr index);
void QAbstractItemDelegate_ConnectSizeHintChanged(QtObjectPtr ptr);
void QAbstractItemDelegate_DisconnectSizeHintChanged(QtObjectPtr ptr);
void QAbstractItemDelegate_UpdateEditorGeometry(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr option, QtObjectPtr index);
void QAbstractItemDelegate_DestroyQAbstractItemDelegate(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif