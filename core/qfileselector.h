#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFileSelector_NewQFileSelector(QtObjectPtr parent);
char* QFileSelector_AllSelectors(QtObjectPtr ptr);
char* QFileSelector_ExtraSelectors(QtObjectPtr ptr);
char* QFileSelector_Select(QtObjectPtr ptr, char* filePath);
char* QFileSelector_Select2(QtObjectPtr ptr, char* filePath);
void QFileSelector_SetExtraSelectors(QtObjectPtr ptr, char* list);
void QFileSelector_DestroyQFileSelector(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif