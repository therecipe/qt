#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGraphicsGridLayout_NewQGraphicsGridLayout(QtObjectPtr parent);
void QGraphicsGridLayout_AddItem2(QtObjectPtr ptr, QtObjectPtr item, int row, int column, int alignment);
void QGraphicsGridLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item, int row, int column, int rowSpan, int columnSpan, int alignment);
int QGraphicsGridLayout_Alignment(QtObjectPtr ptr, QtObjectPtr item);
int QGraphicsGridLayout_ColumnAlignment(QtObjectPtr ptr, int column);
int QGraphicsGridLayout_ColumnCount(QtObjectPtr ptr);
int QGraphicsGridLayout_ColumnStretchFactor(QtObjectPtr ptr, int column);
int QGraphicsGridLayout_Count(QtObjectPtr ptr);
void QGraphicsGridLayout_Invalidate(QtObjectPtr ptr);
QtObjectPtr QGraphicsGridLayout_ItemAt2(QtObjectPtr ptr, int index);
QtObjectPtr QGraphicsGridLayout_ItemAt(QtObjectPtr ptr, int row, int column);
void QGraphicsGridLayout_RemoveAt(QtObjectPtr ptr, int index);
void QGraphicsGridLayout_RemoveItem(QtObjectPtr ptr, QtObjectPtr item);
int QGraphicsGridLayout_RowAlignment(QtObjectPtr ptr, int row);
int QGraphicsGridLayout_RowCount(QtObjectPtr ptr);
int QGraphicsGridLayout_RowStretchFactor(QtObjectPtr ptr, int row);
void QGraphicsGridLayout_SetAlignment(QtObjectPtr ptr, QtObjectPtr item, int alignment);
void QGraphicsGridLayout_SetColumnAlignment(QtObjectPtr ptr, int column, int alignment);
void QGraphicsGridLayout_SetColumnStretchFactor(QtObjectPtr ptr, int column, int stretch);
void QGraphicsGridLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
void QGraphicsGridLayout_SetRowAlignment(QtObjectPtr ptr, int row, int alignment);
void QGraphicsGridLayout_SetRowStretchFactor(QtObjectPtr ptr, int row, int stretch);
void QGraphicsGridLayout_DestroyQGraphicsGridLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif