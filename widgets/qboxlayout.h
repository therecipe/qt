#ifdef __cplusplus
extern "C" {
#endif

int QBoxLayout_Direction(void* ptr);
void* QBoxLayout_NewQBoxLayout(int dir, void* parent);
void QBoxLayout_AddItem(void* ptr, void* item);
void QBoxLayout_AddLayout(void* ptr, void* layout, int stretch);
void QBoxLayout_AddSpacerItem(void* ptr, void* spacerItem);
void QBoxLayout_AddSpacing(void* ptr, int size);
void QBoxLayout_AddStretch(void* ptr, int stretch);
void QBoxLayout_AddStrut(void* ptr, int size);
void QBoxLayout_AddWidget(void* ptr, void* widget, int stretch, int alignment);
int QBoxLayout_Count(void* ptr);
int QBoxLayout_ExpandingDirections(void* ptr);
int QBoxLayout_HasHeightForWidth(void* ptr);
int QBoxLayout_HeightForWidth(void* ptr, int w);
void QBoxLayout_InsertItem(void* ptr, int index, void* item);
void QBoxLayout_InsertLayout(void* ptr, int index, void* layout, int stretch);
void QBoxLayout_InsertSpacerItem(void* ptr, int index, void* spacerItem);
void QBoxLayout_InsertSpacing(void* ptr, int index, int size);
void QBoxLayout_InsertStretch(void* ptr, int index, int stretch);
void QBoxLayout_InsertWidget(void* ptr, int index, void* widget, int stretch, int alignment);
void QBoxLayout_Invalidate(void* ptr);
void* QBoxLayout_ItemAt(void* ptr, int index);
int QBoxLayout_MinimumHeightForWidth(void* ptr, int w);
void QBoxLayout_SetDirection(void* ptr, int direction);
void QBoxLayout_SetGeometry(void* ptr, void* r);
void QBoxLayout_SetSpacing(void* ptr, int spacing);
void QBoxLayout_SetStretch(void* ptr, int index, int stretch);
int QBoxLayout_SetStretchFactor2(void* ptr, void* layout, int stretch);
int QBoxLayout_SetStretchFactor(void* ptr, void* widget, int stretch);
int QBoxLayout_Spacing(void* ptr);
int QBoxLayout_Stretch(void* ptr, int index);
void* QBoxLayout_TakeAt(void* ptr, int index);
void QBoxLayout_DestroyQBoxLayout(void* ptr);

#ifdef __cplusplus
}
#endif