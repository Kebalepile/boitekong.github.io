import { privateJobs } from "../data.js";

export function setupPrivateCompanies(element) {
  const h = document.createElement("h3");
  h.textContent = "Private Companies hiring";

  element.appendChild(h);

  const elems = privateJobs()["blogPosts"].map((p) => {
    const title = p["title"].replace(/is hiring/gi, ""),
      div = document.createElement("div");

    div.classList.add("job-post");
    div.setAttribute("title", title);

    div.innerHTML = `
        <div class="company-logo">
            <img src=${p["imgSrc"]} alt="company logo" />
        </div>
        <p class="title">${title}</p>
        `;

    const dialog = document.getElementById("dialog");
    div.addEventListener("click", (e) => {
      dialog.showModal();
    });
    return div;
  });

  const privateSectorBoard = document.createElement("section");
  privateSectorBoard.classList.add("board");
  privateSectorBoard.appendChild(h);

  const posts = document.createElement("section");
  posts.classList.add("posts");
  for (const e of elems) {
    posts.appendChild(e);
  }
  privateSectorBoard.appendChild(posts);
  element.appendChild(privateSectorBoard);
}
