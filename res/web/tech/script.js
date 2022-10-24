
// array of techs got from api
let allTechList;

update()

function update(){
    httpGetAsync(getTechUrl(), updateTechList);
}

function getTechUrl() {
    return "/tech/all";
}

function updateTechList(response)
{
    // parse
    const list = JSON.parse(response);
    allTechList = list;
    // find root for list
    let listElem = findListElem();
    // fill list
    allTechList.forEach(function (tech) {
        listElem.appendChild(createTechItem(tech))
    })
}

function findListElem()
{
    return document.getElementsByClassName("item-list")[0];
}

function templateTechItemInnerHTML(techItem) {
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

function createTechItem(techItem){
    let res = document.createElement('div')
    res.innerHTML = templateTechItemInnerHTML(techItem)
    res.classList.add('item')
    return res
}