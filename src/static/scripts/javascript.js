let gameInfo = new GameInfo();

gameInfo.canvasLayerZero.addEventListener("click", clickCellFunc);
document.getElementById("stepBtn").addEventListener("click", clickStepBtn);
document.getElementById("runBtn").addEventListener("click", clickRunBtn);
document.getElementById("resetBtn").addEventListener("click", clickResetBtn);
document.getElementById("createBtn").addEventListener("click", clickCreateBtn)
window.addEventListener("load", initGame);
window.setInterval(() => {
    if (gameInfo.running) {
        clickStepBtn();
    }
}, 1000);

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
    console.log("sending %s to %s", data, gameInfo.baseURL + "/set");
}

function clickStepBtn(e) {
    sendToServer("/step", "GET", null)
        .then(response => {
            console.log(response);
            for (let cell in response) {
                fillCell(response[cell].x, response[cell].y, response[cell].state);
            }
        })
        .catch((error) => {
            console.log("Error:", error);
        });
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

function clickCreateBtn() {
    let rowsInput = document.getElementById("rows").value;
    let columnsInput = document.getElementById("columns").value;
    if (inputValidation(rowsInput, columnsInput)) {
        let data = JSON.stringify({ "columns": parseInt(columnsInput), "rows": parseInt(rowsInput) });
        console.log(data);
        sendToServer("/createnewboard", "POST", data)
            .then(response => {
                console.log(response);
                gameInfo.createNewBoard(rowsInput, columnsInput);

                let ctx = gameInfo.canvasLayerOne.getContext("2d");
                ctx.clearRect(0, 0, gameInfo.canvasLayerOne.width, gameInfo.canvasLayerOne.height);
                ctx = gameInfo.canvasLayerZero.getContext("2d");
                ctx.clearRect(0, 0, gameInfo.canvasLayerZero.width, gameInfo.canvasLayerZero.height);
                drawBorders();
            })
            .catch((error) => {
                console.log(error);
            });
    }
}

function inputValidation(rowsInput, columnsInput) {
    if (!isNaN(rowsInput) && !isNaN(columnsInput) && rowsInput > 0 && columnsInput > 0) {
        return true;
    } else {
        alert(`${rowsInput} and ${columnsInput} are not valid input`);
        return false;
    }
}

function initGame() {
    sendToServer("/getboardsize", "GET", null)
        .then(data => {
            gameInfo.createNewBoard(data.rows, data.columns);
            drawBorders();
        })
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

const sendToServerTest = async (url, method, data) => {
    try {
        const response = await fetch(gameInfo.baseURL + url, {
            method: method,
            headers: {
                "Content-Type": "application/json",
            },
            body: data
        });
        console.log(response.status);
        const retData = await response.json();
        console.log(retData);

    } catch (error) {
        console.error(error);
    }
}