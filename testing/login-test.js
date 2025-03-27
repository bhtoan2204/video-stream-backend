import http from "k6/http";
import { check, sleep } from "k6";
import { SharedArray } from "k6/data";
import papaparse from "https://jslib.k6.io/papaparse/5.1.1/index.js";

const users = new SharedArray("user credentials", function () {
  return papaparse.parse(open("./users.csv"), { header: true }).data;
});

export const options = {
  vus: 20,
  iterations: users.length, // run once per user
};

export default function () {
  const user = users[__VU - 1]; // __VU is the current virtual user index (1-based)

  const payload = JSON.stringify({
    email: user.email,
    password: user.password,
  });

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const res = http.post(
    "http://localhost:8080/api/v1/user-service/auth/login",
    payload,
    params
  );

  check(res, {
    "is status 200": (r) => r.status === 200,
    "response has token": (r) =>
      r.json("data.result.access_token") !== undefined,
  });

  sleep(1);
}
