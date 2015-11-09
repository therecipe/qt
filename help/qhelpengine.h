#ifdef __cplusplus
extern "C" {
#endif

void* QHelpEngine_NewQHelpEngine(char* collectionFile, void* parent);
void* QHelpEngine_ContentModel(void* ptr);
void* QHelpEngine_ContentWidget(void* ptr);
void* QHelpEngine_IndexModel(void* ptr);
void* QHelpEngine_IndexWidget(void* ptr);
void* QHelpEngine_SearchEngine(void* ptr);
void QHelpEngine_DestroyQHelpEngine(void* ptr);

#ifdef __cplusplus
}
#endif