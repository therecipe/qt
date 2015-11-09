#ifdef __cplusplus
extern "C" {
#endif

int QAbstractScrollArea_HorizontalScrollBarPolicy(void* ptr);
void QAbstractScrollArea_SetHorizontalScrollBarPolicy(void* ptr, int v);
void QAbstractScrollArea_SetSizeAdjustPolicy(void* ptr, int policy);
void QAbstractScrollArea_SetVerticalScrollBarPolicy(void* ptr, int v);
int QAbstractScrollArea_SizeAdjustPolicy(void* ptr);
int QAbstractScrollArea_VerticalScrollBarPolicy(void* ptr);
void* QAbstractScrollArea_NewQAbstractScrollArea(void* parent);
void QAbstractScrollArea_AddScrollBarWidget(void* ptr, void* widget, int alignment);
void* QAbstractScrollArea_CornerWidget(void* ptr);
void* QAbstractScrollArea_HorizontalScrollBar(void* ptr);
void QAbstractScrollArea_SetCornerWidget(void* ptr, void* widget);
void QAbstractScrollArea_SetHorizontalScrollBar(void* ptr, void* scrollBar);
void QAbstractScrollArea_SetVerticalScrollBar(void* ptr, void* scrollBar);
void QAbstractScrollArea_SetViewport(void* ptr, void* widget);
void QAbstractScrollArea_SetupViewport(void* ptr, void* viewport);
void* QAbstractScrollArea_VerticalScrollBar(void* ptr);
void* QAbstractScrollArea_Viewport(void* ptr);
void QAbstractScrollArea_DestroyQAbstractScrollArea(void* ptr);

#ifdef __cplusplus
}
#endif