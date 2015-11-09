#ifdef __cplusplus
extern "C" {
#endif

void* QTextDecoder_NewQTextDecoder(void* codec);
void* QTextDecoder_NewQTextDecoder2(void* codec, int flags);
void QTextDecoder_DestroyQTextDecoder(void* ptr);

#ifdef __cplusplus
}
#endif