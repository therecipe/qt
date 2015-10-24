#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextDocumentWriter_NewQTextDocumentWriter();
QtObjectPtr QTextDocumentWriter_NewQTextDocumentWriter2(QtObjectPtr device, QtObjectPtr format);
QtObjectPtr QTextDocumentWriter_NewQTextDocumentWriter3(char* fileName, QtObjectPtr format);
QtObjectPtr QTextDocumentWriter_Codec(QtObjectPtr ptr);
QtObjectPtr QTextDocumentWriter_Device(QtObjectPtr ptr);
char* QTextDocumentWriter_FileName(QtObjectPtr ptr);
void QTextDocumentWriter_SetCodec(QtObjectPtr ptr, QtObjectPtr codec);
void QTextDocumentWriter_SetDevice(QtObjectPtr ptr, QtObjectPtr device);
void QTextDocumentWriter_SetFileName(QtObjectPtr ptr, char* fileName);
void QTextDocumentWriter_SetFormat(QtObjectPtr ptr, QtObjectPtr format);
int QTextDocumentWriter_Write(QtObjectPtr ptr, QtObjectPtr document);
int QTextDocumentWriter_Write2(QtObjectPtr ptr, QtObjectPtr fragment);
void QTextDocumentWriter_DestroyQTextDocumentWriter(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif