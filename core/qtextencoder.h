#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextEncoder_NewQTextEncoder(QtObjectPtr codec);
QtObjectPtr QTextEncoder_NewQTextEncoder2(QtObjectPtr codec, int flags);
void QTextEncoder_DestroyQTextEncoder(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif