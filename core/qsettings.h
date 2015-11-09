#ifdef __cplusplus
extern "C" {
#endif

void* QSettings_NewQSettings3(int format, int scope, char* organization, char* application, void* parent);
void* QSettings_NewQSettings5(void* parent);
void* QSettings_NewQSettings2(int scope, char* organization, char* application, void* parent);
void* QSettings_NewQSettings4(char* fileName, int format, void* parent);
void* QSettings_NewQSettings(char* organization, char* application, void* parent);
char* QSettings_AllKeys(void* ptr);
char* QSettings_ApplicationName(void* ptr);
void QSettings_BeginGroup(void* ptr, char* prefix);
int QSettings_BeginReadArray(void* ptr, char* prefix);
void QSettings_BeginWriteArray(void* ptr, char* prefix, int size);
char* QSettings_ChildGroups(void* ptr);
char* QSettings_ChildKeys(void* ptr);
void QSettings_Clear(void* ptr);
int QSettings_Contains(void* ptr, char* key);
int QSettings_QSettings_DefaultFormat();
void QSettings_EndArray(void* ptr);
void QSettings_EndGroup(void* ptr);
int QSettings_FallbacksEnabled(void* ptr);
char* QSettings_FileName(void* ptr);
int QSettings_Format(void* ptr);
char* QSettings_Group(void* ptr);
void* QSettings_IniCodec(void* ptr);
int QSettings_IsWritable(void* ptr);
char* QSettings_OrganizationName(void* ptr);
int QSettings_Scope(void* ptr);
void QSettings_SetArrayIndex(void* ptr, int i);
void QSettings_QSettings_SetDefaultFormat(int format);
void QSettings_SetFallbacksEnabled(void* ptr, int b);
void QSettings_SetIniCodec(void* ptr, void* codec);
void QSettings_SetIniCodec2(void* ptr, char* codecName);
void QSettings_QSettings_SetPath(int format, int scope, char* path);
void QSettings_SetValue(void* ptr, char* key, void* value);
int QSettings_Status(void* ptr);
void QSettings_Sync(void* ptr);
void* QSettings_Value(void* ptr, char* key, void* defaultValue);
void QSettings_DestroyQSettings(void* ptr);

#ifdef __cplusplus
}
#endif