#ifdef __cplusplus
extern "C" {
#endif

int QTextCodec_CanEncode(void* ptr, void* ch);
int QTextCodec_CanEncode2(void* ptr, char* s);
void* QTextCodec_QTextCodec_CodecForHtml2(void* ba);
void* QTextCodec_QTextCodec_CodecForHtml(void* ba, void* defaultCodec);
void* QTextCodec_QTextCodec_CodecForLocale();
void* QTextCodec_QTextCodec_CodecForMib(int mib);
void* QTextCodec_QTextCodec_CodecForName(void* name);
void* QTextCodec_QTextCodec_CodecForName2(char* name);
void* QTextCodec_QTextCodec_CodecForUtfText2(void* ba);
void* QTextCodec_QTextCodec_CodecForUtfText(void* ba, void* defaultCodec);
void* QTextCodec_FromUnicode(void* ptr, char* str);
void* QTextCodec_MakeDecoder(void* ptr, int flags);
void* QTextCodec_MakeEncoder(void* ptr, int flags);
int QTextCodec_MibEnum(void* ptr);
void* QTextCodec_Name(void* ptr);
void QTextCodec_QTextCodec_SetCodecForLocale(void* c);

#ifdef __cplusplus
}
#endif