#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptClass_NewQScriptClass(QtObjectPtr engine);
QtObjectPtr QScriptClass_Engine(QtObjectPtr ptr);
char* QScriptClass_Extension(QtObjectPtr ptr, int extension, char* argument);
char* QScriptClass_Name(QtObjectPtr ptr);
QtObjectPtr QScriptClass_NewIterator(QtObjectPtr ptr, QtObjectPtr object);
int QScriptClass_SupportsExtension(QtObjectPtr ptr, int extension);
void QScriptClass_DestroyQScriptClass(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif