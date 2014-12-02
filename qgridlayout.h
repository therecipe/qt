#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QGridLayout_New_QWidget(QtObjectPtr parent);
QtObjectPtr QGridLayout_New();
void QGridLayout_Destroy(QtObjectPtr ptr);
void QGridLayout_AddItem_QLayoutItem_Int_Int_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr item, int row, int column, int rowSpan, int columnSpan, int alignment);
void QGridLayout_AddLayout_QLayout_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr layout, int row, int column, int alignment);
void QGridLayout_AddLayout_QLayout_Int_Int_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr layout, int row, int column, int rowSpan, int columnSpan, int alignment);
void QGridLayout_AddWidget_QWidget_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr widget, int row, int column, int alignment);
void QGridLayout_AddWidget_QWidget_Int_Int_Int_Int_AlignmentFlag(QtObjectPtr ptr, QtObjectPtr widget, int fromRow, int fromColumn, int rowSpan, int columnSpan, int alignment);
int QGridLayout_ColumnCount(QtObjectPtr ptr);
int QGridLayout_ColumnMinimumWidth_Int(QtObjectPtr ptr, int column);
int QGridLayout_ColumnStretch_Int(QtObjectPtr ptr, int column);
int QGridLayout_HorizontalSpacing(QtObjectPtr ptr);
int QGridLayout_OriginCorner(QtObjectPtr ptr);
int QGridLayout_RowCount(QtObjectPtr ptr);
int QGridLayout_RowMinimumHeight_Int(QtObjectPtr ptr, int row);
int QGridLayout_RowStretch_Int(QtObjectPtr ptr, int row);
void QGridLayout_SetColumnMinimumWidth_Int_Int(QtObjectPtr ptr, int column, int minSize);
void QGridLayout_SetColumnStretch_Int_Int(QtObjectPtr ptr, int column, int stretch);
void QGridLayout_SetHorizontalSpacing_Int(QtObjectPtr ptr, int spacing);
void QGridLayout_SetOriginCorner_Corner(QtObjectPtr ptr, int corner);
void QGridLayout_SetRowMinimumHeight_Int_Int(QtObjectPtr ptr, int row, int minSize);
void QGridLayout_SetRowStretch_Int_Int(QtObjectPtr ptr, int row, int stretch);
void QGridLayout_SetVerticalSpacing_Int(QtObjectPtr ptr, int spacing);
int QGridLayout_VerticalSpacing(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
