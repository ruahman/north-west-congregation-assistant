import request from "utils/src/request";

const { VITE_API_URL: url } = import.meta.env;

export default {
  url,
  getTerritory() {
    request.get(this.url);
  },
  getMessages() {
    throw new Error("TODO: Implement this function");
  },
  postMessage() {
    throw new Error("TODO: Implement this function");
  },
  postReturnVisit() {
    throw new Error("TODO: Implement this function");
  },
  shareReturnVisit() {
    throw new Error("TODO: Implement this function");
  },
};
