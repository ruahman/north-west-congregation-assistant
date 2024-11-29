import { test } from "bun:test";
import request from ".";

test("request get", () => {
  request.get("test/me");
});
