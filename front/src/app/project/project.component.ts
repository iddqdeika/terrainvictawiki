import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';

@Component({
  selector: 'app-project',
  templateUrl: './project.component.html',
  styleUrls: ['./project.component.css']
})
export class ProjectComponent implements OnInit {

  // array of techs got from api
  allList: any[];
  actualList: any[];
  currentFragment: string;

  constructor(private router: Router, private r:ActivatedRoute) {
    this.allList = [];
    this.actualList = [];
    this.currentFragment = "";
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

  update(){
    httpGetAsync(this.getTechUrl(), this.updateTechList.bind(this));
  }

  updateSync(){
    this.updateTechList(httpGetSync(this.getTechUrl()))
  }

  getTechUrl() {
    return "/rest/project/all";
  }

  updateTechList(response: any)
  {
    // parse
    const list = JSON.parse(response);
    this.allList = list;
    this.filterList("");
    //setTimeout(this.goToCurrentFragment.bind(this), 100);
  }

  filterList(name: string)
  {
    console.log("search for " + name.toLowerCase());
    this.actualList = this.allList.filter(
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
