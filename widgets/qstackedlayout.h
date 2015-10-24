#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QStackedLayout_Count(QtObjectPtr ptr);
int QStackedLayout_CurrentIndex(QtObjectPtr ptr);
void QStackedLayout_SetCurrentIndex(QtObjectPtr ptr, int index);
void QStackedLayout_SetCurrentWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QStackedLayout_SetStackingMode(QtObjectPtr ptr, int stackingMode);
int QStackedLayout_StackingMode(QtObjectPtr ptr);
QtObjectPtr QStackedLayout_NewQStackedLayout();
QtObjectPtr QStackedLayout_NewQStackedLayout3(QtObjectPtr parentLayout);
QtObjectPtr QStackedLayout_NewQStackedLayout2(QtObjectPtr parent);
void QStackedLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item);
int QStackedLayout_AddWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QStackedLayout_ConnectCurrentChanged(QtObjectPtr ptr);
void QStackedLayout_DisconnectCurrentChanged(QtObjectPtr ptr);
QtObjectPtr QStackedLayout_CurrentWidget(QtObjectPtr ptr);
int QStackedLayout_HasHeightForWidth(QtObjectPtr ptr);
int QStackedLayout_HeightForWidth(QtObjectPtr ptr, int width);
int QStackedLayout_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget);
QtObjectPtr QStackedLayout_ItemAt(QtObjectPtr ptr, int index);
void QStackedLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
QtObjectPtr QStackedLayout_TakeAt(QtObjectPtr ptr, int index);
QtObjectPtr QStackedLayout_Widget(QtObjectPtr ptr, int index);
void QStackedLayout_ConnectWidgetRemoved(QtObjectPtr ptr);
void QStackedLayout_DisconnectWidgetRemoved(QtObjectPtr ptr);
void QStackedLayout_DestroyQStackedLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif