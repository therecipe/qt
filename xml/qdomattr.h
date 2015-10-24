#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomAttr_NewQDomAttr();
QtObjectPtr QDomAttr_NewQDomAttr2(QtObjectPtr x);
char* QDomAttr_Name(QtObjectPtr ptr);
int QDomAttr_NodeType(QtObjectPtr ptr);
void QDomAttr_SetValue(QtObjectPtr ptr, char* v);
int QDomAttr_Specified(QtObjectPtr ptr);
char* QDomAttr_Value(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif