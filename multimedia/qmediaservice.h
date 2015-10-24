#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QMediaService_ReleaseControl(QtObjectPtr ptr, QtObjectPtr control);
QtObjectPtr QMediaService_RequestControl(QtObjectPtr ptr, char* interfa);
void QMediaService_DestroyQMediaService(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif