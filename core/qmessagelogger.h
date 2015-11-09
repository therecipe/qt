#ifdef __cplusplus
extern "C" {
#endif

void* QMessageLogger_NewQMessageLogger();
void* QMessageLogger_NewQMessageLogger2(char* file, int line, char* function);
void* QMessageLogger_NewQMessageLogger3(char* file, int line, char* function, char* category);

#ifdef __cplusplus
}
#endif