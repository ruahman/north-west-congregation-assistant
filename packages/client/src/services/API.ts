import request from "utils/src/request";

let url = Bun.env.API!;

export default {
  url,
  getTerritory() {
    request.get(this.url);
  },
  getMessages() {},
  postMessage() {},
  postReturnVisit() {},
  shareReturnVisit() {},
};
