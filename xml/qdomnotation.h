#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomNotation_NewQDomNotation();
QtObjectPtr QDomNotation_NewQDomNotation2(QtObjectPtr x);
int QDomNotation_NodeType(QtObjectPtr ptr);
char* QDomNotation_PublicId(QtObjectPtr ptr);
char* QDomNotation_SystemId(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif