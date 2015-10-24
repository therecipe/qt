#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMessageLogger_NewQMessageLogger();
QtObjectPtr QMessageLogger_NewQMessageLogger2(char* file, int line, char* function);
QtObjectPtr QMessageLogger_NewQMessageLogger3(char* file, int line, char* function, char* category);

#ifdef __cplusplus
}
#endif