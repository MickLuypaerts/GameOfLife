let gameInfo = new GameInfo();

gameInfo.canvasLayerZero.addEventListener("click", clickCellFunc);
document.getElementById("stepBtn").addEventListener("click", clickStepBtn);
document.getElementById("runBtn").addEventListener("click", clickRunBtn);
document.getElementById("resetBtn").addEventListener("click", clickResetBtn);
window.addEventListener("load", initGame);
window.setInterval(() => {
    if (gameInfo.running) {
        clickStepBtn();
    }
}, 1000);

function initGame() {
    sendToServer("/getboardsize", "GET", null)
        .then(data => {
            gameInfo.rows = data.rows;
            gameInfo.columns = data.columns;
            for (let i = 0; i < gameInfo.columns; i++) {
                for (let j = 0; j < gameInfo.rows; j++) {
                    gameInfo.cells.set(i + j, 0);
                }
            }
            drawBorders();
        })
}

function clickRunBtn() {
    gameInfo.running = !gameInfo.running;
    document.getElementById("stepBtn").disabled = gameInfo.running;
    let runBtn = document.getElementById("runBtn");
    if (gameInfo.running) {
        runBtn.textContent = "stop";
    } else {
        runBtn.textContent = "run";
    }
}

function drawBorders() {
    if (gameInfo.canvasLayerZero.getContext) {
        let ctx = gameInfo.canvasLayerZero.getContext("2d");
        ctx.fillStyle = "black"
        for (let i = 0; i < gameInfo.columns; i++) {
            ctx.fillRect(gameInfo.cellWidth() * i, 0, 1, gameInfo.canvasHeigth());
        }
        ctx.fillRect(gameInfo.canvasWidth() - 1, 0, 1, gameInfo.canvasHeigth());

        for (let i = 0; i < gameInfo.rows; i++) {
            ctx.fillRect(0, gameInfo.cellHeigth() * i, gameInfo.canvasWidth(), 1)
        }
        ctx.fillRect(0, gameInfo.canvasWidth() - 1, gameInfo.canvasWidth(), 1)
    }
}

function fillCell(x, y, cellState) {
    if (gameInfo.canvasLayerOne.getContext) {
        let ctx = gameInfo.canvasLayerOne.getContext("2d");
        if (cellState == 1) {
            ctx.fillStyle = "red";
            gameInfo.cells.set(x + y, 1)
        } else {
            ctx.fillStyle = "white";
            gameInfo.cells.set(x + y, 0)
        }
        ctx.fillRect(gameInfo.cellWidth() * x, gameInfo.cellHeigth() * y, gameInfo.cellWidth(), gameInfo.cellHeigth());
    }
}

function getMousePos(canvas, evt) {
    let rect = canvas.getBoundingClientRect();
    return {
        x: evt.clientX - rect.left,
        y: evt.clientY - rect.top
    };
}

function getCellCoord(mousePos) {
    let x = Math.floor(mousePos.x / gameInfo.cellWidth());
    let y = Math.floor(mousePos.y / gameInfo.cellHeigth());
    return {
        x: x,
        y: y
    };
}

async function sendToServer(url, method, data) {
    const response = await fetch(gameInfo.baseURL + url, {
        method: method,
        headers: {
            "Content-Type": "application/json",
        },
        body: data
    });
    return response.json();
}

function clickCellFunc(e) {
    let mousePos = getMousePos(gameInfo.canvasLayerOne, e);
    let cellCoord = getCellCoord(mousePos);
    let cellState = +!gameInfo.cells.get(cellCoord.x + cellCoord.y);
    fillCell(cellCoord.x, cellCoord.y, cellState);
    let data = JSON.stringify({ "x": parseInt(cellCoord.x), "y": parseInt(cellCoord.y), "state": parseInt(cellState) });
    sendToServer("/set", "POST", data)
        .then(response => {
            console.log(response);
        })
        .catch((error) => {
            console.log(error);
        });
    console.log("sending %s to %s", data, gameInfo.baseURL+"/set");
}

async function stepCallBack(resp) {
    console.log(resp);
    for (let cell in resp) {
        fillCell(resp[cell].x, resp[cell].y, resp[cell].state);
    }
}

function clickStepBtn(e) {
    sendToServer("/step", "GET", null)
        .then(data => {
            stepCallBack(data);
        })
        .catch((error) => {
            console.log("Error:", error);
        });
}

function clickResetBtn() {
    sendToServer("/resetboard", "GET", null)
        .then(response => {
            console.log(response);
        })
        .catch((error) => {
            console.log("Error:", error);
        });
    let ctx = gameInfo.canvasLayerOne.getContext("2d");
    ctx.clearRect(0, 0, gameInfo.canvasLayerOne.width, gameInfo.canvasLayerOne.height);
}