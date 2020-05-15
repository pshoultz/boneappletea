import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

import { MatSnackBar } from '@angular/material';

@Component({
  selector: 'app-make',
  templateUrl: './make.component.html',
  styleUrls: ['./make.component.scss']
})
export class MakeComponent implements OnInit {

    private data: any = {};

    constructor(
        private api: ApiService,
        private snackbar: MatSnackBar,
    ) { }

    ngOnInit() { }

    make(form: any){
        this.api.MakeBat(form.value)
            .subscribe(data => {
                this.data = data;
            });
    }

    copy(){
        var temp = document.createElement("textarea") as HTMLTextAreaElement;
        debugger
        temp.value = this.data.boneappletea;
        document.body.appendChild(temp);
        temp.focus();
        temp.select();

        var success = document.execCommand('copy');

        if(success) {
            this.snackbar.open("Copied", null, {
                duration: 2000,
                panelClass: ['success']
            });
        }else{
            this.snackbar.open("oh no, error", null, {
                duration: 2000,
                panelClass: ['error']
            });
        }

        document.body.removeChild(temp);
    }

}
