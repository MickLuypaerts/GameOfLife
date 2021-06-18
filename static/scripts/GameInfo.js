class GameInfo {
    constructor() {
        this.rows = 0;
        this.columns = 0;
        this.canvasLayerZero = document.getElementById("board-layer-0");
        this.canvasLayerOne = document.getElementById("board-layer-1");
        this.cells = new Map();
        this.running = false;
        this.baseURL = "http://localhost:3000";
    }
    canvasWidth() {
        return this.canvasLayerOne.clientWidth;
    }
    canvasHeigth() {
        return this.canvasLayerOne.clientHeight;
    }
    cellWidth() {
        return this.canvasWidth() / this.rows;
    }
    cellHeigth() {
        return this.canvasHeigth() / this.columns;
    }
}


