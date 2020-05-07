import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../shared/api.service';

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.scss']
})
export class ContactComponent implements OnInit {

    private data: any = {};

    constructor(
        private api: ApiService
    ) { }

    ngOnInit() {
    }

    contact(form: any){

        debugger
        var address: string = form.address;
        var message: string = form.message;

        this.api.SendEmail(address, message)
            .subscribe((data: any) => {
                console.log(data);
            });
    }

}
