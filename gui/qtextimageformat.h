#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextImageFormat_NewQTextImageFormat();
int QTextImageFormat_IsValid(QtObjectPtr ptr);
char* QTextImageFormat_Name(QtObjectPtr ptr);
void QTextImageFormat_SetName(QtObjectPtr ptr, char* name);

#ifdef __cplusplus
}
#endif