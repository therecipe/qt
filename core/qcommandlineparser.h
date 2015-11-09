#ifdef __cplusplus
extern "C" {
#endif

void* QCommandLineParser_NewQCommandLineParser();
void* QCommandLineParser_AddHelpOption(void* ptr);
int QCommandLineParser_AddOption(void* ptr, void* option);
void QCommandLineParser_AddPositionalArgument(void* ptr, char* name, char* description, char* syntax);
void* QCommandLineParser_AddVersionOption(void* ptr);
char* QCommandLineParser_ApplicationDescription(void* ptr);
void QCommandLineParser_ClearPositionalArguments(void* ptr);
char* QCommandLineParser_ErrorText(void* ptr);
char* QCommandLineParser_HelpText(void* ptr);
int QCommandLineParser_IsSet2(void* ptr, void* option);
int QCommandLineParser_IsSet(void* ptr, char* name);
char* QCommandLineParser_OptionNames(void* ptr);
int QCommandLineParser_Parse(void* ptr, char* arguments);
char* QCommandLineParser_PositionalArguments(void* ptr);
void QCommandLineParser_Process2(void* ptr, void* app);
void QCommandLineParser_Process(void* ptr, char* arguments);
void QCommandLineParser_SetApplicationDescription(void* ptr, char* description);
void QCommandLineParser_SetSingleDashWordOptionMode(void* ptr, int singleDashWordOptionMode);
void QCommandLineParser_ShowHelp(void* ptr, int exitCode);
void QCommandLineParser_ShowVersion(void* ptr);
char* QCommandLineParser_UnknownOptionNames(void* ptr);
char* QCommandLineParser_Value2(void* ptr, void* option);
char* QCommandLineParser_Value(void* ptr, char* optionName);
char* QCommandLineParser_Values2(void* ptr, void* option);
char* QCommandLineParser_Values(void* ptr, char* optionName);
void QCommandLineParser_DestroyQCommandLineParser(void* ptr);

#ifdef __cplusplus
}
#endif