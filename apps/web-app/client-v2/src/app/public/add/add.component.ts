import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.scss']
})
export class AddComponent implements OnInit {

    private data: any = {};

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() {
    }

    add(form: any){
        this.api.AddBat(form.root, form.replacement)
        .subscribe((data: any) => {
            console.log(data);
        });
    }

}
