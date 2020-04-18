import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-front',
  templateUrl: './front.component.html',
  styleUrls: ['./front.component.css']
})
export class FrontComponent implements OnInit {

    public bat: string = "";
    public searchValue: string = "";

    constructor() { }

    ngOnInit() {
    }

    makeBat(){
        this.searchValue;
        debugger
    }

}
