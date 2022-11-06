import {Component, OnInit, ViewEncapsulation} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";

@Component({
  selector: 'app-tech',
  templateUrl: './tech.component.html',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./tech.component.css']
})
export class TechComponent implements OnInit {
  // array of techs got from api
  allTechList: any[];
  actualTechList: any[];

  constructor(private router: Router, private r:ActivatedRoute) {
    this.allTechList = [];
    this.actualTechList = [];
  }

  ngOnInit(): void {
    this.updateSync();
    setTimeout(this.goToCurrentFragment.bind(this), 0);
  }

  goToCurrentFragment()
  {
    const fragment = this.router.parseUrl(this.router.url).fragment;
    if (!fragment){
      return;
    }
    this.router.navigate(['.'], {fragment: fragment, relativeTo: this.r, skipLocationChange: true})
  }

  processSearchInput(event: any)
  {
    this.filterList(event.target.value);
  }

  updateAsync(){
    httpGetAsync(this.getTechUrl(), this.updateTechList.bind(this));
  }

  updateSync(){
    this.updateTechList(httpGetSync(this.getTechUrl()))
  }

  getTechUrl() {
    return "/rest/tech/all";
  }

  updateTechList(response: any)
  {
    // parse
    const list = JSON.parse(response);
    this.allTechList = list;
    this.filterList("");
  }

  filterList(name: string)
  {
    console.log("search for " + name.toLowerCase());
    this.actualTechList = this.allTechList.filter(
      (v, i, arr) => {
        return name == "" || v.Metadata.FriendlyName.toLowerCase().includes(name.toLowerCase())
      }
    );
  }
}

function httpGetAsync(theUrl: any, callback: any)
{
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.onreadystatechange = function() {
    if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
      callback(xmlHttp.responseText);
  }
  xmlHttp.open("GET", theUrl, true); // true for asynchronous
  xmlHttp.send(null);
}

function httpGetSync(theUrl: any)
{
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.open("GET", theUrl, false); // true for asynchronous
  xmlHttp.send(null);
  return xmlHttp.responseText;
}
