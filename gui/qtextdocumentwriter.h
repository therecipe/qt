#ifdef __cplusplus
extern "C" {
#endif

void* QTextDocumentWriter_NewQTextDocumentWriter();
void* QTextDocumentWriter_NewQTextDocumentWriter2(void* device, void* format);
void* QTextDocumentWriter_NewQTextDocumentWriter3(char* fileName, void* format);
void* QTextDocumentWriter_Codec(void* ptr);
void* QTextDocumentWriter_Device(void* ptr);
char* QTextDocumentWriter_FileName(void* ptr);
void* QTextDocumentWriter_Format(void* ptr);
void QTextDocumentWriter_SetCodec(void* ptr, void* codec);
void QTextDocumentWriter_SetDevice(void* ptr, void* device);
void QTextDocumentWriter_SetFileName(void* ptr, char* fileName);
void QTextDocumentWriter_SetFormat(void* ptr, void* format);
int QTextDocumentWriter_Write(void* ptr, void* document);
int QTextDocumentWriter_Write2(void* ptr, void* fragment);
void QTextDocumentWriter_DestroyQTextDocumentWriter(void* ptr);

#ifdef __cplusplus
}
#endif