#ifdef __cplusplus
extern "C" {
#endif

void* QSqlField_NewQSqlField2(void* other);
void QSqlField_Clear(void* ptr);
void* QSqlField_DefaultValue(void* ptr);
int QSqlField_IsAutoValue(void* ptr);
int QSqlField_IsGenerated(void* ptr);
int QSqlField_IsNull(void* ptr);
int QSqlField_IsReadOnly(void* ptr);
int QSqlField_IsValid(void* ptr);
int QSqlField_Length(void* ptr);
char* QSqlField_Name(void* ptr);
int QSqlField_Precision(void* ptr);
int QSqlField_RequiredStatus(void* ptr);
void QSqlField_SetAutoValue(void* ptr, int autoVal);
void QSqlField_SetDefaultValue(void* ptr, void* value);
void QSqlField_SetGenerated(void* ptr, int gen);
void QSqlField_SetLength(void* ptr, int fieldLength);
void QSqlField_SetName(void* ptr, char* name);
void QSqlField_SetPrecision(void* ptr, int precision);
void QSqlField_SetReadOnly(void* ptr, int readOnly);
void QSqlField_SetRequired(void* ptr, int required);
void QSqlField_SetRequiredStatus(void* ptr, int required);
void QSqlField_SetValue(void* ptr, void* value);
void* QSqlField_Value(void* ptr);
void QSqlField_DestroyQSqlField(void* ptr);

#ifdef __cplusplus
}
#endif