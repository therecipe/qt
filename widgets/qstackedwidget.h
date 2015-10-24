#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QStackedWidget_Count(QtObjectPtr ptr);
int QStackedWidget_CurrentIndex(QtObjectPtr ptr);
void QStackedWidget_SetCurrentIndex(QtObjectPtr ptr, int index);
void QStackedWidget_SetCurrentWidget(QtObjectPtr ptr, QtObjectPtr widget);
QtObjectPtr QStackedWidget_NewQStackedWidget(QtObjectPtr parent);
int QStackedWidget_AddWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QStackedWidget_ConnectCurrentChanged(QtObjectPtr ptr);
void QStackedWidget_DisconnectCurrentChanged(QtObjectPtr ptr);
QtObjectPtr QStackedWidget_CurrentWidget(QtObjectPtr ptr);
int QStackedWidget_IndexOf(QtObjectPtr ptr, QtObjectPtr widget);
int QStackedWidget_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget);
void QStackedWidget_RemoveWidget(QtObjectPtr ptr, QtObjectPtr widget);
QtObjectPtr QStackedWidget_Widget(QtObjectPtr ptr, int index);
void QStackedWidget_ConnectWidgetRemoved(QtObjectPtr ptr);
void QStackedWidget_DisconnectWidgetRemoved(QtObjectPtr ptr);
void QStackedWidget_DestroyQStackedWidget(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif