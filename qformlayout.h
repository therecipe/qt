#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QFormLayout_New_QWidget(QtObjectPtr parent);
void QFormLayout_Destroy(QtObjectPtr ptr);
void QFormLayout_AddRow_QWidget_QWidget(QtObjectPtr ptr, QtObjectPtr label, QtObjectPtr field);
void QFormLayout_AddRow_QWidget_QLayout(QtObjectPtr ptr, QtObjectPtr label, QtObjectPtr field);
void QFormLayout_AddRow_String_QWidget(QtObjectPtr ptr, char* labelText, QtObjectPtr field);
void QFormLayout_AddRow_String_QLayout(QtObjectPtr ptr, char* labelText, QtObjectPtr field);
void QFormLayout_AddRow_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QFormLayout_AddRow_QLayout(QtObjectPtr ptr, QtObjectPtr layout);
int QFormLayout_FormAlignment(QtObjectPtr ptr);
int QFormLayout_HorizontalSpacing(QtObjectPtr ptr);
void QFormLayout_InsertRow_Int_QWidget_QWidget(QtObjectPtr ptr, int row, QtObjectPtr label, QtObjectPtr field);
void QFormLayout_InsertRow_Int_QWidget_QLayout(QtObjectPtr ptr, int row, QtObjectPtr label, QtObjectPtr field);
void QFormLayout_InsertRow_Int_String_QWidget(QtObjectPtr ptr, int row, char* labelText, QtObjectPtr field);
void QFormLayout_InsertRow_Int_String_QLayout(QtObjectPtr ptr, int row, char* labelText, QtObjectPtr field);
void QFormLayout_InsertRow_Int_QWidget(QtObjectPtr ptr, int row, QtObjectPtr widget);
void QFormLayout_InsertRow_Int_QLayout(QtObjectPtr ptr, int row, QtObjectPtr layout);
int QFormLayout_LabelAlignment(QtObjectPtr ptr);
QtObjectPtr QFormLayout_LabelForField_QWidget(QtObjectPtr ptr, QtObjectPtr field);
int QFormLayout_RowCount(QtObjectPtr ptr);
void QFormLayout_SetFormAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment);
void QFormLayout_SetHorizontalSpacing_Int(QtObjectPtr ptr, int spacing);
void QFormLayout_SetLabelAlignment_AlignmentFlag(QtObjectPtr ptr, int alignment);
void QFormLayout_SetVerticalSpacing_Int(QtObjectPtr ptr, int spacing);
int QFormLayout_VerticalSpacing(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
