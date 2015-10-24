#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSqlField_NewQSqlField2(QtObjectPtr other);
void QSqlField_Clear(QtObjectPtr ptr);
char* QSqlField_DefaultValue(QtObjectPtr ptr);
int QSqlField_IsAutoValue(QtObjectPtr ptr);
int QSqlField_IsGenerated(QtObjectPtr ptr);
int QSqlField_IsNull(QtObjectPtr ptr);
int QSqlField_IsReadOnly(QtObjectPtr ptr);
int QSqlField_IsValid(QtObjectPtr ptr);
int QSqlField_Length(QtObjectPtr ptr);
char* QSqlField_Name(QtObjectPtr ptr);
int QSqlField_Precision(QtObjectPtr ptr);
int QSqlField_RequiredStatus(QtObjectPtr ptr);
void QSqlField_SetAutoValue(QtObjectPtr ptr, int autoVal);
void QSqlField_SetDefaultValue(QtObjectPtr ptr, char* value);
void QSqlField_SetGenerated(QtObjectPtr ptr, int gen);
void QSqlField_SetLength(QtObjectPtr ptr, int fieldLength);
void QSqlField_SetName(QtObjectPtr ptr, char* name);
void QSqlField_SetPrecision(QtObjectPtr ptr, int precision);
void QSqlField_SetReadOnly(QtObjectPtr ptr, int readOnly);
void QSqlField_SetRequired(QtObjectPtr ptr, int required);
void QSqlField_SetRequiredStatus(QtObjectPtr ptr, int required);
void QSqlField_SetValue(QtObjectPtr ptr, char* value);
char* QSqlField_Value(QtObjectPtr ptr);
void QSqlField_DestroyQSqlField(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif