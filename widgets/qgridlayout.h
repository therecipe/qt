#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGridLayout_HorizontalSpacing(QtObjectPtr ptr);
void QGridLayout_SetHorizontalSpacing(QtObjectPtr ptr, int spacing);
void QGridLayout_SetVerticalSpacing(QtObjectPtr ptr, int spacing);
int QGridLayout_VerticalSpacing(QtObjectPtr ptr);
QtObjectPtr QGridLayout_NewQGridLayout2();
QtObjectPtr QGridLayout_NewQGridLayout(QtObjectPtr parent);
void QGridLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item, int row, int column, int rowSpan, int columnSpan, int alignment);
void QGridLayout_AddLayout(QtObjectPtr ptr, QtObjectPtr layout, int row, int column, int alignment);
void QGridLayout_AddLayout2(QtObjectPtr ptr, QtObjectPtr layout, int row, int column, int rowSpan, int columnSpan, int alignment);
void QGridLayout_AddWidget2(QtObjectPtr ptr, QtObjectPtr widget, int fromRow, int fromColumn, int rowSpan, int columnSpan, int alignment);
void QGridLayout_AddWidget(QtObjectPtr ptr, QtObjectPtr widget, int row, int column, int alignment);
int QGridLayout_ColumnCount(QtObjectPtr ptr);
int QGridLayout_ColumnMinimumWidth(QtObjectPtr ptr, int column);
int QGridLayout_ColumnStretch(QtObjectPtr ptr, int column);
int QGridLayout_Count(QtObjectPtr ptr);
int QGridLayout_ExpandingDirections(QtObjectPtr ptr);
void QGridLayout_GetItemPosition(QtObjectPtr ptr, int index, int row, int column, int rowSpan, int columnSpan);
int QGridLayout_HasHeightForWidth(QtObjectPtr ptr);
int QGridLayout_HeightForWidth(QtObjectPtr ptr, int w);
void QGridLayout_Invalidate(QtObjectPtr ptr);
QtObjectPtr QGridLayout_ItemAt(QtObjectPtr ptr, int index);
QtObjectPtr QGridLayout_ItemAtPosition(QtObjectPtr ptr, int row, int column);
int QGridLayout_MinimumHeightForWidth(QtObjectPtr ptr, int w);
int QGridLayout_OriginCorner(QtObjectPtr ptr);
int QGridLayout_RowCount(QtObjectPtr ptr);
int QGridLayout_RowMinimumHeight(QtObjectPtr ptr, int row);
int QGridLayout_RowStretch(QtObjectPtr ptr, int row);
void QGridLayout_SetColumnMinimumWidth(QtObjectPtr ptr, int column, int minSize);
void QGridLayout_SetColumnStretch(QtObjectPtr ptr, int column, int stretch);
void QGridLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
void QGridLayout_SetOriginCorner(QtObjectPtr ptr, int corner);
void QGridLayout_SetRowMinimumHeight(QtObjectPtr ptr, int row, int minSize);
void QGridLayout_SetRowStretch(QtObjectPtr ptr, int row, int stretch);
void QGridLayout_SetSpacing(QtObjectPtr ptr, int spacing);
int QGridLayout_Spacing(QtObjectPtr ptr);
QtObjectPtr QGridLayout_TakeAt(QtObjectPtr ptr, int index);
void QGridLayout_DestroyQGridLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif