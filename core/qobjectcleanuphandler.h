#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QObjectCleanupHandler_NewQObjectCleanupHandler();
QtObjectPtr QObjectCleanupHandler_Add(QtObjectPtr ptr, QtObjectPtr object);
void QObjectCleanupHandler_Clear(QtObjectPtr ptr);
int QObjectCleanupHandler_IsEmpty(QtObjectPtr ptr);
void QObjectCleanupHandler_DestroyQObjectCleanupHandler(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif