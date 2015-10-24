#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQuickImageProvider_NewQQuickImageProvider(int ty, int flags);
int QQuickImageProvider_Flags(QtObjectPtr ptr);
int QQuickImageProvider_ImageType(QtObjectPtr ptr);
void QQuickImageProvider_DestroyQQuickImageProvider(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif