#ifdef __cplusplus
extern "C" {
#endif

void* QNdefMessage_NewQNdefMessage();
void* QNdefMessage_NewQNdefMessage3(void* message);
void* QNdefMessage_NewQNdefMessage2(void* record);
void* QNdefMessage_ToByteArray(void* ptr);

#ifdef __cplusplus
}
#endif