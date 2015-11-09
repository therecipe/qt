#ifdef __cplusplus
extern "C" {
#endif

void* QXmlAttributes_NewQXmlAttributes();
void QXmlAttributes_DestroyQXmlAttributes(void* ptr);
void QXmlAttributes_Append(void* ptr, char* qName, char* uri, char* localPart, char* value);
void QXmlAttributes_Clear(void* ptr);
int QXmlAttributes_Count(void* ptr);
int QXmlAttributes_Index2(void* ptr, void* qName);
int QXmlAttributes_Index(void* ptr, char* qName);
int QXmlAttributes_Index3(void* ptr, char* uri, char* localPart);
int QXmlAttributes_Length(void* ptr);
char* QXmlAttributes_LocalName(void* ptr, int index);
char* QXmlAttributes_QName(void* ptr, int index);
char* QXmlAttributes_Type2(void* ptr, char* qName);
char* QXmlAttributes_Type3(void* ptr, char* uri, char* localName);
char* QXmlAttributes_Type(void* ptr, int index);
char* QXmlAttributes_Uri(void* ptr, int index);
char* QXmlAttributes_Value3(void* ptr, void* qName);
char* QXmlAttributes_Value2(void* ptr, char* qName);
char* QXmlAttributes_Value4(void* ptr, char* uri, char* localName);
char* QXmlAttributes_Value(void* ptr, int index);

#ifdef __cplusplus
}
#endif