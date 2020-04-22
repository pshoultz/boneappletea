import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

@Component({
  selector: 'app-make',
  templateUrl: './make.component.html',
  styleUrls: ['./make.component.scss']
})
export class MakeComponent implements OnInit {

    private boneappletea: any[] = [];
    private data: any = {
        value: ""
    };

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() { }

    make(form: any){
        console.log(form.value);
    }

    copy(){
        debugger
    }

}
