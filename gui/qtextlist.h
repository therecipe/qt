#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTextList_ItemNumber(QtObjectPtr ptr, QtObjectPtr block);
char* QTextList_ItemText(QtObjectPtr ptr, QtObjectPtr block);
void QTextList_Add(QtObjectPtr ptr, QtObjectPtr block);
int QTextList_Count(QtObjectPtr ptr);
void QTextList_RemoveItem(QtObjectPtr ptr, int i);
void QTextList_SetFormat(QtObjectPtr ptr, QtObjectPtr format);

#ifdef __cplusplus
}
#endif