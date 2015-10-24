#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsLinearLayout_NewQGraphicsLinearLayout(QtObjectPtr parent);
QtObjectPtr QGraphicsLinearLayout_NewQGraphicsLinearLayout2(int orientation, QtObjectPtr parent);
void QGraphicsLinearLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsLinearLayout_AddStretch(QtObjectPtr ptr, int stretch);
int QGraphicsLinearLayout_Alignment(QtObjectPtr ptr, QtObjectPtr item);
int QGraphicsLinearLayout_Count(QtObjectPtr ptr);
void QGraphicsLinearLayout_InsertItem(QtObjectPtr ptr, int index, QtObjectPtr item);
void QGraphicsLinearLayout_InsertStretch(QtObjectPtr ptr, int index, int stretch);
void QGraphicsLinearLayout_Invalidate(QtObjectPtr ptr);
QtObjectPtr QGraphicsLinearLayout_ItemAt(QtObjectPtr ptr, int index);
int QGraphicsLinearLayout_Orientation(QtObjectPtr ptr);
void QGraphicsLinearLayout_RemoveAt(QtObjectPtr ptr, int index);
void QGraphicsLinearLayout_RemoveItem(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsLinearLayout_SetAlignment(QtObjectPtr ptr, QtObjectPtr item, int alignment);
void QGraphicsLinearLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
void QGraphicsLinearLayout_SetOrientation(QtObjectPtr ptr, int orientation);
void QGraphicsLinearLayout_SetStretchFactor(QtObjectPtr ptr, QtObjectPtr item, int stretch);
int QGraphicsLinearLayout_StretchFactor(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsLinearLayout_DestroyQGraphicsLinearLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif