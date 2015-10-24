#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QJSEngine_NewQJSEngine();
QtObjectPtr QJSEngine_NewQJSEngine2(QtObjectPtr parent);
void QJSEngine_CollectGarbage(QtObjectPtr ptr);
void QJSEngine_InstallTranslatorFunctions(QtObjectPtr ptr, QtObjectPtr object);
void QJSEngine_DestroyQJSEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif