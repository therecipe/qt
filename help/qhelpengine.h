#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHelpEngine_NewQHelpEngine(char* collectionFile, QtObjectPtr parent);
QtObjectPtr QHelpEngine_ContentModel(QtObjectPtr ptr);
QtObjectPtr QHelpEngine_ContentWidget(QtObjectPtr ptr);
QtObjectPtr QHelpEngine_IndexModel(QtObjectPtr ptr);
QtObjectPtr QHelpEngine_IndexWidget(QtObjectPtr ptr);
QtObjectPtr QHelpEngine_SearchEngine(QtObjectPtr ptr);
void QHelpEngine_DestroyQHelpEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif