import { otherPrivateJobs, agencyIcons } from "../data.js";

export function setupPrivateSector(element) {
  const h = document.createElement("h3");

  h.textContent =
    "Private Jobs from other companies (which may be smaller or not that famous)";

  const icons = document.createElement("span");

  for (const a of agencyIcons) {
    const img = document.createElement("img");
    img.src = a.src;
    img.title = a.title;
    img.alt = a.title;
    img.classList.add("agency-icon");
    icons.appendChild(img);
  }

  h.classList.add("other-agencies");
  h.appendChild(icons);
  element.appendChild(h);

  const elems = otherPrivateJobs().map((p) => {
    const title = p["jobTitle"],
      div = document.createElement("div");

    div.classList.add("job-post-private");
    div.setAttribute("title", title);

    div.innerHTML = `
       <h3 style="padding:3px;">${title}</h3>
       <br/>
       <hr/>

       <br/>

       <p><strong class="job-field" style="padding:3px;" >${p["jobSpecFields"]}</strong></p>

       ${
         p.province
           ? `<span><p class="province" style="padding:5px;">${p.province}</p></span>`
           : ""
       }

       ${
         p.location
           ? `<span >
              <p class="location" style="padding:3px;">${p.location["region"].replace(",", "")}</p>
              <p class="location" style="padding:3px;">${p.location["city"].replace(",", "")}</p>
           </span>`
           : ""
       }
       <button class="more">more</button>
       `;
       /**
        *   ${
         p.expiryDate
           ? `<p class="expiry-date" style="width:${p.expiryDate.length}ch; padding:3px;">${p.expiryDate}</p>`
           : ""
       }

       ${
         p.startDate
           ? `<p class="start-date" style="width:${
               p.startDate.length + 13
             }ch; padding:3px;">start date: ${p.startDate}</p>`
           : ""
       }


       ${
         p.vacancyType
           ? `<p class="job-type" style="width:${
               p.vacancyType.length + 13
             }ch; padding:3px;" >Vacancy type: ${p.vacancyType}</p>`
           : ""
       }
        *   <br/>

       <section class="details">${p.details.replace(
         /\. /gi,
         ".<br/><br/>"
       )}</section>
       
       <br/>
       <a href=${p.apply} target="_blank">
        <button class="apply">
          apply
        </button>
       </a>
        * **/

    return div;
  });

  const privateSectorBoard = document.createElement("section");
  privateSectorBoard.classList.add("board");
  privateSectorBoard.appendChild(h);

  const posts = document.createElement("section");

  for (const e of elems) {
    posts.appendChild(e);
  }
  privateSectorBoard.appendChild(posts);
  element.appendChild(privateSectorBoard);
}
