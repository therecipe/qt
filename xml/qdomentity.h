#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomEntity_NewQDomEntity();
QtObjectPtr QDomEntity_NewQDomEntity2(QtObjectPtr x);
int QDomEntity_NodeType(QtObjectPtr ptr);
char* QDomEntity_NotationName(QtObjectPtr ptr);
char* QDomEntity_PublicId(QtObjectPtr ptr);
char* QDomEntity_SystemId(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif