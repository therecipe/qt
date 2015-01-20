#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QLayout_SetDefaultConstraint();
int QLayout_SetFixedSize();
int QLayout_SetMinimumSize();
int QLayout_SetMaximumSize();
int QLayout_SetMinAndMaxSize();
int QLayout_SetNoConstraint();
//Public Functions
int QLayout_Activate(QtObjectPtr ptr);
void QLayout_AddWidget_QWidget(QtObjectPtr ptr, QtObjectPtr w);
int QLayout_Count(QtObjectPtr ptr);
int QLayout_IndexOf_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
int QLayout_IsEnabled(QtObjectPtr ptr);
QtObjectPtr QLayout_MenuBar(QtObjectPtr ptr);
QtObjectPtr QLayout_ParentWidget(QtObjectPtr ptr);
void QLayout_RemoveItem_QLayoutItem(QtObjectPtr ptr, QtObjectPtr item);
void QLayout_RemoveWidget_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
int QLayout_SetAlignment_QWidget_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr w, int alignment);
void QLayout_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment);
int QLayout_SetAlignment_QLayout_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr l, int alignment);
void QLayout_SetContentsMargins_Int_Int_Int_Int(QtObjectPtr ptr, int left, int top, int right, int bottom);
void QLayout_SetEnabled_Bool(QtObjectPtr ptr, int enable);
void QLayout_SetMenuBar_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QLayout_SetSizeConstraint_SizeConstraint(QtObjectPtr ptr, int Si);
void QLayout_SetSpacing_Int(QtObjectPtr ptr, int spacing);
int QLayout_SizeConstraint(QtObjectPtr ptr);
int QLayout_Spacing(QtObjectPtr ptr);
void QLayout_Update(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
