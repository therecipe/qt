#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlParseException_NewQXmlParseException(char* name, int c, int l, char* p, char* s);
QtObjectPtr QXmlParseException_NewQXmlParseException2(QtObjectPtr other);
int QXmlParseException_ColumnNumber(QtObjectPtr ptr);
int QXmlParseException_LineNumber(QtObjectPtr ptr);
char* QXmlParseException_Message(QtObjectPtr ptr);
char* QXmlParseException_PublicId(QtObjectPtr ptr);
char* QXmlParseException_SystemId(QtObjectPtr ptr);
void QXmlParseException_DestroyQXmlParseException(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif