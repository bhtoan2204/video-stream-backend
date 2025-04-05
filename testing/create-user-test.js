import http from "k6/http";
import { check, sleep } from "k6";
import { faker } from "https://cdn.skypack.dev/@faker-js/faker";

export const options = {
  vus: 20,
  iterations: 100, // run once per user
};

export default function () {
  const payload = JSON.stringify({
    username: faker.internet.username(),
    password: "@Toan123456",
    email: faker.internet.email(),
    phone: faker.phone.number({ style: "international" }),
    first_name: faker.person.firstName(),
    last_name: faker.person.lastName(),
    birth_date: faker.date
      .birthdate({ min: 18, max: 65, mode: "age" })
      .toISOString()
      .split("T")[0],
    address: faker.location.streetAddress(),
  });

  console.log(payload);

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const res = http.post(
    "http://localhost:8080/api/v1/user-service/users",
    payload,
    params
  );

  if (res.body) {
    check(res, {
      "has data": (r) => r.json("data") !== undefined,
    });
  } else {
    console.error("Response body is null");
  }

  sleep(1);
}
