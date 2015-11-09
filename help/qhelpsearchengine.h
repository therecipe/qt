#ifdef __cplusplus
extern "C" {
#endif

void* QHelpSearchEngine_NewQHelpSearchEngine(void* helpEngine, void* parent);
void QHelpSearchEngine_CancelIndexing(void* ptr);
void QHelpSearchEngine_CancelSearching(void* ptr);
int QHelpSearchEngine_HitCount(void* ptr);
void QHelpSearchEngine_ConnectIndexingFinished(void* ptr);
void QHelpSearchEngine_DisconnectIndexingFinished(void* ptr);
void QHelpSearchEngine_ConnectIndexingStarted(void* ptr);
void QHelpSearchEngine_DisconnectIndexingStarted(void* ptr);
void* QHelpSearchEngine_QueryWidget(void* ptr);
void QHelpSearchEngine_ReindexDocumentation(void* ptr);
void* QHelpSearchEngine_ResultWidget(void* ptr);
void QHelpSearchEngine_ConnectSearchingFinished(void* ptr);
void QHelpSearchEngine_DisconnectSearchingFinished(void* ptr);
void QHelpSearchEngine_ConnectSearchingStarted(void* ptr);
void QHelpSearchEngine_DisconnectSearchingStarted(void* ptr);
void QHelpSearchEngine_DestroyQHelpSearchEngine(void* ptr);

#ifdef __cplusplus
}
#endif