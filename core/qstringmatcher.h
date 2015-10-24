#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStringMatcher_NewQStringMatcher3(QtObjectPtr uc, int length, int cs);
char* QStringMatcher_Pattern(QtObjectPtr ptr);
QtObjectPtr QStringMatcher_NewQStringMatcher();
QtObjectPtr QStringMatcher_NewQStringMatcher2(char* pattern, int cs);
QtObjectPtr QStringMatcher_NewQStringMatcher4(QtObjectPtr other);
int QStringMatcher_CaseSensitivity(QtObjectPtr ptr);
int QStringMatcher_IndexIn2(QtObjectPtr ptr, QtObjectPtr str, int length, int from);
int QStringMatcher_IndexIn(QtObjectPtr ptr, char* str, int from);
void QStringMatcher_SetCaseSensitivity(QtObjectPtr ptr, int cs);
void QStringMatcher_SetPattern(QtObjectPtr ptr, char* pattern);
void QStringMatcher_DestroyQStringMatcher(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif