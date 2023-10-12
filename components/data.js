import propersonnelData from "../backend/database/private/Pro-Personnel.json";
import heithaData from "../backend/database/private/heitha-stuffing-group.json";
import privateSectorData from "../backend/database/public/govpage-private-sector.json";
import publicSectorData from "../backend/database/public/govpage-public-sector.json";

export function publicJobs() {
  return publicSectorData;
}

export function privateJobs() {
  return privateSectorData;
}

export function propersonnelJobs() {
  return propersonnelData;
}
export function heithaJobs() {
  return heithaData;
}

export const stats = {
  govDep: publicSectorData.blogPosts.length,
  privateComp: privateSectorData.blogPosts.length,
  privateSectorOpenings:
    heithaData.blogPosts.length + propersonnelData.blogPosts.length
};
