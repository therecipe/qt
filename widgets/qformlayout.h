#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFormLayout_FieldGrowthPolicy(QtObjectPtr ptr);
int QFormLayout_FormAlignment(QtObjectPtr ptr);
int QFormLayout_HorizontalSpacing(QtObjectPtr ptr);
int QFormLayout_LabelAlignment(QtObjectPtr ptr);
int QFormLayout_RowWrapPolicy(QtObjectPtr ptr);
void QFormLayout_SetFieldGrowthPolicy(QtObjectPtr ptr, int policy);
void QFormLayout_SetFormAlignment(QtObjectPtr ptr, int alignment);
void QFormLayout_SetHorizontalSpacing(QtObjectPtr ptr, int spacing);
void QFormLayout_SetLabelAlignment(QtObjectPtr ptr, int alignment);
void QFormLayout_SetRowWrapPolicy(QtObjectPtr ptr, int policy);
void QFormLayout_SetVerticalSpacing(QtObjectPtr ptr, int spacing);
int QFormLayout_VerticalSpacing(QtObjectPtr ptr);
QtObjectPtr QFormLayout_NewQFormLayout(QtObjectPtr parent);
void QFormLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item);
void QFormLayout_AddRow6(QtObjectPtr ptr, QtObjectPtr layout);
void QFormLayout_AddRow2(QtObjectPtr ptr, QtObjectPtr label, QtObjectPtr field);
void QFormLayout_AddRow(QtObjectPtr ptr, QtObjectPtr label, QtObjectPtr field);
void QFormLayout_AddRow5(QtObjectPtr ptr, QtObjectPtr widget);
void QFormLayout_AddRow4(QtObjectPtr ptr, char* labelText, QtObjectPtr field);
void QFormLayout_AddRow3(QtObjectPtr ptr, char* labelText, QtObjectPtr field);
int QFormLayout_Count(QtObjectPtr ptr);
int QFormLayout_ExpandingDirections(QtObjectPtr ptr);
int QFormLayout_HasHeightForWidth(QtObjectPtr ptr);
int QFormLayout_HeightForWidth(QtObjectPtr ptr, int width);
void QFormLayout_InsertRow6(QtObjectPtr ptr, int row, QtObjectPtr layout);
void QFormLayout_InsertRow2(QtObjectPtr ptr, int row, QtObjectPtr label, QtObjectPtr field);
void QFormLayout_InsertRow(QtObjectPtr ptr, int row, QtObjectPtr label, QtObjectPtr field);
void QFormLayout_InsertRow5(QtObjectPtr ptr, int row, QtObjectPtr widget);
void QFormLayout_InsertRow4(QtObjectPtr ptr, int row, char* labelText, QtObjectPtr field);
void QFormLayout_InsertRow3(QtObjectPtr ptr, int row, char* labelText, QtObjectPtr field);
void QFormLayout_Invalidate(QtObjectPtr ptr);
QtObjectPtr QFormLayout_ItemAt2(QtObjectPtr ptr, int index);
QtObjectPtr QFormLayout_ItemAt(QtObjectPtr ptr, int row, int role);
QtObjectPtr QFormLayout_LabelForField2(QtObjectPtr ptr, QtObjectPtr field);
QtObjectPtr QFormLayout_LabelForField(QtObjectPtr ptr, QtObjectPtr field);
int QFormLayout_RowCount(QtObjectPtr ptr);
void QFormLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
void QFormLayout_SetItem(QtObjectPtr ptr, int row, int role, QtObjectPtr item);
void QFormLayout_SetLayout(QtObjectPtr ptr, int row, int role, QtObjectPtr layout);
void QFormLayout_SetSpacing(QtObjectPtr ptr, int spacing);
void QFormLayout_SetWidget(QtObjectPtr ptr, int row, int role, QtObjectPtr widget);
int QFormLayout_Spacing(QtObjectPtr ptr);
QtObjectPtr QFormLayout_TakeAt(QtObjectPtr ptr, int index);
void QFormLayout_DestroyQFormLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif