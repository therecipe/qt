#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QInputMethodQueryEvent_NewQInputMethodQueryEvent(int queries);
int QInputMethodQueryEvent_Queries(QtObjectPtr ptr);
void QInputMethodQueryEvent_SetValue(QtObjectPtr ptr, int query, char* value);
char* QInputMethodQueryEvent_Value(QtObjectPtr ptr, int query);

#ifdef __cplusplus
}
#endif