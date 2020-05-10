import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

import { MatSnackBar } from '@angular/material';

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.scss']
})
export class AddComponent implements OnInit {

    private data: any = {};

    constructor(
        private api: ApiService,
        private snackbar: MatSnackBar
    ) { }

    ngOnInit() {
    }

    add(form: any){
        if(form.root !== undefined || form.replacement !== undefined){
            this.api.AddBat(form.root, form.replacement)
            .subscribe((data: any) => {
                this.snackbar.open("boneappletea added!", null, {
                    duration: 2000,
                });
            });
        }else{
            this.snackbar.open("FORM IS INCOMPLETE", null, {
                duration: 2000,
                panelClass: ['error']
            });
        }
    }

}
