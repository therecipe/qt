#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QAbstractScrollArea_New_QWidget(QtObjectPtr parent);
void QAbstractScrollArea_Destroy(QtObjectPtr ptr);
void QAbstractScrollArea_AddScrollBarWidget_QWidget_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr widget, int alignment);
QtObjectPtr QAbstractScrollArea_CornerWidget(QtObjectPtr ptr);
int QAbstractScrollArea_HorizontalScrollBarPolicy(QtObjectPtr ptr);
void QAbstractScrollArea_SetCornerWidget_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QAbstractScrollArea_SetHorizontalScrollBarPolicy_ScrollBarPolicy(QtObjectPtr ptr, int ScrollBarPolicy);
void QAbstractScrollArea_SetVerticalScrollBarPolicy_ScrollBarPolicy(QtObjectPtr ptr, int ScrollBarPolicy);
void QAbstractScrollArea_SetViewport_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QAbstractScrollArea_SetupViewport_QWidget(QtObjectPtr ptr, QtObjectPtr viewport);
int QAbstractScrollArea_VerticalScrollBarPolicy(QtObjectPtr ptr);
QtObjectPtr QAbstractScrollArea_Viewport(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
