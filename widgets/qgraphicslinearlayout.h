#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsLinearLayout_NewQGraphicsLinearLayout(void* parent);
void* QGraphicsLinearLayout_NewQGraphicsLinearLayout2(int orientation, void* parent);
void QGraphicsLinearLayout_AddItem(void* ptr, void* item);
void QGraphicsLinearLayout_AddStretch(void* ptr, int stretch);
int QGraphicsLinearLayout_Alignment(void* ptr, void* item);
int QGraphicsLinearLayout_Count(void* ptr);
void QGraphicsLinearLayout_InsertItem(void* ptr, int index, void* item);
void QGraphicsLinearLayout_InsertStretch(void* ptr, int index, int stretch);
void QGraphicsLinearLayout_Invalidate(void* ptr);
void* QGraphicsLinearLayout_ItemAt(void* ptr, int index);
double QGraphicsLinearLayout_ItemSpacing(void* ptr, int index);
int QGraphicsLinearLayout_Orientation(void* ptr);
void QGraphicsLinearLayout_RemoveAt(void* ptr, int index);
void QGraphicsLinearLayout_RemoveItem(void* ptr, void* item);
void QGraphicsLinearLayout_SetAlignment(void* ptr, void* item, int alignment);
void QGraphicsLinearLayout_SetGeometry(void* ptr, void* rect);
void QGraphicsLinearLayout_SetItemSpacing(void* ptr, int index, double spacing);
void QGraphicsLinearLayout_SetOrientation(void* ptr, int orientation);
void QGraphicsLinearLayout_SetSpacing(void* ptr, double spacing);
void QGraphicsLinearLayout_SetStretchFactor(void* ptr, void* item, int stretch);
double QGraphicsLinearLayout_Spacing(void* ptr);
int QGraphicsLinearLayout_StretchFactor(void* ptr, void* item);
void QGraphicsLinearLayout_DestroyQGraphicsLinearLayout(void* ptr);

#ifdef __cplusplus
}
#endif