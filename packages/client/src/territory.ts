import { HereMap } from "@utils/map/here";
import { _ } from "@utils/dom";
import PouchDB from "pouchdb";

// this creates here-map tag
new HereMap();

_.ready(async function () {
  const $ = _.$(document);

  const circuit = $("meta[name=circuit]").get("content");
  const congregation = $("meta[name=congregation]").get("content");
  const territory = $("meta[name=territory]").get("content");

  const db = new PouchDB(`${circuit}/${congregation}/${territory}`);

  try {
    const response = await db.put({
      _id: new Date().toISOString(),
      name: "Sample Document",
      type: "example",
    });
    console.log("Document added:", response);
  } catch (error) {
    console.error("Error adding document:", error);
  }

  // setTimeout(() => {
  //   const $elem = $("here-map");
  //   $elem.set("test", "here-wego after 2 seconds");
  // }, 2000);

  // setTimeout(() => {
  //   const $elem = $("here-map");
  //   const map = $elem[0] as HereMap;
  //   map.test = "here-wego after 4 seconds";
  // }, 4000);
});
