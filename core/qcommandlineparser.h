#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QCommandLineParser_NewQCommandLineParser();
int QCommandLineParser_AddOption(QtObjectPtr ptr, QtObjectPtr option);
void QCommandLineParser_AddPositionalArgument(QtObjectPtr ptr, char* name, char* description, char* syntax);
char* QCommandLineParser_ApplicationDescription(QtObjectPtr ptr);
void QCommandLineParser_ClearPositionalArguments(QtObjectPtr ptr);
char* QCommandLineParser_ErrorText(QtObjectPtr ptr);
char* QCommandLineParser_HelpText(QtObjectPtr ptr);
int QCommandLineParser_IsSet2(QtObjectPtr ptr, QtObjectPtr option);
int QCommandLineParser_IsSet(QtObjectPtr ptr, char* name);
char* QCommandLineParser_OptionNames(QtObjectPtr ptr);
int QCommandLineParser_Parse(QtObjectPtr ptr, char* arguments);
char* QCommandLineParser_PositionalArguments(QtObjectPtr ptr);
void QCommandLineParser_Process2(QtObjectPtr ptr, QtObjectPtr app);
void QCommandLineParser_Process(QtObjectPtr ptr, char* arguments);
void QCommandLineParser_SetApplicationDescription(QtObjectPtr ptr, char* description);
void QCommandLineParser_SetSingleDashWordOptionMode(QtObjectPtr ptr, int singleDashWordOptionMode);
void QCommandLineParser_ShowHelp(QtObjectPtr ptr, int exitCode);
void QCommandLineParser_ShowVersion(QtObjectPtr ptr);
char* QCommandLineParser_UnknownOptionNames(QtObjectPtr ptr);
char* QCommandLineParser_Value2(QtObjectPtr ptr, QtObjectPtr option);
char* QCommandLineParser_Value(QtObjectPtr ptr, char* optionName);
char* QCommandLineParser_Values2(QtObjectPtr ptr, QtObjectPtr option);
char* QCommandLineParser_Values(QtObjectPtr ptr, char* optionName);
void QCommandLineParser_DestroyQCommandLineParser(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif