
import pieces from "./pieces.js";

export default class Square {
    constructor({ board, rank, file,isBlack,index}) {
        this.board =board;
        this.rank = rank;
        this.file = file;
        this.index = index;
        this.element = document.createElement('div');
        this.element.classList.add('square');
 
        if(isBlack) {
            this.element.classList.add('black');
        }
        // this.element.textContent = `${file}${rank}`;
        this.element.setAttribute('data-rank',rank);
        this.element.setAttribute('data-file',file);
        
        this.update()
        

    }

    update() {
        const current = this.board.getSquare(this.index);
        if(current) {
            const image = pieces[`${current.color}${current.type}`];
            this.element.textContent = " ";
            if(image) {
                
                const img = new Image();
                img.src = image;
                img.draggable="true";
                img.id =  `${current.color}${current.type}${this.index}`;
                img.classList.add('piece');
                img.setAttribute('color',current.color);
                img.setAttribute('type',current.type)
                this.element.append(img);
            } else {
                // this.element.textContent = current.type;
            }
            // console.log(current);
            // this.element.textContent= current.type;
        } 
    }

}