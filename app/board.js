
/* eslint-disable import/extensions */
import {Chess} from 'https://cdn.skypack.dev/chess.js'
import Square from "./square.js";

// console.log(Chess);


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