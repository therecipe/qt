#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomProcessingInstruction_NewQDomProcessingInstruction();
QtObjectPtr QDomProcessingInstruction_NewQDomProcessingInstruction2(QtObjectPtr x);
char* QDomProcessingInstruction_Data(QtObjectPtr ptr);
int QDomProcessingInstruction_NodeType(QtObjectPtr ptr);
void QDomProcessingInstruction_SetData(QtObjectPtr ptr, char* d);
char* QDomProcessingInstruction_Target(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif