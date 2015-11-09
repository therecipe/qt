#ifdef __cplusplus
extern "C" {
#endif

int QFormLayout_FieldGrowthPolicy(void* ptr);
int QFormLayout_FormAlignment(void* ptr);
int QFormLayout_HorizontalSpacing(void* ptr);
int QFormLayout_LabelAlignment(void* ptr);
int QFormLayout_RowWrapPolicy(void* ptr);
void QFormLayout_SetFieldGrowthPolicy(void* ptr, int policy);
void QFormLayout_SetFormAlignment(void* ptr, int alignment);
void QFormLayout_SetHorizontalSpacing(void* ptr, int spacing);
void QFormLayout_SetLabelAlignment(void* ptr, int alignment);
void QFormLayout_SetRowWrapPolicy(void* ptr, int policy);
void QFormLayout_SetVerticalSpacing(void* ptr, int spacing);
int QFormLayout_VerticalSpacing(void* ptr);
void* QFormLayout_NewQFormLayout(void* parent);
void QFormLayout_AddItem(void* ptr, void* item);
void QFormLayout_AddRow6(void* ptr, void* layout);
void QFormLayout_AddRow2(void* ptr, void* label, void* field);
void QFormLayout_AddRow(void* ptr, void* label, void* field);
void QFormLayout_AddRow5(void* ptr, void* widget);
void QFormLayout_AddRow4(void* ptr, char* labelText, void* field);
void QFormLayout_AddRow3(void* ptr, char* labelText, void* field);
int QFormLayout_Count(void* ptr);
int QFormLayout_ExpandingDirections(void* ptr);
int QFormLayout_HasHeightForWidth(void* ptr);
int QFormLayout_HeightForWidth(void* ptr, int width);
void QFormLayout_InsertRow6(void* ptr, int row, void* layout);
void QFormLayout_InsertRow2(void* ptr, int row, void* label, void* field);
void QFormLayout_InsertRow(void* ptr, int row, void* label, void* field);
void QFormLayout_InsertRow5(void* ptr, int row, void* widget);
void QFormLayout_InsertRow4(void* ptr, int row, char* labelText, void* field);
void QFormLayout_InsertRow3(void* ptr, int row, char* labelText, void* field);
void QFormLayout_Invalidate(void* ptr);
void* QFormLayout_ItemAt2(void* ptr, int index);
void* QFormLayout_ItemAt(void* ptr, int row, int role);
void* QFormLayout_LabelForField2(void* ptr, void* field);
void* QFormLayout_LabelForField(void* ptr, void* field);
int QFormLayout_RowCount(void* ptr);
void QFormLayout_SetGeometry(void* ptr, void* rect);
void QFormLayout_SetItem(void* ptr, int row, int role, void* item);
void QFormLayout_SetLayout(void* ptr, int row, int role, void* layout);
void QFormLayout_SetSpacing(void* ptr, int spacing);
void QFormLayout_SetWidget(void* ptr, int row, int role, void* widget);
int QFormLayout_Spacing(void* ptr);
void* QFormLayout_TakeAt(void* ptr, int index);
void QFormLayout_DestroyQFormLayout(void* ptr);

#ifdef __cplusplus
}
#endif