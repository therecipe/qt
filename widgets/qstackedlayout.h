#ifdef __cplusplus
extern "C" {
#endif

int QStackedLayout_Count(void* ptr);
int QStackedLayout_CurrentIndex(void* ptr);
void QStackedLayout_SetCurrentIndex(void* ptr, int index);
void QStackedLayout_SetCurrentWidget(void* ptr, void* widget);
void QStackedLayout_SetStackingMode(void* ptr, int stackingMode);
int QStackedLayout_StackingMode(void* ptr);
void* QStackedLayout_NewQStackedLayout();
void* QStackedLayout_NewQStackedLayout3(void* parentLayout);
void* QStackedLayout_NewQStackedLayout2(void* parent);
void QStackedLayout_AddItem(void* ptr, void* item);
int QStackedLayout_AddWidget(void* ptr, void* widget);
void QStackedLayout_ConnectCurrentChanged(void* ptr);
void QStackedLayout_DisconnectCurrentChanged(void* ptr);
void* QStackedLayout_CurrentWidget(void* ptr);
int QStackedLayout_HasHeightForWidth(void* ptr);
int QStackedLayout_HeightForWidth(void* ptr, int width);
int QStackedLayout_InsertWidget(void* ptr, int index, void* widget);
void* QStackedLayout_ItemAt(void* ptr, int index);
void QStackedLayout_SetGeometry(void* ptr, void* rect);
void* QStackedLayout_TakeAt(void* ptr, int index);
void* QStackedLayout_Widget(void* ptr, int index);
void QStackedLayout_ConnectWidgetRemoved(void* ptr);
void QStackedLayout_DisconnectWidgetRemoved(void* ptr);
void QStackedLayout_DestroyQStackedLayout(void* ptr);

#ifdef __cplusplus
}
#endif