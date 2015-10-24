#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTextCodec_CanEncode(QtObjectPtr ptr, QtObjectPtr ch);
int QTextCodec_CanEncode2(QtObjectPtr ptr, char* s);
QtObjectPtr QTextCodec_QTextCodec_CodecForHtml2(QtObjectPtr ba);
QtObjectPtr QTextCodec_QTextCodec_CodecForHtml(QtObjectPtr ba, QtObjectPtr defaultCodec);
QtObjectPtr QTextCodec_QTextCodec_CodecForLocale();
QtObjectPtr QTextCodec_QTextCodec_CodecForMib(int mib);
QtObjectPtr QTextCodec_QTextCodec_CodecForName(QtObjectPtr name);
QtObjectPtr QTextCodec_QTextCodec_CodecForName2(char* name);
QtObjectPtr QTextCodec_QTextCodec_CodecForUtfText2(QtObjectPtr ba);
QtObjectPtr QTextCodec_QTextCodec_CodecForUtfText(QtObjectPtr ba, QtObjectPtr defaultCodec);
QtObjectPtr QTextCodec_MakeDecoder(QtObjectPtr ptr, int flags);
QtObjectPtr QTextCodec_MakeEncoder(QtObjectPtr ptr, int flags);
int QTextCodec_MibEnum(QtObjectPtr ptr);
void QTextCodec_QTextCodec_SetCodecForLocale(QtObjectPtr c);

#ifdef __cplusplus
}
#endif