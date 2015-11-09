#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsGridLayout_NewQGraphicsGridLayout(void* parent);
void QGraphicsGridLayout_AddItem2(void* ptr, void* item, int row, int column, int alignment);
void QGraphicsGridLayout_AddItem(void* ptr, void* item, int row, int column, int rowSpan, int columnSpan, int alignment);
int QGraphicsGridLayout_Alignment(void* ptr, void* item);
int QGraphicsGridLayout_ColumnAlignment(void* ptr, int column);
int QGraphicsGridLayout_ColumnCount(void* ptr);
double QGraphicsGridLayout_ColumnMaximumWidth(void* ptr, int column);
double QGraphicsGridLayout_ColumnMinimumWidth(void* ptr, int column);
double QGraphicsGridLayout_ColumnPreferredWidth(void* ptr, int column);
double QGraphicsGridLayout_ColumnSpacing(void* ptr, int column);
int QGraphicsGridLayout_ColumnStretchFactor(void* ptr, int column);
int QGraphicsGridLayout_Count(void* ptr);
double QGraphicsGridLayout_HorizontalSpacing(void* ptr);
void QGraphicsGridLayout_Invalidate(void* ptr);
void* QGraphicsGridLayout_ItemAt2(void* ptr, int index);
void* QGraphicsGridLayout_ItemAt(void* ptr, int row, int column);
void QGraphicsGridLayout_RemoveAt(void* ptr, int index);
void QGraphicsGridLayout_RemoveItem(void* ptr, void* item);
int QGraphicsGridLayout_RowAlignment(void* ptr, int row);
int QGraphicsGridLayout_RowCount(void* ptr);
double QGraphicsGridLayout_RowMaximumHeight(void* ptr, int row);
double QGraphicsGridLayout_RowMinimumHeight(void* ptr, int row);
double QGraphicsGridLayout_RowPreferredHeight(void* ptr, int row);
double QGraphicsGridLayout_RowSpacing(void* ptr, int row);
int QGraphicsGridLayout_RowStretchFactor(void* ptr, int row);
void QGraphicsGridLayout_SetAlignment(void* ptr, void* item, int alignment);
void QGraphicsGridLayout_SetColumnAlignment(void* ptr, int column, int alignment);
void QGraphicsGridLayout_SetColumnFixedWidth(void* ptr, int column, double width);
void QGraphicsGridLayout_SetColumnMaximumWidth(void* ptr, int column, double width);
void QGraphicsGridLayout_SetColumnMinimumWidth(void* ptr, int column, double width);
void QGraphicsGridLayout_SetColumnPreferredWidth(void* ptr, int column, double width);
void QGraphicsGridLayout_SetColumnSpacing(void* ptr, int column, double spacing);
void QGraphicsGridLayout_SetColumnStretchFactor(void* ptr, int column, int stretch);
void QGraphicsGridLayout_SetGeometry(void* ptr, void* rect);
void QGraphicsGridLayout_SetHorizontalSpacing(void* ptr, double spacing);
void QGraphicsGridLayout_SetRowAlignment(void* ptr, int row, int alignment);
void QGraphicsGridLayout_SetRowFixedHeight(void* ptr, int row, double height);
void QGraphicsGridLayout_SetRowMaximumHeight(void* ptr, int row, double height);
void QGraphicsGridLayout_SetRowMinimumHeight(void* ptr, int row, double height);
void QGraphicsGridLayout_SetRowPreferredHeight(void* ptr, int row, double height);
void QGraphicsGridLayout_SetRowSpacing(void* ptr, int row, double spacing);
void QGraphicsGridLayout_SetRowStretchFactor(void* ptr, int row, int stretch);
void QGraphicsGridLayout_SetSpacing(void* ptr, double spacing);
void QGraphicsGridLayout_SetVerticalSpacing(void* ptr, double spacing);
double QGraphicsGridLayout_VerticalSpacing(void* ptr);
void QGraphicsGridLayout_DestroyQGraphicsGridLayout(void* ptr);

#ifdef __cplusplus
}
#endif