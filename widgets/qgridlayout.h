#ifdef __cplusplus
extern "C" {
#endif

int QGridLayout_HorizontalSpacing(void* ptr);
void QGridLayout_SetHorizontalSpacing(void* ptr, int spacing);
void QGridLayout_SetVerticalSpacing(void* ptr, int spacing);
int QGridLayout_VerticalSpacing(void* ptr);
void* QGridLayout_NewQGridLayout2();
void* QGridLayout_NewQGridLayout(void* parent);
void QGridLayout_AddItem(void* ptr, void* item, int row, int column, int rowSpan, int columnSpan, int alignment);
void QGridLayout_AddLayout(void* ptr, void* layout, int row, int column, int alignment);
void QGridLayout_AddLayout2(void* ptr, void* layout, int row, int column, int rowSpan, int columnSpan, int alignment);
void QGridLayout_AddWidget2(void* ptr, void* widget, int fromRow, int fromColumn, int rowSpan, int columnSpan, int alignment);
void QGridLayout_AddWidget(void* ptr, void* widget, int row, int column, int alignment);
int QGridLayout_ColumnCount(void* ptr);
int QGridLayout_ColumnMinimumWidth(void* ptr, int column);
int QGridLayout_ColumnStretch(void* ptr, int column);
int QGridLayout_Count(void* ptr);
int QGridLayout_ExpandingDirections(void* ptr);
void QGridLayout_GetItemPosition(void* ptr, int index, int row, int column, int rowSpan, int columnSpan);
int QGridLayout_HasHeightForWidth(void* ptr);
int QGridLayout_HeightForWidth(void* ptr, int w);
void QGridLayout_Invalidate(void* ptr);
void* QGridLayout_ItemAt(void* ptr, int index);
void* QGridLayout_ItemAtPosition(void* ptr, int row, int column);
int QGridLayout_MinimumHeightForWidth(void* ptr, int w);
int QGridLayout_OriginCorner(void* ptr);
int QGridLayout_RowCount(void* ptr);
int QGridLayout_RowMinimumHeight(void* ptr, int row);
int QGridLayout_RowStretch(void* ptr, int row);
void QGridLayout_SetColumnMinimumWidth(void* ptr, int column, int minSize);
void QGridLayout_SetColumnStretch(void* ptr, int column, int stretch);
void QGridLayout_SetGeometry(void* ptr, void* rect);
void QGridLayout_SetOriginCorner(void* ptr, int corner);
void QGridLayout_SetRowMinimumHeight(void* ptr, int row, int minSize);
void QGridLayout_SetRowStretch(void* ptr, int row, int stretch);
void QGridLayout_SetSpacing(void* ptr, int spacing);
int QGridLayout_Spacing(void* ptr);
void* QGridLayout_TakeAt(void* ptr, int index);
void QGridLayout_DestroyQGridLayout(void* ptr);

#ifdef __cplusplus
}
#endif