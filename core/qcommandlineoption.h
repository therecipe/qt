#ifdef __cplusplus
extern "C" {
#endif

void* QCommandLineOption_NewQCommandLineOption5(void* other);
void* QCommandLineOption_NewQCommandLineOption(char* name);
void* QCommandLineOption_NewQCommandLineOption3(char* name, char* description, char* valueName, char* defaultValue);
void* QCommandLineOption_NewQCommandLineOption2(char* names);
void* QCommandLineOption_NewQCommandLineOption4(char* names, char* description, char* valueName, char* defaultValue);
char* QCommandLineOption_DefaultValues(void* ptr);
char* QCommandLineOption_Description(void* ptr);
char* QCommandLineOption_Names(void* ptr);
void QCommandLineOption_SetDefaultValue(void* ptr, char* defaultValue);
void QCommandLineOption_SetDefaultValues(void* ptr, char* defaultValues);
void QCommandLineOption_SetDescription(void* ptr, char* description);
void QCommandLineOption_SetValueName(void* ptr, char* valueName);
void QCommandLineOption_Swap(void* ptr, void* other);
char* QCommandLineOption_ValueName(void* ptr);
void QCommandLineOption_DestroyQCommandLineOption(void* ptr);

#ifdef __cplusplus
}
#endif