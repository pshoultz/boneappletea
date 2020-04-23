import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

@Component({
  selector: 'app-make',
  templateUrl: './make.component.html',
  styleUrls: ['./make.component.scss']
})
export class MakeComponent implements OnInit {

    private data: any = {};

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() { }

    make(form: any){
        this.api.MakeBat(form.value)
            .subscribe(data => {
                this.data = data;
            });
    }

    copy(){
        debugger
    }

}
