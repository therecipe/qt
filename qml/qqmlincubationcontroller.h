#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlIncubationController_NewQQmlIncubationController();
QtObjectPtr QQmlIncubationController_Engine(QtObjectPtr ptr);
void QQmlIncubationController_IncubateFor(QtObjectPtr ptr, int msecs);
int QQmlIncubationController_IncubatingObjectCount(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif