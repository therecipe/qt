#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QLayout_SetSizeConstraint(QtObjectPtr ptr, int v);
void QLayout_SetSpacing(QtObjectPtr ptr, int v);
int QLayout_SizeConstraint(QtObjectPtr ptr);
int QLayout_Spacing(QtObjectPtr ptr);
int QLayout_Activate(QtObjectPtr ptr);
void QLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item);
void QLayout_AddWidget(QtObjectPtr ptr, QtObjectPtr w);
int QLayout_ControlTypes(QtObjectPtr ptr);
int QLayout_Count(QtObjectPtr ptr);
int QLayout_ExpandingDirections(QtObjectPtr ptr);
void QLayout_GetContentsMargins(QtObjectPtr ptr, int left, int top, int right, int bottom);
int QLayout_IndexOf(QtObjectPtr ptr, QtObjectPtr widget);
void QLayout_Invalidate(QtObjectPtr ptr);
int QLayout_IsEmpty(QtObjectPtr ptr);
int QLayout_IsEnabled(QtObjectPtr ptr);
QtObjectPtr QLayout_ItemAt(QtObjectPtr ptr, int index);
QtObjectPtr QLayout_Layout(QtObjectPtr ptr);
QtObjectPtr QLayout_MenuBar(QtObjectPtr ptr);
QtObjectPtr QLayout_ParentWidget(QtObjectPtr ptr);
void QLayout_RemoveItem(QtObjectPtr ptr, QtObjectPtr item);
void QLayout_RemoveWidget(QtObjectPtr ptr, QtObjectPtr widget);
QtObjectPtr QLayout_ReplaceWidget(QtObjectPtr ptr, QtObjectPtr from, QtObjectPtr to, int options);
int QLayout_SetAlignment2(QtObjectPtr ptr, QtObjectPtr l, int alignment);
int QLayout_SetAlignment(QtObjectPtr ptr, QtObjectPtr w, int alignment);
void QLayout_SetContentsMargins2(QtObjectPtr ptr, QtObjectPtr margins);
void QLayout_SetContentsMargins(QtObjectPtr ptr, int left, int top, int right, int bottom);
void QLayout_SetEnabled(QtObjectPtr ptr, int enable);
void QLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr r);
void QLayout_SetMenuBar(QtObjectPtr ptr, QtObjectPtr widget);
QtObjectPtr QLayout_TakeAt(QtObjectPtr ptr, int index);
void QLayout_Update(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif