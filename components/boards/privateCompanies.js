import data from "../../backend/database/public/govpage-private-sector.json";

export function setupPrivateCompanies(element) {
  const elems = data["blogPosts"].map((p) => {
    const div = document.createElement("div");
    div.classList.add("job-post")
    div.innerHTML = `
        <div class="company-logo">
            <img src=${p["imgSrc"]} alt="company logo" />
        </div>
        <p class="title">${p["title"].replace(/is hiring/gi, "")}</p>
        `;

    return div;
  });

  for (const e of elems) {
    element.appendChild(e);
  }
}
