
// array of techs got from api
let listData;

update()

function update(){
    httpGetAsync(getAllUrl(), updateList);
}

function getAllUrl() {
    return "/rest/project/all";
}

function updateList(response)
{
    // parse
    const list = JSON.parse(response);
    listData = list;
    // find root for list
    let listElem = findListElem();
    // fill list
    listData.forEach(function (item) {
        listElem.appendChild(createTechItem(item))
    })
}

function findListElem()
{
    return document.getElementsByClassName("item-list")[0];
}

function templateTechItemInnerHTML(item) {
    return `
            <a class="item-href" href="">
                <span class="item-title">
                    ${item.Metadata.FriendlyName}
                </span>
                
            </a>
            <div class="item-content">
                <dl>
                    <dt>cost:</dt><dd>${item.Metadata.ResearchCost}</dd> 
                    <dt>category:</dt><dd>${item.Metadata.TechCategory}</dd>
                    <dt>prerequisites:</dt><dd>${item.Metadata.Prereqs}</dd>
                </dl>
            </div>
            <div class="item-description">
                ${item.Locale.Description}
            </div>
    `;
}

function createTechItem(techItem){
    let res = document.createElement('div')
    res.innerHTML = templateTechItemInnerHTML(techItem)
    res.classList.add('item')
    return res
}