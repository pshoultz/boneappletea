import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

import { MatSnackBar } from '@angular/material';

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.scss']
})
export class ContactComponent implements OnInit {

    private data: any = {};

    constructor(
        private api: ApiService,
        private snackbar: MatSnackBar
    ) { }

    ngOnInit() {
    }

    contact(form: any){

        var address: string = form.address;
        var message: string = form.message;
        debugger

        if(address !== undefined && message !== undefined){
            this.api.SendEmail(address, message)
                .subscribe((data: any) => {
                    console.log(data);
                });
        }else{
            this.snackbar.open("form is incomplete", null, {duration: 2000});
        }
    }

}
