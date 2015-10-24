#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QHelpIndexWidget_ActivateCurrentItem(QtObjectPtr ptr);
void QHelpIndexWidget_FilterIndices(QtObjectPtr ptr, char* filter, char* wildcard);
void QHelpIndexWidget_ConnectLinkActivated(QtObjectPtr ptr);
void QHelpIndexWidget_DisconnectLinkActivated(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif