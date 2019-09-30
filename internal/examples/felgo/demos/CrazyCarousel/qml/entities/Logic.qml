import QtQuick 2.0
import Felgo 3.0

QtObject {
    id: logic

    // properties for handling game speed and bullet/coin timeouts
    // can be used to speed up the game
    property double speedUp: 1.0
    property int minBulletTimeout: 1000
    property int minCoinTimeout: 2000
    property int bulletTimeout: 4000
    property int coinTimeout: 2000

    // increasePoints is called by coins when they are collected
    function increasePoints() {
        gameScene.points++
    }

    // is called by bullets when game is over
    function gameOver() {
        gameWindow.state = "gameover"
    }

    // timer for dynamic bullet creation
    property Timer bulletTimer: Timer {
        id: bulletTimer
        interval : Math.random() * logic.bulletTimeout + logic.minBulletTimeout
        running: gameScene.visible // start running from the beginning, when the scene is loaded

        onTriggered: {
            // calculate bullet properties
            var randomLane = Math.floor(Math.random() * 4 + 1)
            var enemyWidth = gameScene.laneWidth * 0.74
            var laneOffset = gameScene.width * 0.13
            var xPosition = Math.round(laneOffset + (randomLane * enemyWidth) - (enemyWidth / 2))
            var xGoal = randomLane * gameScene.laneWidth - (gameScene.laneWidth / 2)

            var newEntityProperties = {
                    x: xPosition,
                    y: 95, // position at height of enemies,
                    z: 12,
                    goalX: xGoal,
                    startX: xPosition,
                    startY: 95,
                    speed: 3000 / (1 + logic.speedUp / 8 * 1.5) // multiplier between 1 and 2.5
            }

            // add bullet to scene
            entityManager.createEntityFromUrlWithProperties(
                Qt.resolvedUrl("Bullet.qml"),
                newEntityProperties)

            // calculate random interval for next bullet
            interval = Math.random() * logic.bulletTimeout + logic.minBulletTimeout

            // restart the timer
            bulletTimer.restart()
        }
    }

    // timer for dynamic coin creation
    property Timer coinTimer: Timer {
        id: coinTimer
        interval : Math.random() * logic.coinTimeout + logic.minCoinTimeout
        running: gameScene.visible // start running from the beginning, when the scene is loaded

        onTriggered: {
            // calculate coin properties
            var randomLane = Math.floor(Math.random() * 3 + 1)
            var enemyWidth = gameScene.laneWidth * 0.74
            var laneOffset = gameScene.width * 0.13
            var xPosition = Math.round(laneOffset + (randomLane * enemyWidth))
            var xGoal = randomLane * gameScene.laneWidth

            var newEntityProperties = {
                    x: xPosition ,
                    y: 75, // position slightly above enemies
                    z: 12,
                    goalX: xGoal,
                    startX: xPosition,
                    startY: 75,
                    speed: 3000 / (1 + logic.speedUp / 8 * 2) // multiplier between 1 and 3
            }

            // add coin to scene
            entityManager.createEntityFromUrlWithProperties(
                Qt.resolvedUrl("Coin.qml"),
                newEntityProperties)

            // calculate random interval for next coin
            interval = Math.random() * logic.coinTimeout + logic.minCoinTimeout

            // restart the timer
            coinTimer.restart()
        }
    }

    // timer for speeding up the game
    property Timer speedUpTimer: Timer {
        id: speedUpTimer
        interval: 15000 // speed up every 15 seconds
        running: gameScene.visible // timer is only active when gameScene is running

        onTriggered: {
            logic.speedUp += 0.7 // speed up game

            // stop bulletTimer, reduce timeout and restart timer
            bulletTimer.stop()
            logic.bulletTimeout -= 300
            bulletTimer.interval = Math.random() * logic.bulletTimeout +
                    logic.minBulletTimeout
            bulletTimer.restart()

            // stop coinTimer, reduce timeout and restart timer
            coinTimer.stop()
            logic.coinTimeout -= 100
            coinTimer.interval = Math.random() * logic.coinTimeout +
                    logic.minCoinTimeout
            coinTimer.restart()

            // stop after speed up limit is reached
            if(logic.speedUp < 8)
                speedUpTimer.restart()
        }
    }
}
