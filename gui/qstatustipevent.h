#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStatusTipEvent_NewQStatusTipEvent(char* tip);
char* QStatusTipEvent_Tip(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif