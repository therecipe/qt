#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAbstractScrollArea_HorizontalScrollBarPolicy(QtObjectPtr ptr);
void QAbstractScrollArea_SetHorizontalScrollBarPolicy(QtObjectPtr ptr, int v);
void QAbstractScrollArea_SetSizeAdjustPolicy(QtObjectPtr ptr, int policy);
void QAbstractScrollArea_SetVerticalScrollBarPolicy(QtObjectPtr ptr, int v);
int QAbstractScrollArea_SizeAdjustPolicy(QtObjectPtr ptr);
int QAbstractScrollArea_VerticalScrollBarPolicy(QtObjectPtr ptr);
QtObjectPtr QAbstractScrollArea_NewQAbstractScrollArea(QtObjectPtr parent);
void QAbstractScrollArea_AddScrollBarWidget(QtObjectPtr ptr, QtObjectPtr widget, int alignment);
QtObjectPtr QAbstractScrollArea_CornerWidget(QtObjectPtr ptr);
QtObjectPtr QAbstractScrollArea_HorizontalScrollBar(QtObjectPtr ptr);
void QAbstractScrollArea_SetCornerWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QAbstractScrollArea_SetHorizontalScrollBar(QtObjectPtr ptr, QtObjectPtr scrollBar);
void QAbstractScrollArea_SetVerticalScrollBar(QtObjectPtr ptr, QtObjectPtr scrollBar);
void QAbstractScrollArea_SetViewport(QtObjectPtr ptr, QtObjectPtr widget);
void QAbstractScrollArea_SetupViewport(QtObjectPtr ptr, QtObjectPtr viewport);
QtObjectPtr QAbstractScrollArea_VerticalScrollBar(QtObjectPtr ptr);
QtObjectPtr QAbstractScrollArea_Viewport(QtObjectPtr ptr);
void QAbstractScrollArea_DestroyQAbstractScrollArea(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif