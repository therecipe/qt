#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomNodeList_NewQDomNodeList();
QtObjectPtr QDomNodeList_NewQDomNodeList2(QtObjectPtr n);
int QDomNodeList_Count(QtObjectPtr ptr);
int QDomNodeList_IsEmpty(QtObjectPtr ptr);
int QDomNodeList_Length(QtObjectPtr ptr);
int QDomNodeList_Size(QtObjectPtr ptr);
void QDomNodeList_DestroyQDomNodeList(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif