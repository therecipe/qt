#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlItem_NewQXmlItem();
QtObjectPtr QXmlItem_NewQXmlItem4(char* atomicValue);
QtObjectPtr QXmlItem_NewQXmlItem2(QtObjectPtr other);
QtObjectPtr QXmlItem_NewQXmlItem3(QtObjectPtr node);
int QXmlItem_IsNode(QtObjectPtr ptr);
int QXmlItem_IsNull(QtObjectPtr ptr);
void QXmlItem_DestroyQXmlItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif