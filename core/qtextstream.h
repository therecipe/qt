#ifdef __cplusplus
extern "C" {
#endif

int QTextStream_AtEnd(void* ptr);
int QTextStream_AutoDetectUnicode(void* ptr);
void* QTextStream_Codec(void* ptr);
void* QTextStream_Device(void* ptr);
int QTextStream_FieldAlignment(void* ptr);
int QTextStream_FieldWidth(void* ptr);
void QTextStream_Flush(void* ptr);
int QTextStream_GenerateByteOrderMark(void* ptr);
int QTextStream_IntegerBase(void* ptr);
int QTextStream_NumberFlags(void* ptr);
char* QTextStream_ReadAll(void* ptr);
int QTextStream_RealNumberNotation(void* ptr);
int QTextStream_RealNumberPrecision(void* ptr);
void QTextStream_Reset(void* ptr);
void QTextStream_ResetStatus(void* ptr);
void QTextStream_SetAutoDetectUnicode(void* ptr, int enabled);
void QTextStream_SetCodec(void* ptr, void* codec);
void QTextStream_SetCodec2(void* ptr, char* codecName);
void QTextStream_SetDevice(void* ptr, void* device);
void QTextStream_SetFieldAlignment(void* ptr, int mode);
void QTextStream_SetFieldWidth(void* ptr, int width);
void QTextStream_SetGenerateByteOrderMark(void* ptr, int generate);
void QTextStream_SetIntegerBase(void* ptr, int base);
void QTextStream_SetLocale(void* ptr, void* locale);
void QTextStream_SetNumberFlags(void* ptr, int flags);
void QTextStream_SetPadChar(void* ptr, void* ch);
void QTextStream_SetRealNumberNotation(void* ptr, int notation);
void QTextStream_SetRealNumberPrecision(void* ptr, int precision);
void QTextStream_SetStatus(void* ptr, int status);
void QTextStream_SetString(void* ptr, char* stri, int openMode);
void QTextStream_SkipWhiteSpace(void* ptr);
int QTextStream_Status(void* ptr);
char* QTextStream_String(void* ptr);
void QTextStream_DestroyQTextStream(void* ptr);

#ifdef __cplusplus
}
#endif