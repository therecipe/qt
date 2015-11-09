#ifdef __cplusplus
extern "C" {
#endif

void* QByteArrayMatcher_NewQByteArrayMatcher();
void* QByteArrayMatcher_NewQByteArrayMatcher2(void* pattern);
void* QByteArrayMatcher_NewQByteArrayMatcher4(void* other);
void* QByteArrayMatcher_NewQByteArrayMatcher3(char* pattern, int length);
int QByteArrayMatcher_IndexIn(void* ptr, void* ba, int from);
int QByteArrayMatcher_IndexIn2(void* ptr, char* str, int len, int from);
void* QByteArrayMatcher_Pattern(void* ptr);
void QByteArrayMatcher_SetPattern(void* ptr, void* pattern);
void QByteArrayMatcher_DestroyQByteArrayMatcher(void* ptr);

#ifdef __cplusplus
}
#endif