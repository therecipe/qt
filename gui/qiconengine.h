#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QIconEngine_AddFile(QtObjectPtr ptr, char* fileName, QtObjectPtr size, int mode, int state);
void QIconEngine_AddPixmap(QtObjectPtr ptr, QtObjectPtr pixmap, int mode, int state);
QtObjectPtr QIconEngine_Clone(QtObjectPtr ptr);
char* QIconEngine_IconName(QtObjectPtr ptr);
char* QIconEngine_Key(QtObjectPtr ptr);
void QIconEngine_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, int mode, int state);
int QIconEngine_Read(QtObjectPtr ptr, QtObjectPtr in);
int QIconEngine_Write(QtObjectPtr ptr, QtObjectPtr out);
void QIconEngine_DestroyQIconEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif