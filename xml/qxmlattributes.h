#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlAttributes_NewQXmlAttributes();
void QXmlAttributes_DestroyQXmlAttributes(QtObjectPtr ptr);
void QXmlAttributes_Append(QtObjectPtr ptr, char* qName, char* uri, char* localPart, char* value);
void QXmlAttributes_Clear(QtObjectPtr ptr);
int QXmlAttributes_Count(QtObjectPtr ptr);
int QXmlAttributes_Index2(QtObjectPtr ptr, QtObjectPtr qName);
int QXmlAttributes_Index(QtObjectPtr ptr, char* qName);
int QXmlAttributes_Index3(QtObjectPtr ptr, char* uri, char* localPart);
int QXmlAttributes_Length(QtObjectPtr ptr);
char* QXmlAttributes_LocalName(QtObjectPtr ptr, int index);
char* QXmlAttributes_QName(QtObjectPtr ptr, int index);
char* QXmlAttributes_Type2(QtObjectPtr ptr, char* qName);
char* QXmlAttributes_Type3(QtObjectPtr ptr, char* uri, char* localName);
char* QXmlAttributes_Type(QtObjectPtr ptr, int index);
char* QXmlAttributes_Uri(QtObjectPtr ptr, int index);
char* QXmlAttributes_Value3(QtObjectPtr ptr, QtObjectPtr qName);
char* QXmlAttributes_Value2(QtObjectPtr ptr, char* qName);
char* QXmlAttributes_Value4(QtObjectPtr ptr, char* uri, char* localName);
char* QXmlAttributes_Value(QtObjectPtr ptr, int index);

#ifdef __cplusplus
}
#endif