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
            gameInfo.cells.set(cellsToMapKey(x,y), 1)
        } else {
            ctx.fillStyle = "white";
            gameInfo.cells.set(cellsToMapKey(x,y), 0)
        }
        ctx.fillRect(gameInfo.cellWidth() * x, gameInfo.cellHeigth() * y, gameInfo.cellWidth(), gameInfo.cellHeigth());
    }
}

function getCellCoord(mousePos) {
    let x = Math.floor(mousePos.x / gameInfo.cellWidth());
    let y = Math.floor(mousePos.y / gameInfo.cellHeigth());
    return {
        x: x,
        y: y
    };
}

function getMousePos(canvas, evt) {
    let rect = canvas.getBoundingClientRect();
    return {
        x: evt.clientX - rect.left,
        y: evt.clientY - rect.top
    };
}
