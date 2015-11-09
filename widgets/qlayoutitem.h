#ifdef __cplusplus
extern "C" {
#endif

int QLayoutItem_Alignment(void* ptr);
int QLayoutItem_ControlTypes(void* ptr);
int QLayoutItem_ExpandingDirections(void* ptr);
int QLayoutItem_HasHeightForWidth(void* ptr);
int QLayoutItem_HeightForWidth(void* ptr, int w);
void QLayoutItem_Invalidate(void* ptr);
int QLayoutItem_IsEmpty(void* ptr);
void* QLayoutItem_Layout(void* ptr);
int QLayoutItem_MinimumHeightForWidth(void* ptr, int w);
void QLayoutItem_SetAlignment(void* ptr, int alignment);
void QLayoutItem_SetGeometry(void* ptr, void* r);
void* QLayoutItem_SpacerItem(void* ptr);
void* QLayoutItem_Widget(void* ptr);
void QLayoutItem_DestroyQLayoutItem(void* ptr);

#ifdef __cplusplus
}
#endif