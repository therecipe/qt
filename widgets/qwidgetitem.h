#ifdef __cplusplus
extern "C" {
#endif

void* QWidgetItem_NewQWidgetItem(void* widget);
int QWidgetItem_ControlTypes(void* ptr);
int QWidgetItem_ExpandingDirections(void* ptr);
int QWidgetItem_HasHeightForWidth(void* ptr);
int QWidgetItem_HeightForWidth(void* ptr, int w);
int QWidgetItem_IsEmpty(void* ptr);
void QWidgetItem_SetGeometry(void* ptr, void* rect);
void* QWidgetItem_Widget(void* ptr);
void QWidgetItem_DestroyQWidgetItem(void* ptr);

#ifdef __cplusplus
}
#endif