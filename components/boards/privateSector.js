import { otherPrivateJobs } from "../data.js";

export function setupPrivateSector(element) {
  const h = document.createElement("h3");
  h.textContent =
    "Private Jobs from other companies (which may be smaller or not that famous)";

  element.appendChild(h);

  const elems = otherPrivateJobs().map((p) => {
    const title = p["jobTitle"],
      div = document.createElement("div");

    div.classList.add("job-post-private");
    div.setAttribute("title", title);

    div.innerHTML = `
       <h3>${title}</h3>
       <br/>
       <p><strong>${p["jobSpecFields"]}</strong></p>
       ${p.province ? `<p>${p.province}</p>` : ""}
       ${
         p.location
           ? `<p> ${p.location["region"]}\n ${p.location["city"]}</p>`
           : ""
       }
       ${p.expiryDate ? `<p>${p.expiryDate}</p>` : ""}
       ${p.startDate ? `<p> start date: ${p.startDate}</p>` : ""}

       ${p.vacancyType ? `<p> vacancy type: ${p.vacancyType}</p>` : ""}
       <br/>
       <section>${p.details.replace(/\. /gi, ".<br/><br/>")}</section>
       <br/>
       <button id="apply">
        <a href=${p.apply} target="_blank">apply</a>
       </button>
       `;

    return div;
  });

  const privateSectorBoard = document.createElement("section");
  //   privateSectorBoard.classList.add("board");
  privateSectorBoard.appendChild(h);

  const posts = document.createElement("section");
  //   posts.classList.add("posts");
  for (const e of elems) {
    posts.appendChild(e);
  }
  privateSectorBoard.appendChild(posts);
  element.appendChild(privateSectorBoard);
}
