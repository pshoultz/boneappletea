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
                if(data && data.includes("create ok")){
                    this.snackbar.open("boneappletea created!", null, {
                        duration: 2000,
                        panelClass: ['success']
                    });
                }else if(data && data.includes("boneappletea updated")){
                    this.snackbar.open("boneappletea updated!", null, {
                        duration: 2000,
                        panelClass: ['success']
                    });
                }else{
                    this.snackbar.open("boneappletea already exists", null, {
                        duration: 2000,
                        panelClass: ['error']
                    });
                }
            });
        }else{
            this.snackbar.open("FORM IS INCOMPLETE", null, {
                duration: 2000,
                panelClass: ['error']
            });
        }
    }

}
