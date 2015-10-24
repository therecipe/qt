#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTextStream_AtEnd(QtObjectPtr ptr);
int QTextStream_AutoDetectUnicode(QtObjectPtr ptr);
QtObjectPtr QTextStream_Codec(QtObjectPtr ptr);
QtObjectPtr QTextStream_Device(QtObjectPtr ptr);
int QTextStream_FieldAlignment(QtObjectPtr ptr);
int QTextStream_FieldWidth(QtObjectPtr ptr);
void QTextStream_Flush(QtObjectPtr ptr);
int QTextStream_GenerateByteOrderMark(QtObjectPtr ptr);
int QTextStream_IntegerBase(QtObjectPtr ptr);
int QTextStream_NumberFlags(QtObjectPtr ptr);
char* QTextStream_ReadAll(QtObjectPtr ptr);
int QTextStream_RealNumberNotation(QtObjectPtr ptr);
int QTextStream_RealNumberPrecision(QtObjectPtr ptr);
void QTextStream_Reset(QtObjectPtr ptr);
void QTextStream_ResetStatus(QtObjectPtr ptr);
void QTextStream_SetAutoDetectUnicode(QtObjectPtr ptr, int enabled);
void QTextStream_SetCodec(QtObjectPtr ptr, QtObjectPtr codec);
void QTextStream_SetCodec2(QtObjectPtr ptr, char* codecName);
void QTextStream_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
void QTextStream_SetFieldAlignment(QtObjectPtr ptr, int mode);
void QTextStream_SetFieldWidth(QtObjectPtr ptr, int width);
void QTextStream_SetGenerateByteOrderMark(QtObjectPtr ptr, int generate);
void QTextStream_SetIntegerBase(QtObjectPtr ptr, int base);
void QTextStream_SetLocale(QtObjectPtr ptr, QtObjectPtr locale);
void QTextStream_SetNumberFlags(QtObjectPtr ptr, int flags);
void QTextStream_SetPadChar(QtObjectPtr ptr, QtObjectPtr ch);
void QTextStream_SetRealNumberNotation(QtObjectPtr ptr, int notation);
void QTextStream_SetRealNumberPrecision(QtObjectPtr ptr, int precision);
void QTextStream_SetStatus(QtObjectPtr ptr, int status);
void QTextStream_SetString(QtObjectPtr ptr, char* stri, int openMode);
void QTextStream_SkipWhiteSpace(QtObjectPtr ptr);
int QTextStream_Status(QtObjectPtr ptr);
char* QTextStream_String(QtObjectPtr ptr);
void QTextStream_DestroyQTextStream(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif