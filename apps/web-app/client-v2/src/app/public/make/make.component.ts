import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-make',
  templateUrl: './make.component.html',
  styleUrls: ['./make.component.scss']
})
export class MakeComponent implements OnInit {

    private boneappletea: string = "dummy text";

    constructor(
        
    ) { }

    ngOnInit() { }

    make(){
        debugger
    }

    copy(){
        debugger
    }

}
