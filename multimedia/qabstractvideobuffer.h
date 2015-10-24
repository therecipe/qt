#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QAbstractVideoBuffer_Handle(QtObjectPtr ptr);
int QAbstractVideoBuffer_HandleType(QtObjectPtr ptr);
int QAbstractVideoBuffer_MapMode(QtObjectPtr ptr);
void QAbstractVideoBuffer_Release(QtObjectPtr ptr);
void QAbstractVideoBuffer_Unmap(QtObjectPtr ptr);
void QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif