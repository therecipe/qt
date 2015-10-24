#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStyledItemDelegate_NewQStyledItemDelegate(QtObjectPtr parent);
QtObjectPtr QStyledItemDelegate_CreateEditor(QtObjectPtr ptr, QtObjectPtr parent, QtObjectPtr option, QtObjectPtr index);
char* QStyledItemDelegate_DisplayText(QtObjectPtr ptr, char* value, QtObjectPtr locale);
QtObjectPtr QStyledItemDelegate_ItemEditorFactory(QtObjectPtr ptr);
void QStyledItemDelegate_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr index);
void QStyledItemDelegate_SetEditorData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr index);
void QStyledItemDelegate_SetItemEditorFactory(QtObjectPtr ptr, QtObjectPtr factory);
void QStyledItemDelegate_SetModelData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr model, QtObjectPtr index);
void QStyledItemDelegate_UpdateEditorGeometry(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr option, QtObjectPtr index);
void QStyledItemDelegate_DestroyQStyledItemDelegate(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif