#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QCommandLineOption_NewQCommandLineOption5(QtObjectPtr other);
QtObjectPtr QCommandLineOption_NewQCommandLineOption(char* name);
QtObjectPtr QCommandLineOption_NewQCommandLineOption3(char* name, char* description, char* valueName, char* defaultValue);
QtObjectPtr QCommandLineOption_NewQCommandLineOption2(char* names);
QtObjectPtr QCommandLineOption_NewQCommandLineOption4(char* names, char* description, char* valueName, char* defaultValue);
char* QCommandLineOption_DefaultValues(QtObjectPtr ptr);
char* QCommandLineOption_Description(QtObjectPtr ptr);
char* QCommandLineOption_Names(QtObjectPtr ptr);
void QCommandLineOption_SetDefaultValue(QtObjectPtr ptr, char* defaultValue);
void QCommandLineOption_SetDefaultValues(QtObjectPtr ptr, char* defaultValues);
void QCommandLineOption_SetDescription(QtObjectPtr ptr, char* description);
void QCommandLineOption_SetValueName(QtObjectPtr ptr, char* valueName);
void QCommandLineOption_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QCommandLineOption_ValueName(QtObjectPtr ptr);
void QCommandLineOption_DestroyQCommandLineOption(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif