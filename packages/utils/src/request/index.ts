const Request = {
  get(url: string) {
    console.log("get: ", url);
  },
  post(url: string, data: any) {
    throw new Error("TODO: Implement this function");
  },
};

export default Request;
