#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomDocumentType_NewQDomDocumentType();
QtObjectPtr QDomDocumentType_NewQDomDocumentType2(QtObjectPtr n);
char* QDomDocumentType_InternalSubset(QtObjectPtr ptr);
char* QDomDocumentType_Name(QtObjectPtr ptr);
int QDomDocumentType_NodeType(QtObjectPtr ptr);
char* QDomDocumentType_PublicId(QtObjectPtr ptr);
char* QDomDocumentType_SystemId(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif