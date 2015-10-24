#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMimeDatabase_NewQMimeDatabase();
void QMimeDatabase_DestroyQMimeDatabase(QtObjectPtr ptr);
char* QMimeDatabase_SuffixForFileName(QtObjectPtr ptr, char* fileName);

#ifdef __cplusplus
}
#endif