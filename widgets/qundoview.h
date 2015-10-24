#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QUndoView_EmptyLabel(QtObjectPtr ptr);
void QUndoView_SetCleanIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QUndoView_SetEmptyLabel(QtObjectPtr ptr, char* label);
QtObjectPtr QUndoView_NewQUndoView3(QtObjectPtr group, QtObjectPtr parent);
QtObjectPtr QUndoView_NewQUndoView2(QtObjectPtr stack, QtObjectPtr parent);
QtObjectPtr QUndoView_NewQUndoView(QtObjectPtr parent);
QtObjectPtr QUndoView_Group(QtObjectPtr ptr);
void QUndoView_SetGroup(QtObjectPtr ptr, QtObjectPtr group);
void QUndoView_SetStack(QtObjectPtr ptr, QtObjectPtr stack);
QtObjectPtr QUndoView_Stack(QtObjectPtr ptr);
void QUndoView_DestroyQUndoView(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif