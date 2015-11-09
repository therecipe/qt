#ifdef __cplusplus
extern "C" {
#endif

int QStackedWidget_Count(void* ptr);
int QStackedWidget_CurrentIndex(void* ptr);
void QStackedWidget_SetCurrentIndex(void* ptr, int index);
void QStackedWidget_SetCurrentWidget(void* ptr, void* widget);
void* QStackedWidget_NewQStackedWidget(void* parent);
int QStackedWidget_AddWidget(void* ptr, void* widget);
void QStackedWidget_ConnectCurrentChanged(void* ptr);
void QStackedWidget_DisconnectCurrentChanged(void* ptr);
void* QStackedWidget_CurrentWidget(void* ptr);
int QStackedWidget_IndexOf(void* ptr, void* widget);
int QStackedWidget_InsertWidget(void* ptr, int index, void* widget);
void QStackedWidget_RemoveWidget(void* ptr, void* widget);
void* QStackedWidget_Widget(void* ptr, int index);
void QStackedWidget_ConnectWidgetRemoved(void* ptr);
void QStackedWidget_DisconnectWidgetRemoved(void* ptr);
void QStackedWidget_DestroyQStackedWidget(void* ptr);

#ifdef __cplusplus
}
#endif