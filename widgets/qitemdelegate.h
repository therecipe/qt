#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QItemDelegate_HasClipping(QtObjectPtr ptr);
void QItemDelegate_SetClipping(QtObjectPtr ptr, int clip);
QtObjectPtr QItemDelegate_NewQItemDelegate(QtObjectPtr parent);
QtObjectPtr QItemDelegate_CreateEditor(QtObjectPtr ptr, QtObjectPtr parent, QtObjectPtr option, QtObjectPtr index);
QtObjectPtr QItemDelegate_ItemEditorFactory(QtObjectPtr ptr);
void QItemDelegate_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr index);
void QItemDelegate_SetEditorData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr index);
void QItemDelegate_SetItemEditorFactory(QtObjectPtr ptr, QtObjectPtr factory);
void QItemDelegate_SetModelData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr model, QtObjectPtr index);
void QItemDelegate_UpdateEditorGeometry(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr option, QtObjectPtr index);
void QItemDelegate_DestroyQItemDelegate(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif