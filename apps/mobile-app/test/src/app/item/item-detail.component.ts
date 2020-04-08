import { Component, OnInit } from "@angular/core";
import { ActivatedRoute } from "@angular/router";
import { TextField } from "tns-core-modules/ui/text-field";

import { getFile, getImage, getJSON, getString, request, HttpResponse } from "tns-core-modules/http";
import { fromObject, fromObjectRecursive, Observable, PropertyChangeData } from "tns-core-modules/data/observable";


import { Item } from "./item";
import { ItemService } from "./item.service";


@Component({
    selector: "ns-details",
    templateUrl: "./item-detail.component.html"
})

export class ItemDetailComponent implements OnInit {
    item: Item;

    public viewModel = new Observable();
    public value: string = "";

    constructor(
        private itemService: ItemService,
        private route: ActivatedRoute
    ) { }

    ngOnInit(): void {
        const id = +this.route.snapshot.params.id;
        this.item = this.itemService.getItem(id);
        //this.test2();
    }

    onReturnPress(args){
        let textField = <TextField>args.object;
        this.test2(textField.text);
    }

    test2(sentence: string): void{
    console.log("in test 2");
    request({
        url: "http://10.0.2.2:8080/bat",
        method: "GET",
        headers: { "sentence": sentence },
        }).then(response => {
            //var result = response.content.toJSON();
            this.viewModel = response.content.toJSON();
        }, error => {
            console.error(error);
        });
    }

}
