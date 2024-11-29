import { HereMap } from "@utils/map/here";
import { Query } from "@utils/dom";

let $ = Query(document);

// this creates here-map tag
new HereMap();

$(function () {
  setTimeout(() => {
    const $elem = $("here-map");
    $elem.set("test", "here-wego after 2 seconds");
  }, 2000);

  setTimeout(() => {
    const $elem = $("here-map");
    const map = $elem[0] as HereMap;
    map.test = "here-wego after 4 seconds";
  }, 4000);
});
