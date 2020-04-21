import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-front',
  templateUrl: './front.component.html',
  styleUrls: ['./front.component.css']
})
export class FrontComponent implements OnInit {

    public searchValue: string = "";
    public bat: string = "test phrase test phrase";

    constructor() { }

    ngOnInit() {
    }

    makeBat(){
        this.searchValue;
        debugger
    }

}
