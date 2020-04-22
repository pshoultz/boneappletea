import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

    constructor(
        private http: HttpClient 
    ) { }

    GetBats(){
        return this.http.get("http://127.0.0.1:8080/get");
    }

    MakeBat(sentence: string):string{
        let bat: string = "dummy dummy";

        return bat;
    }
}
