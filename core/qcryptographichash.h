#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QCryptographicHash_NewQCryptographicHash(int method);
int QCryptographicHash_AddData2(QtObjectPtr ptr, QtObjectPtr device);
void QCryptographicHash_AddData3(QtObjectPtr ptr, QtObjectPtr data);
void QCryptographicHash_AddData(QtObjectPtr ptr, char* data, int length);
void QCryptographicHash_Reset(QtObjectPtr ptr);
void QCryptographicHash_DestroyQCryptographicHash(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif