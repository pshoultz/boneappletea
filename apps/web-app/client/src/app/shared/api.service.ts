import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})

export class ApiService {

    constructor(
        private http: HttpClient
    ) { }

    GetBats() {
        return this.http.get("http://127.0.0.1:8080/get");
    }

    //NOTE: we should have ids on each BAT so we can find them in the DB faster
    AcceptBat(){
    }

    //NOTE: we should have ids on each BAT so we can find them in the DB faster
    DeleteBat(){
    }
}
