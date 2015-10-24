#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomNamedNodeMap_NewQDomNamedNodeMap();
QtObjectPtr QDomNamedNodeMap_NewQDomNamedNodeMap2(QtObjectPtr n);
int QDomNamedNodeMap_Contains(QtObjectPtr ptr, char* name);
int QDomNamedNodeMap_Count(QtObjectPtr ptr);
int QDomNamedNodeMap_IsEmpty(QtObjectPtr ptr);
int QDomNamedNodeMap_Length(QtObjectPtr ptr);
int QDomNamedNodeMap_Size(QtObjectPtr ptr);
void QDomNamedNodeMap_DestroyQDomNamedNodeMap(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif