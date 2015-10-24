#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDataStream_NewQDataStream3(QtObjectPtr a, int mode);
int QDataStream_AtEnd(QtObjectPtr ptr);
QtObjectPtr QDataStream_NewQDataStream();
QtObjectPtr QDataStream_NewQDataStream2(QtObjectPtr d);
QtObjectPtr QDataStream_NewQDataStream4(QtObjectPtr a);
int QDataStream_ByteOrder(QtObjectPtr ptr);
QtObjectPtr QDataStream_Device(QtObjectPtr ptr);
int QDataStream_FloatingPointPrecision(QtObjectPtr ptr);
int QDataStream_ReadRawData(QtObjectPtr ptr, char* s, int len);
void QDataStream_ResetStatus(QtObjectPtr ptr);
void QDataStream_SetByteOrder(QtObjectPtr ptr, int bo);
void QDataStream_SetDevice(QtObjectPtr ptr, QtObjectPtr d);
void QDataStream_SetFloatingPointPrecision(QtObjectPtr ptr, int precision);
void QDataStream_SetStatus(QtObjectPtr ptr, int status);
void QDataStream_SetVersion(QtObjectPtr ptr, int v);
int QDataStream_SkipRawData(QtObjectPtr ptr, int len);
int QDataStream_Status(QtObjectPtr ptr);
int QDataStream_Version(QtObjectPtr ptr);
int QDataStream_WriteRawData(QtObjectPtr ptr, char* s, int len);
void QDataStream_DestroyQDataStream(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif