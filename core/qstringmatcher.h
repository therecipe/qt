#ifdef __cplusplus
extern "C" {
#endif

void* QStringMatcher_NewQStringMatcher3(void* uc, int length, int cs);
char* QStringMatcher_Pattern(void* ptr);
void* QStringMatcher_NewQStringMatcher();
void* QStringMatcher_NewQStringMatcher2(char* pattern, int cs);
void* QStringMatcher_NewQStringMatcher4(void* other);
int QStringMatcher_CaseSensitivity(void* ptr);
int QStringMatcher_IndexIn2(void* ptr, void* str, int length, int from);
int QStringMatcher_IndexIn(void* ptr, char* str, int from);
void QStringMatcher_SetCaseSensitivity(void* ptr, int cs);
void QStringMatcher_SetPattern(void* ptr, char* pattern);
void QStringMatcher_DestroyQStringMatcher(void* ptr);

#ifdef __cplusplus
}
#endif