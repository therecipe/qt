package main

import (
	"github.com/therecipe/qt/bluetooth"
	"github.com/therecipe/qt/core"
)

type FileTransfer struct {
	core.QObject

	_ func() `constructor:"init"`

	_ float64 `property:"progress"`

	_ func(address, fileName string) `slot:"initTransfer"`
	_ func(int64, int64)             `slot:"updateProgress"`

	manager *bluetooth.QBluetoothTransferManager
	reply   *bluetooth.QBluetoothTransferReply
}

func (f *FileTransfer) init() {
	f.manager = bluetooth.NewQBluetoothTransferManager(f)

	f.ConnectInitTransfer(f.initTransfer)
	f.ConnectUpdateProgress(f.updateProgress)
}

//! [Transfer-1]
func (f *FileTransfer) initTransfer(address, fileName string) {
	println("Begin sharing file:", address, fileName)
	btAddress := bluetooth.NewQBluetoothAddress3(address)
	request := bluetooth.NewQBluetoothTransferRequest(btAddress)
	file := core.NewQFile2(fileName)
	f.reply = f.manager.Put(request, file)
	f.reply.ConnectTransferProgress(f.updateProgress)
}

//! [Transfer-1]

func (f *FileTransfer) updateProgress(transferred, total int64) {
	f.SetProgress(float64(transferred) / float64(total))
}
