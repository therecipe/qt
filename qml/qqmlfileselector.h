#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlFileSelector_NewQQmlFileSelector(QtObjectPtr engine, QtObjectPtr parent);
QtObjectPtr QQmlFileSelector_QQmlFileSelector_Get(QtObjectPtr engine);
void QQmlFileSelector_SetExtraSelectors(QtObjectPtr ptr, char* strin);
void QQmlFileSelector_SetExtraSelectors2(QtObjectPtr ptr, char* strin);
void QQmlFileSelector_SetSelector(QtObjectPtr ptr, QtObjectPtr selector);
void QQmlFileSelector_DestroyQQmlFileSelector(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif