#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextDecoder_NewQTextDecoder(QtObjectPtr codec);
QtObjectPtr QTextDecoder_NewQTextDecoder2(QtObjectPtr codec, int flags);
void QTextDecoder_DestroyQTextDecoder(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif