#ifdef __cplusplus
extern "C" {
#endif

void* QTextEncoder_NewQTextEncoder(void* codec);
void* QTextEncoder_NewQTextEncoder2(void* codec, int flags);
void* QTextEncoder_FromUnicode2(void* ptr, void* uc, int len);
void* QTextEncoder_FromUnicode(void* ptr, char* str);
void QTextEncoder_DestroyQTextEncoder(void* ptr);

#ifdef __cplusplus
}
#endif