#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QByteArrayMatcher_NewQByteArrayMatcher();
QtObjectPtr QByteArrayMatcher_NewQByteArrayMatcher2(QtObjectPtr pattern);
QtObjectPtr QByteArrayMatcher_NewQByteArrayMatcher4(QtObjectPtr other);
QtObjectPtr QByteArrayMatcher_NewQByteArrayMatcher3(char* pattern, int length);
int QByteArrayMatcher_IndexIn(QtObjectPtr ptr, QtObjectPtr ba, int from);
int QByteArrayMatcher_IndexIn2(QtObjectPtr ptr, char* str, int len, int from);
void QByteArrayMatcher_SetPattern(QtObjectPtr ptr, QtObjectPtr pattern);
void QByteArrayMatcher_DestroyQByteArrayMatcher(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif