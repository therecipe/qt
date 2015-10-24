#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QHelpIndexModel_CreateIndex(QtObjectPtr ptr, char* customFilterName);
QtObjectPtr QHelpIndexModel_Filter(QtObjectPtr ptr, char* filter, char* wildcard);
void QHelpIndexModel_ConnectIndexCreated(QtObjectPtr ptr);
void QHelpIndexModel_DisconnectIndexCreated(QtObjectPtr ptr);
void QHelpIndexModel_ConnectIndexCreationStarted(QtObjectPtr ptr);
void QHelpIndexModel_DisconnectIndexCreationStarted(QtObjectPtr ptr);
int QHelpIndexModel_IsCreatingIndex(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif