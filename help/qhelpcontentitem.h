#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHelpContentItem_Child(QtObjectPtr ptr, int row);
int QHelpContentItem_ChildCount(QtObjectPtr ptr);
int QHelpContentItem_ChildPosition(QtObjectPtr ptr, QtObjectPtr child);
QtObjectPtr QHelpContentItem_Parent(QtObjectPtr ptr);
int QHelpContentItem_Row(QtObjectPtr ptr);
char* QHelpContentItem_Title(QtObjectPtr ptr);
char* QHelpContentItem_Url(QtObjectPtr ptr);
void QHelpContentItem_DestroyQHelpContentItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif