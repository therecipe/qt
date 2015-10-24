#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QMacPasteboardMime_ConvertorName(QtObjectPtr ptr);
int QMacPasteboardMime_Count(QtObjectPtr ptr, QtObjectPtr mimeData);
char* QMacPasteboardMime_FlavorFor(QtObjectPtr ptr, char* mime);
char* QMacPasteboardMime_MimeFor(QtObjectPtr ptr, char* flav);
void QMacPasteboardMime_DestroyQMacPasteboardMime(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif