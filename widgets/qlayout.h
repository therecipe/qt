#ifdef __cplusplus
extern "C" {
#endif

void QLayout_SetSizeConstraint(void* ptr, int v);
void QLayout_SetSpacing(void* ptr, int v);
int QLayout_SizeConstraint(void* ptr);
int QLayout_Spacing(void* ptr);
int QLayout_Activate(void* ptr);
void QLayout_AddItem(void* ptr, void* item);
void QLayout_AddWidget(void* ptr, void* w);
int QLayout_ControlTypes(void* ptr);
int QLayout_Count(void* ptr);
int QLayout_ExpandingDirections(void* ptr);
void QLayout_GetContentsMargins(void* ptr, int left, int top, int right, int bottom);
int QLayout_IndexOf(void* ptr, void* widget);
void QLayout_Invalidate(void* ptr);
int QLayout_IsEmpty(void* ptr);
int QLayout_IsEnabled(void* ptr);
void* QLayout_ItemAt(void* ptr, int index);
void* QLayout_Layout(void* ptr);
void* QLayout_MenuBar(void* ptr);
void* QLayout_ParentWidget(void* ptr);
void QLayout_RemoveItem(void* ptr, void* item);
void QLayout_RemoveWidget(void* ptr, void* widget);
void* QLayout_ReplaceWidget(void* ptr, void* from, void* to, int options);
int QLayout_SetAlignment2(void* ptr, void* l, int alignment);
int QLayout_SetAlignment(void* ptr, void* w, int alignment);
void QLayout_SetContentsMargins2(void* ptr, void* margins);
void QLayout_SetContentsMargins(void* ptr, int left, int top, int right, int bottom);
void QLayout_SetEnabled(void* ptr, int enable);
void QLayout_SetGeometry(void* ptr, void* r);
void QLayout_SetMenuBar(void* ptr, void* widget);
void* QLayout_TakeAt(void* ptr, int index);
void QLayout_Update(void* ptr);

#ifdef __cplusplus
}
#endif