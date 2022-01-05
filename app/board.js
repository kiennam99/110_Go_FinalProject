
/* eslint-disable import/extensions */
import {Chess} from 'https://cdn.skypack.dev/chess.js'
import Square from "./square.js";

let socket = new WebSocket("ws://localhost:8899/chess");
let moveAvailable = new Boolean(false);

socket.onopen = function() {
    console.log("Status: Connected\n");
};

socket.onmessage = function (e) {
    if (e.data == "2") {
        console.log("HEY")
    }

    if (e.data == "1") {
        moveAvailable = true;
    } else if (e.data == "0") {
        moveAvailable = false;
    } else {
        moveAvailable = false
    }

    console.log(moveAvailable);
};

const files = ['A','B','C','D','E','F','G','H'];

export default class Board {
    constructor({selector,size} ) {
        this.size = size;
        this.cells = [];
        this.element = document.querySelector(selector);
        this.element.classList.add('board');
        this.init();

        // console.log(this.element)
    }

    init(){
        if(this.size) {
            this.element.style.width = this.size;
            this.element.style.height = this.size;
        } else {
            // const unit = window.innerHeight < window.innerWidth ? vh : vw;
            const size = '90vmin'
            this.element.style.width = size;
            this.element.style.height = size;
        }

        this.chess = new Chess();
        this.board = this.chess.board().flat();

        this.cells = Array.from({ length:64 },(_,index) => {
            const fileNum = index % 8 ;
            const rank = 8-Math.floor(index / 8);
            const file =  files[fileNum] ;
            const isBlack = !(rank % 2 === fileNum %2 ) ;

            const cell = new Square({ 
                board: this,
                rank,
                file,
                isBlack,
                index,
            });
            this.element.appendChild(cell.element);
            return cell;
        });

        // this.chess.reset()
        // console.log(this.chess.board().flat());

    }
    getSquare(index) { 
        return this.board[index];
    }
    

}

document.addEventListener("dragstart",function(event){
    event.dataTransfer.setData("img",event.target.id);
    // event.dataTransfer.setData("index",event.target.index);
});
document.addEventListener("dragend",function(event){
    // event.dataTransfer.setData("")
});

document.addEventListener("drop",function drop(event){
    event.preventDefault();
    
    if(event.target.className == "square" || event.target.className == "square black" ) {
        var data = event.dataTransfer.getData("img");
        var tar = document.getElementById(data);

        var startPos = tar.parentElement.getAttribute('position')
        var endPos = event.target.getAttribute('position')
        socket.send(startPos + " " + endPos);

        setTimeout(function() {
            if (moveAvailable) {
                event.target.append(tar);
            }
        }, 100)
    }
    if( event.target.classList.contains('piece')) {
        var data = event.dataTransfer.getData("img");
        var el = event.target;
        var tar = document.getElementById(data);

        //console.log(tar.getAttribute('color'))
        //console.log(el);

        var startPos = tar.parentElement.getAttribute('position')
        var endPos = event.target.parentElement.getAttribute('position')
        socket.send(startPos + " " + endPos);
        
        if(el != tar && el.getAttribute('color')!= tar.getAttribute('color')){
            setTimeout(function() {
                if (moveAvailable) {
                    el = event.target.parentNode;
                    event.target.remove();
                    // var newimg = 
                    el.append(tar);
                }
            }, 100)
        }

     
    }
    
});
document.addEventListener("dragover", function allowdrop(event){
    event.preventDefault();
});
