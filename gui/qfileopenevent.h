#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFileOpenEvent_OpenFile(QtObjectPtr ptr, QtObjectPtr file, int flags);
char* QFileOpenEvent_File(QtObjectPtr ptr);
char* QFileOpenEvent_Url(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif