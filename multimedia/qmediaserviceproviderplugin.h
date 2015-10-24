#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMediaServiceProviderPlugin_Create(QtObjectPtr ptr, char* key);
void QMediaServiceProviderPlugin_Release(QtObjectPtr ptr, QtObjectPtr service);

#ifdef __cplusplus
}
#endif