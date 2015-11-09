#ifdef __cplusplus
extern "C" {
#endif

void* QDomProcessingInstruction_NewQDomProcessingInstruction();
void* QDomProcessingInstruction_NewQDomProcessingInstruction2(void* x);
char* QDomProcessingInstruction_Data(void* ptr);
int QDomProcessingInstruction_NodeType(void* ptr);
void QDomProcessingInstruction_SetData(void* ptr, char* d);
char* QDomProcessingInstruction_Target(void* ptr);

#ifdef __cplusplus
}
#endif