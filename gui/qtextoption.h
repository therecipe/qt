#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextOption_NewQTextOption3(QtObjectPtr other);
QtObjectPtr QTextOption_NewQTextOption();
QtObjectPtr QTextOption_NewQTextOption2(int alignment);
int QTextOption_Alignment(QtObjectPtr ptr);
int QTextOption_Flags(QtObjectPtr ptr);
void QTextOption_SetAlignment(QtObjectPtr ptr, int alignment);
void QTextOption_SetFlags(QtObjectPtr ptr, int flags);
void QTextOption_SetTextDirection(QtObjectPtr ptr, int direction);
void QTextOption_SetUseDesignMetrics(QtObjectPtr ptr, int enable);
void QTextOption_SetWrapMode(QtObjectPtr ptr, int mode);
int QTextOption_TextDirection(QtObjectPtr ptr);
int QTextOption_UseDesignMetrics(QtObjectPtr ptr);
int QTextOption_WrapMode(QtObjectPtr ptr);
void QTextOption_DestroyQTextOption(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif