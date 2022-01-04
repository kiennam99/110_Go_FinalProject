
/* eslint-disable import/extensions */
import {Chess} from 'https://cdn.skypack.dev/chess.js'
import Square from "./square.js";

let socket = new WebSocket;
let moveAvailable = new Boolean;



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

    console.log(event.target.id);
});
document.addEventListener("dragend",function(event){
    // event.dataTransfer.setData("")
});

document.addEventListener("drop",function drop(event){
    event.preventDefault();
    
    if(event.target.className == "square" || event.target.className == "square black" ) {
        // var index = event.dataTransfer.getData("index");
        var data = event.dataTransfer.getData("img");
        var tar = document.getElementById(data);

        // console.log(document.getElementById(data));
        console.log(tar.id)
        event.target.append(tar);
        
    }
    if( event.target.classList.contains('piece')) {
        var data = event.dataTransfer.getData("img");
        var el = event.target;
        var tar = document.getElementById(data);

        console.log(tar.getAttribute('color'))
        console.log(el);
        
        if(el != tar && el.getAttribute('color')!= tar.getAttribute('color')){
            console.log(tar.id);
            el = event.target.parentNode;
            event.target.remove();
            // var newimg = 
            el.append(tar);
        }

     
    }
    
});
document.addEventListener("dragover", function allowdrop(event){
    event.preventDefault();
});
