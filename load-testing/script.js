import http from "k6/http";
import { sleep } from "k6";

export let options = {
  stages: [
    { duration: "1m", target: 10000 }, // Ramp-up to 10,000 users over 1 minute
  ],
};

export default function () {
  // make a GET request to the target URL
  http.get("http://localhost:8080/");
  // simulate think time
  sleep(1);
}

// You can run the test by executing the following command:
// k6 run k6-test.js
