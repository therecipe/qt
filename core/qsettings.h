#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSettings_NewQSettings3(int format, int scope, char* organization, char* application, QtObjectPtr parent);
QtObjectPtr QSettings_NewQSettings5(QtObjectPtr parent);
QtObjectPtr QSettings_NewQSettings2(int scope, char* organization, char* application, QtObjectPtr parent);
QtObjectPtr QSettings_NewQSettings4(char* fileName, int format, QtObjectPtr parent);
QtObjectPtr QSettings_NewQSettings(char* organization, char* application, QtObjectPtr parent);
char* QSettings_AllKeys(QtObjectPtr ptr);
char* QSettings_ApplicationName(QtObjectPtr ptr);
void QSettings_BeginGroup(QtObjectPtr ptr, char* prefix);
int QSettings_BeginReadArray(QtObjectPtr ptr, char* prefix);
void QSettings_BeginWriteArray(QtObjectPtr ptr, char* prefix, int size);
char* QSettings_ChildGroups(QtObjectPtr ptr);
char* QSettings_ChildKeys(QtObjectPtr ptr);
void QSettings_Clear(QtObjectPtr ptr);
int QSettings_Contains(QtObjectPtr ptr, char* key);
int QSettings_QSettings_DefaultFormat();
void QSettings_EndArray(QtObjectPtr ptr);
void QSettings_EndGroup(QtObjectPtr ptr);
int QSettings_FallbacksEnabled(QtObjectPtr ptr);
char* QSettings_FileName(QtObjectPtr ptr);
int QSettings_Format(QtObjectPtr ptr);
char* QSettings_Group(QtObjectPtr ptr);
QtObjectPtr QSettings_IniCodec(QtObjectPtr ptr);
int QSettings_IsWritable(QtObjectPtr ptr);
char* QSettings_OrganizationName(QtObjectPtr ptr);
int QSettings_Scope(QtObjectPtr ptr);
void QSettings_SetArrayIndex(QtObjectPtr ptr, int i);
void QSettings_QSettings_SetDefaultFormat(int format);
void QSettings_SetFallbacksEnabled(QtObjectPtr ptr, int b);
void QSettings_SetIniCodec(QtObjectPtr ptr, QtObjectPtr codec);
void QSettings_SetIniCodec2(QtObjectPtr ptr, char* codecName);
void QSettings_QSettings_SetPath(int format, int scope, char* path);
void QSettings_SetValue(QtObjectPtr ptr, char* key, char* value);
int QSettings_Status(QtObjectPtr ptr);
void QSettings_Sync(QtObjectPtr ptr);
char* QSettings_Value(QtObjectPtr ptr, char* key, char* defaultValue);
void QSettings_DestroyQSettings(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif