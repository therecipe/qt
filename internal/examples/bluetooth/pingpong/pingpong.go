package main

import (
	"fmt"

	"github.com/therecipe/qt/bluetooth"
	"github.com/therecipe/qt/core"
)

const (
	serviceUuid = "e8e10f95-1a70-4b27-9ccf-02010264e9c9"
	androidUuid = "c9e96402-0102-cf9c-274b-701a950fe1e8"
)

type PingPong struct {
	core.QObject

	_ func() `constructor:"init"`

	_ float64 `property:"ballX"`
	_ float64 `property:"ballY"`
	_ float64 `property:"leftBlockY"`
	_ float64 `property:"rightBlockY"`
	_ bool    `property:"showDialog"`
	_ string  `property:"message"`
	_ float64 `property:"role"`
	_ float64 `property:"leftResult"`
	_ float64 `property:"rightResult"`

	_ func()                                                 `slot:"startGame"`
	_ func()                                                 `slot:"update"`
	_ func(x, y float64)                                     `slot:"setSize"`
	_ func(bX, bY float64)                                   `slot:"updateBall"`
	_ func(lY float64)                                       `slot:"updateLeftBlock"`
	_ func(rY float64)                                       `slot:"updateRightBlock"`
	_ func()                                                 `slot:"startServer"`
	_ func()                                                 `slot:"startClient"`
	_ func()                                                 `slot:"clientConnected"`
	_ func()                                                 `slot:"clientDisconnected"`
	_ func()                                                 `slot:"serverConnected"`
	_ func()                                                 `slot:"serverDisconnected"`
	_ func(bluetooth.QBluetoothSocket__SocketError)          `slot:"socketError"`
	_ func(bluetooth.QBluetoothServer__Error)                `slot:"serverError"`
	_ func(bluetooth.QBluetoothServiceDiscoveryAgent__Error) `slot:"serviceScanError"`
	_ func()                                                 `slot:"done"`
	_ func(b *bluetooth.QBluetoothServiceInfo)               `slot:"addService"`
	_ func()                                                 `slot:"readSocket"`

	_ func() `signal:"ballChanged"`
	_ func() `signal:"resultChanged"`

	m_serverInfo   *bluetooth.QBluetoothServer
	m_serviceInfo  *bluetooth.QBluetoothServiceInfo
	socket         *bluetooth.QBluetoothSocket
	discoveryAgent *bluetooth.QBluetoothServiceDiscoveryAgent

	m_timer        *core.QTimer
	m_serviceFound bool
	interval       float64

	m_direction     float64
	m_ballPreviousX float64
	m_ballPreviousY float64
	m_boardWidth    float64
	m_boardHeight   float64
	m_targetX       float64
	m_targetY       float64
	m_proportionX   float64
	m_proportionY   float64
}

func (p *PingPong) init() {

	//TODO: always init will zero/nil
	p.SetBallX(0)
	p.SetBallY(0)
	p.SetLeftBlockY(0)
	p.SetRightBlockY(0)
	p.SetShowDialog(false)
	p.SetMessage("")
	p.SetRole(0)
	p.SetLeftResult(0)
	p.SetRightResult(0)

	p.interval = 5

	p.m_timer = core.NewQTimer(p)
	p.m_timer.ConnectTimeout(p.update)

	p.ConnectSetMessage(p.setMessage)

	p.ConnectStartGame(p.startGame)
	p.ConnectUpdate(p.update)
	p.ConnectSetSize(p.setSize)
	p.ConnectUpdateBall(p.updateBall)
	p.ConnectUpdateLeftBlock(p.updateLeftBlock)
	p.ConnectUpdateRightBlock(p.updateRightBlock)
	p.ConnectStartServer(p.startServer)
	p.ConnectStartClient(p.startClient)
	p.ConnectClientConnected(p.clientConnected)
	p.ConnectClientDisconnected(p.clientDisconnected)
	p.ConnectServerConnected(p.serverConnected)
	p.ConnectServerDisconnected(p.serverDisconnected)
	p.ConnectSocketError(p.socketError)
	p.ConnectServerError(p.serverError)
	p.ConnectServiceScanError(p.serviceScanError)
	p.ConnectDone(p.done)
	p.ConnectAddService(p.addService)
	p.ConnectReadSocket(p.readSocket)

	p.ConnectBallXChanged(func(float64) { p.BallChanged() })
	p.ConnectBallYChanged(func(float64) { p.BallChanged() })
	p.ConnectLeftResultChanged(func(float64) { p.ResultChanged() })
	p.ConnectRightResultChanged(func(float64) { p.ResultChanged() })
}

func (p *PingPong) startGame() {
	p.SetShowDialog(false)
	p.ShowDialogChanged(p.IsShowDialog())

	//! [Start the game]
	if p.Role() == 1 {
		p.updateDirection()
	}

	p.m_timer.Start(50)
	//! [Start the game]
}

func (p *PingPong) update() {
	size := core.NewQByteArray()
	// Server is only updating the coordinates
	//! [Updating coordinates]
	if p.Role() == 1 {
		p.checkBoundaries()
		p.m_ballPreviousX = p.BallX()
		p.m_ballPreviousY = p.BallY()
		p.SetBallY(p.m_direction*(p.BallX()+p.interval) - p.m_direction*p.BallX() + p.BallY())
		p.SetBallX(p.BallX() + p.interval)

		size.SetNum8(p.BallX(), "g", 6)
		size.Append3(" ")

		size1 := core.NewQByteArray()
		size.SetNum8(p.BallY(), "g", 6)
		size.Append(size1)
		size.Append3(" ")

		size1.SetNum8(p.LeftBlockY(), "g", 6)
		size.Append(size1)
		size.Append3(" \n")
		p.socket.Write2(size.ConstData())
		p.BallChanged()
	} else if p.Role() == 2 {
		size.SetNum8(p.RightBlockY(), "g", 6)
		size.Append3(" \n")
		p.socket.Write2(size.ConstData())
	}
	//! [Updating coordinates]
}

func (p *PingPong) setSize(x, y float64) {
	p.m_boardWidth = x
	p.m_boardHeight = y
	p.m_targetX = p.m_boardWidth
	p.m_targetY = p.m_boardHeight / 2

	p.SetBallX(p.m_boardWidth / 2)
	p.m_ballPreviousX = p.BallX()

	p.SetBallY(p.m_boardHeight - p.m_boardWidth/54)
	p.m_ballPreviousY = p.BallY()

	p.BallChanged()
}

func (p *PingPong) updateBall(bX, bY float64) {
	p.SetBallX(bX)
	p.SetBallY(bY)
}

func (p *PingPong) updateLeftBlock(lY float64) {
	p.SetLeftBlockY(lY)
}

func (p *PingPong) updateRightBlock(rY float64) {
	p.SetRightBlockY(rY)
}

func (p *PingPong) checkBoundaries() {
	ballWidth := p.m_boardWidth / 54
	blockSize := p.m_boardWidth / 27
	blockHeight := p.m_boardHeight / 5
	//! [Checking the boundaries]
	if ((p.BallX() + ballWidth) > (p.m_boardWidth - blockSize)) && ((p.BallY() + ballWidth) < (p.RightBlockY() + blockHeight)) && (p.BallY() > p.RightBlockY()) {
		p.m_targetY = 2*p.BallY() - p.m_ballPreviousY
		p.m_targetX = p.m_ballPreviousX
		p.interval = -5
		p.updateDirection()
	} else if (p.BallX() < blockSize) && ((p.BallY() + ballWidth) < (p.LeftBlockY() + blockHeight)) && (p.BallY() > p.LeftBlockY()) {
		p.m_targetY = 2*p.BallY() - p.m_ballPreviousY
		p.m_targetX = p.m_ballPreviousX
		p.interval = 5
		p.updateDirection()
	} else if p.BallY() < 0 || (p.BallY()+ballWidth > p.m_boardHeight) {
		p.m_targetY = p.m_ballPreviousY
		p.m_targetX = p.BallX() + p.interval
		p.updateDirection()
	} else if (p.BallX() + ballWidth) > p.m_boardWidth { //! [Checking the boundaries]
		p.SetLeftResult(p.LeftResult() + 1)
		p.m_targetX = p.m_boardWidth
		p.m_targetY = p.m_boardHeight / 2
		p.SetBallX(p.m_boardWidth / 2)
		p.m_ballPreviousX = p.BallX()
		p.SetBallY(p.m_boardHeight - p.m_boardWidth/54)
		p.m_ballPreviousY = p.BallY()

		p.updateDirection()
		p.checkResult()

		result := core.NewQByteArray()
		result.Append3("result ")
		res := core.NewQByteArray()
		res.SetNum8(p.LeftResult(), "g", 6)
		result.Append(res)
		result.Append3(" ")
		res.SetNum8(p.RightResult(), "g", 6)
		result.Append(res)
		result.Append3(" \n")
		p.socket.Write3(result)
		println(result.ConstData())
		p.ResultChanged()
	} else if p.BallX() < 0 {
		p.SetRightResult(p.RightResult() + 1)
		p.m_targetX = 0
		p.m_targetY = p.m_boardHeight / 2
		p.SetBallX(p.m_boardWidth / 2)
		p.m_ballPreviousX = p.BallX()
		p.SetBallY(p.m_boardHeight - p.m_boardWidth/54)
		p.m_ballPreviousY = p.BallY()

		p.updateDirection()
		p.checkResult()

		result := core.NewQByteArray()
		result.Append3("result ")
		res := core.NewQByteArray()
		res.SetNum8(p.LeftResult(), "g", 6)
		result.Append(res)
		result.Append3(" ")
		res.SetNum8(p.RightResult(), "g", 6)
		result.Append(res)
		result.Append3(" \n")
		p.socket.Write3(result)
		println(result.ConstData())
		p.ResultChanged()
	}
}

func (p *PingPong) checkResult() {
	if p.RightResult() == 10 && p.Role() == 2 {
		p.SetMessage("Game over. You win!")
		p.m_timer.Stop()

	} else if p.RightResult() == 10 && p.Role() == 1 {
		p.SetMessage("Game over. You lose!")
		p.m_timer.Stop()

	} else if p.LeftResult() == 10 && p.Role() == 1 {
		p.SetMessage("Game over. You win!")
		p.m_timer.Stop()

	} else if p.LeftResult() == 10 && p.Role() == 2 {
		p.SetMessage("Game over. You lose!")
		p.m_timer.Stop()
	}
}

func (p *PingPong) updateDirection() {
	p.m_direction = (p.m_targetY - p.BallY()) / (p.m_targetX - p.BallX())
}

func (p *PingPong) startServer() {
	p.SetMessage("Starting the server")
	//! [Starting the server]
	p.m_serverInfo = bluetooth.NewQBluetoothServer(bluetooth.QBluetoothServiceInfo__RfcommProtocol, nil)

	p.m_serverInfo.ConnectNewConnection(p.clientConnected)
	p.m_serverInfo.ConnectError2(p.serverError)
	p.m_serverInfo.Listen2(bluetooth.NewQBluetoothUuid9(serviceUuid), "PingPong server")

	//! [Starting the server]
	p.SetMessage("Server started, waiting for the client. You are the left player.")
	// m_role is set to 1 if it is a server
	p.SetRole(1)
	p.RoleChanged(p.Role())
}

func (p *PingPong) startClient() {

	//! [Searching for the service]
	p.discoveryAgent = bluetooth.NewQBluetoothServiceDiscoveryAgent2(bluetooth.NewQBluetoothAddress(), nil)

	p.discoveryAgent.ConnectServiceDiscovered(p.addService)
	p.discoveryAgent.ConnectFinished(p.done)
	p.discoveryAgent.ConnectError2(p.serviceScanError)

	p.discoveryAgent.SetUuidFilter2(bluetooth.NewQBluetoothUuid9(serviceUuid))
	p.discoveryAgent.StartDefault(bluetooth.QBluetoothServiceDiscoveryAgent__FullDiscovery) //TODO: register enum for non *Default func call

	//! [Searching for the service]
	p.SetMessage("Starting server discovery. You are the right player")
	// m_role is set to 2 if it is a client
	p.SetRole(2)
	p.RoleChanged(p.Role())
}

func (p *PingPong) clientConnected() {
	//! [Initiating server socket]
	if !p.m_serverInfo.HasPendingConnections() {
		p.SetMessage("FAIL: expected pending server connection")
		return
	}
	p.socket = p.m_serverInfo.NextPendingConnection()
	if p.socket.Pointer() == nil {
		return
	}
	p.socket.SetParent(p)
	p.socket.ConnectReadyRead(p.readSocket)
	p.socket.ConnectDisconnected(p.clientConnected)
	p.socket.ConnectError2(p.socketError)

	//! [Initiating server socket]
	p.SetMessage("Client connected.")

	size := core.NewQByteArray()
	size.SetNum8(p.m_boardWidth, "g", 6)
	size.Append3(" ")

	size1 := core.NewQByteArray()
	size1.SetNum8(p.m_boardHeight, "g", 6)
	size.Append(size1)
	size.Append3(" \n")
	p.socket.Write2(size.ConstData())
}

func (p *PingPong) clientDisconnected() {
	p.SetMessage("Client disconnected")
	p.m_timer.Stop()
}

func (p *PingPong) socketError(bluetooth.QBluetoothSocket__SocketError) {
	p.m_timer.Stop()
}

func (p *PingPong) serverError(bluetooth.QBluetoothServer__Error) {
	p.m_timer.Stop()
}

func (p *PingPong) done() {
	println("Service scan done")
	if !p.m_serviceFound {
		p.SetMessage("PingPong service not found")
	}
}

func (p *PingPong) addService(service *bluetooth.QBluetoothServiceInfo) {
	p.setMessage("Service found. Setting parameters...")
	//! [Connecting the socket]
	p.socket = bluetooth.NewQBluetoothSocket(bluetooth.QBluetoothServiceInfo__RfcommProtocol, nil)

	p.socket.ConnectReadyRead(p.readSocket)
	p.socket.ConnectConnected(p.serverConnected)
	p.socket.ConnectDisconnected(p.serverDisconnected)

	p.socket.ConnectToService(service, 0)
	//! [Connecting the socket]
	p.m_serviceFound = true
}

func (p *PingPong) serviceScanError(err bluetooth.QBluetoothServiceDiscoveryAgent__Error) {
	p.setMessage(fmt.Sprint("Scanning error", err))
}

func (p *PingPong) serverConnected() {
	p.setMessage("Server Connected")
	size := core.NewQByteArray()
	size.SetNum8(p.m_boardWidth, "g", 6)
	size.Append3(" ")
	size1 := core.NewQByteArray()
	size1.SetNum8(p.m_boardHeight, "g", 6)
	size.Append(size1)
	size.Append3(" \n")
	p.socket.Write2(size.ConstData())
}

func (p *PingPong) serverDisconnected() {
	p.setMessage("Server Disconnected")
	p.m_timer.Stop()
}

func (p *PingPong) readSocket() {
	if p.socket.Pointer() == nil {
		return
	}
	sep := " "
	var line *core.QByteArray
	for p.socket.CanReadLine() {
		line = p.socket.ReadLine2(0)
		println(line.ConstData(), line.Length())
		if line.Contains2("result") {
			result := line.Split(sep)
			if len(result) > 2 {
				leftSide := result[1]
				rightSide := result[2]
				p.SetLeftResult(leftSide.ToDouble(false))
				p.SetRightResult(rightSide.ToDouble(false))
				p.ResultChanged()
				p.checkResult()
			}
		}
	}
	if p.m_proportionX == 0 || p.m_proportionY == 0 {
		boardSize := line.Split(sep)
		if len(boardSize) > 1 {
			boardWidth := boardSize[0]
			boardHeight := boardSize[1]
			p.m_proportionX = p.m_boardWidth / boardWidth.ToDouble(false)
			p.m_proportionY = p.m_boardHeight / boardHeight.ToDouble(false)
			p.setMessage("Screen adjusted. Get ready!")

			singleShot := core.NewQTimer(nil)
			singleShot.ConnectTimeout(p.startGame)
			singleShot.SetSingleShot(true)
			singleShot.Start(3000)
		}
	} else if p.Role() == 1 {
		boardSize := line.Split(sep)
		if len(boardSize) > 1 {
			rightBlockY := boardSize[0]
			p.SetRightBlockY(p.m_proportionY * rightBlockY.ToDouble(false))
			p.RightBlockYChanged(p.RightBlockY())
		}
	} else if p.Role() == 2 {
		boardSize := line.Split(sep)
		if len(boardSize) > 2 {
			ballX := boardSize[0]
			ballY := boardSize[1]
			leftBlockY := boardSize[2]
			p.SetBallX(p.m_proportionX * ballX.ToDouble(false))
			p.SetBallY(p.m_proportionY * ballY.ToDouble(false))
			p.SetLeftBlockY(p.m_proportionY * leftBlockY.ToDouble(false))
			p.LeftBlockYChanged(p.LeftBlockY())
			p.BallChanged()
		}
	}
}

func (p *PingPong) setMessage(message string) {
	p.SetShowDialog(true)
	p.SetMessageDefault(message)
	p.ShowDialogChanged(p.IsShowDialog())
}
