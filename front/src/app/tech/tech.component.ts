import {Component, OnInit, ViewEncapsulation} from '@angular/core';

@Component({
  selector: 'app-tech',
  templateUrl: './tech.component.html',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./tech.component.css']
})
export class TechComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    update()
  }

}


// array of techs got from api
let allTechList;

function update(){
  httpGetAsync(getTechUrl(), updateTechList);
}

function getTechUrl() {
  return "/rest/tech/all";
}

function updateTechList(response: any)
{
  // parse
  const list = JSON.parse(response);
  allTechList = list;
  // find root for list
  let listElem = findListElem();
  // fill list
  allTechList.forEach(function (tech: any) {
    listElem.appendChild(createTechItem(tech))
  })
}

function findListElem()
{
  return document.getElementsByClassName("item-list")[0];
}

function templateTechItemInnerHTML(techItem: any) {
  return `
            <a class="item-href" href="">
                <span class="item-title">
                    ${techItem.Metadata.FriendlyName}
                </span>

            </a>
            <div class="item-content">
                <dl>
                    <dt>cost:</dt><dd>${techItem.Metadata.ResearchCost}</dd>
                    <dt>category:</dt><dd>${techItem.Metadata.TechCategory}</dd>
                    <dt>prerequisites:</dt><dd>${techItem.Metadata.Prereqs}</dd>
                </dl>
            </div>
            <div class="item-description">
                ${techItem.Locale.Description}
            </div>
    `;
}

function createTechItem(techItem: any){
  let res = document.createElement('div')
  res.innerHTML = templateTechItemInnerHTML(techItem)
  res.classList.add('item')
  return res
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
