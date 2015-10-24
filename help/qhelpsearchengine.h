#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHelpSearchEngine_NewQHelpSearchEngine(QtObjectPtr helpEngine, QtObjectPtr parent);
void QHelpSearchEngine_CancelIndexing(QtObjectPtr ptr);
void QHelpSearchEngine_CancelSearching(QtObjectPtr ptr);
int QHelpSearchEngine_HitCount(QtObjectPtr ptr);
void QHelpSearchEngine_ConnectIndexingFinished(QtObjectPtr ptr);
void QHelpSearchEngine_DisconnectIndexingFinished(QtObjectPtr ptr);
void QHelpSearchEngine_ConnectIndexingStarted(QtObjectPtr ptr);
void QHelpSearchEngine_DisconnectIndexingStarted(QtObjectPtr ptr);
QtObjectPtr QHelpSearchEngine_QueryWidget(QtObjectPtr ptr);
void QHelpSearchEngine_ReindexDocumentation(QtObjectPtr ptr);
QtObjectPtr QHelpSearchEngine_ResultWidget(QtObjectPtr ptr);
void QHelpSearchEngine_ConnectSearchingFinished(QtObjectPtr ptr);
void QHelpSearchEngine_DisconnectSearchingFinished(QtObjectPtr ptr);
void QHelpSearchEngine_ConnectSearchingStarted(QtObjectPtr ptr);
void QHelpSearchEngine_DisconnectSearchingStarted(QtObjectPtr ptr);
void QHelpSearchEngine_DestroyQHelpSearchEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif