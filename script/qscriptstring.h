#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScriptString_NewQScriptString();
QtObjectPtr QScriptString_NewQScriptString2(QtObjectPtr other);
int QScriptString_IsValid(QtObjectPtr ptr);
char* QScriptString_ToString(QtObjectPtr ptr);
void QScriptString_DestroyQScriptString(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif